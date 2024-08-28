package handlers

import(
	"log"
	"os"
	"crypto/aes"
)

func Decrypt(filename string, password string) {
	// Decrypt the file
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
	aes, err := aes.NewCipher([]byte(password))
	if err != nil {
		log.Fatal(err)
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatal(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText,err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0,0)
	file.Write(plainText)
}
