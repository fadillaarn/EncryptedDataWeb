package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

// UserData struct to hold user data
type UserData struct {
	Name     string
	Email    string
	Phone    string
	Password string
	Address  string
	Role     string
}

func main() {
	// Dummy user data
	userData := UserData{
		Name:     "John Doe",
		Email:    "john@example.com",
		Phone:    "1234567890",
		Password: "supersecret",
		Address:  "123 Main St",
		Role:     "user",
	}

	userDataMap := map[string]string{
		"Name":     userData.Name,
		"Email":    userData.Email,
		"Phone":    userData.Phone,
		"Password": userData.Password,
		"Address":  userData.Address,
		"Role":     userData.Role,
	}
	// Generate a random key and IV for AES-256-CBC
	key, iv, err := generateRandomAESKeyAndIV()
	if err != nil {
		fmt.Println("Error generating random AES key and IV:", err)
		return
	}

	// keyHex := hex.EncodeToString(key)
	// ivHex := hex.EncodeToString(iv)
	// fmt.Println("AES Key (Hex):", keyHex)
	// fmt.Println("AES IV (Hex):", ivHex)

	fields := []string{"Name", "Email", "Phone", "Password", "Address", "Role"}
	// start := time.Now()
	// Map for encrypted data
	encryptedData := make(map[string][]byte)
	for _, fieldName := range fields {
		fieldValue := userDataMap[fieldName]
		encryptedValue, err := AESEncrypted(fieldValue, key, iv)
		if err != nil {
			fmt.Printf("Error during %s encryption: %v\n", fieldName, err)
		} else {
			encryptedData[fieldName] = []byte(encryptedValue)
			fmt.Printf("Encrypted %s: %s\n", fieldName, encryptedValue)
		}
	}
	// elapsed := time.Since(start)
	// elapsedSeconds := float64(elapsed.Microseconds()) / 1000000.0 // 1 million microseconds = 1 second
	// fmt.Printf("Total time for encrypt is: %.20f seconds using \n", elapsedSeconds)
	// TotalTime = TotalTime
	// Map for decrypted data
	decryptedData := make(map[string]string)
	for _, fieldName := range fields {
		encryptedValue := encryptedData[fieldName]
		decryptedValue, err := AESDecrypted(string(encryptedValue), key, iv)
		if err != nil {
			fmt.Printf("Error during %s decryption: %v\n", fieldName, err)
		} else {
			decryptedData[fieldName] = decryptedValue
			fmt.Printf("Decrypted %s: %s\n", fieldName, decryptedValue)
		}
	}

	// _, encryptTime := MeasureTime(func() string {
	// 	for fieldName, fieldValue := range userDataMap {
	// 		encryptedValue, err := AESEncrypted(fieldValue, key, iv)
	// 		if err != nil {
	// 			fmt.Printf("Error during %s encryption: %v\n", fieldName, err)
	// 		} else {
	// 			encryptedData[fieldName] = []byte(encryptedValue)
	// 			fmt.Printf("Encrypted %s: %s\n", fieldName, encryptedValue)
	// 		}
	// 	}
	// 	return ""
	// })
	// fmt.Println("Total Encryption Time:", encryptTime)

	// _, decryptTime := MeasureTime(func() string {
	// 	for fieldName, encryptedValue := range encryptedData {
	// 		decryptedValue, err := AESDecrypted(string(encryptedValue), key, iv)
	// 		if err != nil {
	// 			fmt.Printf("Error during %s decryption: %v\n", fieldName, err)
	// 		} else {
	// 			decryptedData[fieldName] = decryptedValue
	// 			fmt.Printf("Decrypted %s: %s\n", fieldName, decryptedValue)
	// 		}
	// 	}
	// 	return ""
	// })
	// fmt.Println("Total Decryption Time:", decryptTime)
}

func generateRandomAESKeyAndIV() (key []byte, iv []byte, err error) {
	key = make([]byte, 32) // 32 bytes for AES-256
	iv = make([]byte, aes.BlockSize)
	_, err = rand.Read(key)
	if err != nil {
		return nil, nil, err
	}
	_, err = rand.Read(iv)
	if err != nil {
		return nil, nil, err
	}
	
	return key, iv, nil
}

// source : https://medium.com/@thanat.arp/encrypt-decrypt-aes256-cbc-shell-script-golang-node-js-ffb675a05669
// AESEncrypted encrypts given text in AES 256 CBC
func AESEncrypted(plaintext string, key []byte, iv []byte) (string, error) {
	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{byte(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}

// AESDecrypted decrypts given text in AES 256 CBC
func AESDecrypted(encrypted string, key []byte, iv []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("invalid block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return string(ciphertext), nil
}

// PKCS5UnPadding pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return src
	}
	unpadding := int(src[length-1])
	if unpadding >= length {
		return src
	}
	return src[:length-unpadding]
}

func MeasureTime(fn func() string) (result string, elapsed time.Duration) {
	start := time.Now()
	result = fn()
	elapsed = time.Since(start)
	return result, elapsed
}