package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "aji",
		"address": "jakarta",
	}

	//untuk menambah data
	person["title"] = "programmer"

	fmt.Println(person)
	//untuk mengakses seperti dibawah
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "aji"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
