package goofserver

import (
	"fmt"
	"testing"
)

// mocks the logout function with sample data
var c = ChatServer{Port: "3410"}
var none Nothing

//unit test case of RegisterGoofs with valid data
func TestRegisterGoofs_1(t *testing.T) {
	var reply string
	var user = "goof"
	err := c.RegisterGoofs(user, &reply)
	if err != nil {
		t.Error("Error while registering", err)
	} else {
		t.Log("Register successful")
	}
}

//unit test case of RegisterGoofs with empty string
func TestRegisterGoofs_2(t *testing.T) {
	var reply string
	var user = ""
	err := c.RegisterGoofs(user, &reply)
	if err != nil {
		t.Error("Error while registering", err)
	} else {
		t.Log("Register successful")
	}
}

//unit test case of RegisterGoofs with duplicate data
func TestRegisterGoofs_3(t *testing.T) {
	var reply string
	c.Users = []string{"testman", "goofdemo"}
	var user = "testman"
	err := c.RegisterGoofs(user, &reply)
	if err != nil {
		t.Error("Error while registering", err)
	} else {
		t.Log("Register successful")
	}
}

//unit test case of ListGoofs with valid data
func TestListGoofs_1(t *testing.T) {
	var reply []string
	c.Users = []string{"testman", "goofdemo"}
	err := c.ListGoofs(none, &reply)
	if err != nil {
		t.Error("Error listing users:", err)
	} else {
		for i := range reply {
			fmt.Println(reply[i])
		}
	}
}

//unit test case of ListGoofs with invalid data
func TestListGoofs_2(t *testing.T) {
	var reply []string
	c.Users = []string{}
	err := c.ListGoofs(none, &reply)
	if err != nil {
		t.Error("Error listing users:", err)
	} else {
		for i := range reply {
			fmt.Println(reply[i])
		}
	}
}

//unit test case of Logout with valid data
func TestLogout_1(t *testing.T) {
	c.Users = []string{"testman", "goofdemo"}
	var demouser = "goofdemo"
	err := c.Logout(demouser, &none)
	if err != nil {
		t.Error("Could not log out correctly. ")
	} else {
		t.Log("Logout successfull")
	}
}

//unit test case of Logout with invalid data
func TestLogout_2(t *testing.T) {
	c.Users = []string{"testman", "goofdemo"}
	var demouser = "Joseph"
	err := c.Logout(demouser, &none)
	if err != nil {
		t.Error("Could not log out correctly. ")
	} else {
		t.Log("Logout successfull")
	}
}

func TestWhisper_1(t *testing.T) {
	c.Users = []string{"goof1", "goof2", "goof3", "goof4"}
	var demouser = "goof4"
	message := Message{
		User:   demouser,
		Target: "goof1",
		Msg:    "hello this is a dummy communication message",
	}
	err := c.Whisper(message, &none)
	if err != nil {
		t.Error("Unable to send the message to target", err)
	} else {
		t.Log("Message sending test passed")
	}
}
