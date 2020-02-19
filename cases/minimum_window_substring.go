package cases

import "fmt"

/**
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
https://leetcode-cn.com/problems/minimum-window-substring/solution/hua-dong-chuang-kou-suan-fa-tong-yong-si-xiang-by-/
 */

func minWindow(s string, t string) string {

	matched := 0
	rs := ""
	match, need := map[string]int{}, map[string]int{}
	for _, c := range t {
		need[string(c)] += 1
	}
	i, j := 0, 0
	for j < len(s) {
		if matched < len(t) {
			match[string(s[j])] += 1
			if need[string(s[j])] > 0 && match[string(s[j])] == need[string(s[j])] {
				matched += 1
			}
		}
		if matched == len(need) {
			for i < j{
				match[string(s[i])] -= 1
				if need[string(s[i])] > 0  && match[string(s[i])] < need[string(s[i])] {
					matched -= 1
					break
				}
				i += 1
			}
			if rs == "" || len(rs) > j-i+1 {
				rs = s[i: j+1]
			}
			i += 1
		}
		j += 1
	}
	return rs
}

func TestMinWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
}