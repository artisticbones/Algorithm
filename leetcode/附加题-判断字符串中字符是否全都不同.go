/**
  @author: crane
  @since: 2024/12/10
  @desc: 判断字符串中的字符是否全都不同
  问题描述:
		 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
		 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
		 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
**/

package leetcode

// IsUniqueString
//
//	@Description: 判断字符串中的字符是否全都不同
//	@param s string
//	@return bool
func IsUniqueString(s string) bool {
	// 字符串长度为0或者大于3000，则返回false
	if len(s) == 0 || len(s) > 3000 {
		return false
	}

	// 用一个uint64的变量来标记每个字符是否出现过
	var mark1, mark2, mark3, mark4 uint64
	var mark *uint64

	for _, r := range s {
		n := uint64(r)
		// 根据字符的ASCII码值来判断字符属于哪个uint64变量
		if n < 64 {
			mark = &mark1
		} else if n < 128 {
			mark = &mark2
			n -= 64
		} else if n < 192 {
			mark = &mark3
			n -= 128
		} else {
			mark = &mark4
			n -= 192
		}
		// 如果字符已经出现过，则返回false
		if *mark&(1<<n) != 0 {
			return false
		}
		// 标记字符出现过
		*mark = *mark | (1 << n)
	}
	return true
}
