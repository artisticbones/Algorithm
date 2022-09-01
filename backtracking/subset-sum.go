//Given a set of non-negative integers, and a (positive) value sum,
//determine if there is a subset of the given set with sum
//equal to given sum.
//Complexity: O(n*sum)
//references: https://www.geeksforgeeks.org/subset-sum-problem-dp-25/

package backtracking

import (
	"errors"
)

var ErrInvalidPosition = errors.New("invalid position in subset")
var ErrNegativeSum = errors.New("negative sum is not allowed")

func IsSubsetSum(array []int, sum int) (bool, error) {
	// sum can not be negative number
	if sum < 0 {
		return false, ErrNegativeSum
	}
	// create subset matrix
	arraySize := len(array)
	subset := make([][]bool, arraySize+1)
	for i := 0; i <= arraySize; i++ {
		subset[i] = make([]bool, sum+1)
	}

	// sum 0 is always true
	for i := 0; i <= arraySize; i++ {
		subset[i][0] = true
	}

	for i := 1; i <= sum; i++ {
		// empty set is false when sum is not 0
		subset[0][i] = false
	}

	for i := 1; i <= arraySize; i++ {
		for j := 1; j <= sum; j++ {
			if array[i-1] > j {
				subset[i][j] = subset[i-1][j]
			}

			if array[i-1] <= j {
				if j-array[i-1] < 0 || j-array[i-1] > sum {
					//out of bounds
					return false, ErrInvalidPosition
				}

				subset[i][j] = subset[i-1][j] || subset[i-1][j-array[i-1]]
			}
		}
	}
	return subset[arraySize][sum], nil
}
