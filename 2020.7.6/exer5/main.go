package main

import "fmt"

func main()  {
	var array = [][]int{ {0,1,0,0}, {1,1,1,0}, {0,1,0,0}, {1,1,0,0} }

	result := islandPerimeter(array)

	fmt.Println(result)
}

func islandPerimeter(grid [][]int) int {
	for r := 0; r < len(grid); r++{
		for c := 0; c < len(grid[0]); c++{
			if grid[r][c] == 1{
				return dfs(grid, r, c)
			}
		}
	}

	return 0
}

func dfs(grid [][]int, r, c int) int {
	if !(r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])){
		return 1
	}

	if grid[r][c] == 0{
		return 1
	}

	if grid[r][c] != 1{
		return 0
	}

	grid[r][c] = 2

	return dfs(grid, r - 1, c) + dfs(grid, r + 1, c) + dfs(grid, r, c - 1) + dfs(grid, r, c + 1)
}
