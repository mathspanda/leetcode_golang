package main

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[int]bool)
	cols := make(map[int]map[int]bool)
	grids := make(map[int]map[int]bool)

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if rows[i] == nil {
					rows[i] = make(map[int]bool)
				}
				if rows[i][num] == false {
					rows[i][num] = true
				} else {
					return false
				}
				if cols[j] == nil {
					cols[j] = make(map[int]bool)
				}
				if cols[j][num] == false {
					cols[j][num] = true
				} else {
					return false
				}
				gridIndex := i/3*3 + j/3
				if grids[gridIndex] == nil {
					grids[gridIndex] = make(map[int]bool)
				}
				if grids[gridIndex][num] == false {
					grids[gridIndex][num] = true
				} else {
					return false
				}
			}
		}
	}
	return true
}
