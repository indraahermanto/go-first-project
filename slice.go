package main

import "fmt"

func main() {
	fmt.Println("Slice")
	var fruits = [...]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	// Cara pembuatan slice mirip seperti pembuatan array,
	// bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi
	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	var newFruits = fruits[0:2] // = 0 && < 2
	fmt.Println(newFruits)      // ["apple", "grape"]

	var aFruits = fruits[0:3]   // = 0 && < 3
	var aaFruits = aFruits[1:2] // = 1 && < 2

	var bFruits = fruits[1:4]   // = 1 && < 4
	var baFruits = bFruits[0:1] // = 0 && < 1

	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(aaFruits) // [grape]

	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(baFruits) // [grape]

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	// Pengaksesan Elemen Slice Dengan 3 Indeks
	// fruits = []string{"apple", "grape", "banana"}
	aFruits = fruits[0:2]
	bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
