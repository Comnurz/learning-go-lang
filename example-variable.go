package main

import "fmt"

func main() {
	// String declaration
	var first = "initial"
	fmt.Println(first)

	// String declaration with type
	var second string = "initial"
	fmt.Println(second)

	// int declaration
	var intFirst = 1
	fmt.Println(intFirst)

	// int declaration with type
	var intSecond int = 2
	fmt.Println(intSecond)

	// multiple declaration
	var f, b int = 1, 4
	fmt.Println(f, b)

	// shorthand declaration
	strShort := "Hello" // var strShort string = "Hello"
	fmt.Println(strShort)

	// constant val
	const d = 1e11 / 4
	fmt.Println(d)
}
