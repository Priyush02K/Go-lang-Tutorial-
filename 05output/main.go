package main

import (
	"fmt"
)



func main() {
	//print
	var i,j string = "Hello","World"

	fmt.Print(i)
	fmt.Print(j)

	//new line
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	//println
	fmt.Println(i,j)

	//Printf
	var l string = "Hello"
	var m int = 15
  
	fmt.Printf("i has value: %v and type: %T\n", l, m)
	fmt.Printf("j has value: %v and type: %T", m, m)
}
