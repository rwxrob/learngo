package main

import "github.com/rwxrob/learngo/internal/say"

func main() {

	/*
		var first string    // declaration
		var last string     // declaration
		first = "Rob"       // assignment
		last = `Muhlestein` // declaration + assigment
	*/

	/*
		var first, last string // declaration
		first = "Rob"          // assignment
		last = `Muhlestein`    // declaration + assigment
	*/

	// --------- BELOW ARE ONLY IN A FUNCTION BLOCK -----------

	/*
		first := "Rob"       // assignment
		last := `Muhlestein` // declaration + assigment
	*/

	//first, last := "Rob", "Muhlestein" // combined declaration + assignment

	//fmt.Println("Hello, " + first + " " + last + ".") // usage

	say.Hello("Rob", "Muhlestein")
	say.Hi("Rob")
	say.Hello("Doris", "Kapner")
	say.Hi("Doris")
}
