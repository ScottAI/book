package main

import "fmt"

func fb1(n int) int {
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}else {
		return fb1(n-1) + fb1(n-2)
	}

}

func fb2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fb2(n-1) + fb2(n-2)
}

func fb3(n int) int {
	fbMap := make(map[int]int)
	for i := 0;i <= n;i++ {
		var t int
		if i <= 1 {
			t = i
		}else {
			t = fbMap[i-1] + fbMap[i-2]
		}
		fbMap[i] = t
	}
	return fbMap[n]
}

func main() {
	fmt.Println(fb1(50))
	fmt.Println(fb2(50))
	fmt.Println(fb3(50))
}
