package main

import "fmt"

func funcA(a int) int {
	fmt.Println("func A")
	a++
	return a
}

func funcB(a int) int {
	fmt.Println("func B")
	a++
	return a
}

func main() {
	fmt.Println(funcA(funcB(2)))
}
