func lengthOfLongestSubstring(s string) int {
    freq := make(map[byte]int)
    n := len(s)
    ans, start := 0, 0

    for i := 0; i < n; i++ {
        freq[s[i]]++
        for freq[s[i]] > 1 {
            freq[s[start]]--
            start++
        }
        if i-start+1 > ans {
            ans = i - start + 1
        }
    }
    return ans
}
