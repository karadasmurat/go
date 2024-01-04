package networking

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func simpleClient_GET(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func NetworkingBasics() {
	fmt.Println("Networking Basics")
	fmt.Println("-----------------")

	// run the server
	go RunServer()

	// run client
	simpleClient_GET("http://localhost:8080")
	// simpleClient_GET("http://www.google.com/robots.txt")
}
