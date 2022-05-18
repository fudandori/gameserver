package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Request struct {
	Matrix [][]int `json:"matrix"`
	Box    [2]int
}

func (re Request) x() int {
	return re.Box[0]
}

func (re Request) y() int {
	return re.Box[1]
}

func mod(re Request, box [2]int) {
	x := box[0]
	y := box[1]
	height := len(re.Matrix) - 1

	if x < 0 || x > height {
		return
	}

	width := len(re.Matrix[x]) - 1

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

func move(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: move")

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
	box1 := [2]int{data.x() - 1, data.y()}
	box2 := [2]int{data.x(), data.y() - 1}
	box3 := [2]int{data.x() + 1, data.y()}
	box4 := [2]int{data.x(), data.y() + 1}

	mod(data, box1)
	mod(data, box2)
	mod(data, box3)
	mod(data, box4)
	mod(data, data.Box)

	var str, err2 = json.Marshal(&data)
	fmt.Println(str)
	if err2 == nil {
		js := string(str)
		fmt.Fprint(w, js)
	} else {
		fmt.Println(err2)
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: generate")

	const n = 5

	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	var req Request
	req.Matrix = board
	for i := 0; i < 6; i++ {
		mod(req, [2]int{rand.Intn(n - 1), rand.Intn(n - 1)})
	}

	var str, err2 = json.Marshal(&req)
	fmt.Println(str)
	if err2 == nil {
		js := string(str)
		fmt.Fprint(w, js)
	} else {
		fmt.Println(err2)
	}

}

func handleRequests() {
	http.HandleFunc("/calculate", move)
	http.HandleFunc("/generate", generate)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
