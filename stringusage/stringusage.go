package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var newstring string = " a great country"
	if strings.HasPrefix(strings.TrimSpace(newstring), "a") {
		fmt.Println(newstring, "start with a")
	}

	if strings.HasSuffix(strings.TrimSpace(newstring), "try") {
		fmt.Println(newstring, "end with try")
	}
	fmt.Println(strings.Contains(newstring, "a"))
	fmt.Println(strings.Index(newstring, "a"))
	fmt.Println(strings.LastIndex(newstring, "a"))
	fmt.Println(strings.Replace(newstring, "a", "e", -1))
	var stringList = strings.Fields(newstring)
	fmt.Println(stringList)
	fmt.Println(strings.Join(stringList, "-"))
	fmt.Println(strings.ToLower(newstring))
	fmt.Println(strings.ToUpper(newstring))
	var intString = "777.77777"
	f, _ := strconv.ParseFloat(intString, 64)
	fmt.Println(strconv.FormatFloat(f, 'b', 1, 64))
	t := time.Now()
	fmt.Println(t)

}
