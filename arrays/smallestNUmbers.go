package main

import "fmt"

//	func main() {
//		arr := []int{1, 3, 5, 6, 9}
//		second := Largest(arr)
//		fmt.Println("second", second)
//	}
func Smallest(arr []int) int {
	small := arr[0]

	if len(arr) < 2 {
		fmt.Println("Array must have at least two elements")
		return -1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] < small {

			small = arr[i]

		}
	}
	return small

}
func SecondSmallest(arr []int) int {
	small := arr[0]
	second := arr[1]

	if len(arr) < 2 {
		fmt.Println("Array must have at least two elements")
		return -1
	}
	if second < small {
		second, small = small, second
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] < small {
			second = small

			small = arr[i]

		} else if arr[i] < second && arr[i] != small {
			second = arr[i]
		}
	}
	return second

}

