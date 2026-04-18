package pointers

/*
Problem Statement
Given an array of numbers sorted in ascending order and a target sum, find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair) such that they add up to the given target. If no such pair exists return [-1, -1].

Example 1:

Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6
Example 2:

Input: [2, 5, 9, 11], target=11
Output: [0, 2]
Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11
Constraints:

2 <= arr.length <= 104
-109 <= arr[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

Algorithm Walkthrough
Let's walk through the example with input [1, 2, 3, 4, 6] and target 6.

Initialize pointers: left = 0, right = 4
First iteration:
currentSum = 1 + 6 = 7 (greater than target)
Decrement right to 3
Second iteration:
currentSum = 1 + 4 = 5 (less than target)
Increment left to 1
Third iteration:
currentSum = 2 + 4 = 6 (equals target)
Return indices [1, 3]
Here is the visual representation of this algorithm for Example-1:

left[1, 2, 3, 4, 6] right
*/
func FindSum(arr []int, target int) []int {

	// arr:= []int[12,10,4,20,15]
	left, right := 0, len(arr)-1
	for left < right {

		currentsum := arr[left] + arr[right]

		if currentsum == target {

			return []int{left, right}
		}
		if currentsum < target {
			left++
		} else {
			right--
		}

	}
	return []int{-1, -1}

}

/*
finding sun of two consecutive numbers equals two target the retuen 1st index
input: [1, 2, 3, 4, 5], target=5
output: 1
Explanation: The numbers at index 1 and 2 add up to 5: 2+3=5
*/
func SumOfTwoConsecutiveNumbers(arr []int, target int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+arr[i+1] == target {
			return i
		}
	}
	return -1
}
