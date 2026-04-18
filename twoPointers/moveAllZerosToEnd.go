package pointers

/* 2. Move Zeroes
Example 1
Input:  [0,1,0,3,12]
Output: [1,3,12,0,0] */

// func main() {
// 	arr := []int{0, 1, 0, 3, 12}
// 	fmt.Println(MoveZerosAtend(arr))

// }
func MoveZerosAtend(arr []int) []int {
	l := 0
	for r := 0; r < len(arr); r++ {
		if arr[r] != 0 {
			arr[r], arr[l] = arr[l], arr[r]
			l++
		}

	}
	return arr

}

/* 3. Remove Element
Example 1
Input:  nums = [3,2,2,3], val = 3
Output: k = 2
nums = [2,2,_,_] */

// func main() {
// 	arr := []int{3, 2, 2, 3}
// 	val := 3
// 	fmt.Println(RemoveElement(arr, val))

// }
func RemoveElement(arr []int, val int) int {
	l := 0
	for r := 0; r < len(arr); r++ {
		if arr[r] != val {
			arr[l] = arr[r]
			l++
		}

	}
	return l

}

/* 3. Remove Element
Example 1
Input:  nums = [3,2,2,3,4,6], val = 3
Output: [2,2,4,6]
nums = [2,2,_,_] */

// func main() {
// 	arr := []int{3, 2, 2, 3}
// 	val := 3
// 	fmt.Println(RemoveElement(arr, val))

// }
// func main() {
// 	arr := []int{3, 2, 2, 3, 4, 6}
// 	val := 3
// 	k := RemoveElement(arr, val)
// 	//fmt.Println("k", k)
// 	fmt.Println(arr[:k])

// }
func RemoveElementReturnRemaingArry(arr []int, val int) int {
	l := 0
	for r := 0; r < len(arr); r++ {
		if arr[r] != val {
			arr[l] = arr[r]
			l++
		}

	}
	return l

}
