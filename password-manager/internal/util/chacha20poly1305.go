package util

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

// O ChaCha20-Poly1305 exige uma chave de exatamente 32 bytes
func DeriveKey(masterPassword string) []byte {
	sum := sha256.Sum256([]byte(masterPassword))
	return sum[:chacha20poly1305.KeySize]
}

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("criar cifra: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("gerar nonce: %w", err)
	}

	return aead.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("criar cifra: %w", err)
	}

	nonceSize := aead.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("texto cifrado muito curto")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("descriptografar dados: %w", err)
	}

	return plaintext, nil
}
