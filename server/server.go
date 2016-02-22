package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Nothing bool

type Message struct {
	User   string
	Target string
	Msg    string
}

type ChatServer struct {
	port         string
	messageQueue map[string][]string
	users        []string
	shutdown     chan bool
}


func (c *ChatServer) RegisterGoofs(username string, reply *string) error {
	*reply = "Welcome to GOOF TALK\n"
	*reply += "List of GOOFS online:\n"

	c.users = append(c.users, username)
	c.messageQueue[username] = nil

	for _, value := range c.users {
		*reply += value + "\n"
	}

	for k, _ := range c.messageQueue {
		c.messageQueue[k] = append(c.messageQueue[k], username+" has joined.")
	}

	log.Printf("%s has joined the chat.\n", username)

	return nil
}


func (c *ChatServer) ListGoofs(none Nothing, reply *[]string) error {
	*reply = append(*reply, "Current online Goofs:")

	for i := range c.users {
		*reply = append(*reply, c.users[i])
	}

	log.Println("Dumped list of Goofs to client output")

	return nil
}


func parseFlags(cs *ChatServer) {
	flag.StringVar(&cs.port, "port", "3410", "port for chat server to listen on")
	flag.Parse()

	cs.port = ":" + cs.port
}

func RunServer(cs *ChatServer) {
	rpc.Register(cs)
	rpc.HandleHTTP()

	log.Printf("Listening on port %s...\n", cs.port)

	l, err := net.Listen("tcp", cs.port)
	if err != nil {
		log.Panicf("Can't bind port to listen. %q", err)
	}

	go http.Serve(l, nil)
}

func main() {
	cs := new(ChatServer)
	cs.messageQueue = make(map[string][]string)
	cs.shutdown = make(chan bool, 1)

	parseFlags(cs)
	RunServer(cs)

	<-cs.shutdown
}
