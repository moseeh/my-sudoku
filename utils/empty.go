package utilities

func FindEmptyCell(grid *[9][9]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}
