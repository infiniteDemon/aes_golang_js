package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {

	StrKey := "SC1Km1LFVknhbn42YAtxXIY3XB5333RD"
	StrIv := "qoZgNGR9aVg3Ldnq"

	encKey := enHex(StrKey)
	log.Printf("encKey %s", encKey)
	iv := enHex(StrIv)
	log.Printf("iv %s", iv)
	cipherTemp, err := AesEncrypt(`亲爱的，你还好吗?`, deHex(encKey), deHex(iv))
	if err != nil {
		panic(err)
	}
	log.Printf("%s", cipherTemp)

	body, err := AesDecrypt(cipherTemp, deHex(encKey), deHex(iv))
	if err != nil {
		panic(err)
	}
	log.Printf("%s", body)

}

func deHex(text string) []byte {
	decoded, err := hex.DecodeString(text)
	if err != nil {
		panic(err)
	}
	return decoded
}

func enHex(text string) string {
	enStr := hex.EncodeToString([]byte(text))
	return enStr
}

// PKCS5填充方式
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

// Zero填充方式
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{0}, padding)

	return append(ciphertext, padtext...)
}

// PKCS5 反填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Zero反填充
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 加密
func AesEncrypt(encodeStr string, key []byte, iv []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	//encodeBytes = ZeroPadding(encodeBytes, blockSize

	encodeBytes = PKCS5Padding(encodeBytes, blockSize) //PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	hexstr := fmt.Sprintf("%x", crypted)
	return hexstr, nil
	//return base64.StdEncoding.EncodeToString(crypted), nil
}

// 解密
func AesDecrypt(decodeStr string, key []byte, iv []byte) ([]byte, error) {
	//decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)//先解密base64
	decodeBytes, err := hex.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	//origData = ZeroUnPadding(origData) // origData = PKCS5UnPadding(origData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
