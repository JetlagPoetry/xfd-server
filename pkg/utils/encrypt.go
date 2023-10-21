package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pKCS7UnPadding(origData []byte, blockSize int) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length == 0 {
		return nil, errors.New("pkcs7: Data is empty")
	}
	if length%blockSize != 0 {
		return nil, errors.New("pkcs7: Data is not block-aligned")
	}
	ref := bytes.Repeat([]byte{byte(unpadding)}, unpadding)
	if unpadding > blockSize || unpadding == 0 || !bytes.HasSuffix(origData, ref) {
		return nil, errors.New("pkcs7: Invalid padding")
	}
	return origData[:(length - unpadding)], nil
}

func AESEcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func AESEcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(generateKey(key))
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func AESEcbEncodeEncrypt(rawText, key string) string {
	encryptText := AESEcbEncrypt([]byte(rawText), generateKey([]byte(key)))
	return base64.StdEncoding.EncodeToString(encryptText)
}

func AESEcbDecodeDecrypt(rawText, key string) string {
	decodeText, _ := base64.StdEncoding.DecodeString(rawText)
	decryptText := AESEcbDecrypt(decodeText, generateKey([]byte(key)))
	return string(decryptText)
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 32)
	copy(genKey, key)
	for i := 32; i < len(key); {
		for j := 0; j < 32 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func AesDecrypt(data []byte, key []byte, iv string) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pKCS7UnPadding(crypted, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// DecryptByAes Aes 解密
func DecryptByAes(data, pwdKey, iv string) ([]byte, error) {
	key := []byte(pwdKey)
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, key, iv)
}
