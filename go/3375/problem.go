package problem 

func minOperations(nums []int, k int) int {

  // init a var to keep track of the smallest  
  // number in the array 
  var min int = nums[0] 

  // keep a map of unique nums 
  set := make(map[int]int)

  for _, num := range nums {
    set[num]++
    if num < min {
      min = num 
    }
  }

  // if any element of the array is less than k 
  // that means it's impossible to make it equal
  // to k, so we return -1
  if k > min {
    return -1
  } 

  if k == min {
    return len(set)-1
  } 

  return len(set)
}


