```go
func longestPalindrome(words []string) int {
    ans := 0
    n := len(words)
    freq := make(map[[2]byte]int)
    for i := 0 ; i < n ; i++ {
        char1 := byte(words[i][0])
        char2 := byte(words[i][1])
        curr := [2]byte{char1,char2}
        opposite := [2]byte{char2,char1}
        if freq[opposite] > 0 {
            ans += 4
            freq[opposite] -= 1
        }else{
            freq[curr] += 1
        }
    }
    for key,value := range freq {
        if key[0] != key[1] { continue }
        if value > 0 {
            ans += 2
            break
        }
    }
    return ans
}
```
