package pointers

/* 
26. Remove Duplicates from Sorted Array
Easy
Topics
premium lock icon
Companies
Hint
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Consider the number of unique elements in nums to be k​​​​​​​​​​​​​​. After removing duplicates, return the number of unique elements k.

The first k elements of nums should contain the unique numbers in sorted order. The remaining elements beyond index k - 1 can be ignored.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

 

Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
 

Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.



brute force approcah:
Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
we can crete a new array
1-new add it in new arrya
1 is repeated skip
2 is new add it to new array
[1,2]


//2-pointer approach
/*
1. Initialize two pointers: i and j. Set i to 1 and j to 1.
[0,0,1,1,1,2,2,3,3,4]
left=1 -index
right=1
 we have to move  j always if the j value is same as prevous [j-1]number we need to shift j now need t do any thing , if the j and pevis value is not same we have to  put that j value in left place whre it is pontng the aorted non duplicate array
here:
- l=0 r=0 keep l same incerement j
-l=0 r=1 new wlement 
now l=1 r=1 inceremt r
l=1 r=1 icement r
l=1 r=2 new  replce this to to l place 
-l=2 r=2 incremet r
l=2 r=3 new element shift r valu to l place
l=3 r=3 incremet r
l=3 r=4 new elemt
l=4  array ends

 */


func RemoveDuplicatesfromSortedArray(nums []int) int {
	l:=1;
	
	for  r:=1; r<len(nums); r++{
		if nums[r] != nums[r-1]{
			nums[l]=nums[r]
			l ++

		}
		
	}
	return l
}