package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	"golang.org/x/crypto/argon2"
)

func Hash(pw string) string {
	password := []byte(pw)
	s := generateSalt(16)
	h := argon2.IDKey(password, s, 1, 16*1024, 4, 32)
	hashPassword := base64.RawStdEncoding.EncodeToString(h)
	salt := base64.RawStdEncoding.EncodeToString(s)
	key := salt + ":" + hashPassword

	return key
}

func generateSalt(size int) []byte {
	salt := make([]byte, size)
	rand.Read(salt)
	return salt
}

func EqualWithKey(pw string, key string) bool {
	password := []byte(pw)
	k := strings.Split(key, ":")
	salt, _ := base64.RawStdEncoding.DecodeString(k[0])
	h := argon2.IDKey(password, salt, 1, 16*1024, 4, 32)
	if base64.RawStdEncoding.EncodeToString(h) == k[1] {
		return true
	}
	return false
}

func Encrypt(plaintext string) (string, string, error) {
	key := generateSalt(16)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// IV 초기화
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), base64.RawStdEncoding.EncodeToString(key), nil
}

func Decrypt(ciphertext string, key []byte) (string, error) {
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", err
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return string(ciphertextBytes), nil
}
