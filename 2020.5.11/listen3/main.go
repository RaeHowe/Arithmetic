package main

import "fmt"

func main() {
	var array = [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	result := countNegatives(array)

	fmt.Println(result)

}
/*
	给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。 
	请你统计并返回 grid 中 负数 的数目。

	示例 1：
	
	输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
	输出：8
	解释：矩阵中共有 8 个负数。
	示例 2：
	
	输入：grid = [[3,2],[1,0]]
	输出：0
*/

func countNegatives(grid [][]int) int {
	var result = 0

	for i := 0; i < len(grid); i++{ //行
		for j := 0; j < len(grid[i]); j++{ //列
			if grid[i][j] < 0{
				result++
			}
		}
	}

	return result
}