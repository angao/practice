package main

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movie = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey", "Ingrid"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}

	data, err := json.MarshalIndent(movie, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
