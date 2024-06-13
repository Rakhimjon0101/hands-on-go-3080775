// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "as"

func main() {
	// create a local variable "val" and assign it an int value
	var a int = 12

	// print the value of the local variable "val"
	fmt.Println(a)

	// print the value of the package-level variable "val"
	printGloabalVar()

	// update the package-level variable "val" and print it again
	updateGloabalVar()

	// print the pointer address of the local variable "val"
	fmt.Println(&a)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&a) = 14
	fmt.Println(a)
}

func printGloabalVar() {
	fmt.Println(val)
}

func updateGloabalVar() {
	val = "updated"
}
