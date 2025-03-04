package learn_go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Adhitya",
		LastName:  "Adit",
		Age:       25,
		Hobbies:   []string{"Gaming", "Reading", "Watching"},
		Addresses: []Address{
			{
				Street:  "Muncang Blok N",
				City:    "North Jakarta",
				Country: "Indonesia",
			},
		},
	}
	_ = encoder.Encode(customer)
}
