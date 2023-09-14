/*
 * @lc app=leetcode.cn id=1598 lang=golang
 *
 * [1598] 文件夹操作日志搜集器
 */
package leetcode

// @lc code=start
func minOperations(logs []string) int {
	var step int
	for _, dir := range logs {
		switch dir {
		case "./":
		case "../":
			if step == 0 {
				continue
			}
			step--
		default:
			step++
		}
	}
	return step
}

// @lc code=end
