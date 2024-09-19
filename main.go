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
	params.Add("", s.Query)
	return s.BaseURL, nil
	
}


func main (){


	var username string
	fmt.Println("What is your Github Username?: ")
	fmt.Scanf("%s", &username)
	

	
	search := &Searx{
		BaseURL: fmt.Sprintf("https://api.github.com/users/%s/events", username),
	}
	
	resp, err := http.Get(search.BaseURL)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()


	

}

	