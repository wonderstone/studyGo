package main

//var c, python, java bool
//var c, python, java = true, false, "no!"	//If an initializer is present, the type can be omitted;
var i, j int = 1, 2

const Pi = 3.14 // const identifier [type] = value  Constants can be character, string, boolean, or numeric values.
var (           // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

//func add(x, y int) int { // When two or more consecutive named function parameters share a type, type can omit.
func add(x int, y int) int {
	return x + y
}

func split(sum int) (x, y int) { //Named return values
	x = sum * 4 / 9
	y = sum - x
	return // A return statement without arguments returns the named return values.
}
