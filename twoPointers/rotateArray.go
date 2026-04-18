package pointers

/*
	You are given:

An integer N (size of array)
An integer K (number of rotations)
An array of N integers

👉 Task:

Rotate the array to the right by K positions
Return/print the rotated array

👉 Note:

Rotation is circular
Output should be space-separated
*/
// func main() {
// 	arr := []int{10, 20, 30, 40}
// 	k := 2
// 	result := rotateRight(arr, k)
// 	fmt.Println("result:", result)

// }
func rotateRight(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	return append(arr[n-k:], arr[:n-k]...)

}

// left rotation
func rotateLeft(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	return append(arr[k:], arr[:k]...)

}

//correct approach is

func rotateRightInPlace(arr []int, k int) {
	n := len(arr)
	k = k % n
	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)

}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
