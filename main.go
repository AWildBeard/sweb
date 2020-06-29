package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port string
	listenAddress string
	serveDir string
)

func init() {
	flag.StringVar(&port, "p", "8080", "The port that the webserver will bind to")
	flag.StringVar(&listenAddress, "a", "0.0.0.0", "The default listen address on the host that the webserver will bind to. This should only be an ip address")
	flag.StringVar(&serveDir, "d", ".", "The directory to serve content from")
}

func main() {
	flag.Parse()

	listenAddress = listenAddress + ":" + port

	fs := http.FileServer(http.Dir(serveDir))
	http.Handle("/", fs)

	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		fmt.Printf("Error in http server: %v\n", err)
	}
}
