package server

import "net/http"

//SetupServer starts the local static file server as well as the
//GUI file server
func SetupServer(path, port string) *http.Server {
	srv := &http.Server{Addr: port}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
	go srv.ListenAndServe()
	return srv
}
