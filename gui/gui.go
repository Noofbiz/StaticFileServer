package gui

import (
	"log"
	"os"

	astilectron "github.com/asticode/go-astilectron"
)

//StartGUI starts up the astillectron window system.
func StartGUI(port string) {
	p, _ := os.Getwd()
	var a *astilectron.Astilectron
	var err error
	if a, err = astilectron.New(astilectron.Options{
		AppName:            "Static File Server",
		AppIconDefaultPath: p + "gui/assets/icon.png",
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
	if w, err = a.NewWindow("http://localhost"+port, &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(800),
	}); err != nil {
		log.Fatalf("Failed to create new window. Error: %v", err.Error())
	}
	if err = w.Create(); err != nil {
		log.Fatalf("Failed at window.Create(). Error: %v", err.Error())
	}

	a.Wait()
}
