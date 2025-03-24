package problem

import (
	"cmp"
	"slices"
)

func countDays(days int, meetings [][]int) int {

  // sort the meetings first, golang provides the useful 
  // slices.SortFunc and cmp.Compare that work toegether
  slices.SortFunc(meetings, func(a []int, b []int) int {
    return cmp.Compare(a[0], b[0])
  })

  var freeDays int = 0
  var lastDay int = 0

  for _, m := range meetings {
    if lastDay < m[0] {
      freeDays += m[0] - lastDay - 1
    }
    if m[1] > lastDay {
      lastDay = m[1]
    }
  }

  if lastDay < days {
    freeDays += days - lastDay
  }

  return freeDays
}
