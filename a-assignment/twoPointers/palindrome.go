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

// isAlphaNumeric checks if a given character is alphanumeric (a-z, A-Z, or 0-9).
func isAlphaNumeric(c byte) bool {
	// Check if the character is a lowercase letter (a-z).
	// Check if the character is an uppercase letter (A-Z).
	// Check if the character is a digit (0-9).
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


/* 2108. Find First Palindromic String in the Array
Easy
Topics
premium lock icon
Companies
Hint
Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.

 

Example 1:

Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
Example 2:

Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar".
Example 3:

Input: words = ["def","ghi"]
Output: ""
Explanation: There are no palindromic strings, so the empty string is returned.
 

Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists only of lowercase English letters. */
func FindFirstPalindromicString(words []string) string {
	for _, word := range words {
		if IsPalindrome(word) {
			return word
		}
	}
	return ""
}