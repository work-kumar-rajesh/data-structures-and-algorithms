```go
func numEquivDominoPairs(dominoes [][]int) int {
    freq :=  make(map[[2]int]int)
    n := len(dominoes)
    count := 0 
    for i := 0 ; i < n ; i++ {
        val,found := freq[[2]int{dominoes[i][0],dominoes[i][1]}]
        if found { 
            count = count + val 
        }
        val,found = freq[[2]int{dominoes[i][1],dominoes[i][0]}]
        if found && dominoes[i][0] != dominoes[i][1] { 
            count = count + val 
        }
        freq[[2]int{dominoes[i][0],dominoes[i][1]}]++
    }
    return count
}

``