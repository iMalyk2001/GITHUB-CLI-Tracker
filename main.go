package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define the Yield struct
type Yield struct {
	Name  string `json:"username"`
	Event string `json:"type"`
	Date  string `json:"created_at"`
}

// Define the Searx struct
type Searx struct {
	BaseURL string
	Query   string
}

// GetEvents builds the URL for the search query (this function is not necessary for the current functionality)
func (s *Searx) GetEvents() (string, error) {
	// This function is not necessary for your current implementation
	return s.BaseURL, nil
}

func main() {
	var username string
	fmt.Println("What is your Github Username?: ")
	fmt.Scanf("%s", &username)

	// Create a Searx instance with the GitHub events API URL
	search := &Searx{
		BaseURL: fmt.Sprintf("https://api.github.com/users/%s/events", username),
	}

	// Make the HTTP GET request
	resp, err := http.Get(search.BaseURL)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// Define a slice of Yield to hold the unmarshaled data
	var response []Yield

	// Unmarshal the JSON data into the slice of Yield structs
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	// Print the unmarshaled data
	fmt.Println("Events:")
	for _, event := range response {
		fmt.Printf("Actor: %s, Type: %s, Date: %s\n", event.Name, event.Event, event.Date)
	}
}
