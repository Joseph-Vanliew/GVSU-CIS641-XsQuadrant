package server

import (
	"flag"
	"os"
	"time"

	"v/internal/handlers"
	w "v/pkg/webrtc"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

var (
	addr = flag.String("addr", ":"+os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	r := gin.Default()

	// Setup logging middleware
	r.Use(logger.SetLogger())

	// Setup CORS middleware
	r.Use(cors.Default())

	// Load HTML templates (using html/template package)
	r.LoadHTMLGlob("./views/*.html")

	// Define Routes
	r.GET("/", handlers.Welcome)
	r.GET("/room/create", handlers.RoomCreate)
	r.GET("/room/:uuid", handlers.Room)
	r.GET("/room/:uuid/websocket", func(c *gin.Context) {
		handlers.RoomWebsocket(c)
	})
	r.GET("/room/:uuid/chat", handlers.RoomChat)
	r.GET("/room/:uuid/chat/websocket", func(c *gin.Context) {
		handlers.RoomChatWebsocket(c)
	})
	r.GET("/room/:uuid/viewer/websocket", func(c *gin.Context) {
		handlers.RoomViewerWebsocket(c)
	})
	r.GET("/stream/:suuid", handlers.Stream)
	r.GET("/stream/:suuid/websocket", func(c *gin.Context) {
		handlers.StreamWebsocket(c)
	})
	r.GET("/stream/:suuid/chat/websocket", func(c *gin.Context) {
		handlers.StreamChatWebsocket(c)
	})
	r.GET("/stream/:suuid/viewer/websocket", func(c *gin.Context) {
		handlers.StreamViewerWebsocket(c)
	})

	// Static file serving
	r.Static("/", "./assets")

	// Initialize WebRTC rooms and streams
	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)

	// Start background process
	go dispatchKeyFrames()

	// Run the server with or without TLS
	if *cert != "" {
		return r.RunTLS(*addr, *cert, *key)
	}
	return r.Run(*addr)
}

// Background function to dispatch keyframes at a regular interval
func dispatchKeyFrames() {
	for range time.NewTicker(time.Second * 3).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrame()
		}
	}
}
