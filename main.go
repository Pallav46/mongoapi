package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pallav46/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")
	
	r := router.Router()

	fmt.Println("Server is getting started At :- http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))

}
