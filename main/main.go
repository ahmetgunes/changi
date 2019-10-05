package main

import (
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/internal"
)

func main() {
	config := configs.Init()
	internal.Test(config)
}
