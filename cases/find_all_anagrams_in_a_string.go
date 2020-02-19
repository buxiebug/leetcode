package cases

import "fmt"

/**
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例 1:

输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:

输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func findAnagrams(s string, p string) []int {
	need, find := map[string]int{}, map[string]int{}
	matched := 0
	var rs []int
	for i:=0 ; i< len(p); i++{
		need[string(p[i])] += 1
	}

	left, right := 0, 0
	for right < len(s) {
		if matched < len(p) {
			find[string(s[right])] += 1
			if find[string(s[right])] == need[string(s[right])] {
				matched += 1
			}
			right += 1
		}
		if matched == len(need) {
			for left < right {
				find[string(s[left])] -= 1
				if find[string(s[left])] < need[string(s[left])] {
					matched -= 1
					break
				}
				left += 1
			}
			if right - left == len(p) {
				rs = append(rs, left)
			}
			left += 1
		}
	}
	return rs
}

func TestFindAllAnagramsInAString() {
	fmt.Println(findAnagrams("baa", "aa"))
}
