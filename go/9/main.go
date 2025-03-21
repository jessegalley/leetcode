package main

import "fmt"


func isPalindrome(x int) bool {
  // x is negative, negatives cannot be palindromes
  if x < 0 {
    return false
  } 

  // x is between 0 and 9, a single digit number is always a palindrome
  if x >= 0 && x <= 9 {
    return true
  }

  t := x 
  k := 0
  for ; t > 0 ; {
  
    // shift k to the left 
    k = k*10

    // get the leas significant digit of t, and add it to k
    k += t%10

    //shift t to the right, truncating the digit we just added 
    t = t/10 
  }

  return k == x 
}

func main() {

  var testcases = []int{
    -1231,
    12321,
    14434,
    3,
    11,
    -454,
    000,
    0,
  }

  for i, v := range testcases {
    fmt.Printf("%d: %d result:%v\n", i, v, isPalindrome(v))
  }


}
