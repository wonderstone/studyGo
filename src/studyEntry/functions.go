package main

import (
	"errors"
	"fmt"
)

//func add(x, y int) int { // When two or more consecutive named function parameters share a type, type can omit.
func posi_add(x int, y int) (res int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("Should be non-negative numbers")
		return
	}
	return x + y, nil
}

// 不定参数类型
func mul_function(args ...int) {
	args = args[1:]
	for k, v := range args {
		fmt.Println(k, v)
	}
}

func mul_interface_function(args ...interface{}) {
	args = args[1:]
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println(v, " is an int value")
		case string:
			fmt.Println(v, "is a string")
		default:
			fmt.Println(v, " is an unknown type")
		}
	}
}

func function_app() {
	res, err := posi_add(2, 3)
	if err != nil {
		fmt.Println("sth bad happened!")
	} else {
		fmt.Println(res)
	}

	tmp := []int{2, 3, 4, 5, 6}
	mul_function(tmp[1:]...)
	mul_interface_function(3, "ok", "3.14")
	// 匿名函数与闭包
	af := func(a, b int, c float64) bool {
		return a*b < int(c)
	}
	fmt.Println(af(3, 4, 54.6))
	tmpres := func(a, b int, c float64) bool {
		return a*b < int(c)
	}(3, 4, 10.10)
	fmt.Println(tmpres)
}
