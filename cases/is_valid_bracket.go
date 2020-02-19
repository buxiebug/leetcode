package cases

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
https://leetcode-cn.com/problems/valid-parentheses/solution/valid-parentheses-fu-zhu-zhan-fa-by-jin407891080/170276
 */

type Stack struct {
	container []interface{}
}

func (this *Stack) Pop() (interface{}, bool){
	if len(this.container) == 0 {
		return "", false
	}
	rs := this.container[len(this.container) - 1]
	this.container = this.container[:len(this.container) - 1]
	return rs, true
}

func (this *Stack) Push(s interface{}) {
	this.container = append(this.container, s)
}

func (this *Stack) Len() int{
	return len(this.container)
}

func isValid(s string) bool {
	stack := Stack{container:[]interface{}{}}
	bracket := map[string]interface{}{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	for _, ele := range s{
		char := string(ele)
		_, ok := bracket[char]
		if ok {
			stack.Push(string(char))
		} else {
			c, ok := stack.Pop()
			if ok {
				lastChar, ok_ := bracket[c.(string)]
				if ok_ && lastChar == char { continue }
			}
			return false
		}
	}
	if stack.Len() > 0 { return false }
	return true
}

func TestIsValidBracket(){
	fmt.Println(isValid("()"))
}