package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, City, Country string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
