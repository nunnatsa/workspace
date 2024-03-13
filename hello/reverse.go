package hello

func Reverse(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-1-i] = rs[len(rs)-1-i], rs[i]
	}

	return string(rs)
}
