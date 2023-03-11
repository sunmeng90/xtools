package stringer

import "strconv"

// Length return the length of a string
func Length(val string, digits bool) (count int, kind string) {
	if !digits {
		return len(val), "char"
	}
	cnt := 0
	for _, v := range val {
		_, err := strconv.Atoi(string(v))
		if err == nil {
			cnt += 1
		}
	}
	return cnt, "digit"
}
