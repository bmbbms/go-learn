package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 100
	if num < 10 {
		fmt.Println("num < 10")
	} else if num > 111 {
		fmt.Println("num > 111")
	} else {
		fmt.Println("10 < num < 111")
	}

	var oldString string = "777"
	anInt, err1 := strconv.Atoi(oldString)
	if err1 != nil {
		fmt.Println("convert error")
	}
	fmt.Println(anInt)

	if anInt1, err2 := strconv.Atoi(oldString); err2 != nil {
		fmt.Println("error convert")
	} else {
		fmt.Println(anInt1)
	}

	var i int = 100
	switch i {
	case 10:
		fmt.Printf("i=%d", i)
	case 100:
		fmt.Printf("i=%d", i)

	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("g")
		}
		fmt.Println()
	}

	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}
	}

}
