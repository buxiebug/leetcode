package cases

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return s
	}
	maxLen := 1
	maxStart := 0
	dp := make([][]bool, sLen)
	for i:=0; i<sLen; i++{
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
	}
	for j:=1; j<sLen; j++{
		for i:=0; i<j; i++{
			if s[i] == s[j]{
				if j-i < 3{
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] == true {
				if j-i +1 > maxLen{
					maxLen = j-i+1
					maxStart = i
				}
			}
		}
	}
	return s[maxStart: maxStart+maxLen]
}


func TestLongestPalindrome(){
	fmt.Println(longestPalindrome("cbbd"))
}
