package problem

import (
	"reflect"
	"testing"
)

/*
3394. Check if Grid can be Cut into Sections
Medium
Topics
Companies
Hint

You are given an integer n representing the dimensions of an n x n grid, with the origin at the bottom-left corner of the grid. You are also given a 2D array of coordinates rectangles, where rectangles[i] is in the form [startx, starty, endx, endy], representing a rectangle on the grid. Each rectangle is defined as follows:

    (startx, starty): The bottom-left corner of the rectangle.
    (endx, endy): The top-right corner of the rectangle.

Note that the rectangles do not overlap. Your task is to determine if it is possible to make either two horizontal or two vertical cuts on the grid such that:

    Each of the three resulting sections formed by the cuts contains at least one rectangle.
    Every rectangle belongs to exactly one section.

Return true if such cuts can be made; otherwise, return false.
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    dimension int 
    coords [][]int
    want bool
  } {
    {
      "test1",
      5,
      [][]int{{1,0,5,2},{0,2,2,4},{3,2,5,3},{0,4,4,5}},
      true,
    },
    {
      "test2",
      4,
      [][]int{{0,0,1,1},{2,0,3,4},{0,2,2,3},{3,0,4,3}},
      true,
    },
    {
      "test3",
      4,
      [][]int{{0,2,2,4},{1,0,3,2},{2,2,3,4},{3,0,4,2},{3,2,4,4}},
      false,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := checkValidCuts(tc.dimension, tc.coords)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("findAlldimension(%v, %v) = %v; want = %v\n",
          tc.dimension, tc.coords, got, tc.want)
      }
    })
  }
}
