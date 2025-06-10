func solve(nums[]int , k int )  (int) {
    n := len(nums)
    freq := make(map[int]int) 
    freq[0]++
    xor,ans := 0,0
    for i := 0 ; i < n ; i++ {
        xor = xor^nums[i]
        temp := k^xor
        count,ok := freq[temp]
        if ok {
            ans += count
        }
        freq[xor]++
    }
    return ans
}