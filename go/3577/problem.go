package problem

func solve(complexity []int) int {
	const MOD int = 1000000007
	var res int64 = 1
	rootComplexity := complexity[0]

	for i, v := range complexity {
		if i == 0 {
			continue
		}

		if v <= rootComplexity {
			return 0
		}

		res = (res * int64(i)) % int64(MOD)
	}

	return int(res)
}
