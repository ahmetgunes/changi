package main

import (
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal"
	"os"
)

func main() {
	changi.Init()
	env := os.Getenv("ENV")
	if env == "dev" || env == "test" {
		internal.Test()
	} else {
		internal.Listen()
	}
}
