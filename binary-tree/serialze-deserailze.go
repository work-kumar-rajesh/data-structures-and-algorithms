package main

import (
	"strconv"
	"strings"
)

// **Serialization: Convert Tree to String**
func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var res []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, "*") // Use "*" to represent null
			continue
		}

		res = append(res, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}

	return strings.Join(res, ",")
}

// **Deserialization: Convert String back to Tree**
func deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	root := &TreeNode{Val: atoi(values[0])}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if values[i] != "*" {
			node.Left = &TreeNode{Val: atoi(values[i])}
			queue = append(queue, node.Left)
		}
		i++

		if values[i] != "*" {
			node.Right = &TreeNode{Val: atoi(values[i])}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// **Helper function to convert string to int**
func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}


