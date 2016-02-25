package main

import "testing"

// mocks a user logout function
func TestLogout(t *testing.T) {
	c := ChatServer{port: "3410",
		users: []string{"testman", "demo"},
	}
	var none Nothing
	var demouser = "testman"
	err := c.Logout(demouser, &none)
	if err != nil {
		t.Error("Could not log in correctly. ")
	} else {
		t.Log("Logout successfull")
	}
}
