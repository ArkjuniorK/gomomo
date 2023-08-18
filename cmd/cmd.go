package cmd

import "flag"

// Flag set to use by the app
var (
	fs flag.FlagSet

	port = fs.Int("port", 8080, "Port to use by the app")
	host = fs.String("host", "localhost", "Host to use by the app")
)

// This function use to initialize, configure and run the entire app.
func init() {

}
