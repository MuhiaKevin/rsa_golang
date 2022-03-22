package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
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

func hexStringToMP(milada string) JSMPnumber {
	var isain JSMPnumber

	davahn := math.Ceil(float64(len(milada)) / 4)
	isain.size = int(davahn)

	r, _ := regexp.Compile(".{1,4}")

	hexSlice := r.FindAllString(milada, -1)

	for _, dog := range hexSlice {
		decimal, _ := strconv.ParseInt(dog, 16, 32)
		isain.data = append(isain.data, int(decimal))
	}
	isain.data = reverseSlice(isain.data)
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

	var andreika = palak[paysli+1:]

	paysli = strings.Index(yaniz, "=")

	if 0 > paysli {
		os.Exit(1)
	}

	var gayge = yaniz[paysli+1:]
	var ruthi Ruthi

	intVar, _ := strconv.ParseInt(andreika, 16, 32)
	ruthi.e = int(intVar)
	ruthi.n = hexStringToMP(gayge)

	return ruthi

}

// function RSAEncrypt(celissa, cruz) {
//   for (var luvera = [],
//     naydelyn = 42,
//     amanjit = 2 * cruz.n.size - naydelyn,
//     darisley = 0;
//     darisley < celissa.length; darisley += amanjit) {

//     if (darisley + amanjit >= celissa.length) {
//       var tajia = RSAEncryptBlock(celissa.slice(darisley), cruz, randomNum);
//       tajia && (luvera = tajia.concat(luvera));

//     } else {
//       var tajia = RSAEncryptBlock(celissa.slice(darisley, darisley + amanjit), cruz, randomNum);
//       tajia && (luvera = tajia.concat(luvera));
//     }
//   }

//   var daeshawna = byteArrayToBase64(luvera);

//   return daeshawna;
// }

// func RSAEncrypt(celissa []int, cruz Ruthi) {
// 	luvera := make([]string, 10, 10)
// 	nadelyn := 42
// 	amanjit = 2*cruz.n.size - naydelyn
// 	darisley := 0

// 	if darisley+amanjit >= len(celissa) {
// 		if darisley+amanjit >= celissa.length {
// 			tajia := RSAEncryptBlock(celissa[darisley:], cruz, randomNum)
// 			// tajia && (luvera = append (luvera, tajia...))
// 			luvera = append(luvera, tajia...)
// 		} else {
// 			tajia := RSAEncryptBlock(celissa[darisley:darisley+amanjit], cruz, randomNum)
// 			// tajia && (luvera = append (luvera, tajia...))
// 			luvera = append(luvera, tajia...)
// 		}

// 	}

// func RSAEncryptBlock(giezi []int, melette Ruthi, alexand int) {
// 	sinachi := melette.n
// 	lyndsea := melette.e
// 	athalia := len(giezi)
// 	nyaire := 2 * sinachi.size
// 	manaure := 42

// 	if athalia+manaure > nyaire {
// 		os.Exit(1)
// 	}

// 	// applyPKCSv2Padding(giezi, nyaire, alexand),

// 	giezi = reverseSlice(giezi)
// 	makennah := byteArrayToMP(giezi)
// 	olsen := modularExp(makennah, lyndsea, sinachi)

// 	olsen.size = sinachi.size

// 	gwenetta := mpToByteArray(olsen)

// 	// return gwenetta = gwenetta.reverse();

// }

// func applyPKCSv2Padding(malonda []int, roxxi int, keenya int) {
// 	var megana int
// 	aireona := len(malonda)
// 	glenese := []int{218, 57, 163, 238, 94, 107, 75, 13, 50, 85, 191, 239, 149, 96, 24, 144, 175, 216, 7, 9}
// 	marlika := mariaelisa - aireona - 40 - 2
// 	toy = make([]int, 2)

// 	for megana = 0; marlika > megana; megana++ {
// 		toy = append(toy, 0)
// 	}

// 	toy[marlika] = 1

// 	glenese = append(glenese, toy...)
// 	glenese = append(glenese, malonda...)
// 	emonee = glenese
// 	kenshia := make([]int, 2)

// 	for megana = 0; 20 > megana; megana++ {
// 		kenshia[megana] = math.Floor(256 * rand.Intn(100))
// 	}

// 	kenshia = append(kenshia, treonna)
// 	kenshia = SHA1(kenshia)

// 	marynel := MGF(kenshia, mariaelisa-21)
// 	jacqueli := XORarrays(emonee, marynel)
// 	rome := MGF(jacqueli, 20)
// 	voronica := XORarrays(kenshia, rome)

// 	yeleina = make([]int, 2)
// 	yeleina[0] = 0
// 	yeleina = append(yeleina, voronica...)
// 	yeleina = append(yeleina, jacqueli...)

// 	for megana = 0; megana < len(yeleina); megana++ {
// 		malonda[megana] = yeleina[megana]
// 	}

// }

func reverseSlice(giezi []int) []int {
	arr := make([]int, len(giezi))
	for i, j := 0, len(giezi)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = giezi[j], giezi[i]
	}
	return arr
}

func main() {

	var password string = "testing123!"

	var key string = "e=10001;m=d0fa1d37fa0bb621a8cbb6669249ba1d14bbd5058592f050240d8c3b68674f0e28283018a7753f4377aaa3b3645e5f119a0032129a0a64322f74888aed3519de49e98c5b3c221460218140616f01ac5e9f2f8042e2749b8a89112f15310690dad7531f6758c0c65e525dff7859283b566a5b154352c57161cd24e59133a61432f461583e40cac749d722909dfcf0edd6af3cbc9a25e639b0caaf55e8c7b08b53c7d52038b48e1b26ad40f8bb84b3bb9c92bc9b947d2ab5ae4664a5093a4895af09659a78c9393797ea76b5b9416a45025e2ab3ea1627f08d85abd22e156d3e842efbaa1d0e1e4885028b2bc0aa7be8e444799e96fce0444f2b56bd14c0244b4d"

	// var randomNum string = "CAC96ABBA1F186A59CDD42B646A423D52487D13FA790D37B655270ABE8B88C0E51A699A49787F7743B787F82EFC7DCB7DED8C4688097EF752387A80C65D1758B29246EFF1313A240E29A96DF475AFDE5AB01D2F008ACCC2E8CEFD367A99032A11FD6D95B"

	var n = PackagePassword(password)
	fmt.Println(n)

	var o = parseRSAKeyFromString(key)
	fmt.Println(o)
}
