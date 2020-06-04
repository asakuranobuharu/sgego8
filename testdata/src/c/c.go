package c

func f() {
	var v interface{} = "gopher"
	i := v.(int) // want "required two values"
	print(i)

	j, ok := v.(int)
	print(j, ok)

	k := 1
	print(k)
}
