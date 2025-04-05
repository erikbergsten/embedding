package components

import (
	"html/template"
	"log/slog"
	"os"
	"embed"
)

type Code struct {
	Language string
	Text string
}

type Greeting string

type Contact struct {
	Name string
	Address string
}

//go:embed *.html
var fs embed.FS

var Tpl * template.Template
func init() {
	slog.Info("Initializing compoonents")
	var err error
	Tpl, err = template.New("").ParseFS(fs, "*.html")
	if err != nil {
		slog.Error("Error initializing components.", "error", err)
		os.Exit(1)
	}
}

