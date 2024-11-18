package handlers

import (
	"fmt"
	"os"
	"time"
	"v/pkg/chat"
	w "v/pkg/webrtc"

	"crypto/sha256"

	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

// RoomCreate handles redirecting to a new room URL
func RoomCreate(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("/room/%s", guuid.New().String()))
}

// Room handles rendering the room page with WebSocket addresses
func Room(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID is required"})
		return
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	uuid, suuid, _ := createOrGetRoom(uuid)

	c.HTML(http.StatusOK, "peer.html", gin.H{
		"RoomWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/websocket", ws, c.Request.Host, uuid),
		"RoomLink":            fmt.Sprintf("%s://%s/room/%s", c.Request.URL.Scheme, c.Request.Host, uuid),
		"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/chat/websocket", ws, c.Request.Host, uuid),
		"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/room/%s/viewer/websocket", ws, c.Request.Host, uuid),
		"StreamLink":          fmt.Sprintf("%s://%s/stream/%s", c.Request.URL.Scheme, c.Request.Host, suuid),
		"Type":                "room",
	})
}

// RoomWebsocket handles the WebSocket connection for the room
func RoomWebsocket(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)

	// conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 200)
	// if err != nil {
	// 	return
	// }

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all origins, but you can restrict this based on your security needs
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Handle the error appropriately, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	w.RoomConn(conn, room.Peers)
}

// createOrGetRoom creates or retrieves an existing room
func createOrGetRoom(uuid string) (string, string, *w.Room) {
	w.RoomsLock.Lock()
	defer w.RoomsLock.Unlock()

	h := sha256.New()
	h.Write([]byte(uuid))
	suuid := fmt.Sprintf("%x", h.Sum(nil))

	// Check if room exists
	if room := w.Rooms[uuid]; room != nil {
		if _, ok := w.Streams[suuid]; !ok {
			w.Streams[suuid] = room
		}
		return uuid, suuid, room
	}

	// Create a new room and hub
	hub := chat.NewHub()
	p := &w.Peers{}
	p.TrackLocals = make(map[string]*webrtc.TrackLocalStaticRTP)
	room := &w.Room{
		Peers: p,
		Hub:   hub,
	}

	w.Rooms[uuid] = room
	w.Streams[suuid] = room

	go hub.Run()
	return uuid, suuid, room
}

// RoomViewerWebsocket handles WebSocket connection for the room viewer
func RoomViewerWebsocket(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return
	}

	w.RoomsLock.Lock()
	if peer, ok := w.Rooms[uuid]; ok {
		w.RoomsLock.Unlock()
		roomViewerConn(c.Writer, peer.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

// roomViewerConn manages the viewer connection
func roomViewerConn(w http.ResponseWriter, p *w.Peers) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	conn, err := websocket.Upgrade(w, nil, nil, 0, 0)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		select {
		case <-ticker.C:
			// Send number of connections as message
			err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", len(p.Connections))))
			if err != nil {
				return
			}
		}
	}
}

// WebsocketMessage structure for handling WebSocket messages
type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// package handlers

// import (
// 	"fmt"
// 	"os"
// 	"time"
// 	"v/pkg/chat"
// 	w "v/pkg/webrtc"

// 	"crypto/sha256"

// 	guuid "github.com/google/uuid"
// 	"github.com/gorilla/websocket"
// 	"github.com/pion/webrtc/v3"
// )

// func RoomCreate(c *fiber.Ctx) error {
// 	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
// }

// func Room(c *fiber.Ctx) error {
// 	uuid := c.Params("uuid")
// 	if uuid == "" {
// 		c.Status(400)
// 		return nil
// 	}

// 	ws := "ws"
// 	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
// 		ws = "wss"
// 	}

// 	uuid, suuid, _ := createOrGetRoom(uuid)

// 	return c.Render("peer", fiber.Map{
// 		"RoomWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/websocket", ws, c.Hostname(), uuid),
// 		"RoomLink":            fmt.Sprintf("%s://%s/room/%s", c.Protocol(), c.Hostname(), uuid),
// 		"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/chat/websocket", ws, c.Hostname(), uuid),
// 		"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/room/%s/viewer/websocket", ws, c.Hostname(), uuid),
// 		"StreamLink":          fmt.Sprintf("%s://%s/stream/%s", c.Protocol(), c.Hostname(), suuid),
// 		"Type":                "room",
// 	}, "layouts/main")
// }

// func RoomWebsocket(c *websocket.Conn) {
// 	uuid := c.Params("uuid")
// 	if uuid == "" {
// 		return
// 	}

// 	_, _, room := createOrGetRoom(uuid)
// 	w.RoomConn(c, room.Peers)
// }

// func createOrGetRoom(uuid string) (string, string, *w.Room) {
// 	w.RoomsLock.Lock()
// 	defer w.RoomsLock.Unlock()

// 	h := sha256.New()
// 	h.Write([]byte(uuid))
// 	suuid := fmt.Sprintf("%x", h.Sum(nil))

// 	if room := w.Rooms[uuid]; room != nil {
// 		if _, ok := w.Streams[suuid]; !ok {
// 			w.Streams[suuid] = room
// 		}
// 		return uuid, suuid, room
// 	}

// 	hub := chat.NewHub()
// 	p := &w.Peers{}
// 	p.TrackLocals = make(map[string]*webrtc.TrackLocalStaticRTP)
// 	room := &w.Room{
// 		Peers: p,
// 		Hub:   hub,
// 	}

// 	w.Rooms[uuid] = room
// 	w.Streams[suuid] = room

// 	go hub.Run()
// 	return uuid, suuid, room
// }

// func RoomViewerWebsocket(c *websocket.Conn) {
// 	uuid := c.Params("uuid")
// 	if uuid == "" {
// 		return
// 	}

// 	w.RoomsLock.Lock()
// 	if peer, ok := w.Rooms[uuid]; ok {
// 		w.RoomsLock.Unlock()
// 		roomViewerConn(c, peer.Peers)
// 		return
// 	}
// 	w.RoomsLock.Unlock()
// }

// func roomViewerConn(c *websocket.Conn, p *w.Peers) {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()
// 	defer c.Close()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			w, err := c.Conn.NextWriter(websocket.TextMessage)
// 			if err != nil {
// 				return
// 			}
// 			w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
// 		}
// 	}
// }

// type websocketMessage struct {
// 	Event string `json:"event"`
// 	Data  string `json:"data"`
// }
