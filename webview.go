package main

import (
	"github.com/zserge/webview"

	"github.com/PeterBooker/webview/internal/config"
	"github.com/PeterBooker/webview/internal/log"
	"github.com/PeterBooker/webview/internal/server"
)

var (
	version string
	commit  string
	date    string
)

func main() {

	// Create Logger
	l := log.New()

	// Create Config
	c := config.Setup(version, commit, date)

	// Setup server struct to hold all App data
	s := server.New(l, c)

	// Setup HTTP server.
	s.Setup()

	// Open in a 800x600 resizable window
	webview.Open(
		"Webview Boilerplate",
		"http://localhost:8080/",
		800,
		600,
		true,
	)
}
