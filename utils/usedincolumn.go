package utilities

func UsedInColumn(grid *[9][9]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}
