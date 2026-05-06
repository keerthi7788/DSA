package main

/*
Frequency Count (Numbers)
Input
[1,2,2,3,3,3,4]
Output
{1:1, 2:2, 3:3, 4:1}
*/
// You can edit this code!
// Click here and start typing.

// func main() {
// 	arr := []int{1, 2, 2, 3, 3, 3, 4}
// 	value := FrequencyCount(arr)
// 	for k, v := range value {
// 		fmt.Printf("%d:%d ", k, v)
// 	}

// }
func FrequencyCount(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	return freq

}
