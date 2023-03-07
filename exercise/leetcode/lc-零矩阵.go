package leetcode

/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零
*/

func setZeroes(matrix [][]int) {
	flag_col0 := false
	for i, row := range matrix {
		//检查每行第0列的每个元素是否有0
		if matrix[i][0] == 0 {
			flag_col0 = true
		}
		for j := 1; j < len(row); j++ {
			if row[j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if flag_col0 {
			matrix[i][0] = 0
		}
	}
}
