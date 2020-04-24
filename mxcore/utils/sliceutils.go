package utils

func SliceIntGetKey(s []int, index int) interface{} {
	if index <= 0 || index >= len(s) {
		return s[0]
	}
	return s[index]
}