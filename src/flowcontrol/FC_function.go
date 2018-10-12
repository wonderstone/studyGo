package flowcontrol

import "fmt"

func JudgeGrade(marks int) string {
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

func SumAll(startnum, endnum int) int {
	var result int
	i := startnum
	for result < endnum { //The init and post statements are optional.
		//for i :=startnum; i<endnum; i++{
		result += i
	}
	fmt.Println(result)
	return result
}

func Max1(num1, num2 int) int { //func function_name( [parameter list] ) [return_types]

	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
