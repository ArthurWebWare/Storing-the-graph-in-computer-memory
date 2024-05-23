package main

import (
	"fmt"
)

// Graph представляет граф
type Graph struct {
	adjacencyList   map[int][]int
	incidenceMatrix [][]int
	adjacencyMatrix [][]int
}

// Edge представляет ребро графа
type Edge struct {
	start, end int
}

// Функция для отображения первого меню
func displayInitialMenu() {
	fmt.Println("Select the input format:")
	fmt.Println("1 - Incidence Matrix")
	fmt.Println("2 - Adjacency Matrix")
	fmt.Println("3 - Adjacency List")
	fmt.Println("0 - Exit")
}

// Функция для отображения второго меню
func displayOutputMenu() {
	fmt.Println("Select the output format:")
	fmt.Println("1 - Incidence Matrix")
	fmt.Println("2 - Adjacency Matrix")
	fmt.Println("3 - Adjacency List")
	fmt.Println("0 - Exit")
}

func main() {

}
