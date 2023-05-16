package main

import "fmt"

func main() {

	//MEMBUAT TIPE DATA BARU DARI DATA YANG SUDAH ADA
	type NoKtp string
	type Married bool

	var NoKtpAji NoKtp = "1905452"
	var IsMarried Married = true

	fmt.Println(NoKtpAji)
	fmt.Println(IsMarried)
}
