package main

import "fmt"

func trace(func_name string) string {
	fmt.Println("entering:", func_name)
	return func_name

}

func un(func_name string) {
	fmt.Println("leave", func_name)

}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")

}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()

}

func main() {
	b()

	func() {
		for i := 1; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	g := func(a int) {
		fmt.Println("zheshi niming", a)
	}
	g(13)

	func(u string) {
		fmt.Println(u)

	}("hello")

	add2 := add1()
	fmt.Println(add2(4))
	add3 := Adder()
	fmt.Println(add3(1))
	fmt.Println(add3(20))
	fmt.Println(add3(200))
}

func add1() func(a int) int {
	return func(a int) int {
		return a
	}
}

func Adder() func(delta int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
