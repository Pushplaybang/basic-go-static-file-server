package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var dir string

	// parse command line args
	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to server")
	flag.Parse()

	// default to working dir if none is provided
	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	// print to the console
	fmt.Println("Server started, serving " + *path + " directory at port " + *port)

	// listen and serve files
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
