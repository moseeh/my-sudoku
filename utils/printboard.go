package utilities

import "fmt"

func PrintGrid(grid [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%2d ", grid[i][j])
		}
		fmt.Println()
	}
}
