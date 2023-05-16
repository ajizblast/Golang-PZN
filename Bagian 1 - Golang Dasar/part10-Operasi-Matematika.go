package main

import "fmt"

func main() {

	var x = 5 * 5
	fmt.Println(x)

	var a = 20
	var b = 10
	var c = a + b
	var d = a - b

	fmt.Println(c)
	fmt.Println(d)

	//part 2 augmented assignments
	/*
			Operasi Matematika ||| Oprasi Assignment

		a = a + 10				|| a += 10
		a = a - 10				|| a -= 10
		a = a * 10				|| a *= 10
		a = a / 10				|| a /= 10
		a = a % 10				|| a %= 10
	*/

	var i = 15
	i += 10 // i = i + 10
	fmt.Println(i)

	//	Unary Operator
	/*
		Operator ||| Keterangan
		++		|||| a = a + 1
		--		|||| a = a - 1
		-		|||| Negative
		+		|||| Positive
		!		|||| Boolean Kebalikan
	*/

	i++
	fmt.Println(i) // i = i + 1

	var negative = -100
	var positive = +100

	fmt.Println(negative)
	fmt.Println(positive)
}
