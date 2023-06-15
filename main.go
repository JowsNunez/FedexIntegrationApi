package main

import (
	"fedex/handlers"
	"fmt"
	"log"
	"net/http"
)

func createServer(port string) {
	http.HandleFunc("/", handlers.MainHandler)

	log.Fatal(http.ListenAndServe(port, nil))

	fmt.Println("server on port: " + port)

}

func main() {
	fmt.Println("main")
	port := ":3000"
	createServer(port)

}
