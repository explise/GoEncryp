package handlers


import (
	"log"
	"os"
	"crypto/aes"
)


func Encrypt(filename string, password string) {
	// Encrypt the file
	file , err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file
	data := make([]byte,1000)
	n , err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}
	cipherText := gcm.Seal(nonce, nonce, data[:n], nil)
	
	file.Seek(0,0)
	file.Write(cipherText)
}