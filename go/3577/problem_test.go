package problem

import (
	"reflect"
	"testing"
)

/*
 */

func TestProblem(t *testing.T) {

	testCases := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			"test1",
			[]int{1, 2, 3},
			2,
		},
		{
			"test2",
			[]int{3, 3, 3, 4, 4, 4},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solve(tc.complexity)
			ok := reflect.DeepEqual(got, tc.want)
			if !ok {
				t.Errorf("solve(%v) = %v; want = %v\n",
					tc.complexity, got, tc.want)
			}
		})
	}
}
