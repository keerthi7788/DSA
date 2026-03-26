package strings

import "fmt"

//count the cosecutive characters in a string if the achat appera in between and repared the smae again shoud be new.

//s:="AAABBCCAABB"
//output: A3B2C2A2B2   --make sure output formate string.

func main() {
	s := "AAABBCCAABB"
	value := CountCharInString(s)
	fmt.Println("Hello", value)
}
func CountCharInString(s string) string {
	if len(s) == 0 {
		return ""
	}
	count := 1
	result := ""
	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {
			count++
		} else {

			result = result + fmt.Sprintf("%c%d", s[i-1], count)

			count = 1

		}
	}

	result = result + fmt.Sprintf("%c%d", s[len(s)-1], count)

	return result

}

/*Problem

Input: "AABAABBBAA"

Goal: Find the character with the maximum consecutive repeats.

Output: Character 'B' with count 3 (since 'BBB' is the longest run). */

// func main() {
// 	s := "A"
// 	value := CountCharInString(s)
// 	fmt.Println("Hello", value)
// }

func CountLongestConsecutive(s string) string {
	if len(s) == 0 {
		return ""
	}

	count := 1
	maxcount := 1

	char := s[0]

	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {
			count++
		} else {
			if count > maxcount {
				maxcount = count
				char = s[i-1]
			}
			count = 1
		}
	}

	// check the last sequence
	if count > maxcount {
		maxcount = count
		char = s[len(s)-1]
	}

	return fmt.Sprintf("%c%d", char, maxcount)
}

//	func main() {
//		s := "ABCACC"
//		value := CountCharInString(s)
//		fmt.Println("Hello", value)
//	}
//
// Counte max repated char in a sting
func CountMaxRepeatedCharInString(s string) string {
	if len(s) == 0 {
		return ""
	}
	maxcount := 0
	var ch rune
	freq := make(map[rune]int)
	for _, value := range s {
		freq[value]++
	}
	for k, count := range freq {
		if count > maxcount {

			maxcount = count
			ch = k
		}
	}
	return fmt.Sprintf("%c%d", ch, maxcount)

}

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

// You can edit this code!
// Click here and start typing.

// func main() {
// 	s := "0110111"
// 	result := CountMaxConsicutive(s)
// 	for k, v := range result {
// 		fmt.Printf("%c:%d\n", k, v)
// 	}
// }

func CountTheRepeatedcountOfEachChar(s string) map[rune]int {
	count := make(map[rune]int)

	for _, value := range s {

		count[value]++

	}
	return count

}

//output:
//  1:5
// 0:2

// You can edit this code!
// Click here and start typing.

// func main() {
// 	s := []interface{}{0, 1, 1, 0, 1, 1, 1, "a", "a"}
// 	result := CountMaxConsicutive(s)
// 	fmt.Println(result)
// }

func CountEachOccurenceofvalue(s []interface{}) map[interface{}]int {
	count := make(map[interface{}]int)

	for _, value := range s {

		count[value]++

	}
	return count

}

//output :map[a:2 0:2 1:5]
// You can edit this code!
// Click here and start typing.

//Most frequent element [1,1,2,3,3,3] Output 3Most frequent element [1,1,2,3,3,3] Output 3 lets implet mentthis.

// highst repeat value
// output: 1
// func main() {
// 	arr := []int{0, 1, 1, 0, 2, 2, 2}
// 	result := CountMaxConsicutive(arr)
// 	fmt.Println(result)
// }
