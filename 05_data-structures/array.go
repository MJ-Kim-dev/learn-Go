package main

import "fmt"

func main() {
	fruit := []string{"banana", "orange"}
	fruit = append(fruit, "strawberry")
	fmt.Println(fruit)
}
