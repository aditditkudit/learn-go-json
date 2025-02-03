package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       12,
		Hobbies:   []string{"Gaming", "Reading", "Watching"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Doe","Age":12,"Hobbies":["Gaming","Reading","Watching"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Hobbies[0])
	fmt.Println(customer.Hobbies[1])
	fmt.Println(customer.Hobbies[2])
}

func TestJsonComplex(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		Addresses: []Address{
			{
				Street:  "Muncang Blok N",
				City:    "North Jakarta",
				Country: "Indonesia",
			},
			{
				Street:  "Serumpun",
				City:    "Kuala Lumpur",
				Country: "Malaysia",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"","Age":0,"Hobbies":null,"Addresses":[{"Street":"Muncang Blok N","City":"North Jakarta","Country":"Indonesia"},{"Street":"Serumpun","City":"Kuala Lumpur","Country":"Malaysia"}]}`

	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses[0].Street)
	fmt.Println(customer.Addresses[1].Country)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Muncang Blok N","City":"North Jakarta","Country":"Indonesia"},{"Street":"Serumpun","City":"Kuala Lumpur","Country":"Malaysia"}]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

	//var data []Address
	//var err = json.Unmarshal([]byte(jsonString), &data)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(data)
	//fmt.Println(data[0].Street)
}

func TestOnlyJsonArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:  "Muncang Blok N",
			City:    "North Jakarta",
			Country: "Indonesia",
		},
		{
			Street:  "Serumpun",
			City:    "Kuala Lumpur",
			Country: "Malaysia",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
