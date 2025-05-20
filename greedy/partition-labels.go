package main
// https://leetcode.com/problems/partition-labels/description/

func partitionLabels(s string) []int {
    n := len(s)
    pos := make([]int,26)
    for i := 0 ; i < 26 ; i++ {
        pos[i] = -1 
    }
    for i := 0 ; i < n ; i++ {
        pos[s[i]-'a'] = i 
    }
    ans := make([]int,0)
    start,end := 0,0
    for start < n {
        end = pos[s[start]-'a']
        count := 0
        for start < n && start <= end {
            count++
            end = max(end,pos[s[start]-'a'])
            start++
        }
        ans = append(ans,count)
    }
    return ans
}