package main

import (
	sum "dsa/twopointers/a-assignment"
	"fmt"
)

func main() {
	fmt.Println("inside main")

	// arr := []int{2, 7, 11, 15}
	// target := 9
	// result := FindSum(arr, target)
	// result := sum.TwoSum(arr, target)
	dupnum := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := sum.RemoveDuplicatesfromSortedArray(dupnum)
	for i := 0; i < k; i++ {
		fmt.Print(dupnum[i], " ")
	}
	// fmt.Println(result)
	fmt.Println(k)
}
