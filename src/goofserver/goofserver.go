//Package goofserver defines methods which will manage
// connections and all the messages sent by them.
package goofserver

import (
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

//Nothing defines a blank variable
type Nothing bool

//Message defines struct for every message
type Message struct {
	User   string
	Target string
	Msg    string
}

//ChatServer defines struct for each of the ChatServer
type ChatServer struct {
	Port         string
	MessageQueue map[string][]string
	Users        []string
	ShutDown     chan bool
}

// RegisterGoofs function takes a username and registers it with the server
func (c *ChatServer) RegisterGoofs(username string, reply *string) error {
	*reply = "  _____    ____     ____    ______         _______           _   _      \n"
	*reply += "  / ____|  / __ \\   / __ \\  |  ____|       |__   __|         | | | |     \n"
	*reply += " | |  __  | |  | | | |  | | | |__             | |      __ _  | | | | __  \n"
	*reply += " | | |_ | | |  | | | |  | | |  __|            | |     / _` | | | | |/ /  \n"
	*reply += " | |__| | | |__| | | |__| | | |               | |    | (_| | | | |   <   \n"
	*reply += "  \\_____|  \\____/   \\____/  |_|               |_|     \\__,_| |_| |_|\\_\\  v1.0\n"
	*reply += "List of GOOFS online:\n"
	if len(strings.Trim(username, " ")) == 0 {
		err := errors.New("Username cannot be blank.")
		return err
	}

	for _, val := range c.Users {
		if val == username {
			err := errors.New("Username already taken.")
			return err
		}
	}
	c.Users = append(c.Users, username)

	for _, value := range c.Users {
		*reply += value + "\n"
	}

	for k := range c.MessageQueue {
		c.MessageQueue[k] = append(c.MessageQueue[k], "["+username+"] has joined.")
	}

	log.Printf("[%s] has joined the chat.\n", username)

	return nil
}

// CheckMessages does a check every second for new messages for the user
func (c *ChatServer) CheckMessages(username string, reply *[]string) error {
	*reply = c.MessageQueue[username]
	c.MessageQueue[username] = nil
	return nil
}

//ListGoofs function lists all the users in the chat currently
func (c *ChatServer) ListGoofs(none Nothing, reply *[]string) error {
	*reply = append(*reply, "Current online Goofs:")
	if len(c.Users) == 0 {
		err := errors.New("No online users")
		return err
	}
	for i := range c.Users {
		*reply = append(*reply, c.Users[i])
	}

	log.Println("Dumped list of Goofs to client output")

	return nil
}

// Whisper function sends a message to a specific user
func (c *ChatServer) Whisper(msg Message, reply *Nothing) error {
	if len(msg.Msg) > 160 {
		err := errors.New("Maximum length of the message should be less than 160")
		return err
	}
	if queue, ok := c.MessageQueue[msg.Target]; ok {
		m := "[" + msg.User + "] :  " + msg.Msg
		c.MessageQueue[msg.Target] = append(queue, m)
	} else {
		err := errors.New("[" + msg.Target + "] does not exist. use 'list' command to list online Goofs.")
		return err
	}

	*reply = false

	return nil
}

//Logout function logouts a goof out
func (c *ChatServer) Logout(username string, reply *Nothing) error {
	var none Nothing
	var i int
	var val string
	for i, val = range c.Users {
		if val == username {
			c.Users = append(c.Users[:i], c.Users[i+1:]...) //deletes the user from the array(slice)

			for k := range c.MessageQueue {
				c.MessageQueue[k] = append(c.MessageQueue[k], "["+username+"] has left the chat.")
			}
			log.Printf("[%s] has left the chat", username)
			break
		}
	}
	if i == len(c.Users) && val != username {
		err := errors.New("Unable to logout")
		return err
	}
	if len(c.Users) == 0 {
		c.Shutdown(none, &none)
	}
	return nil
}

//Shutdown function shuts the server down if whether evrybody in the chatroom logs out
func (c *ChatServer) Shutdown(nothing Nothing, reply *Nothing) error {
	log.Println("Everybody left the chat. Server is Shutting down...")
	*reply = false
	c.ShutDown <- true
	return nil
}

// ParseFlags function parses the command line arguments
func ParseFlags(cs *ChatServer) {
	flag.StringVar(&cs.Port, "port", "3410", "port for chat server to listen on")
	flag.Parse()
	cs.Port = ":" + cs.Port
}

// RunServer function starts listening to the given or the default port
func RunServer(cs *ChatServer) {
	rpc.Register(cs)
	rpc.HandleHTTP()

	log.Printf("Listening on port %s...\n", cs.Port)

	l, err := net.Listen("tcp", cs.Port)
	if err != nil {
		log.Panicf("Can't bind port to listen. %q", err)
	}

	go http.Serve(l, nil)
}

//StartServer creates new instance of the struct ChatServer and returns it
func StartServer() *ChatServer {
	cs := new(ChatServer)
	return cs
}
