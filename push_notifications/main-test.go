package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type MeowFacts struct {
	Data []string
}

type AnimeQuote struct {
	Anime     string
	Character string
	Quote     string
}

func main() {
	// url := "https://meowfacts.herokuapp.com?count=5"
	url := "https://animechan.xyz/api/random"

	// Create an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var meow MeowFacts
	var quotes AnimeQuote
	// err = json.Unmarshal(body, &meow)
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quotes.Anime)
	fmt.Println(quotes.Character)
	fmt.Println(quotes.Quote)

	// for i := 0; i < len(meow.Data); i++ {
	// 	fmt.Println(meow.Data[i])
	// }
}
