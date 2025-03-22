package problem

import (
	"reflect"
	"testing"
)

/*
You are given an integer n. There is an undirected graph with n vertices, numbered from 0 to n - 1. You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting vertices ai and bi.

Return the number of complete connected components of the graph.

A connected component is a subgraph of a graph in which there exists a path between any two vertices, and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

A connected component is said to be complete if there exists an edge between every pair of its vertices. 
*/  


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nodes int 
    edges [][]int 
    want int  
  } {
    {
      "test1",
      6,
      [][]int{{0,1},{0,2},{1,2},{3,4}},
      3,
    },
    {
      "test2",
      6,
      [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}},
      1,
    },
    {
      "test3",
      5,
      [][]int{{0,1},{1,2},{3,4}},
      1,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := countCompleteComponents(tc.nodes, tc.edges)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("countCompleteComponents(%v, %v) = %v; want = %v\n",
          tc.nodes, tc.edges, got, tc.want)
      }
    })
  }
}
