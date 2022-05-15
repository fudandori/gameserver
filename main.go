package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Matrix [][]int
	Box [2]int
}

func firstRow(data *Request) {

}

func homePage(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	mJson := string(b[:])
	data := Request{}
	json.Unmarshal([]byte(mJson), &data)

	if(data.Box[0] == 0)

	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprint(w, data)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	fmt.Println("Hello World Called")
}

func handleRequests() {
	http.HandleFunc("/calculate", homePage)
	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
