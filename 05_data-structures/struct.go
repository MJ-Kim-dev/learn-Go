package main

import "fmt"

type weather struct {
	humidity int
	status   string
	degree   int
}

func main() {

	today := weather{humidity: 60, status: "cloudy", degree: 26}
	fmt.Println(today)

}
