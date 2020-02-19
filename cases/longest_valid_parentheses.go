package cases

import "fmt"

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func longestValidParentheses(s string) int {
	max := 0
	dp := map[int]int{0 : 0}
	if len(s) > 1 {
		if s[:2] == "()" {dp[1] = 1; max=1} else {dp[1]=0}
	}
	idx := 1
	for idx < len(s){
		if s[idx] == '(' {
			dp[idx] = 0
			idx += 1
			continue
		}
		if s[idx-1] == '(' {
			dp[idx] = dp[idx-2] + 2
		} else {
			if idx-dp[idx-1]-1 >= 0 && s[idx-dp[idx-1]-1] == '(' {
				dp[idx] = dp[idx-1] + dp[idx-dp[idx-1]-2] + 2
			}
		}
		if dp[idx] > max { max=dp[idx]}
		idx += 1
	}
	return max
}

func TestLongestValidParenthesesDP(){
	fmt.Println(longestValidParentheses("(()())"))
}