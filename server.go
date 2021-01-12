package main

import (
	"./controller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controller.AddPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}