// Package goofclient consists of structs and their methods which can be used to
// connect to the host, send or receive messages through the server
package goofclient

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"
)

//Nothing defines a blank variable
type Nothing bool

//Message defines struct for every message
type Message struct {
	User   string
	Target string
	Msg    string
}

//ChatClient defines struct for each of the ChatClient
type ChatClient struct {
	Username string
	Address  string
	Client   *rpc.Client
}

// Global variables to store default port and host.
var (
	DefaultPort = 3410
	DefaultHost = "localhost"
)

// getClientConnection function dials into given host and returns the client variable if success, error otherwise
func (c *ChatClient) getClientConnection() *rpc.Client {
	var err error

	if c.Client == nil {
		c.Client, err = rpc.DialHTTP("tcp", c.Address)
		if err != nil {
			log.Panicf("Error establishing connection with host: %q", err)
		}
	}
	return c.Client
}

// RegisterGoofs function takes a username and registers it with the server
func (c *ChatClient) RegisterGoofs() {
	var reply string
	c.Client = c.getClientConnection()

	err := c.Client.Call("ChatServer.RegisterGoofs", c.Username, &reply)

	if err != nil {
		fmt.Printf("Error registering user: %q\n", err)
		fmt.Println("Enter new GOOF name:")
		fmt.Scanln(&c.Username)
		c.RegisterGoofs()
	} else {
		fmt.Printf("\n %s", reply)
	}
}

// CheckMessages does a check every second for new messages for the user
func (c *ChatClient) CheckMessages() {
	var reply []string
	c.Client = c.getClientConnection()
	for {
		err := c.Client.Call("ChatServer.CheckMessages", c.Username, &reply)
		if err != nil {
			log.Fatalln("Chat has been shutdown. Goodbye.")
		}
		for i := range reply {
			log.Println(reply[i])
		}

		time.Sleep(time.Second)
	}
}

//ListGoofs function lists all the users in the chat currently
func (c *ChatClient) ListGoofs() {
	var reply []string
	var none Nothing
	c.Client = c.getClientConnection()

	err := c.Client.Call("ChatServer.ListGoofs", none, &reply)
	if err != nil {
		log.Printf("Error listing users: %q\n", err)
	}

	for i := range reply {
		fmt.Println(reply[i])
	}
}

// Whisper function sends a message to a specific user
func (c *ChatClient) Whisper(params []string) {
	var reply Nothing
	c.Client = c.getClientConnection()
	target := strings.Replace(params[0], "@", "", 1)
	if len(params) == 2 {
		msg := strings.Join(params[1:], " ")
		message := Message{
			User:   c.Username,
			Target: target,
			Msg:    msg,
		}

		err := c.Client.Call("ChatServer.Whisper", message, &reply)
		if err != nil {
			log.Printf("Error telling users something: %q", err)
		}
	} else {
		log.Println("Usage of whisper: @<username> <your message>")
	}
}


//Logout function logouts a goof out
func (c *ChatClient) Logout() {
	var reply Nothing
	err := c.Client.Call("ChatServer.Logout", c.Username, &reply)
	if err != nil {
		log.Printf("Error logging out: %q\n", err)
	} else {
		log.Println("Logged out Succesfully")
		os.Exit(0)
	}
}

//CreateClientFromFlags function parses the command line arguments
func CreateClientFromFlags() (*ChatClient, error) {
	var c = &ChatClient{}
	var host string

	flag.StringVar(&c.Username, "user", "Goof", "Your username")
	flag.StringVar(&host, "host", "localhost", "The host you want to connect to")

	flag.Parse()
	if c.Username == "Goof" {
		fmt.Println("Enter your Goof ID: ")
		fmt.Scanln(&c.Username)
	}
	if !flag.Parsed() {
		return c, errors.New("Unable to create user from commandline flags. Please try again")
	}

	// Check for the structure of the flag to see if we can make any educated guesses for them
	if len(host) != 0 {

		if strings.HasPrefix(host, ":") { // Begins with a colon means :3410 (just port)
			c.Address = DefaultHost + host
		} else if strings.Contains(host, ":") { // Contains a colon means host:port
			c.Address = host
		} else { // Otherwise, it's just a host
			c.Address = net.JoinHostPort(host, strconv.Itoa(DefaultPort))
		}

	} else {
		c.Address = net.JoinHostPort(DefaultHost, strconv.Itoa(DefaultPort)) // Default to our default port and host
	}

	return c, nil
}

// MainLoop function waits for input from stadard input i.e. keyboard and checks it against list of available functions
func MainLoop(c *ChatClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error: %q\n", err)
		}

		line = strings.TrimSpace(line)
		params := strings.Fields(line)
		if strings.HasPrefix(line, "list") {
			c.ListGoofs()
		} else if strings.HasPrefix(line, "@") {
			c.Whisper(params)
		} else if strings.HasPrefix(line, "logout") {
			c.Logout()
		} else if strings.HasPrefix(line, "help") {
			fmt.Println("Welcome to GOOFtalk help:")
			fmt.Println("List of funcitons, \n1. List all online Goofs : list\n2. Whisper: @<username> <message>\n3.Logout:  logout")
		} else {
			fmt.Println("Invalid function, try 'help' to list all available functions")
		}

	}
}
