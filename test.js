const CryptoJS = require("crypto-js")

const encKey = "5343314b6d314c46566b6e68626e343259417478584959335842353333335244"
const iv = "716f5a674e475239615667334c646e71"
const cipherText = "c5f4a99f2266bda3d101f3c89099eabdebe41bb4bd533114522b49e311e0" +
    "aab362bb3628706869d3183bbab7b193ca7b866c8b285cdd3eb2b361ac28b2fbbd1d0ae89c496d1b" +
    "7f8c4a778e66d5fe88d0"

var encrypt_req = function(key, iv, text) {
    var l = CryptoJS.enc.Utf8.parse(text);
    var e = CryptoJS.enc.Hex.parse(key);
    var a = CryptoJS.AES.encrypt(l, e, {
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
        iv: CryptoJS.enc.Hex.parse(iv)
    })
    return a.ciphertext.toString()
}

var decrypt_req = function(key, iv, text) {
    var e = CryptoJS.enc.Hex.parse(key);
    var WordArray = CryptoJS.enc.Hex.parse(text);
    var text = CryptoJS.enc.Base64.stringify(WordArray);
    var a = CryptoJS.AES.decrypt(text, e, {
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
        iv: CryptoJS.enc.Hex.parse(iv)
    });
    return CryptoJS.enc.Utf8.stringify(a).toString()
}


const txt = `亲爱的，你还好吗?`

const temp = encrypt_req(encKey, iv, txt);

console.log(temp);

console.log(decrypt_req(encKey, iv, temp));