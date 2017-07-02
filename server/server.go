package server

import (
	"context"
	"net/http"
)

//Srv is the static file server
var Srv *http.Server

//SetupServer starts the local static file server as well as the
//GUI file server
func SetupServer(path, port string) {
	Srv = &http.Server{
		Addr:    ":" + port,
		Handler: http.FileServer(http.Dir(path)),
	}
	Srv.SetKeepAlivesEnabled(false)
	go Srv.ListenAndServe()
}

//ChangeRoot changes the root file path from the system that serves
func ChangeRoot(path, port string) {
	Srv.Shutdown(context.Background())
	Srv = &http.Server{
		Addr:    ":" + port,
		Handler: http.FileServer(http.Dir(path)),
	}
	go Srv.ListenAndServe()
}
