package main
// func main() {
// 	arr := []int{1, 2, 2, 3, 4, 5, 4}
// 	fmt.Println(FindDuplicate(arr))

// }

/*
	many duplicate numbers

input: [1, 2, 2, 3, 4, 5, 4]
output: [2, 4]
*/
func FindDuplicateRepeatedNumbers(arr []int) []int {
	freq := make(map[int]int)
	duplicate := []int{}

	for _, v := range arr {
		freq[v]++
	}
	for k, value := range freq {
		if freq[value] != 1 {
			duplicate = append(duplicate, k)
		}
	}
	return duplicate

}

/* single duplicate number
inpiut: [1, 2, 2, 3, 4, 5, 4]
output: 2

*/

func FindDUplicateNUmber(arr []int) int {
	n := len(arr)
	exp := n * (n + 1) / 2
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return sum - exp

}

/*
find the most frequent duplicate number
input: [1, 2, 2, 3, 4, 5, 4]
output: 2
*/
func FindMostFrequentDuplicate(arr []int) (int, int) {
	maxcount := 0
	ele := 0
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++

		if freq[v] > maxcount {
			maxcount = freq[v]
			ele = v

		}
	}
	return ele, maxcount
}