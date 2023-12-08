package main

import "fmt"

func main() {
	usedPrivateFunction()
}

func usedPrivateFunction() {
	fmt.Println("This function is used.")
}

func unusedPrivateFunction() {
	fmt.Println("This function is not used.")
}
