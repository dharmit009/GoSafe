package passutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// the mpassphrase get pushed here in order to generate hash for cipher.
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(passtoencrypt []byte, mpassphrase string) []byte {
	// creates aes cipher
	block, _ := aes.NewCipher([]byte(createHash(mpassphrase)))
	// Galois Counter Mode
	gcm, _ := cipher.NewGCM(block)
	// nonce is a fixed set of characters which are appended to the
	// cipher text.
	nonce := make([]byte, gcm.NonceSize())
	// randomness needed for ciphering text
	io.ReadFull(rand.Reader, nonce)
	// ciphering the plain text
	ciphertext := gcm.Seal(nonce, nonce, passtoencrypt, nil)
	//returning the cipher text
	return ciphertext
}

func decrypt(passtodecrypt []byte, mpassphrase string) []byte {
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
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	file, _ := os.Create(filename)
	defer file.Close()
	file.Write(encrypt(data, mpassphrase))
}
func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

// func main() {
//
// 	mpass := "pass@@"
//
// 	ciphertext := encrypt([]byte("Hello World"), mpass)
// 	fmt.Println("Ciphered Text: ", string(ciphertext))
//
// 	plaintext := decrypt(ciphertext, mpass)
// 	fmt.Println("Plain Text: ", string(plaintext))
// }
