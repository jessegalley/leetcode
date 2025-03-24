package problem

import (
	"reflect"
	"testing"
)

/*
3169. Count Days Without Meetings
Medium
Topics
Companies
Hint

You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1). You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

Return the count of days when the employee is available for work but no meetings are scheduled.

Note: The meetings may overlap.
*/

func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    days int
    meetings [][]int
    want int
  } {
    {
      "test1",
      10,
      [][]int{{5,7},{1,3},{9,10}},
      2,
    },
    {
      "test2",
      5,
      [][]int{{2,4},{1,3}},
      1,
    },
    {
      "test3",
      6,
      [][]int{{1,6}},
      0,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := countDays(tc.days, tc.meetings)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("countDays(%v, %v) = %v; want = %v\n",
          tc.days, tc.meetings, got, tc.want)
      }
    })
  }
}
