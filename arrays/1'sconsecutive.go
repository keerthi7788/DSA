package main

func Find1sconsicutiveNUmber(arr []int) int {
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
