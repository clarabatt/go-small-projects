package ch1

import (
	"fmt"
)

func Map1() {
	var emptyMap = map[string]int{}
	emptyMap["id1"] = 1
	fmt.Println("ID1: ", emptyMap["id1"])
	fmt.Println("Len: ", len(emptyMap))
	emptyMap["id1"]++
	fmt.Println("ID1: ", emptyMap["id1"])

	value, exist := emptyMap["id2"]
	fmt.Println("Value: ", value)
	fmt.Println("Exist?: ", exist)

	value, exist = emptyMap["id1"]
	fmt.Println("Value: ", value)
	fmt.Println("Exist?: ", exist)
}
