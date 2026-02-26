package main

func CountEvenNumbers(arr []int) int {
	count := 0
	for _, num := range arr {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

//sum of two consecutive numbers equals two target the retuen 1st index

func SumofTwoConsecutiveNumbers(arr []int, target int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+arr[i+1] == target {
			return i
		}
	}
	return -1
}