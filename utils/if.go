package utils

func If(isTrue bool,a interface{},b interface{}) interface{} {
	if isTrue {
		return a
	}
	return b
}
