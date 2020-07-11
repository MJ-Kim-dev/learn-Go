package main

import "fmt"

func main() {
	a := 1
	b := a
	c := &a

	fmt.Printf("a=%v, b=%v, &b=%v, c=%v, *c=%v\n", a, b, &b, c, *c)

	a = 2
	fmt.Printf("a=%v, b=%v, &b=%v, c=%v, *c=%v\n", a, b, &b, c, *c)

	*c = 2222
	fmt.Printf("a=%v, b=%v, &b=%v, c=%v, *c=%v\n", a, b, &b, c, *c)
}
