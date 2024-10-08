package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	webport = ":8080"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// setup http server
	http.HandleFunc("/ws", handleWebSocket)
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	fmt.Printf("Starting server  at http://localhost%s\n", webport)
	log.Fatal(http.ListenAndServe(webport, nil))

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			continue
		}

		switch msg["type"] {
		case "join":
			log.Printf("%s joined the session", msg["name"])
		case "offer":
			// Handle offer message
		case "answer":
			// Handle answer message
		case "candidate":
			// Handle ICE candidate message
		}
	}
}

// func CreatePeerConnection() (*webrtc.PeerConnection, error) {

// 	iceServer := []webrtc.ICEServer{
// 		{
// 			URLs: []string{"stun:stun.l.google.com.19302"},
// 		},
// 	}

// 	// Create a new RTCPeerConnection
// 	config := webrtc.Configuration{
// 		ICEServers: iceServer,
// 	}
// 	peerConnection, err := webrtc.NewPeerConnection(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Handle ICE connection state changes
// 	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
// 		fmt.Printf("ICE Connection State has changed: %s\n", state.String())
// 	})

// 	return peerConnection, nil
// }
