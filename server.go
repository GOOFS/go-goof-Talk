package main

import "goofserver"

func main() {
	cs := goofserver.StartServer()
	cs.MessageQueue = make(map[string][]string)
	cs.ShutDown = make(chan bool, 1)

	goofserver.ParseFlags(cs)
	goofserver.RunServer(cs)

	<-cs.ShutDown
}
