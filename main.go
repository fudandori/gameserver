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

func isValid(matrix [][]int, x int, y int) bool {
	height := len(matrix) - 1

	if x < 0 || x > height {
		return false
	}

	width := len(matrix[x]) - 1

	if y < 0 || y > width {
		return false
	}

	return true
}

func switchBoxes(x int, y int, matrix *[][]int) {
	boxes := [][]int{
		{x - 1, y},
		{x, y - 1},
		{x + 1, y},
		{x, y + 1},
		{x, y},
	}

	for _, v := range boxes {
		x := v[0]
		y := v[1]

		if !isValid(*matrix, x, y) {
			return
		}

		a := (*matrix)[x][y]

		if a == 0 {
			(*matrix)[x][y] = 1
			fmt.Printf("changed %d to 1\n", a)
		} else {
			(*matrix)[x][y] = 0
			fmt.Printf("changed %d to 0\n", a)
		}
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

	switchBoxes(data.x(), data.y(), &data.Matrix)

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	const n = 5

	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	var req Request
	req.Matrix = board

	for i := 0; i < 6; i++ {
		x := rand.Intn(n - 1)
		y := rand.Intn(n - 1)
		switchBoxes(x, y, &req.Matrix)
	}

	var str, err2 = json.Marshal(&req)
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
