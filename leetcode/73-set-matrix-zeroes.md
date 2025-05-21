```go
func setZeroes(matrix [][]int)  {
    m,n := len(matrix),len(matrix[0])
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            if matrix[i][j] == 0 {
                for k := 0 ; k < n ; k++ {
                    if matrix[i][k] == 0 { continue }
                    matrix[i][k] = 1e9 + 1
                }
                for k := 0 ; k < m ; k++ {
                    if matrix[k][j] == 0 { continue }
                    matrix[k][j] = 1e9 + 1
                }
            }
        }
    }
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            if matrix[i][j] == 1e9 + 1 {
                matrix[i][j] = 0 
            }
        }
    }
}
```
