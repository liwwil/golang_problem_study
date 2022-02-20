//find the sums of two matrix then print them out as 2d list
package main

import "fmt"

func main() {
	matrix1 := [][]int{ //declare matrix1
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{ //declare matrix2
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	n_rows := len(matrix1) //number of row
	n_columns := len(matrix1[0])
	sum_matrix := [][]int{ // Initial values of sum_matrix
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	for r := 0; r < n_rows; r++ {
		for c := 0; c < n_columns; c++ {
			sum_matrix[r][c] = matrix1[r][c] + matrix2[r][c]
		}
	}
	fmt.Println("Sum of Matrix:", sum_matrix)
}
