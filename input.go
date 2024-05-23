package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Функция для ввода матрицы инцидентности
func inputIncidenceMatrix(scanner *bufio.Scanner) [][]int {
	fmt.Println("Enter the number of vertices and edges (vertices edges):")
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	if len(input) != 2 {
		fmt.Println("Invalid input. Please enter the number of vertices and edges separated by a space.")
		return nil
	}
	numVertices, err1 := strconv.Atoi(input[0])
	numEdges, err2 := strconv.Atoi(input[1])
	if err1 != nil || err2 != nil || numVertices <= 0 || numEdges <= 0 {
		fmt.Println("Invalid input. Please enter valid numbers for vertices and edges.")
		return nil
	}
	matrix := make([][]int, numVertices)
	for i := range matrix {
		matrix[i] = make([]int, numEdges)
	}
	fmt.Println("Enter the incidence matrix:")
	for i := 0; i < numVertices; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		if len(row) != numEdges {
			fmt.Println("Invalid input. Please enter the correct number of edges for each vertex.")
			return nil
		}
		for j := 0; j < numEdges; j++ {
			matrix[i][j], err1 = strconv.Atoi(row[j])
			if err1 != nil {
				fmt.Println("Invalid input. Please enter valid numbers for the incidence matrix.")
				return nil
			}
		}
	}
	return matrix
}
