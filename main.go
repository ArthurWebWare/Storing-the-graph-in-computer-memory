package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var graph Graph
	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayInitialMenu()
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			graph.incidenceMatrix = inputIncidenceMatrix(scanner)
			if graph.incidenceMatrix != nil {
				activeEntity = choice
				outputIncidenceMatrix(graph.incidenceMatrix)
			} else {
				secondMenuActive = false
			}
		case 2:
			graph.adjacencyMatrix = inputAdjacencyMatrix(scanner)
			if graph.adjacencyMatrix != nil {
				activeEntity = choice
				outputAdjacencyMatrix(graph.adjacencyMatrix)
			} else {
				secondMenuActive = false
			}
		case 3:
			graph.adjacencyList = inputAdjacencyList(scanner)
			if graph.adjacencyList != nil {
				activeEntity = choice
				outputAdjacencyList(graph.adjacencyList)
			} else {
				secondMenuActive = false
			}
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			secondMenuActive = false
		}
	}
}
