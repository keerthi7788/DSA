package hasgMap

/*
 350. Intersection of Two Arrays II

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.

Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
func intersect(nums1 []int, nums2 []int) []int {
	freq := make(map[int]int)

	for _, num := range nums1 {
		freq[num]++
	}
	res := []int{}
	for _, v := range nums2 {
		if freq[v] > 0 {
			res = append(res, v)
			freq[v]--
		}
	}
	return res
}

/*
logic:
store the frequency of each element in nums1 in a map. Then,
 iterate through nums2 and check if the current element exists in the map with a frequency greater than 0.
  If it does, add it to the result array and decrement the frequency in the map.
  Finally, return the result array containing the intersection of the two arrays.
*/
