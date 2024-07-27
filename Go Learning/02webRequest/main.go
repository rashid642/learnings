package main

import (
	"fmt"
	"net/http"
	Greeting "webRequest/greeting"
)

const url = "https://rashidaziz.vercel.app/"

func main() {
	Greeting.Greet("Rashid")
	fmt.Println("Hello")
	response, err := http.Get(url)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type : %T", response)
	fmt.Println(response)

	defer response.Body.Close()
}