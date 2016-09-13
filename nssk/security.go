package nssk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
)

// Encrypt AES加密
func Encrypt(content interface{}, key string) string {
	plantText, err := json.Marshal(content)
	if err != nil {
		return ""
	}
	cipherbytes, err := encrypt(plantText, []byte(key))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(cipherbytes)
}

// Decrypt AES解密
func Decrypt(ciphertext string, key string) (interface{}, error) {
	cipherbytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, errors.New("解码失败，无法与服务器建立连接. " + err.Error())
	}
	plantText, err := decrypt(cipherbytes, []byte(key))
	if err != nil {
		return nil, errors.New("解码失败，无法与服务器建立连接. " + err.Error())
	}
	var content interface{}
	if err := json.Unmarshal(plantText, &content); err != nil {
		log.Println(err)
		return nil, err
	}
	return content, nil
}

func encrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = pkcs7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = pkcs7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func pkcs7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func pkcs7UnPadding(text []byte, blockSize int) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}
