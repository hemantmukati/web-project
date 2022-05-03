package main

import (
	"log"
	"net/http"
    "web-project/contoller"
)

func main() {
	log.Println("server started !!!")

	http.HandleFunc("/register", contoller.Register)
	http.HandleFunc("/get", contoller.GETHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

	defer log.Println("server stopped !!!")
}
