package main

import (
	"fmt"

	"./contacts"
)

func main() {
	mycontact := contacts.Dictionary{"mj": "123-456-789"}

	mobile, err := mycontact.Search("mj")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mobile)
	}

	mycontact.Add("sj", "456-789-123")
	fmt.Println(mycontact)

	mycontact.Delete("sj")
	fmt.Println(mycontact)

}
