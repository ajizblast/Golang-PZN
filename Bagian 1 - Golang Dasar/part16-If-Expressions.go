package main

import "fmt"

func main() {
	//Else If Expression
	/*
		- Kadang dalam If, kita butuh membuat beberapa kondisi
		- Kasus seperti ini, kita bisa menggunakan Else If expression
	*/

	//kondisi true
	var name = "aji"
	if name == "aji" {
		fmt.Print("hello aji")
	}

	//kondisi false
	var names = "angga"
	if names == "chahyo" {
		fmt.Println("hello chahyo")
	} else {
		fmt.Println("hei, kenalan lah!")
	}

	// kondisi if else | else if bisa ditambahkan sebanyak mungkin
	var kota = "jakarta"
	if kota == "malang" {
		fmt.Println("disini jakarta uye")
	} else if kota == "bandung" {
		fmt.Println("hello disini bandung")
	} else if kota == "bali" {
		fmt.Println("kadek bali nih senggol dong")
	} else {
		fmt.Println("ini kota padang loh")
	}

	//if short statement
	/*
		- If mendukung short statement sebelum kondisi
		- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan
		terhadap kondisi
	*/

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang bro")
	} else {
		fmt.Println("nama sudah benar")
	}
}
