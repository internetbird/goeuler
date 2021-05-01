package main

import (
	"strconv"
	"strings"

	"github.com/internetbird/goeuler/graphutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler18() {

	inputData := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

	rows := strings.Split(inputData, "\n")
	numsMatrix := [15][15]int{}

	for i := 0; i < len(rows); i++ {
		//	fmt.Printf("\n%d : %s", i, rows[i])
		nums := strings.Split(rows[i], " ")

		for j := 0; j < i+1; j++ {
			numsMatrix[i][j], _ = strconv.Atoi(nums[j])
			numsMatrix[i][j] = -numsMatrix[i][j]
		}
	}

	//Generate a graph from the problem input
	graph := GenerateProblemGraph()
	graph = AddEdges(numsMatrix, graph)

	//And find the shortest path
	sourceVertex := graph.GetVertexByID("s")
	destVertex := graph.GetVertexByID("t")
	distmap := graph.Dijks(sourceVertex, destVertex)

	utils.PrintAnswer(18, distmap[destVertex])
}

func GetVertexID(row int, column int) string {
	return strconv.Itoa(row) + "." + strconv.Itoa(column)
}

func AddEdges(numsMatrix [15][15]int, graph *graphutils.Graph) *graphutils.Graph {

	//Add start edge
	startVertex := graph.GetVertexByID("s")
	topVertexID := GetVertexID(1, 1)
	topVertex := graph.GetVertexByID(topVertexID)

	//Add the start edge
	firstEdge := graphutils.NewEdge(startVertex, topVertex, numsMatrix[0][0])
	startVertex.AddEdges(firstEdge)

	//Add the middle edges
	for i := 2; i <= 15; i++ {
		for j := 1; j <= i; j++ {

			currentVertexID := GetVertexID(i, j)

			currentVertex := graph.GetVertexByID(currentVertexID)
			weight := numsMatrix[i-1][j-1]

			if j == 1 {
				sourceVertexID := GetVertexID(i-1, 1)
				sourceVertex := graph.GetVertexByID(sourceVertexID)

				edge := graphutils.NewEdge(sourceVertex, currentVertex, weight)
				sourceVertex.AddEdges(edge)

			} else if j == i {

				sourceVertexID := GetVertexID(i-1, j-1)
				sourceVertex := graph.GetVertexByID(sourceVertexID)

				edge := graphutils.NewEdge(sourceVertex, currentVertex, weight)
				sourceVertex.AddEdges(edge)

			} else {

				sourceVertexID1 := GetVertexID(i-1, j-1)
				sourceVertex1 := graph.GetVertexByID(sourceVertexID1)

				sourceVertexID2 := GetVertexID(i-1, j)
				sourceVertex2 := graph.GetVertexByID(sourceVertexID2)

				edge1 := graphutils.NewEdge(sourceVertex1, currentVertex, weight)
				edge2 := graphutils.NewEdge(sourceVertex2, currentVertex, weight)

				sourceVertex1.AddEdges(edge1)
				sourceVertex2.AddEdges(edge2)
			}
		}
	}

	//Add end edges
	for i := 1; i <= 15; i++ {
		sourceVertexID := GetVertexID(15, i)
		sourceVertex := graph.GetVertexByID(sourceVertexID)
		endVertex := graph.GetVertexByID("t")
		weight := 0
		edge := graphutils.NewEdge(sourceVertex, endVertex, weight)

		sourceVertex.AddEdges(edge)
	}
	return graph
}

func GenerateProblemGraph() *graphutils.Graph {
	graph := graphutils.NewGraph()

	//add start vertex
	startvertex := graphutils.NewVertex("s")
	graph.AddVertices(startvertex)

	for i := 1; i <= 15; i++ {
		for j := 1; j <= i; j++ {
			vertexID := GetVertexID(i, j)
			vertex := graphutils.NewVertex(vertexID)
			graph.AddVertices(vertex)
		}
	}

	//Add end vertex
	endvertex := graphutils.NewVertex("t")
	graph.AddVertices(endvertex)

	return graph
}
