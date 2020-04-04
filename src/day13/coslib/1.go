package count

//method 1
func Count(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

//method 2
func Counts(s string) int {
	n := 0
	for _, _ = range s {
		n++
	}
	return n
}
