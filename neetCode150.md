```go
func twoSum(nums []int, target int) []int {
    freq := make(map[int]int)
    for i,val := range nums {
        if pos,exists := freq[target-val] ; exists {
            return []int{i,pos}
        }
        freq[val] = i 
    }
    return []int{}
}
```

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans, curr *ListNode
	ans = &ListNode{Val: 0, Next: curr}
	curr = ans
	var carry, sum int = 0, 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		switch {
		case l1 == nil:
			sum = l2.Val + carry
		case l2 == nil:
			sum = l1.Val + carry
		default:
			sum = l1.Val + l2.Val + carry
		}
		carry = 0
		if sum > 9 {
			carry = 1
			sum = sum % 10
		}
		curr.Next = &ListNode{Val: sum}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: 1}
	}
	return ans.Next
}
```

```go
func lengthOfLongestSubstring(s string) int {
    freq , n := make(map[byte]int) , len(s)
    if n == 0 { return 0 }
    ans,curr,start := 1,0,0
    for i := 0 ; i < n ; i++ {
        curr++
        freq[s[i]]++
        if freq[s[i]] > 1 {
            for start < i {
                freq[s[start]]--
                curr--
                if s[start] == s[i] {
                    start++
                    break 
                }
                start++ 
            }
        }
        ans = max(ans,curr)
    }
    return ans
}
```

``` go
// 4 median of two sorted array didnt get it
```

```go
//best approach time o(n*N) space o(1)
func longestPalindrome(s string) string {
    n := len(s)
    if n == 0 {
        return ""
    }
    
    start, maxLen := 0, 0
    
    expand := func(left, right int) {
        for left >= 0 && right < n && s[left] == s[right] {
            left--
            right++
        }
        left++
        right--
        if right-left+1 > maxLen {
            start = left
            maxLen = right - left + 1
        }
    }
    
    for i := 0; i < n; i++ {
        expand(i, i)   // Odd length palindromes
        expand(i, i+1) // Even length palindromes
    }
    
    return s[start : start+maxLen]
}

//mine approach ime o(n*N) space o(n*n)

func longestPalindrome(s string) string {
    n := len(s)
    if n == 0 {
        return ""
    }

    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }

    maxStart, maxEnd := 0, 0

    for length := 2; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            if s[i] == s[j] {
                if length == 2 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            if dp[i][j] {
                maxStart = i
                maxEnd = j
            }
        }
    }

    return s[maxStart : maxEnd+1]  
}


```

```go
//need to understand again 
func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1
    
    for _, coin := range coins { 
        for j := coin; j <= amount; j++ { 
            dp[j] += dp[j - coin]
        }
    }
    return dp[amount]
}
```

```go
func reverse(x int) int {
    if x == 0 { return 0}
    var curr,multiplier int =   0,1e15
    for x/multiplier == 0{
        multiplier = multiplier/10
    }
    for x != 0 {
        curr = curr + (x%10)*multiplier 
        multiplier /= 10
        x /= 10
    }
    mina := (2<<30)*-1
    maxa := (2<<30)-1
    if curr < mina || curr > maxa { return 0 }
    return curr
}
```

```go
func helper(node *TreeNode,maxa int) int {
    if node == nil { return 0 }
    count := 0 
    if node.Val >= maxa { count++ }
    count += helper(node.Left,max(maxa,node.Val))
    count += helper(node.Right,max(maxa,node.Val))
    return count
}
func goodNodes(root *TreeNode) int {
    return helper(root,-1e9)
}
```

```go
func helper(i, j, n, m int, s, p string) bool {
    if j == m { 
        return i == n 
    }
    match := i < n && (s[i] == p[j] || p[j] == '.')
    if j+1 < m && p[j+1] == '*' { // If next char is '*'
        return helper(i, j+2, n, m, s, p) || // Skip "p[j]*"
               (match && helper(i+1, j, n, m, s, p)) // Use "p[j]*" at least once
    }
    if match {
        return helper(i+1, j+1, n, m, s, p)
    }
    return false
}
func isMatch(s string, p string) bool {
    return helper(0, 0, len(s), len(p), s, p)
}
```

```go
func maxArea(height []int) int {
    n := len(height)
    maxa,i,j := 0,0,n-1
    for i <= j {
        width := j-i
        length := min(height[i],height[j])
        maxa = max(maxa,length*width)
        if height[i]>height[j]{
            j--
        }else{
            i++
        }
    }
    return maxa
}
```

```go
type pos struct{
    x int 
    y int
}

var dx []int = []int{ 0 , 0 , 1, -1 }
var dy []int = []int{ 1 , -1 , 0, 0 }

func orangesRotting(grid [][]int) int {
    n,m := len(grid),len(grid[0])
    queue := list.New()
    totalOranges,rottenOranges,time := 0,0,-1
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < m ; j++ {
            if grid[i][j] == 1 || grid[i][j] == 2 { totalOranges++ }
            if grid[i][j] == 2 { 
                rottenOranges++
                queue.PushBack(pos{i,j}) 
            }
        }
    }
    if totalOranges == 0 { return 0 }
    for queue.Len()>0 {
        time++
        sz := queue.Len()
        for k := 0 ; k < sz ; k++ {
            curr := queue.Remove(queue.Front()).(pos)
            x,y := curr.x,curr.y
            for i := 0 ; i < 4 ; i++ {
                nx,ny := x+dx[i],y+dy[i]
                if nx < 0 || ny < 0 || nx == n || ny == m { continue }
                if grid[nx][ny] == 0 || grid[nx][ny] == 2 { continue }
                grid[nx][ny] = 2
                rottenOranges++
                queue.PushBack(pos{nx,ny})
            }
        }
    }
    if totalOranges != rottenOranges { return -1 }
    return time
}
```

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums) 
    n := len(nums)
    res := [][]int{}
    for i := 0; i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left, right := i+1, n-1 
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left-1] { left++ }
                for left < right && nums[right] == nums[right+1] { right-- }
            } else if sum < 0 {
                left++ 
            } else {
                right-- 
            }
        }
    }
    return res
}

```

```go
func helper(idx,n int,s,curr string,ans *[]string,options map[byte][]byte){
    if idx == n {
        if curr != "" { *ans = append(*ans,curr) }
        return 
    }
    for i := 0 ; i < len(options[s[idx]]) ; i++ {
        temp := curr + string(options[s[idx]][i])
        helper(idx+1,n,s,temp,ans,options)
        temp = temp[:len(temp)-1]
    }
}
func letterCombinations(digits string) []string {
    n := len(digits)
    options := make(map[byte][]byte)
    options['2'] = []byte{'a','b','c'}
    options['3'] = []byte{'d','e','f'}
    options['4'] = []byte{'g','h','i'}
    options['5'] = []byte{'j','k','l'}
    options['6'] = []byte{'m','n','o'}
    options['7'] = []byte{'p','q','r','s'}
    options['8'] = []byte{'t','u','v'}
    options['9'] = []byte{'w','x','y','z'}
    ans := make([]string,0)
    helper(0,n,digits,"",&ans,options)
    return ans
}
```

```go
func lengthOfLinkedList(head *ListNode) int {
    temp,count := head,0
    for temp != nil {
        count++
        temp = temp.Next
    }
    return count
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    k := lengthOfLinkedList(head) - n
    if k == 0 {
        toDelete := head
        head = head.Next
        toDelete.Next = nil 
        return head
    }
    temp := head
    for k > 1 {
        temp = temp.Next
        k--
    }
    toDelete := temp.Next
    temp.Next = temp.Next.Next
    toDelete.Next = nil 
    return head
}
```

```go
func isValid(s string) bool {
    n := len(s)
    stack := make([]byte,0)
    for i := 0 ; i < n ; i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '['{
            stack = append(stack,s[i])
            continue
        }
        sz := len(stack)
        if sz == 0 { return false }
        if s[i] == ')' && stack[sz-1] != '(' { return false }
        if s[i] == '}' && stack[sz-1] != '{' { return false }
        if s[i] == ']' && stack[sz-1] != '[' { return false }
        stack = stack[:sz-1]
    }
    if len(stack) != 0 { return false }
    return true
}
```

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    curr := &ListNode{Val: 0}  
    ans := curr  

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next  
    }
    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }
    return ans.Next  
}
```

```go
func helper(idx,openingCount int,curr string,ans *[]string){
    if idx == 0 {
        if openingCount == 0 { *ans = append(*ans,curr) } 
        return 
    }
    if openingCount > 0 {
        temp := curr + string(')')
        helper(idx-1,openingCount-1,temp,ans)
    }
    temp := curr + string('(')
    helper(idx-1,openingCount+1,temp,ans)
}

func generateParenthesis(n int) []string {
    ans := make([]string,0)
    n = 2*n 
    helper(n,0,"",&ans)
    return ans
}
```

```go
func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    prev, curr := (*ListNode)(nil), head
    for curr != nil {
        next := curr.Next  
        curr.Next = prev   
        prev = curr       
        curr = next
    }
    return prev  
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 || head == nil || head.Next == nil { return head }
    count := k - 1 
    start,curr := head,head
    for curr.Next != nil && count > 0  {
        curr = curr.Next
        count--
    }
    if count != 0 { return start }
    newNext := reverseKGroup(curr.Next,k)
    curr.Next = nil
    newHead := reverse(start)
    head.Next = newNext
    return newHead
}
```

```go
func helper(node *TreeNode,maxa *int) int {
    if node == nil { return 0 }
    left := helper(node.Left,maxa)
    right := helper(node.Right,maxa)
    *maxa = max(*maxa,left,right,1+left+right)
    return 1+max(left,right)
}
func diameterOfBinaryTree(root *TreeNode) int {
    maxa := 0 
    helper(root,&maxa)
    return maxa-1
}
```

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

```go
func helper(idx, target int, candidates, curr []int, ans *[][]int) {
    if target == 0 {
        *ans = append(*ans, append([]int(nil), curr...)) 
        return 
    }
    if idx == -1 { return }
    
    helper(idx-1, target, candidates, curr, ans)
    
    if target >= candidates[idx] {
        curr = append(curr, candidates[idx]) 
        helper(idx, target-candidates[idx], candidates, curr, ans) 
    }
}

func combinationSum(candidates []int, target int) [][]int {
    n := len(candidates)

    sort.Slice(candidates, func(i, j int) bool {
        return candidates[i] < candidates[j]
    })
    
    ans := make([][]int, 0)
    helper(n-1, target, candidates, []int{}, &ans)
    return ans
}
```

```go
func solve(ind, target int, arr []int, ans *[][]int, curr *[]int) {
    if target == 0 {
        *ans = append(*ans, append([]int(nil), *curr...))
        return
    }
    for i := ind; i < len(arr); i++ {
        if i > ind && arr[i] == arr[i-1] {
            continue
        }
        if arr[i] > target {
            break
        }
        *curr = append(*curr, arr[i])
        solve(i+1, target-arr[i], arr, ans, curr)
        *curr = (*curr)[:len(*curr)-1]
    }
}

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var ans [][]int
    var curr []int
    solve(0, target, candidates, &ans, &curr)
    return ans
}
```
```go
func trap(height []int) int {
    n := len(height)
    preMax := make([]int,n)
    postMax := make([]int,n)
    curr := 0 
    for i := 0 ; i < n ; i++ {
        preMax[i] = curr
        curr = max(curr,height[i])
    }
    curr = 0 
    for i := n-1 ; i >= 0  ; i-- {
        postMax[i] = curr
        curr = max(curr,height[i])
    }
    ans := 0 
    for i := 0 ; i < n ; i++ {
        ans = ans + max(0,min(preMax[i],postMax[i]) - height[i])
    }
    return ans
}
```

```go
func jump(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }

    jumps, farthest, end := 0, 0, 0
    for i := 0; i < n-1; i++ {
        farthest = max(farthest, i+nums[i])
        
        if i == end {
            jumps++
            end = farthest
            if end >= n-1 {
                break
            }
        }
    }
    return jumps
}
```

```go
func helper(idx,n int,nums,vis,curr []int,ans *[][]int){
    if idx == n {
        *ans = append(*ans, append([]int(nil), curr...))
		return
    }
    for i := 0 ;  i < n ; i++ {
        if vis[i] == 1 { continue }
        curr = append(curr,nums[i])
        vis[i] = 1 
        helper(idx+1,n,nums,vis,curr,ans)
        curr = curr[:len(curr)-1]
        vis[i]= 0
    }
}
func permute(nums []int) [][]int {
    n := len(nums)
    vis := make([]int,n)
    ans := make([][]int,0)
    helper(0,n,nums,vis,[]int{},&ans)
    return ans
}
```

```go
func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0 ; i < n ; i++ {
        for j := i+1 ;  j < n ; j++ {
            matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
        }
    }
    for i := 0 ; i < n ; i++ {
        start,end := 0,n-1
        for start < end {
            matrix[i][start],matrix[i][end] = matrix[i][end],matrix[i][start]
            start++
            end--
        }
    }
}
```

```go
func groupAnagrams(strs []string) [][]string {
    freq := make(map[[26]int][]string)
    
    for _, str := range strs {
        var count [26]int
        for _, ch := range str {
            count[ch-'a']++ 
        }
        freq[count] = append(freq[count], str)
    }

    ans := make([][]string, 0, len(freq))
    for _, val := range freq {
        ans = append(ans, val)
    }
    return ans
}
```

```go
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1 
    }
    if n < 0 {
        x = 1 / x 
        n = -n   
    }
    result := 1.0
    base := x
    for n > 0 {
        if n%2 == 1 { result *= base }
        base *= base 
        n /= 2        
    }
    return result
}
```

```go 
func maxSubArray(nums []int) int {
    n := len(nums)
    maxa,curr := nums[0],nums[0]
    for i := 1 ; i < n ; i++ {
        curr = max(nums[i],curr+nums[i])
        maxa = max(maxa,curr)
    }
    return maxa
}
```

```go
func merge(intervals [][]int) [][]int {
   n := len(intervals)
   sort.Slice(intervals,func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
   }) 
   ans := make([][]int,0)
   curr := []int{intervals[0][0],intervals[0][1]}
   for i := 1 ; i < n ; i++ {
        if intervals[i][0] <= curr[1] {
            curr[1] = max(curr[1],intervals[i][1]) 
        }else{
            ans = append(ans,append([]int(nil),curr...))
            curr = intervals[i]
        }
    }
    ans = append(ans,append([]int(nil),curr...))
    return ans
}
```

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    if isSameTree(root, subRoot) {
        return true
    } else {
        return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
    }
}

func isSameTree(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }

    if root1 == nil || root2 == nil {
        return false
    }

    if root1.Val == root2.Val {
        return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
    } else {
        return false
    }

}
```

```go
func searchMatrix(matrix [][]int, target int) bool {
    m,n := len(matrix),len(matrix[0])
    i,j := 0,n-1
    for i < m && j >= 0 {
        if matrix[i][j] == target { return true }
        if matrix[i][j] > target{
            j--
        }else{
            i++
        }
    }
    return false
}
```

```go
func isValid(base, curr map[byte]int) bool {
    for key, val := range base {
        if found, exists := curr[key]; !exists || found < val {
            return false
        }
    }
    return true
}

func minWindow(s string, t string) string {
    n := len(s)
    freq := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        freq[t[i]]++
    }

    curr := make(map[byte]int)
    start, end := -1, -1
    i, j := 0, 0

    for j < n {
        curr[s[j]]++ 
        for isValid(freq, curr) {
            if start == -1 || (end-start+1) > (j-i+1) {
                start = i
                end = j
            }
            curr[s[i]]-- 
            if curr[s[i]] == 0 {
                delete(curr, s[i])
            }
            i++
        }
        j++
    }
    if start == -1 {
        return ""
    }
    return s[start : end+1]
}
```

```go
func helper(idx,n int,nums,curr []int,ans *[][]int){
    if idx == n {
        *ans = append(*ans, append([]int(nil), curr...))
        return
    }
    helper(idx+1,n,nums,curr,ans)
    curr = append(curr,nums[idx])
    helper(idx+1,n,nums,curr,ans)
}
func subsets(nums []int) [][]int {
    n := len(nums)
    ans := make([][]int,0)
    helper(0,n,nums,[]int{},&ans)
    return ans
}
```

```go
var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func helper(i, j, m, n, idx int, board [][]byte, word string) bool {
    if idx == len(word) {
        return true
    }
    if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[idx] {
        return false
    }

    temp := board[i][j]
    board[i][j] = '#' // Mark as visited

    for k := 0; k < 4; k++ {
        nx, ny := i+dx[k], j+dy[k]
        if helper(nx, ny, m, n, idx+1, board, word) {
            return true
        }
    }

    board[i][j] = temp 
    return false
}

func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if helper(i, j, m, n, 0, board, word) {
                return true
            }
        }
    }
    return false
}
```



