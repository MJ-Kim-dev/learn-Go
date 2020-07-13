package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go Count10("mj", c)
	go Count10("sj", c)

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func Count10(name string, c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "Count..", i)
		time.Sleep(time.Second)
	}
	c <- true
}
