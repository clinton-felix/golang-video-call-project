package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/template/html"
	"github.com/gofiber/websocket/v2"
)

// defining the address, certificate and key variabes
var (
	addr = flag.String("addr, :", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")
)

func Run() {
	flag.Parse()

	if *addr == ":"{
		*addr = ":8080"
	}

	// Defining fiber engine to access app.
	// setting up the logger and the cors
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())


	// Here we are defining the routes and the actions that
	// should be performed once a user hits those routes using fiber
	app.Get("/", handlers.Welcome)					// root route runs the welcome handler
	app.Get("/room/create", handlers.RoomCreate)	// route creates a room
	app.Get("/room/:uuid", handlers.Room)			// takes the user to the unique room
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		handshakeTimeout : 10 * time.Second,
	}))				
	app.Get("/room/:uuid/chat", handlers.RoomChat)	// chat endpoint
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))	// realtime chatting connection route
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)		// stream route
	app.Get("/stream/:ssuid/websocket", ha)			
	app.Get("/stream/:ssuid/chat/websocket", )		
	app.Get("/stream/:ssuid/viewer/websocket", )	
}