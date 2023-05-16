package main

import "fmt"

func main() {

	//Switch Expression
	/*
		-Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
		-Switch expression sangat sederhana dibandingkan if
		-Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable
	*/

	name := "Aji"

	switch name {
	case "Aji":
		fmt.Print("Hello Aji")
		fmt.Print("Hello Aji")
	case "Angga":
		fmt.Print("Hello Angga")
		fmt.Print("Hello Angga")
	default:
		fmt.Print("Kenalan lah!")
		fmt.Print("Kenalan lah!")
	}

	//Switch dengan Short Statement
	/*
			-Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek
		kondisinya
	*/

	//switch length := len(name); length < 5 {
	//case true:
	//	fmt.Print("Nama terlalu panjang")
	//case false:
	//	fmt.Print("Nama terlalu pendek")
	//}

	//Switch Tanpa Kondisi
	/*
		- Kondisi di switch expression tidak wajib
		- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut
		di setiap case nya

	*/

	length := len(name)
	switch {
	case length > 10:
		fmt.Print("Nama terlalu panjang")
	case length > 5:
		fmt.Print("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
