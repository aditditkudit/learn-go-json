package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Price    int64  `json:"price"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Tas",
		ImageUrl: "tas.jpeg",
		Price:    10000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Tas","image_url":"tas.jpeg","price":10000}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
