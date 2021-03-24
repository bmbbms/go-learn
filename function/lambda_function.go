package main

import (
	"fmt"
	"time"
)

const LIM = 141000000

var fibs [LIM]uint64

func main() {
	var g int
	var endnum = 40
	fmt.Println(g)
	start := time.Now()
	for m := 1; m < endnum; m++ {
		fibonacci(m)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	for m := 1; m < endnum; m++ {
		fmt.Println(m, fib2(m))
	}
	end = time.Now()
	fmt.Println(end.Sub(start))

}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return

}

func fib2(n int) (res uint64) {

	if n <= 1 {
		res = 1
		return
	} else {
		res = fib2(n-1) + fib2(n-2)
	}

	return
}
