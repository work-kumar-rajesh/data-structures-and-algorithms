package main 

 type ListNode struct {
      Val int
      Next *ListNode
 }

func nextLargerNodes(head *ListNode) []int {
    if head == nil {
        return []int{}
    }
    ans := []int{}
    stack := [][2]int{} 
    curr := head
    index := 0

    for curr != nil {
        ans = append(ans, 0)
        for len(stack) > 0 && stack[len(stack)-1][0] < curr.Val {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[top[1]] = curr.Val
        }
        stack = append(stack, [2]int{curr.Val, index})
        index++
        curr = curr.Next
    }
    return ans
}
