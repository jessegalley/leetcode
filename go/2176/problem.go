package problem

// import "fmt"

func countPairs(nums []int, k int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				if (i*j)%k == 0 {
          // fmt.Printf("i:%d; j:%d\n", i, j)
					ans++
				}
			}
		}
	}
	return ans
}
