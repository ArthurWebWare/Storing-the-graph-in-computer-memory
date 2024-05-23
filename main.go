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
	secondMenuActive := true
	var activeEntity int

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

		for secondMenuActive {
			displayOutputMenu()
			scanner.Scan()
			outputChoice, _ := strconv.Atoi(scanner.Text())

			switch outputChoice {
			case 1:
				//Incidence Matrix
				if activeEntity != outputChoice {
					activeEntity = outputChoice
					switch getActiveEntityId(graph) {
					case 2:
						graph.incidenceMatrix = adjacencyListToIncidenceMatrix(
							adjacencyMatrixToAdjacencyList(graph.adjacencyMatrix),
						)
					case 3:
						graph.incidenceMatrix = adjacencyListToIncidenceMatrix(graph.adjacencyList)
					}
				}

				clearInactiveEntities(activeEntity, &graph)

				outputIncidenceMatrix(graph.incidenceMatrix)
			case 2:
				//Adjacency Matrix
				if activeEntity != outputChoice {
					activeEntity = outputChoice
					switch getActiveEntityId(graph) {
					case 1:
						graph.adjacencyMatrix = incidenceMatrixToAdjacencyMatrix(graph.incidenceMatrix)
					case 3:
						graph.adjacencyMatrix = adjacencyListToAdjacencyMatrix(graph.adjacencyList)
					}
				}

				clearInactiveEntities(activeEntity, &graph)

				outputAdjacencyMatrix(graph.adjacencyMatrix)
			case 3:
				//Adjacency Matrix
				if activeEntity != outputChoice {
					activeEntity = outputChoice
					switch getActiveEntityId(graph) {
					case 1:
						graph.adjacencyList = incidenceMatrixToAdjacencyList(graph.incidenceMatrix)
					case 2:
						graph.adjacencyList = adjacencyMatrixToAdjacencyList(graph.adjacencyMatrix)
					}
				}

				clearInactiveEntities(activeEntity, &graph)

				outputAdjacencyList(graph.adjacencyList)
			case 0:
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
}

func getActiveEntityId(graph Graph) int {
	if len(graph.incidenceMatrix) != 0 {
		return 1
	} else if len(graph.adjacencyMatrix) != 0 {
		return 2
	} else if len(graph.adjacencyList) != 0 {
		return 3
	} else {
		return 0
	}
}

func clearInactiveEntities(activeEntityId int, graph *Graph) {
	switch activeEntityId {
	case 1:
		graph.adjacencyMatrix = nil
		graph.adjacencyList = nil
	case 2:
		graph.incidenceMatrix = nil
		graph.adjacencyList = nil
	case 3:
		graph.incidenceMatrix = nil
		graph.adjacencyMatrix = nil
	}
}
