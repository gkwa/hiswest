package test3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func RunTest3() {
	resp, err := http.Get("https://httpbin.org/headers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Create a new JSON decoder linked to stdout
	decoder := json.NewDecoder(resp.Body)

	// Create a map to store the decoded JSON
	var data map[string]interface{}

	// Decode the JSON into the map
	if err := decoder.Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "error decoding JSON: %v", err)
		return
	}

	// Create a new JSON encoder linked to stdout
	encoder := json.NewEncoder(os.Stdout)

	// Encode the decoded JSON
	if err := encoder.Encode(data); err != nil {
		fmt.Fprintf(os.Stderr, "error encoding JSON: %v", err)
		return
	}
}
