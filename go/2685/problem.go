package problem


var (
  adj [][]int 
  vis []bool
)

func countCompleteComponents(n int, edges [][]int) int {
  adj = make([][]int, n)
  vis = make([]bool, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

  ans := 0
  // traverse each node
  for i := 0; i < n; i++ {
    // if the node is visited, skip
    if vis[i] {
      continue
    }
    // call dfs
    result := dfs(i)
    if result[0] * (result[0] - 1) == result[1]{
      ans++
    }
  }

	return ans
}

func dfs (v int) []int  {
  vis[v] = true 
  nCnt := 1
  eCnt := len(adj[v])
  for _, neighbour := range adj[v] {
    if vis[neighbour]{
      continue
    }
    result := dfs(neighbour)
    nCnt += result[0]
    eCnt += result[1]
  }
  return []int{nCnt, eCnt}
}

