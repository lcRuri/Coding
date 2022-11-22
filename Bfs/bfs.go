package Bfs

import (
	"math"
)

//UpdateMatrix 542. 01 矩阵 https://leetcode.cn/problems/01-matrix/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func UpdateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	q := make([][]int, 0)
	row := len(mat)
	col := len(mat[0])

	for i, _ := range res {
		res[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
				res[i][j] = 0
			} else {
				res[i][j] = math.MaxInt32
			}

		}
	}

	for len(q) > 0 {
		r := q[0][0]
		c := q[0][1]
		q = q[1:]
		for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			rr := r + v[0]
			cc := c + v[1]

			if rr >= 0 && rr < row && cc >= 0 && cc < col && res[rr][cc] == math.MaxInt32 {
				if res[r][c]+1 < res[rr][cc] {
					res[rr][cc] = res[r][c] + 1
				}

				q = append(q, []int{rr, cc})
			}
		}
	}

	return res
}

//OrangesRotting 994. 腐烂的橘子 https://leetcode.cn/problems/rotting-oranges/description/
func OrangesRotting(grid [][]int) int {
	cnts, cost, m, n, queue := 0, 0, len(grid), len(grid[0]), [][]int{}
	for i, row := range grid {
		for j, v := range row {
			if v > 0 {
				cnts++
				if v == 2 {
					queue = append(queue, []int{i, j})
				}
			}
		}
	}
	for len(queue) > 0 {
		l := len(queue)
		cnts -= l
		for i := 0; i < l; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				nx, ny := point[0]+d[0], point[1]+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		if len(queue) > 0 {
			cost++
		}
	}
	if cnts == 0 {
		return cost
	}
	return -1
}

func NumIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	//q:=make([][]int,0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				res++
			}
		}
	}

	return res
}

func bfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	if x-1 >= 0 && grid[x-1][y] == '1' {
		bfs(grid, x-1, y)
	}
	if x+1 < len(grid) && grid[x+1][y] == '1' {
		bfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		bfs(grid, x, y-1)
	}
	if y+1 < len(grid[0]) && grid[x][y+1] == '1' {
		bfs(grid, x, y+1)
	}

	return
}

var flag = make(map[int]bool, 0)

func FindCircleNum(isConnected [][]int) int {
	//纯种sb题
	//n := len(isConnected)
	//res := 0
	//flag = map[int]bool{}
	//t := 0
	//for i := 0; i < n; i++ {
	//
	//	if isConnected[0][i] == 1 {
	//		t++
	//	}
	//	if t == n {
	//		return 1
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	if flag[i] == true {
	//		continue
	//	}
	//	k := 0
	//	for j := 0; j < n; j++ {
	//		if j != i && isConnected[i][j] == 1 {
	//			k = 1
	//			flag[i] = true
	//			bfs1(isConnected, i, j)
	//			res++
	//		}
	//		if len(flag) == n {
	//			return 1
	//		}
	//	}
	//
	//	if k == 0 {
	//		res++
	//	}
	//}
	//
	//return res

	return 0
}

//func bfs1(isConnected [][]int, i, j int) {
//
//	flag[j] = true
//	if len(flag) == len(isConnected) {
//		return
//	}
//	for idx, v := range isConnected[j] {
//		if v == 1 && idx != i && idx != j {
//			bfs1(isConnected, j, idx)
//		}
//
//	}
//
//	return
//}
