package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["jan"] = 50
	chicken["feb"] = 40

	fmt.Println("jan", chicken["jan"])
	fmt.Println("feb", chicken["feb"])

	// inisialisasi
	fmt.Println("\nInisialisasi")
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken2["januari"])

	// for
	fmt.Println("\nFor (Pengulangan)")
	chicken = map[string]int{
		"jan": 50,
		"feb": 30,
		"mar": 40,
		"apr": 10,
	}

	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}

	// delete item map
	fmt.Println("\nDelete item map")
	fmt.Println(len(chicken)) // menghitung jumlah map
	fmt.Println(chicken)

	delete(chicken, "mar")
	fmt.Println(len(chicken))
	fmt.Println(chicken)

}
