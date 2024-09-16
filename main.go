package main

import (
	"net/http"
	"net/url"
	"log"
	"fmt"

)
type Searx struct {
	BaseURL string
	Query string

}

// BuildURL builds the URL for the search query

func (s *Searx) GetEvents() (string, error) {
	params:= url.Values{}
	params.Add("query", s.Query)
	return s.BaseURL + "?" + params.Encode(), nil
	
}


func main (){
	var username string
	fmt.Println("What is your Github Username?: ")
	fmt.Scanf("%s", &username)
	
	
	search := &Searx{
		BaseURL: fmt.Sprintf("https://api.github.com/users/%s/events", username),
	}
	
	eventsURL, err := search.GetEvents()
	if err != nil {
		log.Fatal("Error building the URL: ", err)
	}

	fmt.Println("GitHub Events URL:", eventsURL)

	// Perform an HTTP GET request (example)
	resp, err := http.Get(eventsURL)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	defer resp.Body.Close()


}
