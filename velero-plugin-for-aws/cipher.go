package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(plainBytes []byte) ([]byte, error) {

	key  := []byte("aler,amz3daps.f9hgandkal4dsxk3d0")
        // Creating block of algorithm
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
        }

        // Creating GCM mode
    gcm, err := cipher.NewGCM(block)
    if err != nil {
		return nil, err
        }

        // Generating random nonce
    nonce := make([]byte, gcm.NonceSize())
	//nonce := make([]byte, t.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	//cipherBytes := t.gcm.Seal(nonce, nonce, plainBytes, nil)
	cipherBytes := gcm.Seal(nonce, nonce, plainBytes, nil)
	return cipherBytes, nil
}

	func Decrypt(cipherBytes []byte) ([]byte, error) {

	key  := []byte("aler,amz3daps.f9hgandkal4dsxk3d0")
    
        // Creating block of algorithm
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
        }

        // Creating GCM mode
    gcm, err := cipher.NewGCM(block)
    if err != nil {
		return nil, err
        }

        // Generating random nonce
	//nonceSize := t.gcm.NonceSize()
	nonceSize := gcm.NonceSize()
	nonce, cipherBytesWithoutNounce := cipherBytes[:nonceSize], cipherBytes[nonceSize:]

	//plainBytes, err := t.gcm.Open(nil, nonce, cipherBytesWithoutNounce, nil)
	plainBytes, err := gcm.Open(nil, nonce, cipherBytesWithoutNounce, nil)
	
	if err != nil {
		return nil, err
	}

	return plainBytes, nil
}
