// Package main implements a WebRTC signaling server using the Gin framework for HTTP handling.
// It allows users to connect and exchange SDP (Session Description Protocol) to establish media connections.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"xsface/controllers"
	"xsface/initializers"
	"xsface/middleware"

	"xsface/signaling" // Custom package for signaling

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"  // HTTP framework for routing and middleware
	"github.com/pion/rtcp"      // For handling RTCP packets like PLI
	"github.com/pion/webrtc/v2" // WebRTC implementation
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// Interval for sending Picture Loss Indication (PLI) to request keyframes
const (
	rtcpPLIInterval = time.Second * 3
)

// Sdp represents the SDP payload used to describe media communication sessions
type Sdp struct {
	Sdp string
}

func main() {
	// Open the log file for writing logs. It creates or appends to "info.log"
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err) // Terminate if the log file can't be opened
	}
	defer file.Close()      // Ensure the file is closed when the program exits
	log.SetOutput(file)     // Set the output of the logger to the log file
	router := gin.Default() // Create a new Gin router instance for handling HTTP requests

	//testing out end points
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Set trusted proxies (e.g., "127.0.0.1", "10.0.0.0/8")
	router.SetTrustedProxies([]string{"127.0.0.1", "10.0.0.0/8", "localhost"})
	router.Use(cors.Default())

	// Map to store channels associated with each peer's tracks
	peerConnectionMap := make(map[string]chan *webrtc.Track)

	m := webrtc.MediaEngine{}
	// Register VP8 codec for video compression
	// This codec ensures that video data can be compressed and decompressed properly
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))

	// Create a new WebRTC API instance with the configured media engine
	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	// Configuration for the peer connection, specifying ICE servers for NAT traversal
	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // Public STUN server for ICE
			},
		},
	}

	// Define an HTTP POST route for handling SDP exchanges
	router.POST("/webrtc/sdp/m/:meetingId/c/:userId/p/:peerId/s/:isSender", func(c *gin.Context) {
		// Parse the URL parameters and the JSON payload
		isSender, _ := strconv.ParseBool(c.Param("isSender"))
		userID := c.Param("userId")
		peerID := c.Param("peerId")

		log.Println("******userID", userID)
		log.Println("******peerID", peerID)

		var session Sdp
		if err := c.ShouldBindJSON(&session); err != nil {
			// Return a 400 status if there's an error binding the JSON payload
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Decode the SDP from the payload
		offer := webrtc.SessionDescription{}
		signaling.Decode(session.Sdp, &offer)

		// Create a new peer connection with the configuration
		peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
		if err != nil {
			log.Fatal(err) // Log and terminate if creating the peer connection fails
		}

		// Determine if the user is a sender or receiver, and set up tracks accordingly
		if !isSender {
			recieveTrack(peerConnection, peerConnectionMap, peerID)
		} else {
			createTrack(peerConnection, peerConnectionMap, userID)
		}

		// Set the remote description from the SDP offer
		peerConnection.SetRemoteDescription(offer)

		// Create an SDP answer for the peer
		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			log.Fatal(err)
		}

		// Set the local description for the connection
		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			log.Fatal(err)
		}

		// Return the encoded SDP answer as a JSON response
		c.JSON(http.StatusOK, Sdp{Sdp: signaling.Encode(answer)})
	})

	router.Run(":8080") // Start the HTTP server on port 8080
}

// receiveTrack sets up a track for receiving media from a peer.
// If the peer connects before the user, it creates a channel to hold the track and waits for the peer to add it.
// If the peer connects later, it uses the existing channel to add the track.
func recieveTrack(peerConnection *webrtc.PeerConnection,
	peerConnectionMap map[string]chan *webrtc.Track,
	peerID string) {

	// Check if a channel already exists for this peer; if not, create one
	if _, ok := peerConnectionMap[peerID]; !ok {
		peerConnectionMap[peerID] = make(chan *webrtc.Track, 1)
	}
	// Retrieve the local track from the channel and add it to the peer connection
	localTrack := <-peerConnectionMap[peerID]
	peerConnection.AddTrack(localTrack)
}

// createTrack sets up a track for sending media to a peer.
// It creates a new track and either stores it in a new channel or uses an existing one if the peer is already listening.
func createTrack(peerConnection *webrtc.PeerConnection,
	peerConnectionMap map[string]chan *webrtc.Track,
	currentUserID string) {

	// Add a transceiver for video (RTP) to the peer connection
	if _, err := peerConnection.AddTransceiver(webrtc.RTPCodecTypeVideo); err != nil {
		log.Fatal(err) // Terminate if adding the transceiver fails
	}

	// Set a handler for when a new remote track is received
	peerConnection.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		// Send Picture Loss Indications (PLI) periodically to request keyframes from the sender
		go func() {
			ticker := time.NewTicker(rtcpPLIInterval)
			for range ticker.C {
				if rtcpSendErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: remoteTrack.SSRC()}}); rtcpSendErr != nil {
					fmt.Println(rtcpSendErr) // Log any errors related to RTCP packet sending
				}
			}
		}()

		// Create a new local track for streaming the received video to other peers
		localTrack, newTrackErr := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "video", "pion")
		if newTrackErr != nil {
			log.Fatal(newTrackErr) // Terminate if the track creation fails
		}

		// Create a channel to store the local track
		localTrackChan := make(chan *webrtc.Track, 1)
		localTrackChan <- localTrack

		// If a channel already exists for this user, send the track; otherwise, create one
		if existingChan, ok := peerConnectionMap[currentUserID]; ok {
			existingChan <- localTrack
		} else {
			peerConnectionMap[currentUserID] = localTrackChan
		}

		// Buffer for receiving RTP packets from the remote track
		rtpBuf := make([]byte, 1400)
		for { // Loop to continuously read RTP packets from the publisher
			i, readErr := remoteTrack.Read(rtpBuf)
			if readErr != nil {
				log.Fatal(readErr) // Terminate if there's an error reading the track
			}

			// Write RTP packets to the local track, unless there are no subscribers
			if _, err := localTrack.Write(rtpBuf[:i]); err != nil && err != io.ErrClosedPipe {
				log.Fatal(err) // Terminate if writing to the track fails
			}
		}
	})
}
