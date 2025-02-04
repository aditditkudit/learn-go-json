package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJSON(t *testing.T) {
	jsonString := `{"id":"1","name":"Tas","image_url":"tas.jpeg","price":10000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapJSONEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "1",
		"name":      "Tas",
		"image_url": "tas.jpeg",
		"price":     10000,
	}

	jsonString, _ := json.Marshal(product)
	fmt.Println(string(jsonString))
}
