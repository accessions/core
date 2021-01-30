package utils

func Shift (i []interface{}) (interface{}, []interface{}) {
	 x, i := i[0], i[1:]
	 return x, i
}

func Pop(i []interface{}) (interface{}, []interface{}) {
	x, i := i[len(i)-1], i[:len(i)-1]
	return x, i
}

func Push(i []interface{}, n interface{}) []interface{} {
	return append(i, n)
}

func Delete(i []interface{}, n int32) []interface{} {
	return append(i[:n], i[n+1:])
}

func Append(a []interface{}, b[]interface{}) []interface{} {
	return append(a, b...)
}
