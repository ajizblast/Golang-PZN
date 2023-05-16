package main

import "fmt"

func main() {
	//Operator Perbandingan
	/*
		Operator ||| Keterangan
		>		|||| Lebih Dari
		<		|||| Kurang dari
		>=		|||| Lebih Dari Sama Dengan
		<=		|||| Kurang Dari Sama Dengan
		==		|||| Sama Dengan
		!=		|||| Tidak Sama Dengan
	*/
	var name1 = "Chahyo"
	var name2 = "Aji"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
