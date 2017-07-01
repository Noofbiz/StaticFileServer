package main

import (
	"github.com/Noofbiz/StaticFileServer/gui"
	"github.com/Noofbiz/StaticFileServer/server"
)

func main() {
	path, port := readConfig()

	srv := server.SetupServer(path, port)

	gui.StartGUI(srv)
}
