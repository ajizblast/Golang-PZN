package main

import "fmt"

func main() {

	//part 1 membuat variabel dengan tipe data string
	var name string

	name = "Chahyo Purnomo Aji"
	fmt.Println(name)

	name = "huhaha"
	fmt.Println(name)

	//part 2 membuat variabel tanpa menyertakan tipe datanya
	var names = "angga ramadhani kurniawan"
	fmt.Println(names)

	names = "huhehe"
	fmt.Println(names)

	//integer
	var age = 50
	fmt.Println(age)

	//	part 3 membuat variabel tanpa menulis var tapi dibantu dengan :=
	dimana := "Indonesia"
	fmt.Println(dimana)
	//lalu untuk mengubahnya hanya cukup = tanpa :
	dimana = "Malaysia"
	fmt.Println(dimana)

	// part 4 deklarasi multiple variabel
	var (
		firstName = "Chahyo Purnomo"
		lastName  = "Aji"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
