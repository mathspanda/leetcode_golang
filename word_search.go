package main

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	for i, row := range board {
		for j, col := range row {
			if col == word[0] {
				visited := initVisited(board)
				visited[i][j] = true
				if has(board, visited, i, j, word, 1) {
					return true
				}
			}
		}
	}
	return false
}

func initVisited(board [][]byte) [][]bool {
	visited := make([][]bool, len(board))
	for i, b := range board {
		visited[i] = make([]bool, len(b))
	}
	return visited
}

func has(board [][]byte, visited [][]bool, rowIndex, colIndex int, word string, wordIndex int) bool {
	if wordIndex == len(word) {
		return true
	}
	direction := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for i := 0; i < len(direction); i++ {
		tempRowIndex := rowIndex + direction[i][0]
		tempColIndex := colIndex + direction[i][1]
		if tempRowIndex < 0 || tempRowIndex >= len(board) {
			continue
		}
		if tempColIndex < 0 || tempColIndex >= len(board[tempRowIndex]) {
			continue
		}
		if !visited[tempRowIndex][tempColIndex] && word[wordIndex] == board[tempRowIndex][tempColIndex] {
			visited[tempRowIndex][tempColIndex] = true
			if has(board, visited, tempRowIndex, tempColIndex, word, wordIndex+1) {
				return true
			} else {
				visited[tempRowIndex][tempColIndex] = false
			}
		}
	}
	return false
}
