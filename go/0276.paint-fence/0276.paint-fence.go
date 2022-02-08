package leetcode

/* 1. Tabulation */
func numWays(n, k int) int {
	// base cases
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	// create table
	ways := make([]int, n)
	ways[0], ways[1] = k, k*k

	for i := 2; i < n; i++ {
		/*
		  Total(n-1) * (k-1) ways of use different color than a prev post
		  Total(n-2) * (k-2) ways of use same color as a prev post
		*/
		ways[i] = (k - 1) * (ways[i-1] + ways[i-2])
	}

	return ways[n-1]
}

/* 2. Recursion with memorization */
func numWays2(n, k int) int {
	memo := make(map[int]int)

	return total(n, k, memo)
}

func total(n, k int, memo map[int]int) int {
	// base cases
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	// check if we've already calculated nth case
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	/*
	  Total(n-1) * (k-1) ways of use different color than a prev post
	  Total(n-2) * (k-2) ways of use same color as a prev post
	*/
	res := (k - 1) * (total(n-1, k, memo) + total(n-2, k, memo))

	memo[n] = res

	return res
}
