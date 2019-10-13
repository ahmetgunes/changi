package main

import (
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal"
	"github.com/ahmetgunes/changi/internal/web"
	"os"
)

func main() {
	changi.Init()
	env := os.Getenv("ENV")
	if env == "dev" || env == "test" {
		internal.Test()
	} else {
		web.Listen()
	}
}
