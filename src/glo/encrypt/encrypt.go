package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// AesEncryptString 用户密码加密函数
func AesEncryptString(orginString, aeskey []byte) (string, error) {
	xpass, err := AesEncrypt(orginString, aeskey)
	if err != nil {
		return "", err
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	return pass64, nil
}

// AesDecryptString 用户密码解密函数
func AesDecryptString(pass64 string, aeskey []byte) ([]byte, error) {
	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		return nil, err
	}
	return tpass, nil
}
