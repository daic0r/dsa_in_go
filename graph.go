package mst

import (
   "container/heap"
)

type Edge struct {
   to int
   weight int
}

type CompleteEdge struct {
   from int
   edge Edge
}

type Graph map[int][]Edge

func (g *Graph) Prims() Graph {
   ret := Graph{}
   pq := EdgeHeap{}
   heap.Init(&pq)
   graph := *g
   seen := make(map[int]bool)
   // Pick random edge
   var cur_node int
   for from := range graph {
      cur_node = from
      break
   }
   cur_node = 1

   first := true
   for first || pq.Len() > 0 {
      seen[cur_node] = true
      for _, edge := range graph[cur_node] {
         if !seen[edge.to] {
            heap.Push(&pq, CompleteEdge{ cur_node, edge })
         }
      }
      top := heap.Pop(&pq).(CompleteEdge)
      if seen[top.edge.to] {
         continue
      }
      ret[top.from] = append(ret[top.from], top.edge)
      cur_node = top.edge.to
      first = false
   }

   // for from, edges := range ret {
   //    for _, edge := range edges {
   //       fmt.Printf("(%v -> %v, %v)\n", from, edge.to, edge.weight)
   //    }
   // }

   return ret
}
