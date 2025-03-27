package problem


func minimumIndex(nums []int) int {
  freq := make(map[int]int)
  max := 0
  x := -1

  // find the element with maximum freq 
  for _, n := range nums {
    freq[n]++
    if freq[n] > max {
      max = freq[n]
      x = n
    } 
  }

  // iterate through slice and find the first 
  // candidate, returning it 
  cur := 0
  for i, n := range nums {
    if n == x {
      cur++
    }

    if (cur * 2 > (i + 1)) {
      if ((max - cur) * 2 > len(nums)-(i+1)) {
        return i 
      }
    }

  }

  return -1
}
