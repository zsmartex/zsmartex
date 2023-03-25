package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

var key = []byte("password")

func getSalt() []byte {
	hasher := md5.New()
	hasher.Write(key)
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}

func Encrypt(value string) string {
	block, _ := aes.NewCipher(getSalt())
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(value), nil)
	base64Str := base64.StdEncoding.EncodeToString(ciphertext)

	return base64Str
}

func Decrypt(value string) string {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher(getSalt())
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
