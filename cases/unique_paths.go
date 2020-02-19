package cases

import "fmt"

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func uniquePaths(m int, n int) int {
	dp := map[int]map[int]int{}
	i, j := 0, 0
	for i < m {
		dp[i] = map[int]int{0: 1}
		i += 1
	}
	for j < n {
		dp[0][j] = 1
		j += 1
	}
	i, j = 1, 1
	for i < m {
		j = 1
		for j<n {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			j += 1
		}
		i += 1
	}
	return dp[m-1][n-1]
}

func TestUniquePaths() {
	fmt.Println(uniquePaths(7, 3))
}