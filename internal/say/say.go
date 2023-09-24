package say

import "fmt"

const Finalpunct string = "."

// func Hello(name string, last string) {

func Hello(first, last string) {
	fmt.Println("Hello, " + first + " " + last + Finalpunct)
}

func Hi(first string) {
	fmt.Println("Hi, " + first + Finalpunct)
}
