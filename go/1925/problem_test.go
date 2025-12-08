package problem

import (
	"reflect"
	"testing"
)

/*
 */

func TestProblem(t *testing.T) {

	testCases := []struct {
		name  string
		input int
		want  int
	}{
		{
			"test1",
			5,
			2,
		},
		{
			"test2",
			10,
			4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solve(tc.input)
			ok := reflect.DeepEqual(got, tc.want)
			if !ok {
				t.Errorf("solve(%v) = %v; want = %v\n",
					tc.input, got, tc.want)
			}
		})
	}
}
