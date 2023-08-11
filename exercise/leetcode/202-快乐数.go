package leetcode

func splitNumAndSumOfSquares(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

func isHappy(n int) bool {
	sumOfSquares := make(map[int]struct{}, 1)
	for {
		t := splitNumAndSumOfSquares(n)
		if t == 1 {
			return true
		}
		if _, ok := sumOfSquares[t]; ok {
			return false
		} else {
			sumOfSquares[t] = struct{}{}
			n = t
		}
	}
}
