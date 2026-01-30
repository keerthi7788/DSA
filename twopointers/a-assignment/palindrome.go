package pointers

import "fmt"

/*
Code
Testcase
Testcase
Test Result
125. Valid Palindrome
Easy

125. A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

soln:
What is a Palindrome?

A palindrome is a word, number, or phrase that reads the same forward and backward.

Examples of plain words:
Word	Forward	Backward	Palindrome?
madam	madam	madam	  Yes
racecar	racecar	racecar	  Yes
hello	hello	olleh	  No

eg:
Input: s = "A man, a plan, a canal: Panama"
*/
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			fmt.Println("Comparing:", toLower(s[left]), toLower(s[right]))
			left++

		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// helper function
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
