package handlers

import (
	"net/http"
	"v/pkg/chat"
	w "v/pkg/webrtc"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// RoomChat handles rendering the chat page
func RoomChat(c *gin.Context) {
	// Render the chat page (assuming you have a layout system set up in Gin)
	c.HTML(http.StatusOK, "chat.html", gin.H{})
}

// RoomChatWebsocket handles the WebSocket connection for the room's chat
func RoomChatWebsocket(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return
	}

	// Locking to prevent race condition
	w.RoomsLock.Lock()
	room := w.Rooms[uuid]
	w.RoomsLock.Unlock()

	if room == nil || room.Hub == nil {
		return
	}

	// Use WebSocket connection from the request context
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 0, 0)
	if err != nil {
		return
	}

	// Handle the chat connection for the peer
	chat.PeerChatConn(conn, room.Hub)
}

// StreamChatWebsocket handles WebSocket connection for the stream's chat
func StreamChatWebsocket(c *gin.Context) {
	suuid := c.Param("suuid")
	if suuid == "" {
		return
	}

	// Locking to prevent race condition
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		if stream.Hub == nil {
			hub := chat.NewHub()
			stream.Hub = hub
			go hub.Run()
		}

		// Use WebSocket connection from the request context
		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 0, 0)
		if err != nil {
			return
		}

		// Handle the chat connection for the peer
		chat.PeerChatConn(conn, stream.Hub)
		return
	}
	w.RoomsLock.Unlock()
}
