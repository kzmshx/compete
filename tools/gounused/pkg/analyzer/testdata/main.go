package main

import "fmt"

func main() {
	fmt.Println(usedPrivateConst)
	fmt.Println(UsedPublicConst)

	fmt.Println(usedPrivateVar)
	fmt.Println(UsedPublicVar)

	usedPrivateFunction()
	UsedPublicFunction()

	usedPrivateVal := usedPrivate{}
	usedPrivateVal.UsedPublicMethod()

	usedPublicVal := UsedPublic{}
	usedPublicVal.UsedPublicMethod()
}

const usedPrivateConst = "Used private const."

const unusedPrivateConst = "Unused private const."

const UsedPublicConst = "Used public const."

const UnusedPublicConst = "Unused public const."

var usedPrivateVar = "Used private var."

var unusedPrivateVar = "Unused private var."

var UsedPublicVar = "Used public var."

var UnusedPublicVar = "Unused public var."

func UsedPublicFunction() {
	fmt.Println("Used public function.")
}

func UnusedPublicFunction() {
	fmt.Println("Unused public function.")
}

func usedPrivateFunction() {
	fmt.Println("Used private function.")
}

func unusedPrivateFunction() {
	fmt.Println("Unused private function.")
}

type usedPrivate struct{}

func (s usedPrivate) UsedPublicMethod() {
	fmt.Println("Used public method of used private struct.")
	s.privateMethod()
}

func (s usedPrivate) UnusedPublicMethod() {
	fmt.Println("Unused public method of used private struct.")
	s.privateMethod()
}

func (s usedPrivate) privateMethod() {
	fmt.Println("Private method of used private struct.")
}

type unusedPrivate struct{}

func (s unusedPrivate) PublicMethod() {
	fmt.Println("Public method of unused private struct.")
	s.privateMethod()
}

func (s unusedPrivate) privateMethod() {
	fmt.Println("Private method of unused private struct.")
}

type UsedPublic struct{}

func (s UsedPublic) UsedPublicMethod() {
	fmt.Println("Used public method of used public struct.")
	s.privateMethod()
}

func (s UsedPublic) UnusedPublicMethod() {
	fmt.Println("Unused public method of used public struct.")
	s.privateMethod()
}

func (s UsedPublic) privateMethod() {
	fmt.Println("Private method of used public struct.")
}

type UnusedPublic struct{}

func (s UnusedPublic) PublicMethod() {
	fmt.Println("Public method of unused public struct.")
	s.privateMethod()
}

func (s UnusedPublic) privateMethod() {
	fmt.Println("Private method of used public struct.")
}
