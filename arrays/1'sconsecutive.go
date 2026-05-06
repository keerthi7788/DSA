package main

// Problem

// Given a binary arry, find the maximum number of consecutive 1s.

// Example:

// Input:  {0, 1, 1, 0, 1, 1, 1}
// Output: 3

// Because the longest continuous sequence of 1s is:

// 111

//	func main() {
//		arr := []int{0, 1, 1, 0, 1, 1, 1}
//		second := CountMaxConsicutive(arr)
//		fmt.Println("second", second)
//	}
func Find1sconsicutiveNumber(arr []int) int {
	count := 0
	maxcount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			count++
			if count > maxcount {
				maxcount = count
			}
		} else {
			count = 0
		}
	}
	return maxcount
}

func CountMaxConsicutive(arr []int) int {
	count := 0
	maxcount := 0

	for _, value := range arr {
		if value == 0 {
			count++
			if count > maxcount {
				maxcount = count
			}
		} else {
			count = 0
		}
	}
	return maxcount

}

// \Problem

// Given a binary string, find the maximum number of consecutive 1s.

// Example:

// Input:  "0110111"
// Output: 3

// Because the longest continuous sequence of 1s is:

// func main() {
// 	s := "0110111"
// 	result := CountMaxConsicutiveString(s)
// 	for k, v := range result {
// 		fmt.Printf("%c:%d\n", k, v)
// 	}
// }

func CountMaxConsicutiveString(s string) int {
	count := 0
	maxcount := 0

	for _, value := range s {
		if value == '1' {
			count++
			if count > maxcount {
				maxcount = count
			}
		} else {
			count = 0
		}
	}
	return maxcount

}


//output:
//  1:5
// 0:2

// func main() {
// 	s := []interface{}{0, 1, 1, 0, 1, 1, 1, "a", "a"}
// 	result := CountTheRepeatedcountOfEachChar(s)
// 	fmt.Println(result)
// }

func CountTheRepeatedcountOfEachChar(s string) map[rune]int {
	count := make(map[rune]int)

	for _, value := range s {

		count[value]++

	}
	return count

}

//output :map[a:2 0:2 1:5]
func CountEachOccurenceofvalue(s []interface{}) map[interface{}]int {
	count := make(map[interface{}]int)

	for _, value := range s {

		count[value]++

	}
	return count

}


//Most frequent element [1,1,2,3,3,3] Output 3Most frequent element [1,1,2,3,3,3] Output 3 lets implet mentthis.

// highst repeat value
// output: 1
// func main() {
// 	arr := []int{0, 1, 1, 0, 2, 2, 2}
// 	result := CountThehighestRepeatedValue(arr)
// 	fmt.Println(result)
// }

func CountThehighestRepeatedValue(arr []int) int {
	freq := make(map[int]int)
	maxcount := 0
	result := 0
	for _, value := range arr {

		freq[value]++

	}
	for i, v := range freq {
		if v > maxcount {
			maxcount = v
			result = i
		}
	}
	return result

}

//sum of two consecuive , return the 1st index
// func main() {
// 	arr := []int{0, 1, 2, 3, 4, 5}
// 	target := 5
// 	result := SumOfTwoConsecutive(arr, target)
// 	fmt.Println(result)
// }

func SumOfTwoConsecutive(arr []int, target int) int {

	for i := 0; i < len(arr)-1; i++ {

		if arr[i]+arr[i+1] == target {
			return i
		}

	}
	return -1
}
