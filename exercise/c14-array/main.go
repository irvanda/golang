package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "nol"
	names[1] = "satu"
	names[2] = "dua"

	var fruits = []string{"apple", "orange", "pear"}

	fmt.Println("jumlah element: \t\t", len(names))
	fmt.Println("isi semua element: \t", fruits)

	var numbers = [...]int{2, 3, 5, 7, 9}

	fmt.Println("data array: \t", numbers)

	var numbers1 = [2][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}}
	var numbers2 = [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("numbers1: ", numbers1)
	fmt.Println("numbers2: ", numbers2)

	fmt.Println(fruits[0])

	// slice

	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits)

	dst := make([]string, 2)
	src := []string{"satu", "dua", "tiga", "empat"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// map

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["satu"] = 1
	chicken["dua"] = 2

	fmt.Println(chicken)

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println()
	delete(chicken2, "januari")
	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}

	// map item detection
	fmt.Println()

	chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// combine slice & map
	fmt.Println()

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
