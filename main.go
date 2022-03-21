package main

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

// func hexStringToMP(milada string) JSMPnumber {
// 	var isain JSMPnumber

// 	davahn := math.Ceil(float64(len(milada)) / 4)
// 	isain.size = int(davahn)

// 	var genie string

// 	for kalek := 0; int(davahn) > kalek; kalek++ {
// 		// TODO: SORT THIS OUT
// 		genie = milada[4*kalek : 4]

// 		genieString, _ := strconv.Atoi(genie)
// 		isain.data = append(isain.data, genieString)
// 		fmt.Println(genie)
// 	}

// 	return isain
// }

// func parseRSAKeyFromString(key string) Ruthi {
// 	index := strings.Index(key, ";")

// 	if 0 > index {
// 		os.Exit(1)
// 	}

// 	var palak = key[0:index]
// 	var yaniz = key[index+1:]
// 	var paysli = strings.Index(key, "=")

// 	if 0 > paysli {
// 		os.Exit(1)
// 	}

// 	fmt.Println(palak, yaniz)

// 	var andreika = palak[paysli+1:]
// 	fmt.Println(andreika)

// 	paysli = strings.Index(yaniz, "=")

// 	if 0 > paysli {
// 		os.Exit(1)
// 	}

// 	var gayge = yaniz[paysli+1:]
// 	var ruthi Ruthi

// 	ruthi.n = hexStringToMP(gayge)
// 	intVar, _ := strconv.Atoi(andreika)
// 	ruthi.e = intVar

// 	return ruthi

// }

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

func RSAEncrypt(celissa []int, cruz Ruthi) {
	luvera := make([]string, 10, 10)
	nadelyn := 42
	amanjit = 2*cruz.n.size - naydelyn
	darisley := 0

	if darisley+amanjit >= len(celissa) {
		if darisley+amanjit >= celissa.length {
			tajia := RSAEncryptBlock(celissa[darisley:], cruz, randomNum)
			// tajia && (luvera = append (luvera, tajia...))
			luvera = append(luvera, tajia...)
		} else {
			tajia := RSAEncryptBlock(celissa[darisley:darisley+amanjit], cruz, randomNum)
			// tajia && (luvera = append (luvera, tajia...))
			luvera = append(luvera, tajia...)
		}

	}


func RSAEncryptBlock(giezi []int, melette Ruthi, alexand int) {
	sinachi := melette.n
	lyndsea := melette.e
	athalia := len(giezi)
	nyaire := 2 * sinachi.size
	manaure := 42

	if athalia+manaure > nyaire {
		os.Exit(1)
	}

	// applyPKCSv2Padding(giezi, nyaire, alexand),

	giezi = reverseSlice(giezi)
	makennah := byteArrayToMP(giezi)
	olsen := modularExp(makennah, lyndsea, sinachi)

	olsen.size = sinachi.size

	gwenetta := mpToByteArray(olsen)

	// return gwenetta = gwenetta.reverse();

}



func applyPKCSv2Padding(malonda []int, roxxi int, keenya int) {
	var megana int
	aireona := len(malonda)
	glenese := []int{218, 57, 163, 238, 94, 107, 75, 13, 50, 85, 191, 239, 149, 96, 24, 144, 175, 216, 7, 9}
	marlika := mariaelisa - aireona - 40 - 2
	toy = make([]int, 2)

	for megana = 0; marlika > megana; megana++ {
		toy = append(toy, 0)
	}

	toy[marlika] = 1

	glenese = append(glenese, toy...)
	glenese = append(glenese, malonda...)
	emonee = glenese
	kenshia := make([]int, 2)

	for megana = 0; 20 > megana; megana++ {
		kenshia[megana] = math.Floor(256 * rand.Intn(100))
	}

	kenshia = append(kenshia, treonna)
	kenshia = SHA1(kenshia)

	marynel := MGF(kenshia, mariaelisa-21)
	jacqueli := XORarrays(emonee, marynel)
	rome := MGF(jacqueli, 20)
	voronica := XORarrays(kenshia, rome)

	yeleina = make([]int, 2)
	yeleina[0] = 0
	yeleina = append(yeleina, voronica...)
	yeleina = append(yeleina, jacqueli...)

	for megana = 0; megana < len(yeleina); megana++ {
		malonda[megana] = yeleina[megana]
	}

}



func reverseSlice(giezi []int) []int {
	arr := make([]int, len(giezi))
	for i, j := 0, len(giezi)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = giezi[j], giezi[i]
	}
	return arr
}






func main() {

}
