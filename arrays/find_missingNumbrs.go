func main() {
	num := []int{3, 0, 1}
	fmt.Println(FindMissingNumber(num))

}

//Question:
//Array contains numbers 0..n with one missing.
//Example nums = [3,0,1]
//Output = 2

func FindMissingNumber(num []int) int {
	n := len(num)
	xor := n

	for i := 0; i < n; i++ {
		xor ^= i ^ num[i]
	}
	return xor

}

// You can edit this code!
// Click here and start typing

func main() {
	num := []int{3, 0, 1, 4, 6, 2}
	fmt.Println(FindMissingNumber(num))

}

//Question:
//Array contains numbers 0..n with one missing.
//Example nums = [3,0,1,2,5]
//Output = 4

func FindMissingNumber(num []int) int {
	n := len(num)
	expected := n * (n + 1) / 2
	sum := 0
	for _, v := range num {
		sum = sum + v
	}
	return expected - sum

}
