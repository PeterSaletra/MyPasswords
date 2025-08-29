package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func deriveKeyPBKDF2(password string, salt string, iterations, keyLen int) string {
	if iterations == 0 {
		iterations = 100_000
	}

	if keyLen == 0 {
		keyLen = 32
	}
	derived_key := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLen, sha256.New)
	return string(derived_key)
}

func GenerateIV(lenght int) (string, error) {
	if lenght == 0 {
		lenght = 16
	}

	iv := make([]byte, lenght)
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("Error while generating VI for passowrd %w", err)
	}

	return hex.EncodeToString(iv), nil
}

// Maybe change it to (key* MasterKey) EncryptAESCBC ...
// TODO: Think about it maybe it makes more sens lol

func EncryptAESCBC(plaintext string, key string, iv string) (string, error) {
	bKey := []byte(key)
	bIV := []byte(iv)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", fmt.Errorf("Error while encrypting password: %w", err)
	}

	bPlainText := []byte(plaintext)
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlainText)

	return hex.EncodeToString(ciphertext), nil
}

// Maybe change it to (key* MasterKey) DecryptAESCBC ...
// TODO: Think about it maybe it makes more sens lol

func DecryptAESCBC(ciphertext string, key string, iv string) (string, error) {
	bKey := []byte(key)
	bIV := []byte(iv)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", fmt.Errorf("Error while encrypting password: %w", err)
	}

	bCiphertext := []byte(ciphertext)
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(plaintext, bCiphertext)

	return hex.EncodeToString(plaintext), nil
}

func DeriveMasterHash(password string, username string) string {
	return deriveKeyPBKDF2(password, username, 100_000, 32)
}

func DeriveMasterKey(master_hash string, password string) string {
	return deriveKeyPBKDF2(master_hash, password, 100_000, 32)
}
