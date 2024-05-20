package utilities

func IsSafe(grid *[9][9]int, row, col, num int) bool {
	return !UsedInRow(grid, row, num) && !UsedInColumn(grid, col, num) && !UsedInSubgrid(grid, row-row%3, col-col%3, num)
}
