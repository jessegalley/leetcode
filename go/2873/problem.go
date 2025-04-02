package problem 


func maximumTripletValue(nums []int) int64 {
    
  var max int
  var maxDiff int 
  var ans int 

  for _,num := range nums {
    cur := maxDiff * num  
    if cur > ans {
      ans = cur
    }
    if num > max {
      max = num
    }
    curDiff := max - num 
    if curDiff > maxDiff {
      maxDiff = curDiff
    }
  }


  return int64(ans)
}

