package utils

func Sequence(start int, end int) []int {
	seq := make([]int, end-start+1)

	for i := range seq {
		seq[i] = start + i
	}

	return seq
}
