package main

//partition dp type

// Given an array arr[] which represents dimensions of sequence of matrices where the ith matrix has the dimensions
//  (arr[i-1] x arr[i]) for i>=1., find the most efficient way to multiply these matrices together. The efficient way is
//  the one that involves the least number of multiplications.

// Examples:

// Input: arr[] = [2, 1, 3, 4]
// Output: 20
// Explanation: There are 3 matrices of dimensions 2 × 1, 1 × 3, and 3 × 4, Let the input 3 matrices be M1, M2, and M3. There are two ways to multiply: ((M1 x M2) x M3) and (M1 x (M2 x M3)), note that the result of (M1 x M2) is a 2 x 3 matrix and result of (M2 x M3) is a 1 x 4 matrix.
// ((M1 x M2) x M3)  requires (2 x 1 x 3)  + (0) +  (2 x 3 x 4) = 30
// (M1 x (M2 x M3))  requires (0)  + (1 x 3 x 4) +  (2 x 1 x 4) = 20.
// The minimum of these two is 20.

const INF = int(^uint(0) >> 1) // Maximum integer value

// ---------------- Recursive Method ----------------

func matrixMultiplicationRecursive(arr []int, N int) int {
	return f(arr, 1, N-1)
}

func f(arr []int, i, j int) int {
	if i == j {
		return 0
	}
	mini := INF
	for k := i; k <= j-1; k++ {
		cost := f(arr, i, k) + f(arr, k+1, j) + arr[i-1]*arr[k]*arr[j]
		if cost < mini {
			mini = cost
		}
	}
	return mini
}

// ---------------- Iterative Method ----------------

// matrixMultiplicationIterative computes the minimum multiplication cost using bottom-up DP.
func matrixMultiplicationIterative(arr []int, N int) int {
	// dp[i][j] will hold the minimum cost of multiplying matrices i..j.
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	// For a single matrix (i == j), the cost is 0.
	// L is the chain length.
	for L := 2; L < N; L++ {
		for i := 1; i <= N-L; i++ {
			j := i + L - 1
			dp[i][j] = INF
			for k := i; k < j; k++ {
				cost := dp[i][k] + dp[k+1][j] + arr[i-1]*arr[k]*arr[j]
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}
	return dp[1][N-1]
}
