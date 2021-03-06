package gui

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Noofbiz/StaticFileServer/configuration"
	"github.com/Noofbiz/StaticFileServer/server"
	astilectron "github.com/asticode/go-astilectron"
)

//StartGUI starts up the astillectron window system.
func StartGUI(path, port string) {
	p, _ := os.Getwd()
	var a *astilectron.Astilectron
	var err error
	if a, err = astilectron.New(astilectron.Options{
		AppName:            "Static File Server",
		AppIconDefaultPath: filepath.Join(p, "assets", "icon.png"),
		BaseDirectoryPath:  p,
	}); err != nil {
		log.Fatalf("Failed to create new astillectron. Error: %v", err.Error())
	}
	defer a.Close()
	a.HandleSignals()

	if err = a.Start(); err != nil {
		log.Fatalf("Failed to start. Error: %v", err.Error())
	}

	var w *astilectron.Window
	if w, err = a.NewWindow(filepath.Join(p, "assets", "html", "gui.html"), &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(800),
	}); err != nil {
		log.Fatalf("Failed to create new window. Error: %v", err.Error())
	}
	if err = w.Create(); err != nil {
		log.Fatalf("Failed at window.Create(). Error: %v", err.Error())
	}

	msg := []string{path, port}
	w.SendMessage(msg)

	w.OnMessage(func(e *astilectron.EventMessage) interface{} {
		var m string
		if err = e.Unmarshal(&m); err != nil {
			log.Fatalf("Recieved improper message from gui. Error: %v", err.Error())
		}
		pathport := strings.Split(m, "<a-o>")
		path = pathport[0]
		port = pathport[1]
		configuration.UpdateConfig(path, port)
		server.ChangeRoot(path, port)
		return nil
	})

	a.Wait()
}
