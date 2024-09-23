package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Define structs to match the JSON structure
type Keys struct {
	N int `json:"n"`
	K int `json:"k"`
}

type Item struct {
	Base  string `json:"base"`
	Value string `json:"value"`
}

type Data struct {
	Keys Keys            `json:"keys"`
	Nums map[string]Item `json:"-"` // We will manually unmarshal these fields
}

func main() {
	// Open the JSON file
	file, err := os.Open("testcase1.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Declare a Data struct to hold the parsed JSON
	var data Data

	// First, unmarshal the known part of the JSON
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("Failed to unmarshal keys: %s", err)
	}

	// Now, manually handle the rest of the numbered keys
	// Using a map[string]interface{} for flexibility
	var temp map[string]interface{}
	if err := json.Unmarshal(byteValue, &temp); err != nil {
		log.Fatalf("Failed to unmarshal temp map: %s", err)
	}

	data.Nums = make(map[string]Item)

	// Iterate through the map and process numbered keys
	for key, value := range temp {
		if key != "keys" { // Skip the "keys" object
			if valMap, ok := value.(map[string]interface{}); ok {
				base := valMap["base"].(string)
				val := valMap["value"].(string)
				data.Nums[key] = Item{Base: base, Value: val}
			}
		}
	}

	// Output the data
	fmt.Printf("Keys: N = %d, K = %d\n", data.Keys.N, data.Keys.K)
	for num, item := range data.Nums {
		fmt.Printf("Num %s: Base = %s, Value = %s\n", num, item.Base, item.Value)
	}
}
