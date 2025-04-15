package main

import "container/list"

// Populating Next Right Pointers in Each Node
// You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
// The binary tree has the following definition:

// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.

// Before Connecting Next Pointers:

//         1 -> nil
//       /   \
//      2     3 -> nil
//     / \   / \
//    4   5 6   7 -> nil
// After Connecting Next Pointers:

//         1 -> nil
//       /   \
//      2  -> 3 -> nil
//     / \   / \
//    4-> 5->6->7 -> nil

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// time O(N)
// space O(N)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		sz := queue.Len()
		temp := make([]*Node, 0)
		for i := 0; i < sz; i++ {
			curr := queue.Remove(queue.Front()).(*Node)
			temp = append(temp, curr)
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		for j := 0; j < len(temp)-1; j++ {
			temp[j].Next = temp[j+1]
		}
		temp[len(temp)-1].Next = nil
	}
	return root
}

// time O(N)
// space O(1)

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost := root           // Start from the leftmost node of each level
	for leftmost.Left != nil { // If there's no left child, we're done
		curr := leftmost
		for curr != nil {
			curr.Left.Next = curr.Right
			if curr.Next != nil {
				curr.Right.Next = curr.Next.Left
			}
			curr = curr.Next // Move to next node in the current level
		}
		leftmost = leftmost.Left // Move to the next level
	}

	return root
}

// recursive code to undestand  above o(1) space  approach ( this itself takes o(n) though)
// func connect(root *Node) *Node {
//     if root == nil {
//         return nil
//     }

//     if root.Left != nil {
//         root.Left.Next = root.Right
//         if root.Next != nil {
//             root.Right.Next = root.Next.Left
//         }
//         connect(root.Left)
//         connect(root.Right)
//     }

//     return root
// }
