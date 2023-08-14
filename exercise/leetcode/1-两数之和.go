package leetcode

func twoSum(nums []int, target int) []int {
	indexMap := map[int]int{}
	for i, num := range nums {
		if index, ok := indexMap[target-num]; ok {
			return []int{index, i}
		}
		indexMap[num] = i
	}
	return nil
}
