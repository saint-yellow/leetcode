// LeetCode 主站 Problem Nr. 63: 不同路径 II

package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return method1(obstacleGrid)
}

func method1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{1},
		{0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}