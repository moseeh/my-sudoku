package main

import (
	"fmt"
	"os"
	utilities "sudoku/utils"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error: not enough arguments")
		return
	}

	sudokugrid := [9][9]int{}

	for i := 1; i <= 9; i++ {
		row := os.Args[i]
		if len(row) != 9 {
			fmt.Println("Error not enough rows at", i)
			return
		}

		for j := 0; j < 9; j++ {
			if row[j] == '.' {
				sudokugrid[i-1][j] = 0
			} else {
				sudokugrid[i-1][j] = int(row[j] - '0')
			}
		}
	}

	if utilities.SolveSudoku(&sudokugrid) {
		utilities.PrintGrid(sudokugrid)
	} else {
		fmt.Println("Error")
	}
}
