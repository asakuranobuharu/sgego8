package a

func f() {
	print(gopher())
}

func gopher() string {  // want "function is gopher"
	return "hello gopher"
}