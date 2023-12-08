package main

import "fmt"

func main() {
	UsedPublicFunction()
}

func UsedPublicFunction() {
	fmt.Println("This function is used.")
}

func UnusedPublicFunction() {
	fmt.Println("This function is not used.")
}
