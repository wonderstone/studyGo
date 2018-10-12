package collection

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Map_func() string {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
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

func Array_func(size int) (x [5]float32, y [5]float32) {
	var balence1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	x = balence1
	y = balance2

	return
}

func Slice_func() ([]float32, []int) {
	slice01 := []float32{1.0, 2.2, 3.2}
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	var slice02 []int = make([]int, 3, 5)

	slice02 = append(slice02, 0)
	printSlice(slice02)

	return slice01, slice02
}
