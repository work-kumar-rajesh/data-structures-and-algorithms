package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func deepCopyList(head *Node) *Node {
	if head == nil {
		return nil
	}

	freq := make(map[*Node]*Node)

	var helper func(*Node) *Node
	helper = func(curr *Node) *Node {
		if curr == nil {
			return nil
		}

		if deepCopy, exists := freq[curr]; exists {
			return deepCopy
		}
		deepCopy := &Node{Val: curr.Val}
		freq[curr] = deepCopy
		deepCopy.Next = helper(curr.Next)
		deepCopy.Random = helper(curr.Random)

		return deepCopy
	}

	return helper(head)
}


func subsetSum(idx,currSum int, nums []int,dp [][]int) bool {
	if currSum == 0 { return true }
	if idx == -1 { return false }
	if dp[idx][currSum] != -1 && dp[idx][currSum] == 1 { return true }
	if dp[idx][currSum] != -1 && dp[idx][currSum] == 0 { return false }
	ok := false
	if nums[idx] <= currSum {
		ok =  subsetSum(idx-1,currSum-nums[idx],nums,dp) 
	}
	ok = ok ||  subsetSum(idx-1,currSum,nums,dp)
	if ok { 
		dp[idx][currSum] = 1 
	}else{
		dp[idx][currSum] = 0 
	}
	return ok
}


func isPossible(nums []int) bool {
	n := len(nums)
	sum := 0 
	for i := 0 ; i < n ; i++ { sum += nums[i] }
	if sum%2 != 0 { return false }
	sum = sum/2
	dp := make([][]int,n)
	for i := 0 ; i < n ; i++ {
		dp[i] = make([]int,sum+1)
		for j := 0 ; j <= sum ; j++ {
			dp[i][j] = -1
		}
	}
	return subsetSum(n-1,sum,nums,dp)
}