package others

import (
	"fmt"
)

var weight int

func Test() {
	var n, k, x int
	if _, err := fmt.Scan(&n, &k, &x); err != nil {
		return
	}
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&a[i]); err != nil {
			return
		}
	}
	weight = x * a[0]
	tmp := weight
	rest := 0
	for i := 0; i < n; i++ {
		if i+1 <= n-1 {
			if a[i] > a[i+1] {
				tmp = weight
				x = tmp / a[i+1]
				if x > k {
					x = k
				}
				rest = tmp - x*a[i+1]
			} else {
				weight = x*a[i] + rest
			}
		}
	}
	if a[n-1] < a[n-2] {
		weight = x*a[n-2] + rest
	} else {
		weight = x*a[n-1] + rest
	}
	fmt.Println(weight)
}
