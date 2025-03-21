package problem

func maximumCandies(candies []int, k int64) int {
	// im not sure if there's something more elegant than
	// simply brute forcing this, so we'll try to make it
	// a bit quicker by using a binary search

	// left will be the minimum number of candies each child
	// could get, which will aways be 0
	var left int = 0
	// right is the maximum amount that would be possible, which
	// is always the biggest number in the set
	var right int = sliceMax(candies)

	// binary search
	for left < right {
		// mid is the number of candies to dsitribute per child
		mid := (left + right + 1) >> 1

		// count how many kids can get mid candies
		var totalPortions int64 = 0
		for _, pile := range candies {
			totalPortions += int64(pile / mid)
		}

		// if the number of portions available is at least k
		// mid can be a porential soluton, so we move left
		// up to mid, otherwise, decrease right pointer
		if totalPortions >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

// go doesn't have a built in max() that works with slices
func sliceMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
