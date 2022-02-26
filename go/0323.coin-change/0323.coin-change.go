package coinChange

import "math"

// Dynamic Programming (Bottom up)
func coinChange(coins []int, amount int) int {
	max := amount + 1
	counts := make([]int, amount+1)
	for i := range counts {
		counts[i] = max
	}
	counts[0] = 0

	for cur := 1; cur <= amount; cur++ {
		for _, coin := range coins {
			if coin > cur {
				continue
			}
			counts[cur] = min(counts[cur], counts[cur-coin]+1)
		}
	}

	if counts[amount] == amount+1 {
		return -1
	} else {
		return counts[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Dynamic Programming (Top down)
func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	counts := make([]int, amount)

	return helper(coins, amount, counts)
}

func helper(coins []int, remain int, counts []int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}

	if counts[remain-1] != 0 {
		return counts[remain-1]
	}

	min := math.MaxInt32
	for _, coin := range coins {
		res := helper(coins, remain-coin, counts)
		if res >= 0 && res < min {
			min = res + 1
		}
	}
	if min == math.MaxInt32 {
		counts[remain-1] = -1
	} else {
		counts[remain-1] = min
	}
	return counts[remain-1]
}
