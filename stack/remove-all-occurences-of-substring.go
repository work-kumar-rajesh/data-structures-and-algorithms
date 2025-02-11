package main

// Given two strings s and part, perform the following operation on s until all occurrences of the substring part are removed:
// Find the leftmost occurrence of the substring part and remove it from s.
// Return s after removing all occurrences of part.
// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "daabcbaabcbc", part = "abc"
// Output: "dab"
// Explanation: The following operations are done:
// - s = "daabcbaabcbc", remove "abc" starting at index 2, so s = "dabaabcbc".
// - s = "dabaabcbc", remove "abc" starting at index 4, so s = "dababc".
// - s = "dababc", remove "abc" starting at index 3, so s = "dab".
// Now s has no occurrences of "abc".

func removeOccurrences(s string, part string) string {
	n := len(s)
	m := len(part)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		if len(stack) < m {
			continue
		}
		temp := stack[(len(stack) - m):]
		if string(temp) == part {
			stack = stack[:len(stack)-m]
		}
	}
	ans := ""
	for i := 0; i < len(stack); i++ {
		ans += string(stack[i])
	}
	return ans
}
