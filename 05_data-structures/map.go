package main

import "fmt"

func main() {

	//map[key_type]value_type

	map1 := map[string]int{"mj": 28, "sj": 31}
	map2 := make(map[int]int)

	fmt.Println("1.", map1, map2)

	map1["name"] = 10

	fmt.Println("2.", map1, map2)

	fmt.Println("3.", map1["there is no this key"], map2[1234])

	retVal, retCode := map1["there is no this key"]
	fmt.Println("4.", retVal, retCode)

	if retVal, retCode := map1["mj"]; retCode {
		fmt.Println("5.", retVal)
	}

	delete(map1, "name")

	for key, value := range map1 {
		fmt.Println(key, value)
	}

}
