```go
func buildArray(nums []int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        val := nums[nums[i]] % n
        nums[i] = nums[i] + n*val
    }
    for i := 0; i < n; i++ {
        nums[i] = nums[i] / n
    }
    return nums
}
```