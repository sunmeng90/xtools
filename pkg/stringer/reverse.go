package stringer

// Reverse a string
func Reverse(val string) string {
	if len(val) <= 1 {
		return val
	}
	rns := []rune(val)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
