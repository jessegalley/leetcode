package problem

func solve(n int) int {
	// Given an integer n, return the number of square triples such that 1 <= a, b, c <= n
	pythCnt := 0

	// we need to brute force this but instead of _pure_ brute force, we can make it a bit
	// more efficient by only iterating over a+b where a and b < c.
	for c := 3; c <= n; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if (a*a)+(b*b) == (c * c) {
					// any time we have a match, iterate by two because (a,b) and (b,a) are the same
					pythCnt = pythCnt + 2
				}
			}
		}
	}

	return pythCnt
}
