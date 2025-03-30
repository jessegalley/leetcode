package problem

func partitionLabels(s string) []int {
  ans := make([]int, 0)
  lastIdx := make(map[rune]int)

  for i, c := range s {
    lastIdx[c] = i
  }

  start := 0 
  max := 0
  for i, c := range s {
    if lastIdx[c] > max {
      max = lastIdx[c]
    } 
    if i == max {
      ans = append(ans, len(s[start:i+1]))
      start = i+1
    }
  }

  return ans
}
