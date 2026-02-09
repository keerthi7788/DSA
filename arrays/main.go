package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", Sum(nums))
	n := 4
	fmt.Println("factorial", Factorial(n))
	largest, _ := FindLargestNumber(nums)
	fmt.Println("largestnum:", largest)
}
