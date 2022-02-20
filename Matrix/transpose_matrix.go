//Transpose Matrix 3x3
package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	for r := 0; r < len(m); r++ {
		fmt.Println(m[r])
	}
	fmt.Println("/n")
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			result[j][i] = m[i][j]

		}
	}
	for res := 0; res < len(result); res++ {
		fmt.Println(result[res])
	}
}
