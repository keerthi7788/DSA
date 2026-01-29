package main

import (
	"fmt"
)

func main() {
	fmt.Println("inside main")

	arr := []int{1, 2, 3, 4, 6}
	target := 8
	result := FindSum(arr, target)
	fmt.Println(result)
}
