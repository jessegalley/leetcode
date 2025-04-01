package problem

var (
	numQuestions int 
  cache map[int]int64
	questionData [][]int
)

func mostPoints(questions [][]int) int64 {
 	numQuestions = len(questions)
	cache = make(map[int]int64) 
	questionData = questions 
	return dfs(0)    
}

func dfs(i int) int64 {
	if (i >= numQuestions) { 
		return 0 
	}

  cachedResult, ok := cache[i]
  if ok {
    return cachedResult 
  }

	points := questionData[i][0]
	skipped := questionData[i][1] 
	
	solvedPoints := int64(points) + dfs(i+skipped+1)
  skippedPoints := dfs(i+1)
  if solvedPoints > skippedPoints {
    cache[i] = solvedPoints  
    return cache[i]
  } else {
    cache[i] = skippedPoints 
    return cache[i]
  }
}
