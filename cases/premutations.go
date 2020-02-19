package cases

import "fmt"

/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
 */

func DFSPre(nums[]int, depth int, res*[][]int, visited map[int]interface{}, path[]int) {
	if depth == len(nums) {
		*res = append(*res, path)
		return
	}
	for _, num := range nums {
		_, ok := visited[num]
		if ok {continue}
		visited[num] = nil
		DFSPre(nums, depth+1, res, visited, append(path, num))
		delete(visited, num)
	}
}

func permute(nums []int) [][]int {
	rs := &[][]int{}
	visited := map[int]interface{}{}
	var path []int
	DFSPre(nums, 0, rs, visited, path)
	return *rs
}

func TestPremute(){
	path := []int{1, 2, 3, 4, 5 }
	t1 := path[:3]
	t1 = append(t1, []int{10}...)
	fmt.Println(path)
	fmt.Printf("%p %p\n", path,t1 )
	t1 = append(t1, []int{10, 12, 13}...)
	fmt.Println(path)
	fmt.Printf("%p %p\n", path,t1 )
	t := []int{5, 4, 6, 2}
	rs := permute(t)
	fmt.Println(rs)
}