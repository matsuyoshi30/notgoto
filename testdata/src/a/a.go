package a

func f() {
	i := 0
L:
	i++
	if i < 5 {
		goto L // want `goto statement found`
	}
	println(i)
}

func g() {
	println("goto")
}
