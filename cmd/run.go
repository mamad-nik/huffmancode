package main

import (
	"fmt"

	huffman "github.com/mamad-nik/huffman"
	builder "github.com/mamad-nik/huffman/builder"
	code "github.com/mamad-nik/huffman/code"
	p "github.com/mamad-nik/huffman/preprocess"
)

func getinput() (input string) {
	fmt.Print("please enter the text you want code: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("error in reading input, please try again")
	}
	return input
}

func main() {
	var pq huffman.Pqueue

	input := getinput()
	nodes, characters := p.Process(input)

	fmt.Println("tree nodes: ")
	tree := builder.Builder(&pq, nodes...)

	code := code.Get(tree, input, characters)
	fmt.Println("code: ", code)

}
