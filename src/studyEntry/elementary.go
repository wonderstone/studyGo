package main

import (
	"fmt"
	"reflect"
)

// var variableName type
// var vname1,vname2,vname3 type
var vname1 int // 全局声明与初始化分开，初始化就只能在函数应用处执行
// var vname1,vname2,vname3 type = v1,v2,v3
// var vname1,vname2,vname3 = v1,v2,v3 //If an initializer is present, the type can be omitted

var vname3, vname4 = 3, "four"

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a string
	b bool
)

const Pi = 3.14 // const identifier [type] = value  Constants can be character, string, boolean, or numeric values.
const (
	c0 = iota // const关键字出现时被重置为0，在下一个const出现前，每出现一次iota，代表的数字就自增1
	c1 = iota
	c2 = iota * 3
)

func judgeGrade(marks int) string {
	var grade string
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

	switch {
	case marks >= 10, marks <= 50:
		fmt.Printf("优秀!\n")
	case marks > 50:
		fmt.Printf("完美!\n")
	default:
		fmt.Printf("不及格!\n")
	}
	return grade
}

func sumAll(startnum, endnum int) int {
	var result int
	i := startnum
	for result < endnum { //The init and post statements are optional.
		//for i :=startnum; i<endnum; i++{
		result += i
	}
	fmt.Println(result)
	return result
}

func max(num1, num2 int) int { //func function_name( [parameter list] ) [return_types]
	/* 声明局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func map_func() string {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	/*使用键输出地图值 */
	for country, capital := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country], capital)
	}
	delete(countryCapitalMap, "France")
	for country, capital := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country], capital)
	}

	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
	return "ok"
}

func array_func() (x [5]float32, y [6]float32) {
	var balence1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0, 87.98}
	x = balence1
	y = balance2

	for i := 0; i < len(balence1); i++ {
		fmt.Println("Elementary", i, "of array is ", balence1[i])
	}

	for ii, vv := range balance2 {
		fmt.Println("Elementary", ii, "of array is ", vv)
	}
	return
}

func slice_func() ([]float32, []int) {
	slice01 := []float32{1.0, 2.2, 3.2}
	slice00 := []float32{3.2, 5.0}

	copy(slice00, slice01)
	fmt.Print(slice00)
	copy(slice01, slice00)
	fmt.Print(slice01)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	var slice02 []int = make([]int, 3, 5)

	slice02 = append(slice02, 0)
	printSlice(slice02)
	slice02 = append(slice02, numbers...)
	printSlice(slice02)

	return slice01, slice02
}

func variable_app() {
	vname1 = 1
	var vname2 int
	vname2 = 2
	// vname1,vname2,vname3 := v1,v2,v3 //只能用在函数里
	vname5, vname6 := 5, 6.0
	a, b = "a", true
	fmt.Println(reflect.TypeOf(vname1), reflect.TypeOf(vname2), reflect.TypeOf(vname3), reflect.TypeOf(vname4),
		reflect.TypeOf(vname5), reflect.TypeOf(vname6))
	fmt.Println(vname1, vname2, vname3, vname4, vname5, vname6)
	// 采用%T获得变量类型
	fmt.Printf("The types for variables %T,%T,%T,%T,%T,%T.\n", vname1, vname2, vname3, vname4, vname5, vname6)
	fmt.Printf("The variable a is \"%s\".\n", a)
	fmt.Printf("The variable b is \"%t\".\n", b)
	fmt.Printf("The constant Pi,c0,c1,c2 is \"%f, %d, %d,%d \".\n", Pi, c0, c1, c2)
	// 函数作为参数
	fmt.Println(judgeGrade(sumAll(max(30, 50), 100)))
	//
	fmt.Println(map_func())
	fmt.Println(array_func())
	fmt.Println(slice_func())

}
