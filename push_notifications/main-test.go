package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

type Color struct {
	Status string `json:"status"`
	Base   struct {
		Hex struct {
			Value string
		} `json:"hex"`
		Rgb struct {
			Value string
		} `json:"rgb"`
		Hsl struct {
			Value string
		} `json:"hsl"`
	} `json:"base"`
	Error struct {
		Message string
	} `json:"error"`
}

func main() {
	// url := "https://meowfacts.herokuapp.com?count=5"
	colors := []string{"y", "blue", "green"}
	// Create an HTTP GET request
	for i := 0; i < len(colors); i++ {
		url := "https://color.serialif.com/" + colors[i]
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
		var color Color
		// err = json.Unmarshal(body, &meow)
		err = json.Unmarshal(body, &color)
		if err != nil {
			log.Fatal(err)
		}
		if color.Status == "error" {
			fmt.Println(color.Error.Message)
		} else {
			fmt.Println(color.Status)
			fmt.Println(color.Base.Hex.Value)
			fmt.Println(color.Base.Rgb.Value)
			fmt.Println(color.Base.Hsl.Value)
			fmt.Println()
			time.Sleep(3 * time.Second)
		}
	}

	// for i := 0; i < len(meow.Data); i++ {
	// 	fmt.Println(meow.Data[i])
	// }
}
