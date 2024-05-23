package main

import "fmt"

// Функция для вывода матрицы инцидентности
func outputIncidenceMatrix(matrix [][]int) {
	fmt.Println("Incidence Matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// Функция для вывода матрицы смежности
func outputAdjacencyMatrix(matrix [][]int) {
	fmt.Println("Adjacency Matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// Функция для вывода списка смежности
func outputAdjacencyList(adjacencyList map[int][]int) {
	fmt.Println("Adjacency List:")
	for vertex, neighbors := range adjacencyList {
		fmt.Printf("%d: ", vertex)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}
