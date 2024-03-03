package main

import (
	"biblioteca/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Run API")

	r := server.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))

}
