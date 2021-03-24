package main

import (
	"fmt"
	"math/rand"
)

type options struct {
	name string
	pwd  string
	uid  int
}

func main() {

	n := 0
	reply := &n
	multlynum(6, 5, reply)
	fmt.Println(*reply)
	fmt.Println(n)
	fmt.Println(add(1, 2, 3, 4, 5))
	slices := []int{1, 2, 5, 6, 7}
	fmt.Println(add(slices...))

	var op_new options
	userinfo("cart", &op_new)
	fmt.Println(op_new)
	fmt.Printf("%#v", op_new)
	fmt.Printf("%+v", op_new)

	typecheck("1", 12312, "80000")

}

func multlynum(i int, i2 int, reply *int) {
	*reply = i * i2

}

func add(slice ...int) (result int) {
	result = 0
	if len(slice) > 0 {
		for _, num := range slice {
			result = result + num
		}

	}
	return result
}

func userinfo(card string, op *options) {
	if card != "" {
		op.name = card
		op.pwd = card
		op.uid = rand.Int()

	}
}

func typecheck(values ...interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println(value)
		case float64:
			fmt.Println(value)
		case string:
			fmt.Println(value)
		case bool:
			fmt.Println(value)
		default:
			fmt.Println(value)
		}
	}
}
