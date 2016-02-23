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
	*reply = "  _____    ____     ____    ______         _______           _   _      \n"
	*reply += "  / ____|  / __ \\   / __ \\  |  ____|       |__   __|         | | | |     \n"
	*reply += " | |  __  | |  | | | |  | | | |__             | |      __ _  | | | | __  \n"
	*reply += " | | |_ | | |  | | | |  | | |  __|            | |     / _` | | | | |/ /  \n"
	*reply += " | |__| | | |__| | | |__| | | |               | |    | (_| | | | |   <   \n"
	*reply += "  \\_____|  \\____/   \\____/  |_|               |_|     \\__,_| |_| |_|\\_\\  v1.0\n"
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
func (c *ChatServer) Logout(username string, reply *[]string) error {
	for i, val := range c.users {
		if val == username {
			c.users = append(c.users[:i], c.users[i+1:]...) //deletes the user from the array(slice)
			break
		}
	}
	log.Printf("%s has left the chat", username)
	return nil

}
func (c *ChatServer) Shutdown(nothing Nothing, reply *Nothing) error {
          var rep []string 
          for _,val := range c.users {

          c.Logout(val, &rep)
        }
	log.Println("Server shutdown...Goodbye.")
	*reply = false
	c.shutdown <- true

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
