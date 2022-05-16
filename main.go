package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Matrix [][]byte
	Box    [2]byte
}

func (re Request) x() byte {
	return re.Box[0]
}

func (re Request) y() byte {
	return re.Box[1]
}

func mod(re Request, box [2]byte) {
	x := box[0]
	y := box[1]
	height := byte(len(re.Matrix) - 1)

	if x < 0 || x > height {
		return
	}

	width := byte(len(re.Matrix[x]) - 1)

	if y < 0 || y > width {
		return
	}

	a := re.Matrix[x][y]

	if a == 0 {
		re.Matrix[x][y] = 1
		fmt.Printf("changed %d to 1\n", a)
	} else {
		re.Matrix[x][y] = 0
		fmt.Printf("changed %d to 0\n", a)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	mJson := string(b[:])
	data := Request{}
	json.Unmarshal([]byte(mJson), &data)

	fmt.Println(mJson)
	fmt.Println(data)
	box1 := [2]byte{data.x() - 1, data.y()}
	box2 := [2]byte{data.x(), data.y() - 1}
	box3 := [2]byte{data.x() + 1, data.y()}
	box4 := [2]byte{data.x(), data.y() + 1}

	mod(data, box1)
	mod(data, box2)
	mod(data, box3)
	mod(data, box4)
	mod(data, data.Box)

	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprint(w, data.Matrix)
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
