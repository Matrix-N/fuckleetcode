package fuckleetcode

import "strconv"

func addBinary(a, b string) string {
	result := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}

	if carry > 0 {
		result = "1" + result
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
