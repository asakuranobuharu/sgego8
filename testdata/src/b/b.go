package b

func f() {
	print(gopher())  // want "call gopher function"
}

func gopher() string {  // want "function is gopher"
	return "hello gopher"
}