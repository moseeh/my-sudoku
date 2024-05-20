package utilities

func SolveSudoku(grid *[9][9]int) bool {
	emptyRow, emptyCol := FindEmptyCell(grid)
	if emptyRow == -1 && emptyCol == -1 {
		return true
	}

	for num := 1; num <= 9; num++ {
		if IsSafe(grid, emptyRow, emptyCol, num) {
			grid[emptyRow][emptyCol] = num

			if SolveSudoku(grid) {
				return true
			}

			grid[emptyRow][emptyCol] = 0
		}
	}

	return false
}
