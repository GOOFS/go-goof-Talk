package main

import (
	"goofclient"
	"log"
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
	// Listen for messages
	goofclient.MainLoop(client)

}
