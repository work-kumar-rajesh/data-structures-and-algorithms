func longestSubarray(arr []int, k int) int {
	mp := make(map[int]int)
	res := 0
	prefSum := 0

	for i := 0; i < len(arr); i++ {
		prefSum += arr[i]

		if prefSum == k {
			res = i + 1
		} else if index, exists := mp[prefSum-k]; exists {
			if i - index > res {
				res = i - index
			}
		}

		// Store only first occurrence of prefSum
		if _, exists := mp[prefSum]; !exists {
			mp[prefSum] = i
		}
	}

	return res
}