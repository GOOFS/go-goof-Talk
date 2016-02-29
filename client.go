package main

import (
	"goofclient"
	"log"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	// Set MAX PROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start by parsing any flags given to the program
	client, err := goofclient.CreateClientFromFlags()
	if err != nil {
		log.Panicf("Error creating client from flags: %q", err)
	}

	client.RegisterGoofs()
	go client.CheckMessages()

	//Capture ctrl+c and logout if pressed
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		_ = <-sigc
		client.Logout()
	}()

	// Listen for messages
	goofclient.MainLoop(client)
}
