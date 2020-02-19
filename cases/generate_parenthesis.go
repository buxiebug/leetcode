package cases

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
 */

func generateParenthesis(n int) []string {
	dp := map[int][]string{0: {""}, 1: {"()"}}
	i := 2
	for i <= n {
		var tmp []string
		j := 0
		for j < i {
			leftV, _:= dp[j]
			rightV, _:= dp[i-1-j]
			for _, leftItem := range leftV {
				for _, rightItem := range rightV {
					tmp = append(tmp, "(" + leftItem + ")" + rightItem)
				}
			}
			j += 1
		}
		dp[i] = tmp
		i += 1
	}
	return dp[n]
}
