package stringutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
)

// Obtener el hash md5 hexadecimal de un string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Obtener el hash sha-256 hexadecimal de un string
func GetSHA256Hash(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Strrev(s string) string {

	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []byte(s)
	result := []byte{}

	// Add bytes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func Base64_encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64_decode(s string) string {
	var r []byte
	r, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	} else {
		return string(r)
	}
}

// Leer un archivo y devolver el contenido como string
func ReadFile(filename string) (string, error) {
	var c []byte
	var err error
	c, err = ioutil.ReadFile(filename)
	return string(c), err
}

func Encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = "!ENC$" + base64.URLEncoding.EncodeToString(cipherText)
	return
}

func Decrypt(key []byte, securemess string) (decodedmess string, err error) {

	if securemess[0:5] != "!ENC$" {
		err = errors.New("Not encrypted string")
		return
	}

	cipherText, err := base64.URLEncoding.DecodeString(securemess[5:])
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}

func Is_Encrypted(texto string) bool {
	return texto[0:5] == "!ENC$"
}
