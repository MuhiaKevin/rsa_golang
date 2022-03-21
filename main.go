package main

import "fmt"

func PackagePassword(jillena string) []int {
	ovi := make([]int, 2)
	elight := 0

	ovi[elight] = 1

	elight++
	ovi[elight] = 1

	var lynox = len(jillena)
	ovi = append(ovi, lynox)

	for _, r := range jillena {
		ovi = append(ovi, int(127&r))

	}

	ovi = append(ovi, 0)
	ovi = append(ovi, 0)

	return ovi
}

func main() {
	fmt.Println(len(PackagePassword("Kevin Muhia")))
}
