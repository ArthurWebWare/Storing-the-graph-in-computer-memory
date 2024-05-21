# Laboratory Work №1: Graph Storage in Computer Memory

## Description

This laboratory work is dedicated to studying various methods of storing a graph in computer memory. During the course of this work, procedures will be developed for inputting a graph, converting it between different forms of representation, and displaying information about the graph.

## Task

### 1. Graph Input

Develop procedures for inputting a graph in the following forms:
- Incidence Matrix
- Adjacency Matrix
- Adjacency List

Ensure the ability to correct the entered data.

### 2. Conversion Between Graph Storage Forms

Develop procedures for converting between different graph storage forms:
- From adjacency matrix to adjacency list and back
- From incidence matrix to adjacency list and back

### 3. Graph Manipulation Program

Using the procedures mentioned above, develop a program that performs the following functions:
- Input a graph in any of the three forms (as requested by the user) with the ability to correct the data
- Store the entered graph in the computer's memory as an adjacency list
- Display information about the graph in any of the three forms (as requested by the user) on the screen

## Project Structure

- `main.go` — the main file of the program containing the implementation of all procedures and the main loop for interacting with the user
- `input.go` — module for inputting the graph in different forms
- `conversion.go` — module for converting the graph between different forms of representation
- `output.go` — module for displaying the graph in various forms

## Usage Examples

### Graph Input

The user can choose the form of input for the graph (incidence matrix, adjacency matrix, or adjacency list) and enter the data. The program will provide the ability to correct the entered data.

### Graph Conversion

The user can convert the graph from one form of representation to another using the appropriate procedures.

### Graph Information Output

The user can request the display of graph information in any of the three forms on the screen.

## Running the Project

To run the project, execute the following command:

```bash
go run main.go
```
## Conclusion

In this laboratory work, various methods for storing and processing graphs in computer memory will be developed and tested. 
The implementation of all tasks will allow consolidating theoretical knowledge through practice and improving programming skills.
