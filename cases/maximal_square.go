package cases

/**
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximalSquare(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	for i:=0; i<len(matrix); i++{
		dp[i] = make([]int, len(matrix[0]))
	}
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == '0' { dp[i][j] = 0; continue}
			left, top, leftTop := 0, 0, 0
			if i-1 >= 0 {
				top = dp[i-1][j]
			}
			if j-1 >= 0 {
				left = dp[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				leftTop = dp[i-1][j-1]
			}
			min := top
			if left < min {min=left}
			if leftTop < min {min=leftTop}
			dp[i][j] = min + 1
			size := dp[i][j] * dp[i][j]
			if size > max { max = size}
		}
	}
	return max
}