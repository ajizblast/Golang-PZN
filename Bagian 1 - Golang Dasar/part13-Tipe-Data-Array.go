package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Chahyo"
	names[1] = "Purnomo"
	names[2] = "Aji"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		70,
		80,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[1])

	fmt.Println(len(values))
	fmt.Println(len(names))

}
