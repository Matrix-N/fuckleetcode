package fuckleetcode

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := range result {
		result[i] = onesCount(i)
	}
	return result
}

func onesCount(x int) (result int) {
	for ; x > 0; x &= (x - 1) {
		result++
	}

	return result
}
