package server

import (
	"flag"
	"fmt"
	"os"
	"time"
	"v/backend/initializers"
	"v/backend/routes"

	"v/internal/handlers"
	w "v/pkg/webrtc"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

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

	// logging
	r.Use(logger.SetLogger())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define Routes
	routes.RegisterRoutes(r)

	r.GET("api/room/create", handlers.RoomCreate)
	r.GET("api/room/:uuid", handlers.Room)
	r.GET("api/room/:uuid/websocket", func(c *gin.Context) {
		handlers.RoomWebsocket(c)
	})
	r.GET("api/room/:uuid/chat", handlers.RoomChat)
	r.GET("api/room/:uuid/chat/websocket", func(c *gin.Context) {
		handlers.RoomChatWebsocket(c)
	})
	r.GET("api/room/:uuid/viewer/websocket", func(c *gin.Context) {
		handlers.RoomViewerWebsocket(c)
	})
	r.GET("api/stream/:suuid", handlers.Stream)
	r.GET("api/stream/:suuid/websocket", func(c *gin.Context) {
		handlers.StreamWebsocket(c)
	})
	r.GET("api/stream/:suuid/chat/websocket", func(c *gin.Context) {
		handlers.StreamChatWebsocket(c)
	})
	r.GET("api/stream/:suuid/viewer/websocket", func(c *gin.Context) {
		handlers.StreamViewerWebsocket(c)
	})

	r.Static("/static", "./frontend/dist")

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	// Initialize WebRTC rooms and streams
	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)
	fmt.Print("Initialized Rooms")

	// Start background process
	go dispatchKeyFrames()
	fmt.Print("Started KeyFrames Thread")

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
