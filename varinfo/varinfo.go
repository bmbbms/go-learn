package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var a, b *int
	var (
		c = 100

		e = 100
	)
	a = &c
	b = &e
	fmt.Println(&a,&b)


	var (
    HOME = os.Getenv("HOME")
    USER = os.Getenv("USER")
    GOROOT = os.Getenv("GOROOT")
)
	fmt.Println(HOME,USER,GOROOT)


	var goes string = runtime.GOOS

	fmt.Println("the operator system is : ", goes)
}
