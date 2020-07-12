package main

import (
	"fmt"
	"time"
)

func main() {
	go Count10("mj")
	Count10("sj")
}

func Count10(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "Count..", i)
		time.Sleep(time.Second)
	}
}
