// All material is licensed under the GNU Free Documentation License
// https://github.com/gobridge/gotraining/blob/master/LICENSE

// https://play.golang.org/p/Bn-7NAh9di

/*
An interface is a reference type who's header is a two word value. The
first word represents the type of the value and the second is the data
itself.

Something like this:
type interface struct {
       Type uintptr     // points to the type of the interface implementation
       Data uintptr     // holds the data for the interface's receiver
}

interface{} represents the empty set of methods and is satisfied by any
value at all, since any value has zero or more methods.
*/

// Sample program to show how the empty interface works.
package main

import (
	"encoding/json"
	"fmt"
)

// User is a sample struct.
type User struct {
	Name  string
	Email string
}

// JSONString converts any value into a JSON string.
func JSONString(value interface{}) string {
	data, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		return ""
	}

	return string(data)
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type User.
	user := User{
		Name:  "Henry Ford",
		Email: "henry@ford.com",
	}

	fmt.Println(JSONString(&user))
}
