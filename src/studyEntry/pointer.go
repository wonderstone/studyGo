package main

import "fmt"

const MAX int = 3

//指针作为函数参数
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func swap_app() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func ptr() {
	// 00 常规用法 ***************************************************************//
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 01 空指针  ***************************************************************//
	var ptr *int
	fmt.Printf("空指针 ptr 的值为 : %x\n", ptr) //一个指针被定义后没有分配到任何变量时，它的值为 nil

	// 02 指针数组对比  ***************************************************************//
	aa := []int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("aa[%d] = %d\n", i, aa[i])
	}

	aaa := []int{10, 100, 200}
	var i01 int
	var ptr01 [MAX]*int

	for i01 = 0; i01 < MAX; i01++ {
		ptr01[i01] = &aaa[i01] /* 整数地址赋值给指针数组 */
		fmt.Printf("aaa[%d] = %d\n", i01, *ptr01[i01])
	}

	// 03 指向指针的指针  ***************************************************************//
	var a02 int
	var ptr02 *int
	var pptr **int //指向指针的指针变量声明格式

	a02 = 3000

	/* 指针 ptr 地址 */
	ptr02 = &a02

	/* 指向指针 ptr 地址 */
	pptr = &ptr02

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a02 = %d\n", a02)
	fmt.Printf("指针变量 *ptr02 = %d\n", *ptr02)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) //访问指向指针的指针变量值需要使用两个 * 号

	// 04 指针作为函数参数  ***************************************************************//
	swap_app()
}
