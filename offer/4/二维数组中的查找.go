/**
  @author: crane
  @since: 2024/11/20
  @desc: 剑指 offer 第 4 题: 二维数组中的查找
**/

package _4

// find
//
//	 @Description: 二维数组中的查找,通过手动演算发现规律,从右上角开始查找。如果当前元素等于要查找的元素，则返回 true。
//		如果当前元素大于要查找的元素，则剔除当前元素所在列。如果当前元素小于要查找的元素，则剔除当前元素所在行。
//	 @param matrix [][]int
//	 @param rows int
//	 @param columns int
//	 @param target int
//	 @return bool
func find(matrix [][]int, rows, columns int, target int) bool {
	var isFind = false
	if matrix != nil && rows > 0 && columns > 0 {
		row, column := 0, columns-1
		for row < rows && column >= 0 {
			if matrix[row][column] == target {
				isFind = true
				break
			} else if matrix[row][column] > target {
				column--
			} else {
				row++
			}
		}
	}
	return isFind
}
