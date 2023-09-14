package leetcode

/*
给你一幅由N*N矩阵表示的图像,其中每个像素的大小为4字节.请你设计一种算法,将图像旋转 90 度。
不占用额外内存空间能否做到？
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
