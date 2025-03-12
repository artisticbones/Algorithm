package leetcode

import "strconv"

var div = func(a, b int) int { return a / b }

var add = func(a, b int) int { return a + b }

var mul = func(a, b int) int { return a * b }

var sub = func(a, b int) int { return a - b }

var operation = map[string]func(a, b int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func evalRPN(tokens []string) int {
	var stack []int

	for i := 0; i < len(tokens); i++ {
		if op, exists := operation[tokens[i]]; !exists {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, op(a, b))
		}
	}
	return stack[0]
}
