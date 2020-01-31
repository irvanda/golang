package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	name := new(string)
	*name = "tes"

	var decimalNumber float32 = 2.62

	var exist bool = false

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Printf("%.3f\n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)

	for i := 0; i < 5; i++ {
		fmt.Println("For-normal", i)
	}

	var i = 0

	for i < 5 {
		fmt.Println("For-with-argument", i)
		i++
	}

	i = 0

	for {
		fmt.Println("For-no-argument", i)

		i++
		if i == 5 {
			break
		}
	}

	for i = 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		if i == 4 {
			break
		}

		fmt.Println("Angka", i)
	}
}
