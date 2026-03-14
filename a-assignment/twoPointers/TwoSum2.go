package pointers

import "fmt"

/*
167. Two Sum II - Input Array Is Sorted
Medium
Topics

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Constraints:

2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.

Brute force solution:
we can manually compare with each 2 pairs
like take 1st element and compare with all other elements

eg: Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
[2,7] =9
[2,11] = 13
[2,15] = 17
tc: O(n^2)

	for(i:=0;i<len(numbers);i++){
	  for(j:=i+1;j<len(numbers);j++){
	  if(numbers[i]+numbers[j]===target){
	  return [i,j]
	  }
	  }
	}

the outer loop runs tc n and inner loop runs tc 2 : n*n=n^2

--so to reduce the tc we will go with 2-pointers
-define left is the 1st element and right is the last element

-we will check the sum of the elements at these
- we will check 3 conditions
- take two elements and check whether sum is greater than target and less than target
- if less than target we know that we need to move the left pointer to the right
- if greater than target we know that we need to move the right pointer to the left
-if eqals it is the

eg: Input: numbers = [2,3,4], target = 6
Output: [1,3]
left=2
right:= 4
numbers[left]+numbers[right] =current
2+4=6
current> target
*/
func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		fmt.Println("left:", left, "right:", right, "current:", numbers[left]+numbers[right])
		currentsun := numbers[left] + numbers[right]
		fmt.Println("currentsun:", currentsun, "target", target)
		if currentsun == target {
			return []int{left+1, right+1}
		} else if currentsun < target {
			left++
		}else{
		right--
	}
 }
	return nil
}
