package problem

import (
	// "fmt"

	// "github.com/davecgh/go-spew/spew"
)


const INF int = 10e10+1 // int bigger than max time to represent inf 
const MOD int = 1e9 + 7

func countPaths(n int, roads [][]int) int {
	//intialized adjacency matrix, distance, ways, and visited arrays 
	adj := make([][]int, n)
	dists := make([]int, n)
	ways := make([]int, n)
	vis :=   make([]bool, n)

  // populate adj graph and dists with infinity 
	for u := range adj {
    adj[u] = make([]int, n)
    for v := 0; v < len(adj); v++ {
			adj[u][v] = INF
		}
		dists[u] = INF
	}

  // populate actual road data 
  for _, road := range roads {
		adj[road[0]][road[1]] = road[2]
		adj[road[1]][road[0]] = road[2]
	}

  // set start point  
  adj[0][0] = 0
  dists[0] = 0
	ways[0] = 1

  // printAdjMatrix(adj, n)
  // spew.Dump(dists)
  // spew.Dump(ways)

  // impl Dijkstra's algorithm
  for i := 0; i < n; i++ {
    cur := -1
    // find unvisited node with shortest dist 
    for j := 0; j < n; j++ {
      // fmt.Println(j)
      if !vis[j] && (cur == -1 || dists[j] < dists[cur]) {
        cur = j
      }
    }
    // fmt.Printf("cur: %d\n", cur)
    vis[cur] = true 

    // update dists and count of ways for neighbours  
    for j := 0; j < n; j++ {
      if j == cur {
        // skip this node if its the current one
        continue 
      }

      newDist :=  dists[cur] + adj[cur][j]

      // if there's a shorter path, update dists
      if newDist < dists[j] {
        dists[j] = newDist
        ways[j] = ways[cur]
      } else if newDist == dists[j] {
        // also increment ways to count the number of same dists 
        ways[j] = (ways[j] + ways[cur]) % MOD
      }
    }

  }
  
  return ways[n - 1]
}
//
// func printAdjMatrix(adj [][]int, n int) {
//     // Print the header row
//     fmt.Print("  |")
//     for i := 0; i < n; i++ {
//         fmt.Printf(" %d ", i)
//     }
//     fmt.Println()
//     
//     // Print the separator row
//     fmt.Print("--+")
//     for i := 0; i < n; i++ {
//         fmt.Print("---")
//     }
//     fmt.Println()
//     
//     // Print each row of the matrix
//     for i := 0; i < n; i++ {
//         fmt.Printf("%d | ", i)
//         for j := 0; j < n; j++ {
//             if adj[i][j] == INF {
//                 fmt.Print(" ∞ ")
//             } else {
//                 fmt.Printf(" %d ", adj[i][j])
//             }
//         }
//         fmt.Println()
//     }
// }
