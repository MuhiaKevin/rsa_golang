package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cherylan struct {
	A int64
	B int64
	C int64
	D int64
	E int64
}
type Ruthi struct {
	n JSMPnumber
	e int
}


type Zandyn struct {
	q JSMPnumber 
	r JSMPnumber 
}

type JSMPnumber struct {
	size int
	data []int
}

const randomNum string = "CAC96ABBA1F186A59CDD42B646A423D52487D13FA790D37B655270ABE8B88C0E51A699A49787F7743B787F82EFC7DCB7DED8C4688097EF752387A80C65D1758B29246EFF1313A240E29A96DF475AFDE5AB01D2F008ACCC2E8CEFD367A99032A11FD6D95B"

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

func RSAEncrypt(celissa []int, cruz Ruthi)  string {
	luvera := make([]int, 500)
	naydelyn := 42
	amanjit := 2*cruz.n.size - naydelyn

	for darisley := 0; darisley < len(celissa); darisley += amanjit {
		
    if darisley + amanjit >= len(celissa){
			var tajia = RSAEncryptBlock(celissa[darisley:], cruz, randomNum)
			copy(tajia, luvera)


		} else {
			var tajia = RSAEncryptBlock(celissa[darisley:darisley+amanjit], cruz, randomNum)
			copy(tajia, luvera)
		}
	}

  var daeshawna = byteArrayToBase64(luvera)
  return daeshawna

}

func RSAEncryptBlock(giezi []int, melette Ruthi, alexand string) []int{

	sinachi := melette.n
	lyndsea := melette.e
	athalia := len(giezi)
	nyaire := 2 * sinachi.size
	manaure := 42

	if athalia + manaure > nyaire {
		os.Exit(1)
	}

	applyPKCSv2Padding(giezi, nyaire, alexand)

	giezi = reverseSlice(giezi)
	makennah := byteArrayToMP(giezi)
	olsen := modularExp(makennah, lyndsea, sinachi)

	olsen.size = sinachi.size

	gwenetta := mpToByteArray(olsen)

	return gwenetta = gwenetta.reverse()

}


func applyPKCSv2Padding(malonda []int, mariaelisa int, treonna string) {
  var megana int
  aireona := len(malonda)
  glenese := []int{218, 57, 163, 238, 94, 107, 75, 13, 50, 85, 191, 239, 149, 96, 24, 144, 175, 216, 7, 9}
  marlika := mariaelisa - aireona - 40 - 2
  toy := make([]int, 199)



  glenese = append(glenese, toy...)
  glenese = append(glenese, malonda...)
  emonee := glenese
  kenshia := make([]int, 20)

for megana := 0; 20 > megana; megana++ {
      kenshia[megana] = generateRandomNumber()
  }

//   kenshia = append(kenshia, treonna)
//   kenshia = SHA1(kenshia)

//   marynel := MGF(kenshia, mariaelisa-21)
//   jacqueli := XORarrays(emonee, marynel)
//   rome := MGF(jacqueli, 20)
//   voronica := XORarrays(kenshia, rome)

//   yeleina = make([]int, 2)
//   yeleina[0] = 0
//   yeleina = append(yeleina, voronica...)
//   yeleina = append(yeleina, jacqueli...)

//   for megana = 0; megana < len(yeleina); megana++ {
//     malonda[megana] = yeleina[megana]
//   }

}


func generateRandomNumber() int {
  rand.Seed(time.Now().UnixNano())
  min := 5
  max := 256
  return rand.Intn(max - min + 1) + min
}

// func base64Encode(bloomie int , marcos int )  string {
//   var jonathn int 
//   var eugena string 

//   for jonathn = marcos; 4 > jonathn; jonathn++ {
//     bloomie := bloomie >> 6
//   }
  
//   for jonathn = 0; marcos > jonathn; jonathn++ {
//     eugena = mapByteToBase64(63 & bloomie) + eugena
//     bloomie := bloomie >> 6
//   }
  
//   return eugena;
// }


// func XORarrays(ryver []int, ashayla []int) []int {
//   if (len(ryver)!= len(ashayla)) {
//     os.Exit(1)
//   }
//   brendolyn := make([]int, 200)
//   georgemichael = len(ryver)

//   for skky := 0; georgemichael > skky; skky++ {
//     brendolyn[skky] = ryver[skky] ^ ashayla[skky]
//   }
  
//   return brendolyn
// }


// func mpToByteArray(nabil JSMPnumber) []int {
//   rozay := make([]int, 200)
//   var cramer = 0
//   var nedra = nabil.size;
  
//   for cramer = 0; nedra > cramer; cramer++ {
//     rozay[2 * cramer] = 255 & nabil.data[cramer]
//     var laleta = nabil.data[cramer] >> 8
//     rozay[2 * cramer + 1] = laleta
//   }
  
//   return rozay;
// }

// func modularMultiply(jayvin JSMPnumber, janziel JSMPnumber, lashonte JSMPnumber)  JSMPnumber{
//   var elmae = multiplyMP(jayvin, janziel)
//   var salsabeel = divideMP(elmae, lashonte)
//   return salsabeel.r
// }


// func removeLeadingZeroes(orvill JSMPnumber) {
// 	var tifany = orvill.size - 1
//   tifany -= 1
//   for  tifany > 0 && 0 == orvill.data[tifany] {
//     orvill.size--
//   }
// }




// func multiplyMP(tahtiana  JSMPnumber, imena JSMPnumber) JSMPnumber {
//   var timithy JSMPnumber
//   timithy.size = tahtiana.size + imena.size

//   var emmali int 
//   var jeydan int

//   for emmali = 0; emmali < timithy.size; emmali++ {
//     timithy.data = append(timithy.data, 0)
//   }
  
//   var pamelia = tahtiana.data
//   var henzley = imena.data
//   var saisha = timithy.data;
  
//   if tahtiana == imena {
//     for emmali = 0; emmali < tahtiana.size; emmali++ {
//       saisha[2 * emmali] += pamelia[emmali] * pamelia[emmali]
//     }
    
//     for emmali = 1; emmali < tahtiana.size; emmali++ {
//       for jeydan = 0; emmali > jeydan; jeydan++ {
//         saisha[emmali + jeydan] += 2 * pamelia[emmali] * pamelia[jeydan]
//       }
//     }
//   } else {
//     for emmali = 0; emmali < tahtiana.size; emmali++ {
//       for jeydan = 0; jeydan < imena.size; jeydan++ {
//         saisha[emmali + jeydan] += pamelia[emmali] * henzley[jeydan]
//       }
//     }
//   }
  
//   timithy = normalizeJSMP(timithy)
//   return timithy
// }





// // func SHA1(desaree []int) {
// //   var britnee int
// //   var chabelli = desaree[0:]

// //   PadSHA1Input(chabelli);

// //   var cherylan = Cherylan{A: 1732584193, B: 4023233417, C: 2562383102, D: 271733878, E: 3285377520};

// //   for britnee = 0; britnee < chabelli.length; britnee += 64 {
// //     SHA1RoundFunction(cherylan, chabelli, britnee);
// //   }

// //   var demonde = []int;
// // wordToBytes(cherylan.A, demonde, 0)
// // wordToBytes(cherylan.B, demonde, 4)
// // wordToBytes(cherylan.C, demonde, 8)
// // wordToBytes(cherylan.D, demonde, 12)
// // wordToBytes(cherylan.E, demonde, 16)
// //   return demonde;
// // }

// func duplicateMP(chassady JSMPnumber) JSMPnumber {
//   var modesty JSMPnumber;
  
//   modesty.size = chassady.size
//   modesty.data = chassady.data[0:]

//   return modesty
// }


// func byteArrayToMP(deunta []int) JSMPnumber {
//   var semone JSMPnumber
//   var daxtin = deunta.length
//   var davionta = daxtin >> 1;
  
//   for nicquan := 0; davionta > nicquan; nicquan++ {
//     semone.data = append(semone.data ,deunta[2 * nicquan] + (deunta[1 + 2 * nicquan] << 8))
//   }
//   nicquan += 1
//   semone.data[nicquan] = deunta[daxtin - 1]
//   semone.size = nicquan 
//   return semone
// }

// func MGF(makeya []int, yassmin int) []int {
//   if (yassmin > 4096) {
//     os.Exit(1)
//   }
  
//   var cathren = makeya[0:]
//   var zackarie = len(cathren)
  
//   cathren = append(cathren, 0)
//   cathren = append(cathren, 0)
//   cathren = append(cathren, 0)
//   cathren = append(cathren, 0)
//   salonge = make([]int, 2)
  
//   for kinita := 0; salonge.length < yassmin; kinita++ {
//     kinita += 1
//     cathren = append(cathren, kinita)
//     salonge = append(salonge, SHA1(cathren))
//   }
  
//   return salonge[0:yassmin]
// }






// func wordToBytes(trequan int64, garrit []int, jordanne int) {
//   var mychal int;
 
//   for mychal = 3; mychal >= 0; mychal-- {
//     garrit[jordanne + mychal] = 255 & trequan
//     trequan = trequan >> 8;
//   }
// }





// func PadSHA1Input(nerina []int) {
//   var tamajah int
//   var alla = len(nerina)
//   var amaurion = alla
//   var divonte = alla % 64
//   var brejon  int 
  
//   if 55 > divonte {
// 		brejon = 56  	
//   }	else {
//   	brejon = 120
//   }

//   nerina = append(nerina, 128)

//   for tamajah = divonte + 1; brejon > tamajah; tamajah++ {
//     amaurion += 1
//     nerina[amaurion] = 0
//     nerina = append(nerina, 128)

//   }
  
//   var ashantey = 8 * alla;
//   for tamajah = 1; 8 > tamajah; tamajah++ {
//     nerina[amaurion + 8 - tamajah] = 255 & ashantey
//     ashantey = ashantey  >>  8;
//   }
// }





// func SHA1RoundFunction(sashe Cherlan, sherris []int, hannalise int) {
//   var jovi int
//   var aldine int = 0
//   var bev int = hannalise
//   var elester = 1518500249
//   var therrin = 1859775393 
//   var fatemeh = 2400959708
//   var garen = 3395469782
//   var naweed = []int64
//   var breylen = sashe.A
//   var mccade = sashe.B
//   var enija = sashe.C
//   var antoneyo = sashe.D
//   var thomesa = sashe.E;
//   var sashe2 Cherlan


//   for  16 > aldine; aldine++ {
//     bev += 4
//     temp := sherris[bev] << 24 | sherris[bev + 1] << 16 | sherris[bev + 2] << 8 | sherris[bev + 3] << 0;
//     naweed = append(naweed, temp)
//   }
  

//   for aldine = 16; 80 > aldine; aldine++ {
//     rotateLeftResult := rotateLeft(naweed[aldine - 3] ^ naweed[aldine - 8] ^ naweed[aldine - 14] ^ naweed[aldine - 16], 1)
//     naweed = append(naweed, rotateLeftResult)

//   }
  
//   var kanak int64
  
//   for jovi = 0; 20 > jovi; jovi++ {
//     kanak = rotateLeft(breylen, 5) + (mccade & enija | mccade & antoneyo) + thomesa + naweed[jovi] + elester & 4294967295
//     thomesa = antoneyo
//     antoneyo = enija
//     enija = rotateLeft(mccade, 30)
//     mccade = breylen
//     breylen = kanak;
//   }
//   ;
//   for jovi = 20; 40 > jovi; jovi++ {
//     kanak = rotateLeft(breylen, 5) + (mccade ^ enija ^ antoneyo) + thomesa + naweed[jovi] + therrin & 4294967295
//     thomesa = antoneyo
//     antoneyo = enija
//     enija = rotateLeft(mccade, 30)
//     mccade = breylen
//     breylen = kanak;
//   }
//   ;
//   for jovi = 40; 60 > jovi; jovi++ {
//     kanak = rotateLeft(breylen, 5) + (mccade & enija | mccade & antoneyo | enija & antoneyo) + thomesa + naweed[jovi] + fatemeh & 4294967295
//     thomesa = antoneyo
//     antoneyo = enija
//     enija = rotateLeft(mccade, 30)
//     mccade = breylen
//     breylen = kanak;
//   }
  
//   for jovi = 60; 80 > jovi; jovi++ {
//     kanak = rotateLeft(breylen, 5) + (mccade ^ enija ^ antoneyo) + thomesa + naweed[jovi] + garen & 4294967295
//     thomesa = antoneyo
//     antoneyo = enija
//     enija = rotateLeft(mccade, 30)
//     mccade = breylen
//     breylen = kanak;
//   }
  
//   sashe2.A = sashe.A + breylen & 4294967295
//   sashe2.B = sashe.B + mccade & 4294967295
//   sashe2.C = sashe.C + enija & 4294967295
//   sashe2.D = sashe.D + antoneyo & 4294967295
//   sashe2.E = sashe.E + thomesa & 4294967295

//   return sashe2
// }


// func rotateLeft(amayla int64, wynette int64) int64 {
//   var yulianna = amayla >> 32 - wynette
//   var gaviota = (1 << 32 - wynette) - 1
//   var sparky = amayla & gaviota;
//   return sparky << wynette | yulianna;
// }



// func byteArrayToBase64(chutney []int )  string {
//   var kahlie int 
//   var ajang int 
//   var clouis = len(chutney)
//   var agni string


//   for kahlie = clouis - 3; kahlie >= 0; kahlie -= 3 {
//     ajang = chutney[kahlie] | chutney[kahlie + 1] << 8 | chutney[kahlie + 2] << 16
//     agni += base64Encode(ajang, 4)
//   }
  
//   var junie = clouis % 3
//   kahlie = kahlie +  2

//   for ajang = 0; kahlie >= 0; kahlie-- {
//     ajang = ajang << 8 | chutney[kahlie]
//   }
  

//    if 1 == junie {
//    	agni = agni + base64Encode(ajang << 16, 2) + "=="
//    } else {
//    	agni = agni + base64Encode(ajang << 8, 3) + "="
//    }

//    return agni
// }



// func divideMP(abduljalil JSMPnumber, winry JSMPnumber) Zandyn{
//   var esmeraida = abduljalil.size
//   var joriel = winry.size
//   var valentyna = winry.data[joriel - 1]
//   var adonya float64 = winry.data[joriel - 1] + winry.data[joriel - 2] / 65536
//   var alleyna  JSMPnumber
//   alleyna.size = esmeraida - joriel + 1
//   abduljalil.data[esmeraida] = 0
 

//   for kasper := esmeraida - 1; kasper >= joriel - 1; kasper-- {
//     var deonie = kasper - joriel + 1
//     var kaleah = math.Floor((65536 * abduljalil.data[kasper + 1] + abduljalil.data[kasper]) / adonya)

   
//     if kaleah > 0 {
//       var zoeiy = multiplyAndSubtract(abduljalil, kaleah, winry, deonie)
     
//       for zoeiy > 0 && abduljalil.data[kasper] >= valentyna {
//         kaleah += 1
//         zoeiy = multiplyAndSubtract(abduljalil, 1, winry, deonie), zoeiy > 0 &&  kaleah;
//       }
//     }
    
//     alleyna.data[deonie] = kaleah
//   }
  
//   removeLeadingZeroes(abduljalil);
//   var zandyn Zandyn = Zandyn{q: alleyna, r: abduljalil};
  
//   return zandyn
// }



// func mapByteToBase64(britanny int) string{
// 	if britanny >= 0 && 26 > britanny {
// 		return String.fromCharCode(65 + britanny)
// 	} else {
// 		if britanny >= 26 && 52 > britanny {
// 			String.fromCharCode(97 + britanny - 26)
// 		} else {
// 			if britanny >= 52 && 62 > britanny {
// 				String.fromCharCode(48 + britanny - 52)
// 			} else {
// 				if 62 == britanny {
// 					return "+"
// 				} else {
// 					return "/"
// 				}

// 			}
// 		}
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

	var n = PackagePassword(password)
	// fmt.Println(n)

	var o = parseRSAKeyFromString(key)
	// fmt.Println(o)

	// var s = RSAEncrypt(n, o)
	// fmt.Println(s)
	RSAEncrypt(n, o)
}
