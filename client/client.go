package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Nothing bool

type Message struct {
	User   string
	Target string
	Msg    string
}

type ChatClient struct {
	Username string
	Address  string
	Client   *rpc.Client
}

// Globals/Constants
var (
	DEFAULT_PORT = 3410
	DEFAULT_HOST = "localhost"
)

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

// Register takes a username and registers it with the server
func (c *ChatClient) RegisterGoofs() {
	var reply string
	c.Client = c.getClientConnection()

	err := c.Client.Call("ChatServer.RegisterGoofs", c.Username, &reply)
	if err != nil {
		log.Printf("Error registering user: %q", err)
	} else {
		fmt.Printf("\n %s", reply)
	}
}

// List lists all the users in the chat currently
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

// logout a Goof
func (c *ChatClient) Logout() {
	var reply []string
	err := c.Client.Call("ChatServer.Logout", c.Username, &reply)
	if err != nil {
		log.Printf("Error logging out: %q\n", err)
	} else {
		log.Println("Logged out Succesfully")
		os.Exit(0)
	}

}

// Parse the command list arguments
func createClientFromFlags() (*ChatClient, error) {
	var c *ChatClient = &ChatClient{}
	var host string

	flag.StringVar(&c.Username, "user", "Goof", "Your username")
	flag.StringVar(&host, "host", "localhost", "The host you want to connect to")

	flag.Parse()

	if !flag.Parsed() {
		return c, errors.New("Unable to create user from commandline flags. Please try again")
	}

	// Check for the structure of the flag to see if we can make any educated guesses for them
	if len(host) != 0 {

		if strings.HasPrefix(host, ":") { // Begins with a colon means :3410 (just port)
			c.Address = DEFAULT_HOST + host
		} else if strings.Contains(host, ":") { // Contains a colon means host:port
			c.Address = host
		} else { // Otherwise, it's just a host
			c.Address = net.JoinHostPort(host, strconv.Itoa(DEFAULT_PORT))
		}

	} else {
		c.Address = net.JoinHostPort(DEFAULT_HOST, strconv.Itoa(DEFAULT_PORT)) // Default to our default port and host
	}

	return c, nil
}

func mainLoop(c *ChatClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error: %q\n", err)
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "listGoofs") {
			c.ListGoofs()
		} else if strings.HasPrefix(line, "logout") {
			c.Logout()
		} else if strings.HasPrefix(line, "help") {
			fmt.Println("Welcome to GOOFtalk help:")
			fmt.Println("List of funcitons, \n1. listGoofs\n2. logout")
		} else {
			fmt.Println("Invalid function, try 'help' to list all available functions")
		}

	}
}

func main() {
	// Set MAX PROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start by parsing any flags given to the program
	client, err := createClientFromFlags()
	if err != nil {
		log.Panicf("Error creating client from flags: %q", err)
	}

	client.RegisterGoofs()

	// Listen for messages
	mainLoop(client)
}
