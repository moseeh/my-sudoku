package utilities

func UsedInSubgrid(grid *[9][9]int, startRow, startCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[startRow+row][startCol+col] == num {
				return true
			}
		}
	}
	return false
}
