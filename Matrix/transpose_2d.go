//Transpose Matrix for 2 x 4

package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
	}
	result := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	for r := 0; r < len(m); r++ {
		fmt.Println(m[r])
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			result[j][i] = m[i][j]
		}
	}

	for res := 0; res < len(result); res++ {
		fmt.Println(result[res])
	}
}
