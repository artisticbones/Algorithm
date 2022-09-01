package leetcode

func finalPrices(prices []int) []int {
	ret := make([]int, len(prices))
	discount := make(map[int]int, len(prices))
	for index := range discount {
		discount[index] = 0
	}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount[i] = prices[j]
				break
			}
		}
	}
	for index, v := range prices {
		ret[index] = v - discount[index]
	}
	return ret
}
