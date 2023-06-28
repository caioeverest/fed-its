package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// Encrypt takes a secret and content and encrypts the content with AES
func Encrypt(secret, content string) (result string, err error) {
	var (
		c         cipher.Block
		out       = make([]byte, len(content))
		plaintext = []byte(content)
		key       = []byte(secret)
	)
	if c, err = aes.NewCipher(key); err != nil {
		return
	}
	c.Encrypt(out, plaintext)
	return hex.EncodeToString(out), nil
}
