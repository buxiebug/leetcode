package cases

import "fmt"

/**
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//https://leetcode-cn.com/problems/number-of-islands/solution/dfs-bfs-bing-cha-ji-python-dai-ma-java-dai-ma-by-l/
// 并查集　知乎

type UnionFind struct {
	count  int
	rank   map[int]int
	parent map[int]int
}

func (uf *UnionFind) find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	for p != uf.parent[p] {
		p = uf.parent[p]
		uf.parent[p] = root
	}
	return p
}

func (uf *UnionFind) union(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else {
		uf.parent[pRoot] = qRoot
		if uf.rank[pRoot] == uf.rank[qRoot] {
			uf.rank[qRoot]++
		}
	}
	uf.count--
}

func (uf *UnionFind) getCount() int {
	return uf.count
}

func buildUnionFind(n int) *UnionFind {
	uf := &UnionFind{count: n, rank: map[int]int{}, parent: map[int]int{}}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func getIndex(x, y, width int) int {
	return x*width + y
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	uf := buildUnionFind(len(grid)*len(grid[0]) + 1)
	dummyNode := len(grid) * len(grid[0]) //索引是从０开始的
	width := len(grid[0])

	for i := 0; i < len(grid); i++ {
		for j := 0; j < width; j++ {
			right := j + 1
			bottom := i + 1
			if grid[i][j] == '0' {
				uf.union(getIndex(i, j, width), dummyNode)
				continue
			}
			if bottom < len(grid) && grid[bottom][j] == '1' {
				uf.union(getIndex(i, j, width), getIndex(bottom, j, width))
			}
			if right < len(grid[0]) && grid[i][right] == '1' {
				uf.union(getIndex(i, j, width), getIndex(i, right, width))
			}
		}
	}
	return uf.getCount() - 1
}

func TestNumIslands() {
	// grid := [][]byte{
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'1', '1', '0', '0', '0'},
	// 	[]byte{'0', '0', '1', '0', '0'},
	// 	[]byte{'0', '0', '0', '1', '1'},
	// }
	// grid := [][]byte{
	// 	[]byte{'1', '1', '1'},
	// 	[]byte{'0', '1', '0'},
	// 	[]byte{'1', '1', '1'},
	// }
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '1', '1', '1'},
		[]byte{'0', '0', '0', '0', '0', '0', '1'},
		[]byte{'1', '1', '1', '1', '1', '0', '1'},
		[]byte{'1', '0', '0', '0', '1', '0', '1'},
		[]byte{'1', '0', '1', '0', '1', '0', '1'},
		[]byte{'1', '0', '1', '1', '1', '0', '1'},
		[]byte{'1', '1', '1', '1', '1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
