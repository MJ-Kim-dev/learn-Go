package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return true
}

func whenCanIDrink3(age int) bool {
	switch age {
	case 17:
		return false
	case 18:
		return true
	}
	return false
}

func main() {

	ret := superAdd(1, 2, 3, 0, 5, 6)
	fmt.Println(ret)

	age := 19
	fmt.Println("Can I drink?", age, canIDrink(age))
	fmt.Println("Can I drink? (v2)", age, canIDrink2(age))
	fmt.Println("When can I drink? (v3)", age, whenCanIDrink3(age))
}
