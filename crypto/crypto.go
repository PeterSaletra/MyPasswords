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

type Keys struct {
	master_hash string
	Master_key  string
}

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

func GenerateIV(length int) ([]byte, error) {
	if length == 0 {
		length = 16
	}

	iv := make([]byte, length)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("Error while generating IV for password: %w", err)
	}

	return iv, nil
}

func (keys *Keys) EncryptAESCBC(plaintext string, iv []byte) (string, error) {
	bKey := []byte(keys.master_hash)

	if len(iv) != aes.BlockSize {
		return "", fmt.Errorf("IV length must be %d bytes, got %d ,IV: %x", aes.BlockSize, len(iv), iv)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", fmt.Errorf("Error while encrypting password: %w", err)
	}

	bPlainText := []byte(plaintext)
	padding := aes.BlockSize - len(bPlainText)%aes.BlockSize
	for i := 0; i < padding; i++ {
		bPlainText = append(bPlainText, byte(padding))
	}

	ciphertext := make([]byte, len(bPlainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlainText)

	return hex.EncodeToString(ciphertext), nil
}

func (keys *Keys) DecryptAESCBC(ciphertext string, iv []byte) (string, error) {
	bKey := []byte(keys.master_hash)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", fmt.Errorf("Error while encrypting password: %w", err)
	}

	bCiphertext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("error decoding ciphertext: %w", err)
	}
	plaintext := make([]byte, len(bCiphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, bCiphertext)

	if len(plaintext) > 0 {
		padding := int(plaintext[len(plaintext)-1])
		if padding > 0 && padding <= aes.BlockSize && padding <= len(plaintext) {
			plaintext = plaintext[:len(plaintext)-padding]
		}
	}

	return string(plaintext), nil
}

func (keys *Keys) DeriveMasterHash(password string, username string) {
	keys.master_hash = deriveKeyPBKDF2(password, username, 100_000, 32)
}

func (keys *Keys) DeriveMasterKey(password string) {
	keys.Master_key = deriveKeyPBKDF2(keys.master_hash, password, 100_000, 32)
}
