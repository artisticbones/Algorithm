package leetcode

func generateMatrix(n int) [][]int {
	left, right, up, bottom := 0, n-1, 0, n-1
	num, target := 1, n*n

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for num <= target {
		for i := left; i <= right; i++ {
			matrix[up][i] = num
			num++
		} // 从左到右
		up++
		for i := up; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		} //从上到下
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		} // 从右到左
		bottom--
		for i := bottom; i >= up; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
