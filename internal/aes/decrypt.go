package aes

import (
	caes "crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// Decrypt takes a secret and a content string and returns the plaintext string.
// The ciphertext string is hex-decoded before decryption.
func Decrypt(secret, content string) (result string, err error) {
	var (
		ciphertext []byte
		c          cipher.Block
		key        = []byte(secret)
	)

	if ciphertext, err = hex.DecodeString(content); err != nil {
		return
	}
	if c, err = caes.NewCipher(key); err != nil {
		return
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	return string(pt[:]), nil
}
