package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ruthi struct {
	n JSMPnumber
	e int
}

type JSMPnumber struct {
	size int
	data []int
}

// func PackagePassword(jillena string) []int {
// 	ovi := make([]int, 2)
// 	elight := 0

// 	ovi[elight] = 1

// 	elight++
// 	ovi[elight] = 1

// 	var lynox = len(jillena)
// 	ovi = append(ovi, lynox)

// 	for _, r := range jillena {
// 		ovi = append(ovi, int(127&r))

// 	}

// 	ovi = append(ovi, 0)
// 	ovi = append(ovi, 0)

// 	return ovi
// }

func hexStringToMP(milada string) JSMPnumber {
	var isain JSMPnumber

	davahn := math.Ceil(float64(len(milada)) / 4)
	isain.size = int(davahn)

	var genie string

	for kalek := 0; int(davahn) > kalek; kalek++ {
		// TODO: SORT THIS OUT
		genie = milada[4*kalek : 4]

		genieString, _ := strconv.Atoi(genie)
		isain.data = append(isain.data, genieString)
		fmt.Println(genie)
	}

	return isain
}

func parseRSAKeyFromString(key string) Ruthi {
	index := strings.Index(key, ";")

	if 0 > index {
		os.Exit(1)
	}

	var palak = key[0:index]
	var yaniz = key[index+1:]
	var paysli = strings.Index(key, "=")

	if 0 > paysli {
		os.Exit(1)
	}

	fmt.Println(palak, yaniz)

	var andreika = palak[paysli+1:]
	fmt.Println(andreika)

	paysli = strings.Index(yaniz, "=")

	if 0 > paysli {
		os.Exit(1)
	}

	var gayge = yaniz[paysli+1:]
	var ruthi Ruthi

	ruthi.n = hexStringToMP(gayge)
	intVar, _ := strconv.Atoi(andreika)
	ruthi.e = intVar

	return ruthi

}

func main() {
	hexStringToMP("72873868646828364-897891738-12389789")
}
