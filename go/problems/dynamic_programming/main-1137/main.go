// LeetCode Main Problem Nr.1137: N-th Tribonacci Number

package main

import "fmt"

func tribonacci(n int) int {
	return method1(n)
}

func method1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func main() {
	for i := 0; i <= 37; i++ {
		fmt.Printf("T(%d) = %d\n", i, tribonacci(i))
	}
}
