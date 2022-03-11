package main

import "fmt"

func main() {
	board := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]string) bool {

	arr := make(map[string]string)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == "." {
				continue
			}

			row := fmt.Sprintf("%s+%d+%s", "row", i, num)
			column := fmt.Sprintf("%s+%d+%s", "column", j, num)
			box := fmt.Sprintf("%s+%d+%s", "box", (i/3)*3+j/3, num)

			if _, isExist := arr[row]; isExist {
				return false
			}

			if _, isExist := arr[column]; isExist {
				return false
			}

			if _, isExist := arr[box]; isExist {
				return false
			}

			arr[row] = row
			arr[column] = column
			arr[box] = box
		}
	}

	return true
}
