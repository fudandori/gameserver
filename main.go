package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	//value := r.URL.Query()["number"]

	//fmt.Fprintf(w, value[0])
	fmt.Println("Endpoint Hit: homePage")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	fmt.Println("Hello World Called")
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
