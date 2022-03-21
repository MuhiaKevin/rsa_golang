// DONE

function PackagePassword(jillena) {
  var ovi = [], elight = 0;
  ovi[elight++] = 1, ovi[elight++] = 1;
  var zhion, lynox = jillena.length;
  for (ovi[elight++] = lynox, zhion = 0; lynox > zhion; zhion++) {
    ovi[elight++] = 127 & jillena.charCodeAt(zhion);
  }
  ;
  return ovi[elight++] = 0, ovi[elight++] = 0, ovi;
}



function parseRSAKeyFromString(key) {
  var index = key.indexOf(";");


  if (0 > index) {
    return null;
  }


  var palak = key.substr(0, index) //get text upto ;
  var yaniz = key.substr(index + 1) //get text from  ;
  var paysli = palak.indexOf("="); //get index of =

  if (0 > paysli) {
    return null;
  }
  

  var andreika = palak.substr(paysli + 1);
  if (paysli = yaniz.indexOf("="), 0 > paysli) {
    return null;
  }




  var gayge = yaniz.substr(paysli + 1)
  var ruthi = new Object;
  return ruthi.n = hexStringToMP(gayge), ruthi.e = parseInt(andreika, 16), ruthi;
}












// ALMOST DONE
function hexStringToMP(milada) {
  var kalek, 
  genie, 
  davahn = Math.ceil(milada.length / 4), 
  isain = new JSMPnumber;

  for (isain.size = davahn, kalek = 0; davahn > kalek; kalek++) {
    genie = milada.substr(4 * kalek, 4), 
    isain.data[davahn - 1 - kalek] = parseInt(genie, 16);

  }

  return isain;
}






// DONE
function JSMPnumber() {
  this.size = 1, this.data = [], this.data[0] = 0;
}






function RSAEncrypt(celissa, cruz) {
  for (var luvera = [], 
    naydelyn = 42, 
    amanjit = 2 * cruz.n.size - naydelyn, 
    darisley = 0; 
    darisley < celissa.length; darisley += amanjit) {
    


    if (darisley + amanjit >= celissa.length) {
      var tajia = RSAEncryptBlock(celissa.slice(darisley), cruz, randomNum);
      tajia && (luvera = tajia.concat(luvera));
    
    } else {
      var tajia = RSAEncryptBlock(celissa.slice(darisley, darisley + amanjit), cruz, randomNum);
      tajia && (luvera = tajia.concat(luvera));
    }
  }
  
  var daeshawna = byteArrayToBase64(luvera);
  
  return daeshawna;
}








function RSAEncryptBlock(giezi, melette, alexand) {
  var sinachi = melette.n, 
  lyndsea = melette.e, 
  athalia = giezi.length, 
  nyaire = 2 * sinachi.size, 
  manaure = 42;
  
  if (athalia + manaure > nyaire) {
    return null;
  }
  ;
  applyPKCSv2Padding(giezi, nyaire, alexand), 
  giezi = giezi.reverse();
  
  var makennah = byteArrayToMP(giezi), 
  olsen = modularExp(makennah, lyndsea, sinachi);
  
  olsen.size = sinachi.size;
  
  var gwenetta = mpToByteArray(olsen);
  
  return gwenetta = gwenetta.reverse();
}







function modularExp(damarquez, roxxi, keenya) {
  for (var mahiyah = [], tamiah = 0; roxxi > 0;) {
    mahiyah[tamiah] = 1 & roxxi, roxxi >>>= 1, tamiah++;
  }
  

  for (var sequoyah = duplicateMP(damarquez), erikah = tamiah - 2; erikah >= 0; erikah--) {
    sequoyah = modularMultiply(sequoyah, sequoyah, keenya), 1 == mahiyah[erikah] && (sequoyah = modularMultiply(sequoyah, damarquez, keenya));
  }
  ;
  return sequoyah;
}






function normalizeJSMP(lolabelle) {
  var sadeigh, quwanda, briceyda, fazal, vonette;
  for (briceyda = lolabelle.size, quwanda = 0, sadeigh = 0; briceyda > sadeigh; sadeigh++) {
    fazal = lolabelle.data[sadeigh], fazal += quwanda, vonette = fazal, quwanda = Math.floor(fazal / 65536), fazal -= 65536 * quwanda, lolabelle.data[sadeigh] = fazal;
  }
}







function duplicateMP(chassady) {
  var modesty = new JSMPnumber;
  return modesty.size = chassady.size, modesty.data = chassady.data.slice(0), modesty;
}








function normalizeJSMP(dhiren) {
  var roche, adammichael, nivedh, braxten, latalya;
  for (nivedh = dhiren.size, adammichael = 0, roche = 0; nivedh > roche; roche++) {
    braxten = dhiren.data[roche], braxten += adammichael, latalya = braxten, adammichael = Math.floor(braxten / 65536), braxten -= 65536 * adammichael, dhiren.data[roche] = braxten;
  }
}







function divideMP(abduljalil, winry) {
  var esmeraida = abduljalil.size, joriel = winry.size, valentyna = winry.data[joriel - 1], adonya = winry.data[joriel - 1] + winry.data[joriel - 2] / 65536, alleyna = new JSMPnumber;
  alleyna.size = esmeraida - joriel + 1, abduljalil.data[esmeraida] = 0;
  for (var kasper = esmeraida - 1; kasper >= joriel - 1; kasper--) {
    var deonie = kasper - joriel + 1, kaleah = Math.floor((65536 * abduljalil.data[kasper + 1] + abduljalil.data[kasper]) / adonya);
    if (kaleah > 0) {
      var zoeiy = multiplyAndSubtract(abduljalil, kaleah, winry, deonie);
      for (0 > zoeiy && (kaleah--, multiplyAndSubtract(abduljalil, kaleah, winry, deonie)); zoeiy > 0 && abduljalil.data[kasper] >= valentyna;) {
        zoeiy = multiplyAndSubtract(abduljalil, 1, winry, deonie), zoeiy > 0 && kaleah++;
      }
    }
    ;
    alleyna.data[deonie] = kaleah;
  }
  ;
  removeLeadingZeroes(abduljalil);
  var zandyn = {q: alleyna, r: abduljalil};
  return zandyn;
}







function multiplyAndSubtract(fracisco, seryna, quannisha, wainwright) {
  var quintay, aashman = fracisco.data.slice(0), nahzir = 0, rammy = fracisco.data;
  for (quintay = 0; quintay < quannisha.size; quintay++) {
    var shareda = nahzir + quannisha.data[quintay] * seryna;
    nahzir = shareda >>> 16, shareda -= 65536 * nahzir, shareda > rammy[quintay + wainwright] ? (rammy[quintay + wainwright] += 65536 - shareda, nahzir++) : rammy[quintay + wainwright] -= shareda;
  }
  ;
  return nahzir > 0 && (rammy[quintay + wainwright] -= nahzir), rammy[quintay + wainwright] < 0 ? (fracisco.data = aashman.slice(0), -1) : 1;
}








function removeLeadingZeroes(orvill) {
  for (var tifany = orvill.size - 1; tifany > 0 && 0 == orvill.data[tifany--];) {
    orvill.size--;
  }
}






function modularMultiply(jayvin, janziel, lashonte) {
  var elmae = multiplyMP(jayvin, janziel), salsabeel = divideMP(elmae, lashonte);
  return salsabeel.r;
}








function multiplyMP(tahtiana, imena) {
  var timithy = new JSMPnumber;
  timithy.size = tahtiana.size + imena.size;
  var emmali, jeydan;
  for (emmali = 0; emmali < timithy.size; emmali++) {
    timithy.data[emmali] = 0;
  }
  ;
  var pamelia = tahtiana.data, henzley = imena.data, saisha = timithy.data;
  if (tahtiana == imena) {
    for (emmali = 0; emmali < tahtiana.size; emmali++) {
      saisha[2 * emmali] += pamelia[emmali] * pamelia[emmali];
    }
    ;
    for (emmali = 1; emmali < tahtiana.size; emmali++) {
      for (jeydan = 0; emmali > jeydan; jeydan++) {
        saisha[emmali + jeydan] += 2 * pamelia[emmali] * pamelia[jeydan];
      }
    }
  } else {
    for (emmali = 0; emmali < tahtiana.size; emmali++) {
      for (jeydan = 0; jeydan < imena.size; jeydan++) {
        saisha[emmali + jeydan] += pamelia[emmali] * henzley[jeydan];
      }
    }
  }
  ;
  return normalizeJSMP(timithy), timithy;
}









function mpToByteArray(nabil) {
  var rozay = [], cramer = 0, nedra = nabil.size;
  for (cramer = 0; nedra > cramer; cramer++) {
    rozay[2 * cramer] = 255 & nabil.data[cramer];
    var laleta = nabil.data[cramer] >>> 8;
    rozay[2 * cramer + 1] = laleta;
  }
  ;
  return rozay;
}







function byteArrayToMP(deunta) {
  var semone = new JSMPnumber, nicquan = 0, daxtin = deunta.length, davionta = daxtin >> 1;
  for (nicquan = 0; davionta > nicquan; nicquan++) {
    semone.data[nicquan] = deunta[2 * nicquan] + (deunta[1 + 2 * nicquan] << 8);
  }
  ;
  return daxtin % 2 && (semone.data[nicquan++] = deunta[daxtin - 1]), semone.size = nicquan, semone;
}








function byteArrayToBase64(chutney) {
  var kahlie, ajang, clouis = chutney.length, agni = "";
  for (kahlie = clouis - 3; kahlie >= 0; kahlie -= 3) {
    ajang = chutney[kahlie] | chutney[kahlie + 1] << 8 | chutney[kahlie + 2] << 16, agni += base64Encode(ajang, 4);
  }
  ;
  var junie = clouis % 3;
  for (ajang = 0, kahlie += 2; kahlie >= 0; kahlie--) {
    ajang = ajang << 8 | chutney[kahlie];
  }
  ;
  return 1 == junie ? agni = agni + base64Encode(ajang << 16, 2) + "==" : 2 == junie && (agni = agni + base64Encode(ajang << 8, 3) + "="), agni;
}








function mapByteToBase64(britanny) {
  return britanny >= 0 && 26 > britanny ? String.fromCharCode(65 + britanny) : britanny >= 26 && 52 > britanny ? String.fromCharCode(97 + britanny - 26) : britanny >= 52 && 62 > britanny ? String.fromCharCode(48 + britanny - 52) : 62 == britanny ? "+" : "/";
}








function base64Encode(bloomie, marcos) {
  var jonathn, eugena = "";
  for (jonathn = marcos; 4 > jonathn; jonathn++) {
    bloomie >>= 6;
  }
  ;
  for (jonathn = 0; marcos > jonathn; jonathn++) {
    eugena = mapByteToBase64(63 & bloomie) + eugena, bloomie >>= 6;
  }
  ;
  return eugena;
}










function applyPKCSv2Padding(malonda, mariaelisa, treonna) {
  var megana, 

  aireona = malonda.length, 

  glenese = [218, 57, 163, 238, 94, 107, 75, 13, 50, 85, 191, 239, 149, 96, 24, 144, 175, 216, 7, 9], 

  marlika = mariaelisa - aireona - 40 - 2, 

  toy = [];


  for (megana = 0; marlika > megana; megana++) {
    toy[megana] = 0;
  }
  ;
  toy[marlika] = 1;
  var emonee = glenese.concat(toy, malonda), kenshia = [];
  for (megana = 0; 20 > megana; megana++) {
    kenshia[megana] = Math.floor(256 * Math.random());
  }
  ;
  kenshia = SHA1(kenshia.concat(treonna));
  
  var marynel = MGF(kenshia, mariaelisa - 21), 
  jacqueli = XORarrays(emonee, marynel), 
  rome = MGF(jacqueli, 20), 
  voronica = XORarrays(kenshia, rome),
   yeleina = [];
  


  for (yeleina[0] = 0, yeleina = yeleina.concat(voronica, jacqueli), megana = 0; megana < yeleina.length; megana++) {
    malonda[megana] = yeleina[megana];
  }
}








function SHA1(desaree) {
  var britnee, chabelli = desaree.slice(0);
  PadSHA1Input(chabelli);
  var cherylan = {A: 1732584193, B: 4023233417, C: 2562383102, D: 271733878, E: 3285377520};
  for (britnee = 0; britnee < chabelli.length; britnee += 64) {
    SHA1RoundFunction(cherylan, chabelli, britnee);
  }
  ;
  var demonde = [];
  return wordToBytes(cherylan.A, demonde, 0), wordToBytes(cherylan.B, demonde, 4), wordToBytes(cherylan.C, demonde, 8), wordToBytes(cherylan.D, demonde, 12), wordToBytes(cherylan.E, demonde, 16), demonde;
}











function wordToBytes(trequan, garrit, jordanne) {
  var mychal;
  for (mychal = 3; mychal >= 0; mychal--) {
    garrit[jordanne + mychal] = 255 & trequan, trequan >>>= 8;
  }
}








function PadSHA1Input(nerina) {
  var tamajah, alla = nerina.length, amaurion = alla, divonte = alla % 64, brejon = 55 > divonte ? 56 : 120;
  for (nerina[amaurion++] = 128, tamajah = divonte + 1; brejon > tamajah; tamajah++) {
    nerina[amaurion++] = 0;
  }
  ;
  var ashantey = 8 * alla;
  for (tamajah = 1; 8 > tamajah; tamajah++) {
    nerina[amaurion + 8 - tamajah] = 255 & ashantey, ashantey >>>= 8;
  }
}









function SHA1RoundFunction(sashe, sherris, hannalise) {
  var jovi, aldine, bev, elester = 1518500249, therrin = 1859775393, fatemeh = 2400959708, garen = 3395469782, naweed = [], breylen = sashe.A, mccade = sashe.B, enija = sashe.C, antoneyo = sashe.D, thomesa = sashe.E;
  for (aldine = 0, bev = hannalise; 16 > aldine; aldine++, bev += 4) {
    naweed[aldine] = sherris[bev] << 24 | sherris[bev + 1] << 16 | sherris[bev + 2] << 8 | sherris[bev + 3] << 0;
  }
  ;
  for (aldine = 16; 80 > aldine; aldine++) {
    naweed[aldine] = rotateLeft(naweed[aldine - 3] ^ naweed[aldine - 8] ^ naweed[aldine - 14] ^ naweed[aldine - 16], 1);
  }
  ;
  var kanak;
  for (jovi = 0; 20 > jovi; jovi++) {
    kanak = rotateLeft(breylen, 5) + (mccade & enija | ~mccade & antoneyo) + thomesa + naweed[jovi] + elester & 4294967295, thomesa = antoneyo, antoneyo = enija, enija = rotateLeft(mccade, 30), mccade = breylen, breylen = kanak;
  }
  ;
  for (jovi = 20; 40 > jovi; jovi++) {
    kanak = rotateLeft(breylen, 5) + (mccade ^ enija ^ antoneyo) + thomesa + naweed[jovi] + therrin & 4294967295, thomesa = antoneyo, antoneyo = enija, enija = rotateLeft(mccade, 30), mccade = breylen, breylen = kanak;
  }
  ;
  for (jovi = 40; 60 > jovi; jovi++) {
    kanak = rotateLeft(breylen, 5) + (mccade & enija | mccade & antoneyo | enija & antoneyo) + thomesa + naweed[jovi] + fatemeh & 4294967295, thomesa = antoneyo, antoneyo = enija, enija = rotateLeft(mccade, 30), mccade = breylen, breylen = kanak;
  }
  ;
  for (jovi = 60; 80 > jovi; jovi++) {
    kanak = rotateLeft(breylen, 5) + (mccade ^ enija ^ antoneyo) + thomesa + naweed[jovi] + garen & 4294967295, thomesa = antoneyo, antoneyo = enija, enija = rotateLeft(mccade, 30), mccade = breylen, breylen = kanak;
  }
  ;
  sashe.A = sashe.A + breylen & 4294967295, sashe.B = sashe.B + mccade & 4294967295, sashe.C = sashe.C + enija & 4294967295, sashe.D = sashe.D + antoneyo & 4294967295, sashe.E = sashe.E + thomesa & 4294967295;
}






function rotateLeft(amayla, wynette) {
  var yulianna = amayla >>> 32 - wynette, gaviota = (1 << 32 - wynette) - 1, sparky = amayla & gaviota;
  return sparky << wynette | yulianna;
}








function MGF(makeya, yassmin) {
  if (yassmin > 4096) {
    return null;
  }
  ;
  var cathren = makeya.slice(0), zackarie = cathren.length;
  cathren[zackarie++] = 0, cathren[zackarie++] = 0, cathren[zackarie++] = 0, cathren[zackarie] = 0;
  for (var kinita = 0, salonge = []; salonge.length < yassmin;) {
    cathren[zackarie] = kinita++, salonge = salonge.concat(SHA1(cathren));
  }
  ;
  return salonge.slice(0, yassmin);
}







function XORarrays(ryver, ashayla) {
  if (ryver.length != ashayla.length) {
    return null;
  }
  ;
  for (var brendolyn = [], georgemichael = ryver.length, skky = 0; georgemichael > skky; skky++) {
    brendolyn[skky] = ryver[skky] ^ ashayla[skky];
  }
  ;
  return brendolyn;
}


module.exports = {
	parseRSAKeyFromString,
	hexStringToMP,
	JSMPnumber, 
	RSAEncrypt, 
	RSAEncryptBlock, 
	modularExp, 
	normalizeJSMP, 
	duplicateMP, 
	normalizeJSMP, 
	divideMP, 
	// multiplyAndSubtract
	// removeLeadingZeroes
	// modularMultiply
	// multiplyMP
	// mpToByteArray
	// byteArrayToMP
	// byteArrayToBase64
	// mapByteToBase64
	// base64Encode
	// applyPKCSv2Padding
	// SHA1
	// wordToBytes
	// PadSHA1Input
	// SHA1RoundFunction
	// rotateLeft
	// MGF
	// XORarrays
}
