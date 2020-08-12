package main

import "fmt"

func main() {
	fmt.Println("Array")
	var names [4]string
	names[0] = "somad"
	names[1] = "aldo"
	names[2] = "sule"
	names[3] = "darman"

	fmt.Println(names)

	// inisialisasi
	fmt.Println("\nInisialisasi")
	var fruits = [4]string{
		"apple", "orange", "watermelon", "peer",
	}
	fmt.Println(fruits)

	// inisial tanpa nilai awal array
	fmt.Println("\nInisial tanpa nilai awal")
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(numbers)

	// array multidimensi
	fmt.Println("\nMultidimensi array")
	var numbers1 = [3][2]int{
		[2]int{
			3, 2,
		},
		[2]int{
			2, 3,
		},
		[2]int{
			6, 5,
		},
	}

	var numbers2 = [3][2]int{
		{
			3, 2,
		}, {
			2, 3,
		}, {
			6, 5,
		},
	}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// For Array
	fmt.Println("\nFor array")
	var fruits2 = [1]string{0: "test"}

	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits2[i])
	}

}
