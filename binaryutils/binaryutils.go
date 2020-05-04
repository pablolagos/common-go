package binaryutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Reverse(data []byte) []byte {

	result := []byte{}

	// Add bytes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	// Return new binary data.
	return result
}

func Encrypt(key []byte, message []byte) (result []byte, err error) {

	if len(key) == 0 {
		return result, errors.New("Empty key")
	}

	fixedkey := fixKey(key)

	block, err := aes.NewCipher(fixedkey)
	if err != nil {
		return result, err
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(message))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return result, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], message)

	result = []byte(base64.StdEncoding.EncodeToString(cipherText))
	return result, nil
}

func Decrypt(key []byte, encryptedMessage []byte) (message []byte, err error) {

	if len(key) == 0 {
		return message, errors.New("Empty key")
	}
	fixedkey := fixKey(key)

	cipherText, err := base64.StdEncoding.DecodeString(string(encryptedMessage))
	if err != nil {
		return message, err
	}

	block, err := aes.NewCipher(fixedkey)
	if err != nil {
		return message, err
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

	return cipherText, nil
}

func fixKey(key []byte) []byte {
	newkey := bytes.Repeat(key, 32/len(key)+1)
	return newkey[0:32]
}
