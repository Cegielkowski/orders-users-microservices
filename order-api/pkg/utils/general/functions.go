package general

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/errors"
	"os"
	"strconv"
)

// Stringify takes a uint32 and transform in string.
func Stringify(number uint32) string {
	return strconv.FormatUint(uint64(number), 10)
}

// ReadConfig Reads Settings file.
func ReadConfigJson(configStruct interface{}, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configStruct)
	if err != nil {
		return err
	}

	return nil
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		errors.Log(err)
		return []byte(s)
	}
	return data
}

// Encrypt a string.
func Encrypt(key []byte, text string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		errors.Log(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}

// Decrypt a string.
func Decrypt(key []byte, text string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		errors.Log(err)
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}
