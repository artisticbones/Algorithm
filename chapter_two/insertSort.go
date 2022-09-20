package chapter_two

func InsertSort(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 1 {
		return a
	}
	for i := 1; i < len(a); i++ {
		j := i - 1
		for j >= 0 && a[j] > a[i] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = a[i]
	}
	return a
}
