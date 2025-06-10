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