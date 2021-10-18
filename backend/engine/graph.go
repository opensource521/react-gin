package engine

type Graph struct {
	size int         // Number of points in graph
	adj  [101]Vector // Array of vector. Each vector is an array of adjacent vertexes of one
}

// Add edge to graph
func (g *Graph) AddEdge(u int, v int) {
	(*g).adj[u].Push(v)
	(*g).adj[v].Push(u)
}

// Set size of graph
func (g *Graph) SetSize(s int) {
	(*g).size = s
}

// Get size of graph
func (g *Graph) GetSize() int {
	return (*g).size
}

// Check if two points are connected in graph
func (g *Graph) HasEdge(u int, v int) bool {
	return (*g).adj[u].Has(v)
}

// Get adjacent vertexes of one
func (g *Graph) GetAdjacentVertexes(u int) Vector {
	return (*g).adj[u]
}