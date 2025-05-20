```go
func isZeroArray(nums []int, queries [][]int) bool {
    n := len(nums)
    prefix := make([]int,n+1)
    for i := 0 ; i < len(queries) ; i++ {
        start,end := queries[i][0],queries[i][1]
        prefix[start] +=1 
        prefix[end+1] -=1
    }
    curr := 0 
    for i := 0 ; i < n ; i++ {
        curr = curr + prefix[i]
        switch{
            case nums[i] == 0 :
            continue
            case nums[i] > 0 && nums[i] <= curr :
            continue
            case nums[i] < 0 && nums[i]+curr >= 0 :
            continue
            default:
            return false
        }
    }
    return true
}
```
