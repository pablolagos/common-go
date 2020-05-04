package binaryutils

import (
	"bytes"
	"testing"
)

func Test_encrypt_decrypt(t *testing.T) {

	/* Encrypt/decrypt */
	t.Log("Encryption/Decryption test")
	t.Log("==========================")
	key := []byte("clave secreta")
	a := []byte("{}")
	encrypted, err := Encrypt(key, a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Encrypted:", encrypted)
	b, err := Decrypt(key, encrypted)
	if err != nil {
		t.Error(err)
		return
	}
	if bytes.Compare(a, b) != 0 {
		t.Error("a != b")
		return
	}
	t.Log("Encryption/Decription Successfull")

	/* empty key */

	t.Log("")
	t.Log("Empty Key")
	t.Log("=========")
	var emptyKey []byte
	_, err = Encrypt(emptyKey, a)
	if err == nil {
		t.Error("Not error")
		return
	}
	t.Log("Test OK")

	t.Log("")
	t.Log("Empty Data")
	t.Log("=========")
	var emptyData []byte
	encrypted, err = Encrypt(key, emptyData)
	t.Log(encrypted, err)
	b, err = Decrypt(key, encrypted)
	t.Log(b, err)

}
