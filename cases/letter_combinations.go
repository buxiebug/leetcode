package cases

import "fmt"

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

phone keyboard

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
DFS
 */

func DFS(digits string, tmp string, idx int, rs*[]string, charMap map[string][]string) {
	if idx == len(digits) {
		*rs = append(*rs, tmp)
		return
	}
	char := fmt.Sprintf("%c", digits[idx])
	words, _ := charMap[char]
	for _, word := range words {
		DFS(digits, tmp + word, idx+1, rs, charMap)
	}
}

func letterCombinations(digits string) []string {
	rs := &[]string{}
	if len(digits) == 0{
		return *rs
	}
	charMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	DFS(digits, "", 0, rs, charMap)
	return *rs
}


func TestLetterCombinations(){
	rs := letterCombinations("23")
	fmt.Println(rs)
}
