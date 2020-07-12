package main

import (
	"fmt"

	"./members"
)

func main() {
	member := members.AddMember("mj", "000-123-1234", 100)
	fmt.Println(member)

	member.SetAddress("Canada")
	fmt.Println(member)

	fmt.Println(member.ExtMemberShip())
	fmt.Println(member.ExtMemberShip())
	fmt.Println(member.ExtMemberShip())
	fmt.Println(member.ExtMemberShip())

	fmt.Println(member)
}
