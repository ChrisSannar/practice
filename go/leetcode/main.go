package main

import (
	"fmt"
)

func main() {
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrixB := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	}

	rowsA, colsA := len(matrixA), len(matrixA[0])
	colsB := len(matrixB[0])

	matrixC := make([][]int, rowsA)
	for i := range matrixC {
		matrixC[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	fmt.Println("A:")
	printMatrix(matrixA)
	fmt.Println("B:")
	printMatrix(matrixB)
	fmt.Println("A * B:")
	printMatrix(matrixC)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
