package cases

import "fmt"

/**
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	num1 := len(word1)
	num2 := len(word2)
	i, j := 0, 0
	for i <= num1 {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
		i += 1
	}
	for j <= num2 {
		dp[0][j] = j
		j += 1
	}
	i, j = 1, 1
	for i <= num1 {
		j = 1
		for j <= num2 {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := dp[i-1][j-1]
				if dp[i-1][j] < min { min = dp[i-1][j]}
				if dp[i][j-1] < min { min = dp[i][j-1]}
				dp[i][j] += min + 1
			}
			j += 1
		}
		i += 1
	}
	return dp[num1][num2]
}

func TestEditDistance() {
	s1 := "horse"
	s2 := "ros"
	fmt.Println(minDistance(s1, s2))
}
