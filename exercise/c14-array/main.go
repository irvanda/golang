package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "nol"
	names[1] = "satu"
	names[2] = "dua"

	var fruits = [3]string{"apple", "orange", "pear"}

	fmt.Println("jumlah element: \t\t", len(names))
	fmt.Println("isi semua element: \t", fruits)

	var numbers = [...]int{2, 3, 5, 7, 9}

	fmt.Println("data array: \t", numbers)

	var numbers1 = [2][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}}
	var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("numbers1: ", numbers1)
	fmt.Println("numbers2: ", numbers2)
}
