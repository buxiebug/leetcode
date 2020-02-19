package cases

import "fmt"

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func explore(board [][]byte, visited map[int]map[int]bool, x int, y int, wordCount int, word string) bool{
	if wordCount == len(word) {
		return true
	}
	if wordCount >= len(word) {
		return false
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]){
		return false
	}
	if word[wordCount] != board[x][y] {
		return false
	}
	_, leftOk := visited[x][y-1]
	_, rightOk := visited[x][y+1]
	_, topOk := visited[x-1][y]
	_, bottomOk := visited[x+1][y]
	visited[x][y] = true
	if !leftOk {
		rsLeft := explore(board, visited, x, y-1, wordCount+1, word)
		if rsLeft { return rsLeft }
	}
	if !rightOk {
		rsRight := explore(board, visited, x, y+1, wordCount+1, word)
		if rsRight { return rsRight }
	}
	if !topOk {
		rsTop := explore(board, visited, x-1, y, wordCount+1, word)
		if rsTop { return rsTop }
	}
	if !bottomOk {
		rsBottom := explore(board, visited, x+1, y, wordCount+1, word)
		if rsBottom { return rsBottom }
	}
	delete(visited[x], y)
	return false
}

func exist(board [][]byte, word string) bool {
	visited := map[int]map[int]bool{}
	i, j := 0, 0
	for i < len(board) {
		visited[i] = map[int]bool{}
		i += 1
	}
	i = 0
	for i < len(board) {
		j = 0
		for j < len(board[0]) {
			if explore(board, visited, i, j, 0, word) {
				return true
			}
			j += 1
		}
		i += 1
	}
	return false
}

func TestWordSearch() {
	//board := [][]byte{
	//	[]byte{'A', 'B', 'C', 'E'},
	//	[]byte{'S', 'F', 'C', 'S'},
	//	[]byte{'A', 'D', 'E', 'E'},
	//}
	//fmt.Println(exist(board, "ABCCED"))
	//board := [][]byte{
	//	[]byte{'A', 'B'},
	//	[]byte{'C', 'D'},
	//}
	//fmt.Println(exist(board, "CDBA"))
	board := [][]byte{
		[]byte{'b', 'a', 'a', 'b', 'a', 'b'},
		[]byte{'a', 'b', 'a', 'a', 'a', 'a'},
		[]byte{'a', 'b', 'a', 'a', 'a', 'b'},
		[]byte{'a', 'b', 'a', 'b', 'b', 'a'},
		[]byte{'a', 'a', 'b', 'b', 'a', 'b'},
		[]byte{'a', 'a', 'b', 'b', 'b', 'a'},
		[]byte{'a', 'a', 'b', 'a', 'a', 'b'},
	}
	fmt.Println(exist(board, "aabbbbabbaababaaaabababbaaba"))
}