package main

import "fmt"

func FindLargestNumber(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	largest := nums[0]
	for _, num := range nums {
		if num > largest {
			largest = num
		}

	}
	return largest, true
}
func FindSecondLargestNumber(num []int) (int, bool) {
	if len(num) < 2 {
		return 0, false
	}
	largest := num[0]
	secondLargest := num[1]

	if secondLargest > largest {
		largest, secondLargest = secondLargest, largest
	}
	for i := 2; i < len(num); i++ {
		if num[i] > largest {
			secondLargest = largest
			largest = num[i]
		} else if num[i] > secondLargest {
			secondLargest = num[i]
		}
	}
	return secondLargest, true
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func PositiveNegative(n int) {
	if n > 0 {
		fmt.Println("positive num")
	} else {
		fmt.Println("neagtive number")
	}

}
