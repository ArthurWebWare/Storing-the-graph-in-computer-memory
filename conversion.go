package main

import "fmt"

// Функция для преобразования из матрицы инцидентности в список смежности
func incidenceMatrixToAdjacencyList(matrix [][]int) map[int][]int {
	adjacencyList := make(map[int][]int)
	numVertices := len(matrix)
	if numVertices == 0 {
		fmt.Println("Error: Incidence matrix has no vertices.")
		return adjacencyList
	}
	numEdges := len(matrix[0])

	for j := 0; j < numEdges; j++ {
		vertices := []int{}
		for i := 0; i < numVertices; i++ {
			if matrix[i][j] == 1 {
				vertices = append(vertices, i)
			}
		}
		if len(vertices) == 2 {
			start, end := vertices[0], vertices[1]
			adjacencyList[start] = append(adjacencyList[start], end)
			adjacencyList[end] = append(adjacencyList[end], start)
		} else {
			fmt.Printf("Error: Edge %d does not connect exactly two vertices.\n", j)
		}
	}
	return adjacencyList
}

// Функция для преобразования из матрицы смежности в список смежности
func adjacencyMatrixToAdjacencyList(matrix [][]int) map[int][]int {
	adjacencyList := make(map[int][]int)
	numVertices := len(matrix)
	for i := 0; i < numVertices; i++ {
		for j := 0; j < numVertices; j++ {
			if matrix[i][j] == 1 {
				adjacencyList[i] = append(adjacencyList[i], j)
			}
		}
	}
	return adjacencyList
}

// Функция для преобразования из списка смежности в матрицу инцидентности
func adjacencyListToIncidenceMatrix(adjacencyList map[int][]int) [][]int {
	edges := []Edge{}
	for start, neighbors := range adjacencyList {
		for _, end := range neighbors {
			if start < end {
				edges = append(edges, Edge{start, end})
			}
		}
	}
	numVertices := len(adjacencyList)
	numEdges := len(edges)
	matrix := make([][]int, numVertices)
	for i := range matrix {
		matrix[i] = make([]int, numEdges)
	}
	for j, edge := range edges {
		matrix[edge.start][j] = 1
		matrix[edge.end][j] = 1
	}
	return matrix
}

// Функция для преобразования из списка смежности в матрицу смежности
func adjacencyListToAdjacencyMatrix(adjacencyList map[int][]int) [][]int {
	numVertices := len(adjacencyList)
	matrix := make([][]int, numVertices)
	for i := range matrix {
		matrix[i] = make([]int, numVertices)
	}
	for start, neighbors := range adjacencyList {
		for _, end := range neighbors {
			matrix[start][end] = 1
		}
	}
	return matrix
}

// Функция для преобразования из матрицы инцидентности в матрицу смежности
func incidenceMatrixToAdjacencyMatrix(incidenceMatrix [][]int) [][]int {
	numVertices := len(incidenceMatrix)
	if numVertices == 0 {
		return nil
	}
	numEdges := len(incidenceMatrix[0])

	// Инициализация матрицы смежности нулями
	adjacencyMatrix := make([][]int, numVertices)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, numVertices)
	}

	// Заполнение матрицы смежности на основе матрицы инцидентности
	for j := 0; j < numEdges; j++ {
		start, end := -1, -1
		for i := 0; i < numVertices; i++ {
			if incidenceMatrix[i][j] == 1 {
				if start == -1 {
					start = i
				} else {
					end = i
					break
				}
			}
		}
		if start != -1 && end != -1 {
			adjacencyMatrix[start][end] = 1
			adjacencyMatrix[end][start] = 1
		}
	}

	return adjacencyMatrix
}
