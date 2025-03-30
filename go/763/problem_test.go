package problem

import (
	"reflect"
	"testing"
)

/*
763. Partition Labels
Medium
Topics
Companies
Hint

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part. For example, the string "ababcc" can be partitioned into ["abab", "cc"], but partitions such as ["aba", "bcc"] or ["ab", "ab", "cc"] are invalid.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    s string 
    want []int
  } {
    {
      "test1",
      "ababcbacadefegdehijhklij",
      []int{9, 7, 8},
    },
    {
      "test2",
      "eccbbbbdec",
      []int{10},
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := partitionLabels(tc.s)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("partitionLabels(%v) = %v; want = %v\n",
        tc.s, got, tc.want)
      }
    })
  }
}
