package server

import (
	"html/template"
	"net/http"
)

//SetupServer starts the local static file server as well as the
//GUI file server
func SetupServer(path, port string) {
	http.HandleFunc("/gui", guiHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
	go http.ListenAndServe(port, nil)
}

func guiHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("gui/assets/html/*"))
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
