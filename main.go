package main

import (
	"fmt"
	"os"

	"HashtableRepo/comm"
	"HashtableRepo/hashtable"
)

func main() {
	fmt.Printf("Hashtable repository.\n")

	//testHashtable()

	testComm()
}

func testComm() {

	args := os.Args[1:]
	if len(args) > 0 {
		comm.Create(args[0])
	} else {
		comm.Create("client")
	}
}

func testHashtable() {

	var p hashtable.Pair

	hashtable.Init(5)
	hashtable.Put(p.New(1, "one"))
	hashtable.Print()

	if hashtable.ContainsKey(1) {
		fmt.Printf("Is contains key of 1\n")
	} else {
		fmt.Printf("Is NOT contains key of 1\n")
	}

	if hashtable.ContainsKey(2) {
		fmt.Printf("Is CONTAINS key of 2\n")
	} else {
		fmt.Printf("Is not contains key of 2\n")
	}

	if hashtable.ContainsValue("one") {
		fmt.Printf("Is contains value of one\n")
	} else {
		fmt.Printf("Is NOT contains value of one\n")
	}

	if hashtable.ContainsValue("two") {
		fmt.Printf("Is CONTAINS value of two\n")
	} else {
		fmt.Printf("Is not contains value of two\n")
	}

	fmt.Printf("Size of hashtable: %d\n", hashtable.Size())

	if hashtable.IsEmpty() {
		fmt.Println("Is EMPTY hashtable")
	} else {
		fmt.Println("Is not empty hashtable")
	}

	hashtable.Remove(1)
	hashtable.Print()
}
