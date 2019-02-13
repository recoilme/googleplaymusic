package main

import "github.com/zserge/webview"

func main() {
	// Open GooglePlayMusic in a 960x600 resizable window
	w := webview.New(webview.Settings{
		Title:     "GooglePlayMusic",
		URL:       "https://play.google.com/music/", //listen?u=0&nomobilenative=1#/uqg",
		Debug:     false,
		Width:     960,
		Height:    600,
		Resizable: true,
	})
	defer w.Exit()
	w.Run()

}
