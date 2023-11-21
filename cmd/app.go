package cmd

import (
	"flag"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app"
)

// Flag set to use by the app
var (
	fs flag.FlagSet

	host = fs.String("host", "localhost", "Host to use by the app")
	port = fs.String("port", "8080", "Port to use by the app")

	//httpPort = fs.String("httpport", "8080", "http transport port")
	//grpcPort = fs.String("grpcport", "8081", "gRPC transport port")
	//amqpPort = fs.String("amqpport", "8082", "AMQP transport port")
)

// This function use to init, configure and run the entire app.
func init() {
	app.New(*host, *port).Run()
}
