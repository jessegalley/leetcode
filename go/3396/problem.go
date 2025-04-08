package problem

import "fmt"

func minimumOperations(nums []int) int {

	// hash map to remember duplicates
	h := make(map[int]bool)

	//traverse the array backwards
	for i := len(nums) - 1; i >= 0; i-- {
    fmt.Printf("i:%d;", nums[i])
		_, ok := h[nums[i]]
		if !ok {
      fmt.Printf("unfound:%d\n",nums[i])
			h[nums[i]] = true
			continue
		}

    fmt.Printf("found:%d\n",nums[i])
		return (i / 3) + 1
	}

	return 0
}
