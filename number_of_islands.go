package main

func numIslands(grid [][]byte) int {
	num := 0
	for rowIndex, row := range grid {
		for colIndex := range row {
			if grid[rowIndex][colIndex] == '1' {
				num++
				dfs(rowIndex, colIndex, grid)
			}
		}
	}
	return num
}

func dfs(rowIndex, colIndex int, grid [][]byte) {
	if rowIndex < 0 || colIndex < 0 || rowIndex >= len(grid) || colIndex >= len(grid[rowIndex]) {
		return
	}
	if grid[rowIndex][colIndex] != '1' {
		return
	}
	grid[rowIndex][colIndex] = '0'
	dfs(rowIndex+1, colIndex, grid)
	dfs(rowIndex-1, colIndex, grid)
	dfs(rowIndex, colIndex+1, grid)
	dfs(rowIndex, colIndex-1, grid)
}
