package main

import (
	"video-transcription/config"
	"video-transcription/ginreg"
)

func main() {
	// Set environment variables
	config.LoadConfig()

	//Create a channel to wait for the server to finish
	done := make(chan bool)

	// Start the server in a goroutine
	go func() {
		ginreg.GinInit()
		done <- true
	}()

	// Wait for the server to finish
	<-done
}
