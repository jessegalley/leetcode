package problem

// import "fmt"


func canPartition(nums []int) bool {
  
  // if the sum of nums is odd then it's impossible
  // that we can have have a split
  var sum int 
  for _, num := range nums {
    sum += num 
  }
  if sum % 2 != 0 {
    return false
  }

  // the sum is even so the target of the split will be 
  // the half of the sum 
  q := sum / 2

  // initalize a bool slice of possible subset sums 
  f := make([]bool, q+1)
  f[0] = true 

  // iterate over every element to check if subset sums 
  // are possible
  for _, num := range nums {
    for j := q; j >= num; j-- {
      f[j] = f[j] || f[j - num]
    }
    // fmt.Printf("num: %d; f: %v\n", num, f)
  }

  return f[q]
}


