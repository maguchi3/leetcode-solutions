package utils

func sequence(s int, e int) []int {
	seq := make([]int, e-s+1)

	for i := range seq {
		seq[i] = s + i
	}

	return seq
}
