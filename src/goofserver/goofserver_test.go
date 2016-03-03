// Unit tests for the package goofserver which
// mocks certain conditions and tests the functions in goofserver package
package goofserver

import (
	"fmt"
	"testing"
)

// Define a port number for server
var c = ChatServer{Port: "3410"}

// A dummy empty variable
var none Nothing

//TestRegisterGoofs_1 test mocks registration of new user with valid name
// Expected result: PASS
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

//TestRegisterGoofs_1 test mocks registration of new user with empty string
//Expected result: FAIL
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

//TestRegisterGoofs_1 test mocks registration of new user with existing username
//Expected result: FAIL
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

//TestListGoofs_1 test mocks listing of online users
//Expected result: PASS
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

//TestListGoofs_1 test mocks listing of online users when no users present
//Expected result: FAIL
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

//TestLogout_1 test mocks a user logging out
//Expected result: PASS
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

//TestLogout_1 test mocks a user logging out when given user is not online
//Expected result: FAIL
func TestLogout_2(t *testing.T) {
	c.Users = []string{"testman", "goofdemo"}
	var demouser = "nouser"
	err := c.Logout(demouser, &none)
	if err != nil {
		t.Error("Could not log out correctly. ")
	} else {
		t.Log("Logout successfull")
	}
}

//TestWhisper_1 test tries to send a message to online user
//Expected result: PASS
func TestWhisper_1(t *testing.T) {
	c.Users = []string{"goof1", "goof2", "goof3", "goof4"}
	var demouser = "goof4"
	var target = "goof3"
	message := Message{
		User:   demouser,
		Target: target,
		Msg:    "hello this is a dummy communication message",
	}
	demomsg := []string{"hello"}
	c.MessageQueue = make(map[string][]string, len(demomsg))
	c.MessageQueue[target] = demomsg
	err := c.Whisper(message, &none)
	if err != nil {
		t.Error("Unable to send the message to target", err)
	} else {
		t.Log("Message sending test passed")
	}
}

//TestWhisper_2 test tries to send a message to an non existant user
//Expected result: FAIL
func TestWhisper_2(t *testing.T) {
	c.Users = []string{"goof1", "goof2", "goof3", "goof4"}
	var demouser = "goof4"
	var target = "goof6"
	message := Message{
		User:   demouser,
		Target: target,
		Msg:    "hello this is a dummy communication message",
	}
	err := c.Whisper(message, &none)
	if err != nil {
		t.Error("Unable to send the message to target", err)
	} else {
		t.Log("Message sending test passed")
	}
}

//TestWhisper_3 test tries to send a message with more than 160 characters
//Expected result: FAIL
func TestWhisper_3(t *testing.T) {
	c.Users = []string{"goof1", "goof2", "goof3", "goof4"}
	var demouser = "goof4"
	var target = "goof2"
	message := Message{
		User:   demouser,
		Target: target,
		Msg:    "This is a message to test whipser function. This suppose to have more than 160 characters.This is a message to test whipser function. And this is more than one sixty characters",
	}
	demomsg := []string{"hello"}
	c.MessageQueue = make(map[string][]string, len(demomsg))
	c.MessageQueue[target] = demomsg
	err := c.Whisper(message, &none)
	if err != nil {
		t.Error("Unable to send the message to target", err)
	} else {
		t.Log("Message sending test passed")
	}
}
