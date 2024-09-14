package main

import "fmt"



func main() {
	var name string
	
	fmt.Printf("What is your name: ")
	fmt.Scanf("%v\n", &name)

	ptr := &name
	*ptr = "Louda"


	fmt.Printf("Name = %T", ptr)
}