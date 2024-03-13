package mst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getGraph() Graph {
   graph := Graph{}
   graph[1] = append(graph[1], Edge{ to: 6, weight: 10 })
   graph[1] = append(graph[1], Edge{ to: 2, weight: 28 })
   graph[2] = append(graph[2], Edge{ to: 3, weight: 16 })
   graph[2] = append(graph[2], Edge{ to: 7, weight: 14 })
   graph[2] = append(graph[2], Edge{ to: 1, weight: 28 })
   graph[3] = append(graph[3], Edge{ to: 4, weight: 12 })
   graph[3] = append(graph[3], Edge{ to: 2, weight: 16 })
   graph[4] = append(graph[4], Edge{ to: 7, weight: 18 })
   graph[4] = append(graph[4], Edge{ to: 5, weight: 22 })
   graph[4] = append(graph[4], Edge{ to: 3, weight: 12 })
   graph[5] = append(graph[5], Edge{ to: 7, weight: 24 })
   graph[5] = append(graph[5], Edge{ to: 6, weight: 25 })
   graph[5] = append(graph[5], Edge{ to: 4, weight: 22 })
   graph[6] = append(graph[6], Edge{ to: 1, weight: 10 })
   graph[6] = append(graph[6], Edge{ to: 5, weight: 25 })
   graph[7] = append(graph[7], Edge{ to: 2, weight: 14 })
   graph[7] = append(graph[7], Edge{ to: 4, weight: 18 })
   graph[7] = append(graph[7], Edge{ to: 5, weight: 24 })
   return graph
}

func verifyGraph(t *testing.T, mst Graph) {
   assert.Equal(t, len(mst), 6)
   assert.Equal(t, mst[1][0], Edge{ to: 6, weight: 10 })
   assert.Equal(t, mst[6][0], Edge{ to: 5, weight: 25 })
   assert.Equal(t, mst[5][0], Edge{ to: 4, weight: 22 })
   assert.Equal(t, mst[4][0], Edge{ to: 3, weight: 12 })
   assert.Equal(t, mst[3][0], Edge{ to: 2, weight: 16 })
   assert.Equal(t, mst[2][0], Edge{ to: 7, weight: 14 })
}

func TestPrims(t *testing.T) {
   graph := getGraph()

   mst := graph.Prims()

   verifyGraph(t, mst)
}
