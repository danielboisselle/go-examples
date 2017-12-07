package main

import (
	"flag"

	"github.com/mccurdyc/go-examples/example2/transports/http/service"
)

// global variables
var (
	serverHost = flag.String("host", "localhost", "server host")
	serverPort = flag.Int("port", 8080, "server port")
)

// init gets called before main()
func init() {
	flag.Parse()
}

func main() {
	// pass in the value of serverHost and serverPort
	s := service.NewService(*serverHost, *serverPort)
	s.Start()
}
