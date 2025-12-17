package sorting

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		for j := i - 1; j <= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
			arr[j] = key

		}
	}
}
