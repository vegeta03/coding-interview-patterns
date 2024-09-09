//go:build ignore
// +build ignore

package main

import (
	"container/heap"
	"fmt"
)

func modify(s []int) {
	s[0] = 100
}

// Heaps
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Node represents a node in a binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	if n == nil {
		fmt.Println("Node is nil")
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}

	n.Left.InOrderTraversal()

	fmt.Print(n.Value, " ")

	n.Right.InOrderTraversal()
}

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	vertices int           // Total number of vertices in the graph
	adjList  map[int][]int // Adjacency list: key is vertex, value is slice of adjacent vertices
}

// NewGraph creates and initializes a new Graph with a given number of vertices
// It returns a pointer to the newly created Graph
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int), // Initialize the adjacency list as an empty map
	}
}

// AddEdge adds an edge between vertices v and w
// Since this is an undirected graph, the edge is added for both vertices
func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w) // Add w to v's list of adjacent vertices
	g.adjList[w] = append(g.adjList[w], v) // Add v to w's list of adjacent vertices
}

// BFS performs a Breadth-First Search traversal of the graph starting from the given vertex
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool) // Map to keep track of visited vertices
	queue := []int{start}         // Initialize queue with the starting vertiex

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			visited[vertex] = true
			fmt.Print(vertex, " ")
		}

		for _, neighbor := range g.adjList[vertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	// Creating a int slice
	numbers := []int{1, 2, 3, 5, 6}
	fmt.Printf("slice: %[1]v\n", numbers)

	// Appending to a slice
	numbers = append(numbers, 7, 8)
	fmt.Printf("Appending to a slice: %[1]v\n", numbers)

	// Slicing
	fmt.Printf("Slicing [1:4]: %[1]v\n", numbers[1:4])

	// Length and capacity
	fmt.Printf("Length: %[1]v, Capacity: %[2]v\n", len(numbers), cap(numbers))

	// Slice as a function parameter
	nums := []int{1, 2, 3}
	modify(nums)
	fmt.Println(nums)

	// Creating a map
	ages := make(map[string]int)
	// adding key/value pairs
	ages["Alice"] = 30
	ages["Neo"] = 25

	// Accessing values
	fmt.Println("Alice age: ", ages["Alice"])

	// Checking if a key exists
	age, ok := ages["Smith"]
	if !ok {
		fmt.Println("Smith does not exist")
	} else {
		fmt.Println("Smith age: ", age)
	}

	// Deleting a key-value pair
	delete(ages, "Neo")
	fmt.Printf("Ages %v\n", ages)

	ages["Neo"] = 25
	ages["Smith"] = 31

	// Iterating Over a Map
	for name, age := range ages {
		fmt.Printf("%[1]s is %[2]d years old\n", name, age)
	}

	// Create a new IntHeap with initial values
	h := &IntHeap{2, 7, 5}

	// Initialize the heap
	// This call is necessary to establish the heap invariants
	heap.Init(h)
	fmt.Printf("Values in the heap h: %v\n", h)

	// Push a new element onto the heap
	heap.Push(h, 1)
	fmt.Printf("Values in the heap h: %v\n", h)

	// Print the minimum element
	// In a min-heap, the minimum element is always at index 0
	fmt.Printf("minimum: %d\n", (*h)[0])

	// Pop and print all elements form the heap
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Printf("\n")
	fmt.Printf("Values in the heap h: %v\n", h)

	// Create the root node of the binary search tree with value 5
	bstRoot := &Node{Value: 5}

	bstRoot.Insert(3)
	bstRoot.Insert(7)
	bstRoot.Insert(1)

	fmt.Println("BST InOrderTraversal")
	bstRoot.InOrderTraversal()

	fmt.Printf("\n")

	g := NewGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	fmt.Println("BFS starting from vertex 2:")
	g.BFS(2) // Perform BFS starting from vertex 2

	fmt.Printf("\n")
}
