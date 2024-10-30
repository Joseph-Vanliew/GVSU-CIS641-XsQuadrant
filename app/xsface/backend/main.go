// Package main implements a WebRTC signaling server using the Gin framework for HTTP handling.
// It allows users to connect and exchange SDP (Session Description Protocol) to establish media connections.
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"xsface/initializers"

	"xsface/signaling" // Custom package for signaling

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

const (
	rtcpPLIInterval = time.Second * 3
)

type Sdp struct {
	Sdp string
}

type PeerTracks struct {
	AudioChannel chan *webrtc.Track
	VideoChannel chan *webrtc.Track
}

func main() {

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1", "10.0.0.0/8", "localhost"})
	router.Use(cors.Default())

	peerConnectionMap := make(map[string]*PeerTracks)

	m := webrtc.MediaEngine{}
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))
	m.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))

	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	router.POST("/webrtc/meeting/:meetingID/peer/:peerID/isAdmin/:isAdmin", func(c *gin.Context) {
		isAdmin, _ := strconv.ParseBool(c.Param("isAdmin"))
		peerID := c.Param("peerID")

		var session Sdp
		if err := c.ShouldBindJSON(&session); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		offer := webrtc.SessionDescription{}
		signaling.Decode(session.Sdp, &offer)

		peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
		if err != nil {
			log.Fatal(err)
		}

		if !isAdmin {
			recieveTrack(peerConnection, peerConnectionMap, peerID)
		} else {
			createTrack(peerConnection, peerConnectionMap, peerID)
		}

		peerConnection.SetRemoteDescription(offer)

		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			log.Fatal(err)
		}

		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, Sdp{Sdp: signaling.Encode(answer)})
	})

	router.Run(":8080")
}

// Adjusted recieveTrack function with logging and transceiver check
func recieveTrack(peerConnection *webrtc.PeerConnection, peerConnectionMap map[string]*PeerTracks, peerID string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if peerTracks, ok := peerConnectionMap[peerID]; ok {
				// Add video track to peer connection
				if videoTrack := <-peerTracks.VideoChannel; videoTrack != nil {
					if _, err := peerConnection.AddTrack(videoTrack); err != nil {
						log.Printf("Error adding video track for peer %s: %v", peerID, err)
					} else {
						log.Printf("Video Track added successfully to peer %s", peerID)
					}
				}

				// Add audio track to peer connection
				if audioTrack := <-peerTracks.AudioChannel; audioTrack != nil {
					if _, err := peerConnection.AddTrack(audioTrack); err != nil {
						log.Printf("Error adding audio track for peer %s: %v", peerID, err)
					} else {
						log.Printf("Audio Track added successfully to peer %s", peerID)
					}
				}
				return
			} else {
				log.Printf("Channel not ready for peer %s. Retrying...", peerID)
			}
		}
	}
}

func createTrack(peerConnection *webrtc.PeerConnection, peerConnectionMap map[string]*PeerTracks, currentPeerID string) {

	// Initialize PeerTracks if it doesn't exist for the currentPeerID
	if _, ok := peerConnectionMap[currentPeerID]; !ok {
		peerConnectionMap[currentPeerID] = &PeerTracks{
			AudioChannel: make(chan *webrtc.Track, 1),
			VideoChannel: make(chan *webrtc.Track, 1),
		}
	}

	// // Video track setup
	// videoTrack, newVideoTrackErr := peerConnection.NewTrack(webrtc.DefaultPayloadTypeVP8, rand.Uint32(), "video", "pion-video")
	// if newVideoTrackErr != nil {
	// 	log.Fatal(newVideoTrackErr)
	// }
	// peerConnectionMap[currentPeerID].VideoChannel <- videoTrack

	// // Audio track setup
	// audioTrack, newAudioTrackErr := peerConnection.NewTrack(webrtc.DefaultPayloadTypeOpus, rand.Uint32(), "audio", "pion-audio")
	// if newAudioTrackErr != nil {
	// 	log.Fatal(newAudioTrackErr)
	// }
	// peerConnectionMap[currentPeerID].AudioChannel <- audioTrack

	// Explicitly add transceivers for bidirectional communication
	// if _, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
	// 	log.Fatal("Failed to add video transceiver:", err)
	// }
	// if _, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
	// 	log.Fatal("Failed to add audio transceiver:", err)
	// }

	_, err := peerConnection.AddTransceiverFromKind(
		webrtc.RTPCodecTypeVideo,
		webrtc.RtpTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionSendrecv,
		},
	)
	if err != nil {
		log.Fatal("Failed to add video transceiver:", err)
	}

	_, err = peerConnection.AddTransceiverFromKind(
		webrtc.RTPCodecTypeAudio,
		webrtc.RtpTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionSendrecv,
		},
	)
	if err != nil {
		log.Fatal("Failed to add audio transceiver:", err)
	}

	peerConnection.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		log.Println("**********on track received")
		go func() {
			ticker := time.NewTicker(rtcpPLIInterval)
			defer ticker.Stop() // Ensure ticker is stopped to prevent resource leaks
			for range ticker.C {
				if rtcpSendErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: remoteTrack.SSRC()}}); rtcpSendErr != nil {
					log.Println("RTCP PLI Error:", rtcpSendErr)
					// return
				}
			}
		}()

		var videoTrack *webrtc.Track
		var audioTrack *webrtc.Track
		var newVideoTrackErr error
		var newAudioTrackErr error

		switch remoteTrack.Kind() {
		case webrtc.RTPCodecTypeVideo:

			// Video track setup
			videoTrack, newVideoTrackErr = peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "video", "pion")
			if newVideoTrackErr != nil {
				log.Fatal(newVideoTrackErr)
			}
			peerConnectionMap[currentPeerID].VideoChannel <- videoTrack

			log.Println("********** withing video track")
		case webrtc.RTPCodecTypeAudio:
			// Audio track setup
			audioTrack, newAudioTrackErr = peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "audio", "pion")
			if newAudioTrackErr != nil {
				log.Fatal(newAudioTrackErr)
			}
			peerConnectionMap[currentPeerID].AudioChannel <- audioTrack
			log.Println("********** withing audio track")

		}

		rtpBuf := make([]byte, 1400)
		for {
			i, readErr := remoteTrack.Read(rtpBuf)
			if readErr != nil {
				log.Printf("Error reading track: %v", readErr)
				continue
			}
			log.Println("********** after remoteTrack.Read(rtpBuf)")

			// Only write if track is not nil and has not been closed
			switch remoteTrack.Kind() {
			case webrtc.RTPCodecTypeVideo:
				log.Println("********** wrting video outer")

				if videoTrack != nil {
					log.Println("********** wrting video inner")

					if _, err := videoTrack.Write(rtpBuf[:i]); err != nil {
						log.Println("Error writing video track:", err)
						continue
					} else {
						log.Println("Video outside")
					}
					log.Println("Video written successfully.")

				}
			case webrtc.RTPCodecTypeAudio:
				log.Println("********** wrting audio outer")

				if audioTrack != nil {
					log.Println("********** wrting audio inner")

					if _, err := audioTrack.Write(rtpBuf[:i]); err != nil {
						log.Println("Error writing audio track:", err)
						continue
					} else {
						log.Println("Audio written successfully.")
					}
					log.Println("Audio outside")

				}
			}
		}
	})
}
