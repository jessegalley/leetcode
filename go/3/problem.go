package problem


func solve(s string) int {
  // make a hash table to store frequency of chars 
	freq := make(map[rune]int)
  maxSize := 0

  l := 0
  // set up right pointer as index of loop
	for r, c := range s {
    // range returns runes when iterating over slice  
		freq[c]++

    // if duplicate rune, then loop and advance left pointer until its not dup anymore 
    for freq[c] > 1 {
      // accessing a slice from string directly returns byte, not rune 
      freq[rune(s[l])]--
      l++
    }

    maxSize = max(r - l + 1, maxSize)
	}

	return maxSize 
}
