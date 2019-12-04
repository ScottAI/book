package testTest

func fb1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fb1(n-1) + fb1(n-2)
}

func fb2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	return fb2(n-1) + fb2(n-2)
}
