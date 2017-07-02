package main

import (
	"github.com/Noofbiz/StaticFileServer/configuration"
	"github.com/Noofbiz/StaticFileServer/gui"
	"github.com/Noofbiz/StaticFileServer/server"
)

func main() {
	path, port := configuration.ReadConfig()

	server.SetupServer(path, port)

	gui.StartGUI(path, port)
}
