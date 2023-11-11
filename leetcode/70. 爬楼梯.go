package leetcode

//func climbStairs(n int) int {
//	p, q, ans := 0, 0, 1
//	for i := 1; i <= n; i++ {
//		p = q
//		q = ans
//		ans = p + q
//	}
//	return ans
//}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
