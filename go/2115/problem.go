package problem 

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
  
  answer := []string{}
  deg := make(map[string]int)
  graph := make(map[string][]string)

  for i,recipe := range recipes {
    // construct dependancy graph
    for _,ingredient := range ingredients[i] {
      graph[ingredient] = append(graph[ingredient], recipe)
    } 
    // construct in degrees 
    deg[recipe] = len(ingredients[i])
  }

  // make dequeue of ingredients we have 
  q := supplies
  // process until no more supplies available 
  for ; len(q) >0 ; {
    var x string  
    // shift first supply item from queue
    x, q = q[0], q[1:]
    // find recipies that use this supply item from dep grah
    for _,r := range graph[x] {
      // decrement indegrees to satisfy dep 
      deg[r]--
      // if this recipe has all deps satisfied, include it in 
      // answer and also add it to the queue in case it's a 
      // dep for something else 
      if deg[r] <= 0 {
        answer = append(answer, r)
        q = append(q, r)
      }
    } 
  }
  
  return answer
}

