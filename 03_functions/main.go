package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func add(a int, b int) int {
	return a + b
}

func repeat(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(keyword string) (length int, upper string) {
	defer fmt.Println("lenAndUpper's defer")
	fmt.Println("I'm in lenAndUpper func")

	length = len(keyword)
	upper = strings.ToUpper(keyword)
	return
}

func main() {
	fmt.Println(multiply(1, 2))
	fmt.Println(add(2, 2))

	repeat("Hi", "How", "are", "you", "!")

	retLen, retKw := lenAndUpper("learngo")
	fmt.Println(retLen, retKw)

}
