package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"strings"
)

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func PKCS7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func PKCS7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return b[:len(b)-n], nil
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	dec, _ := base64.StdEncoding.DecodeString(data)
	return string(dec)
}

func MD5Hash(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func AESEncrypt(key string, text string) (string, error) {
	/* Convert key hex  */
	cipherKey := []byte(key)
	/* Convert text */
	cipherText := []byte(text)
	/* Create new cipher */
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}
	/* Padding text */
	blockSize := block.BlockSize()
	origData, err := PKCS7Pad(cipherText, blockSize)
	if err != nil {
		return "", err
	}
	/* Create iv */
	cipherText = make([]byte, aes.BlockSize+len(origData))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	/* */
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText[aes.BlockSize:], origData)
	return string(cipherText), nil
}

func PasswordHash(password, salt string) string {
	h := sha512.New()
	h.Write([]byte(MD5Hash(password) + salt))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
