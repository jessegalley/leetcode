package problem

import (
	// "fmt"
	"math"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// makes a linked list from an int slice and returns the head node
func makeLinkedList(nums []int) *ListNode { 

  _ = spew.Config // debug remove alter

  var head, curr  *ListNode

  head = &ListNode{} 
  curr = head 

  l :=  len(nums)
  for i, v := range nums {
    curr.Val = v
    if i == (l-1) {
      curr.Next = nil
    } else {
      curr.Next = &ListNode{}
    }
    curr = curr.Next
  } 

  return head 
}

func listToNums(l *ListNode) []int {
  var nums []int 
  for l != nil {
    nums = append(nums, l.Val)
    l = l.Next
  }
  
  return  nums
}

// converts a linked list back to its integer representation 
func listToInt(l *ListNode) int {
  num := 0
  i := 0
  for l != nil {
    num += l.Val * int(math.Pow(10, float64(i)))
    // fmt.Printf("l: %v; l.Val: %d; num: %d\n", l, l.Val, num)
    l = l.Next
    i++
  }
  
  return num 
}

func TestProblem(t *testing.T) {

  // tl := makeLinkedList([]int{1,2,3})
  // spew.Dump(tl)
  // num1 := listToInt(tl)
  // spew.Dump(num1)
  // s1 := listToNums(tl)
  // spew.Dump(s1)

  testCases := []struct {
    name string 
    l1 *ListNode
    l2 *ListNode
    want *ListNode
  } {
    {
      "test1",
      makeLinkedList([]int{2,4,3}),
      makeLinkedList([]int{5,6,4}),
      makeLinkedList([]int{7,0,8}),
    },
    {
      "test2",
      makeLinkedList([]int{0}),
      makeLinkedList([]int{0}),
      makeLinkedList([]int{0}),
    },
    {
      "test3",
      makeLinkedList([]int{9,9,9,9,9,9,9}),
      makeLinkedList([]int{9,9,9,9}),
      makeLinkedList([]int{8,9,9,9,0,0,0,1}),
    },

  }

  // spew.Dump(listToInt(testCases[1].l1))
  // spew.Dump(listToInt(testCases[1].l2))

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      // fmt.Println("L1: ")
      // spew.Dump(tc.l1)
      // fmt.Println("L2: ")
      // spew.Dump(tc.l2)
      got := solve(tc.l1, tc.l2)
      // fmt.Println("got:")
      // spew.Dump(got)
      gotInt := listToInt(got)
      wantInt := listToInt(tc.want)
      ok := reflect.DeepEqual(gotInt, wantInt)
      if !ok {
        t.Errorf("solve(%v, %v) = %v; want = %v\n",
        tc.l1, tc.l2, gotInt, wantInt)
      }
    })
  }


}
