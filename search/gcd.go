package search

func Gcd(p, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return Gcd(q, r)
}
