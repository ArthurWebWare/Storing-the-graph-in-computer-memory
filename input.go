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

// Функция для ввода матрицы смежности
func inputAdjacencyMatrix(scanner *bufio.Scanner) [][]int {
	fmt.Println("Enter the number of vertices:")
	scanner.Scan()
	numVertices, err := strconv.Atoi(scanner.Text())
	if err != nil || numVertices <= 0 {
		fmt.Println("Invalid input. Please enter a valid number of vertices.")
		return nil
	}
	matrix := make([][]int, numVertices)
	for i := range matrix {
		matrix[i] = make([]int, numVertices)
	}
	fmt.Println("Enter the adjacency matrix:")
	for i := 0; i < numVertices; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		if len(row) != numVertices {
			fmt.Println("Invalid input. Please enter the correct number of vertices for each row.")
			return nil
		}
		for j := 0; j < numVertices; j++ {
			matrix[i][j], err = strconv.Atoi(row[j])
			if err != nil {
				fmt.Println("Invalid input. Please enter valid numbers for the adjacency matrix.")
				return nil
			}
		}
	}
	return matrix
}

// Функция для ввода списка смежности
func inputAdjacencyList(scanner *bufio.Scanner) map[int][]int {
	fmt.Println("Enter the number of vertices:")
	scanner.Scan()
	numVertices, err := strconv.Atoi(scanner.Text())
	if err != nil || numVertices <= 0 {
		fmt.Println("Invalid input. Please enter a valid number of vertices.")
		return nil
	}
	adjacencyList := make(map[int][]int)
	fmt.Println("Enter the adjacency list (vertex: neighbors):")
	for i := 0; i < numVertices; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), ":")
		if len(input) != 2 {
			fmt.Println("Invalid input. Please enter data in the format 'vertex: neighbors'.")
			return nil
		}
		vertex, err1 := strconv.Atoi(strings.TrimSpace(input[0]))
		neighbors := strings.Fields(input[1])
		if err1 != nil {
			fmt.Println("Invalid input. Please enter a valid vertex number.")
			return nil
		}
		for _, neighbor := range neighbors {
			n, err2 := strconv.Atoi(neighbor)
			if err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers for neighbors.")
				return nil
			}
			adjacencyList[vertex] = append(adjacencyList[vertex], n)
		}
	}
	return adjacencyList
}
