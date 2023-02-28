package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// the mpassphrase get pushed here in order to generate hash for cipher.
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(passtoencrypt []byte, mpassphrase string) ([]byte, error) {
	// creates aes cipher
	block, err := aes.NewCipher([]byte(createHash(mpassphrase)))
	check(err)
	// Galois Counter Mode
	gcm, err := cipher.NewGCM(block)
	check(err)
	// nonce is a fixed set of characters which are appended to the
	// cipher text.
	nonce := make([]byte, gcm.NonceSize())
	// randomness needed for ciphering text
	io.ReadFull(rand.Reader, nonce)
	// ciphering the plain text
	ciphertext := gcm.Seal(nonce, nonce, passtoencrypt, nil)
	//returning the cipher text
	return ciphertext, nil
}

func decrypt(passtodecrypt []byte, mpassphrase string) ([]byte, error) {
	// creates aes cipher
	block, _ := aes.NewCipher([]byte(createHash(mpassphrase)))
	// galieous control mode
	gcm, _ := cipher.NewGCM(block)
	// getting nonceSize to separate encrypted passtodecrypt from cipher text.
	nonceSize := gcm.NonceSize()
	// sperating nonce text from ciphertext
	nonce, ciphertext := passtodecrypt[:nonceSize], passtodecrypt[nonceSize:]
	// ciphering the plain text
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	//returning the cipher text
	return plaintext, nil
}

func encryptFile(filename string, data []byte, passphrase string) {
	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	ciphertext, err := encrypt(data, passphrase)
	check(err)

	fmt.Println("Ciphered Text: ", string(ciphertext))

	_, err = file.Write(ciphertext)
	check(err)
}

func decryptFile(filename string, passphrase string) ([]byte, error) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	stat, err := file.Stat()
	check(err)

	ciphertext := make([]byte, stat.Size())
	_, err = file.Read(ciphertext)
	check(err)

	plaintext, err := decrypt(ciphertext, passphrase)
	check(err)

	return plaintext, nil
}

// func main() {
//
// 	mpass := "pass@@"
// 	var hash string = createHash(mpass)
//
// 	ciphertext, _ := encrypt([]byte(mpass), hash)
// 	fmt.Println("Ciphered Text: ", string(ciphertext))
//
// 	plaintext, _ := decrypt([]byte(ciphertext), hash)
// 	fmt.Println("Plain Text: ", string(plaintext))
//
// 	encryptFile("mpass.enc", []byte(mpass), mpass)
// 	out, _ := decryptFile("mpass.enc", mpass)
//
// 	fmt.Println("Out: ", string(out))
// }
