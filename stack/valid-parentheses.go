package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func isValid(s string) bool {
	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		sz := len(stack)
		if sz == 0 {
			return false
		}
		if s[i] == ')' && stack[sz-1] != '(' {
			return false
		}
		if s[i] == '}' && stack[sz-1] != '{' {
			return false
		}
		if s[i] == ']' && stack[sz-1] != '[' {
			return false
		}
		stack = stack[:sz-1]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
