func fourSum(nums []int, target int) [][]int {
    var res [][]int
    
    if len(nums) < 4 {return res}
    
    sort.Ints(nums)
    
    m := make(map[[4]int]struct{})
    
    for i := 0; i < len(nums) - 3; i++ {
        for j := i + 1; j < len(nums); j++ {
            left, right := j + 1, len(nums) - 1
            
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]

                if sum == target {
                    val := [4]int{nums[i], nums[j], nums[left], nums[right]}
                    
                    if _, ok := m[val]; !ok {
                        res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
                        m[val] = struct{}{}
                    }
                    
                    left++
                    right--
                    
                    for left < right && nums[left] == nums[left-1] {
                        left++
                    }

                    for left < right && nums[right] == nums[right + 1] {
                        right--
                    }
                } else if sum > target {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    
    return res
}