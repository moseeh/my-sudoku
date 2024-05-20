package utilities

func UsedInRow(grid *[9][9]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}
