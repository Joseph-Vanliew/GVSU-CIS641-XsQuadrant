package handlers

import (
	"fmt"
	"os"
	"time"
	w "v/pkg/webrtc"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Stream handles the streaming page rendering
func Stream(c *gin.Context) {
	suuid := c.Param("suuid")
	if suuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "suuid is required"})
		return
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	w.RoomsLock.Lock()
	if _, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		c.HTML(http.StatusOK, "stream.html", gin.H{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", ws, c.Request.Host, suuid),
			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Request.Host, suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", ws, c.Request.Host, suuid),
			"Type":                "stream",
		})
		return
	}
	w.RoomsLock.Unlock()

	c.HTML(http.StatusOK, "stream.html", gin.H{
		"NoStream": "true",
		"Leave":    "true",
	})
}

// StreamWebsocket handles WebSocket connection for the stream
func StreamWebsocket(c *gin.Context) {
	suuid := c.Param("suuid")
	if suuid == "" {
		return
	}

	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 0, 0)
		if err != nil {
			return
		}
		w.StreamConn(conn, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

// StreamViewerWebsocket handles WebSocket connection for the stream viewer
func StreamViewerWebsocket(c *gin.Context) {
	suuid := c.Param("suuid")
	if suuid == "" {
		return
	}

	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 0, 0)
		if err != nil {
			return
		}
		viewerConn(conn, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

// viewerConn manages the connection for stream viewers
func viewerConn(c *websocket.Conn, p *w.Peers) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer c.Close()

	for {
		select {
		case <-ticker.C:
			w, err := c.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
		}
	}
}

// package handlers

// import (
// 	"fmt"
// 	"os"
// 	"time"
// 	w "v/pkg/webrtc"

// )

// func Stream(c *fiber.Ctx) error {
// 	suuid := c.Params("suuid")
// 	if suuid == "" {
// 		c.Status(400)
// 		return nil
// 	}

// 	ws := "ws"
// 	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
// 		ws = "wss"
// 	}

// 	w.RoomsLock.Lock()
// 	if _, ok := w.Streams[suuid]; ok {
// 		w.RoomsLock.Unlock()
// 		return c.Render("stream", fiber.Map{
// 			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", ws, c.Hostname(), suuid),
// 			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
// 			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", ws, c.Hostname(), suuid),
// 			"Type":                "stream",
// 		}, "layouts/main")
// 	}
// 	w.RoomsLock.Unlock()

// 	return c.Render("stream", fiber.Map{
// 		"NoStream": "true",
// 		"Leave":    "true",
// 	}, "layouts/main")
// }

// func StreamWebsocket(c *websocket.Conn) {
// 	suuid := c.Params("suuid")
// 	if suuid == "" {
// 		return
// 	}

// 	w.RoomsLock.Lock()
// 	if stream, ok := w.Streams[suuid]; ok {
// 		w.RoomsLock.Unlock()
// 		w.StreamConn(c, stream.Peers)
// 		return
// 	}
// 	w.RoomsLock.Unlock()
// }

// func StreamViewerWebsocket(c *websocket.Conn) {
// 	suuid := c.Params("suuid")
// 	if suuid == "" {
// 		return
// 	}

// 	w.RoomsLock.Lock()
// 	if stream, ok := w.Streams[suuid]; ok {
// 		w.RoomsLock.Unlock()
// 		viewerConn(c, stream.Peers)
// 		return
// 	}
// 	w.RoomsLock.Unlock()
// }

// func viewerConn(c *websocket.Conn, p *w.Peers) {
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
