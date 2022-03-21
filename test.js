const RsaModule = require("./rsa.js")



let password = "asjkdhjkashd0asdjaskdu899qqasdj0;23898"
let key = "4564654987-487987;454465489"


let n = RsaModule.PackagePassword(password)
let o = RsaModule.parseRSAKeyFromString(key)
// let s = RsaModule.RSAEncrypt(n, o)
console.log(s)