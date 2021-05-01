package graphutils

import "strconv"

type Graph struct {
	VertexArray []*Vertex
}

type Vertex struct {
	Id      string
	Visited bool
	AdjEdge []*Edge
}

type Edge struct {
	Source      *Vertex
	Destination *Vertex
	Weight      int
}

func NewGraph() *Graph {
	return &Graph{
		make([]*Vertex, 0),
	}
}

func NewVertex(input_id string) *Vertex {
	return &Vertex{
		Id:      input_id,
		Visited: false,
		AdjEdge: make([]*Edge, 0),
	}
}

func NewEdge(source, destination *Vertex, weight int) *Edge {
	return &Edge{
		source,
		destination,
		weight,
	}
}

func StrToInt(input_str string) int {
	result, err := strconv.Atoi(input_str)
	if err != nil {
		panic("failed to convert string")
	}
	return result
}

func (G *Graph) AddVertices(more ...*Vertex) {
	for _, vertex := range more {
		G.VertexArray = append(G.VertexArray, vertex)
	}
}

func (A *Vertex) AddEdges(more ...*Edge) {
	for _, edge := range more {
		A.AdjEdge = append(A.AdjEdge, edge)
	}
}

//Find the node with the id, or create it.
func (G *Graph) GetOrConst(id string) *Vertex {
	vertex := G.GetVertexByID(id)
	if vertex == nil {
		vertex = NewVertex(id)
		G.AddVertices(vertex)
	}
	return vertex
}

func (G *Graph) GetVertexByID(id string) *Vertex {
	for _, vertex := range G.VertexArray {
		if vertex.Id == id {
			return vertex
		}
	}
	return nil
}

func (A *Vertex) GetAdEdg() chan *Edge {
	edgechan := make(chan *Edge)

	go func() {
		defer close(edgechan)
		for _, edge := range A.AdjEdge {
			edgechan <- edge
		}
	}()

	return edgechan
}

const MAXWEIGHT = 1000000

type MinDistanceFromSource map[*Vertex]int

func (G *Graph) Dijks(StartSource, TargetSource *Vertex) MinDistanceFromSource {
	D := make(MinDistanceFromSource)
	for _, vertex := range G.VertexArray {
		D[vertex] = MAXWEIGHT
	}
	D[StartSource] = 0

	for edge := range StartSource.GetAdEdg() {
		D[edge.Destination] = edge.Weight
	}
	CalculateDistance(StartSource, TargetSource, D)
	return D
}

func CalculateDistance(StartSource, TargetSource *Vertex, D MinDistanceFromSource) {
	for edge := range StartSource.GetAdEdg() {
		if D[edge.Destination] > D[edge.Source]+edge.Weight {
			D[edge.Destination] = D[edge.Source] + edge.Weight
		} else if D[edge.Destination] < D[edge.Source]+edge.Weight {
			continue
		}
		CalculateDistance(edge.Destination, TargetSource, D)
	}
}
