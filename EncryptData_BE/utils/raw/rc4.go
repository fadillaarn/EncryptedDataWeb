package main

import (
	"crypto/rand"
	"crypto/rc4"
	"fmt"
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

	// Generate a random RC4 key
	key, err := GenerateRandomRC4Key(32) // 32 bytes key
	if err != nil {
		fmt.Println("Error generating random RC4 key:", err)
		return
	}

	// Print RC4 key in hexadecimal format
	// keyHex := hex.EncodeToString(key)
	// fmt.Println("RC4 Key (Hex):", keyHex)

	// Encrypt and display each data field using RC4
	encryptedData := make(map[string][]byte)
	for fieldName, fieldValue := range map[string]string{
		"Name":     userData.Name,
		"Email":    userData.Email,
		"Phone":    userData.Phone,
		"Password": userData.Password,
		"Address":  userData.Address,
		"Role":     userData.Role,
	} {
		encryptedValue, err := EncryptRC4([]byte(fieldValue), key)
		if err != nil {
			fmt.Printf("Error during %s encryption: %v\n", fieldName, err)
		} else {
			encryptedData[fieldName] = encryptedValue
			fmt.Printf("Encrypted %s: %x\n", fieldName, encryptedValue)
		}
	}

	// Decrypt and display each data field using RC4
	decryptedData := make(map[string]string)
	for fieldName, encryptedValue := range encryptedData {
		decryptedValue, err := DecryptRC4(encryptedValue, key)
		if err != nil {
			fmt.Printf("Error during %s decryption: %v\n", fieldName, err)
		} else {
			decryptedData[fieldName] = string(decryptedValue)
			fmt.Printf("Decrypted %s: %s\n", fieldName, decryptedData[fieldName])
		}
	}
}

func GenerateRandomRC4Key(keyLength int) ([]byte, error) {
	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// EncryptRC4 encrypts data using RC4
func EncryptRC4(data []byte, key []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	return encrypted, nil
}

// DecryptRC4 decrypts data using RC4
func DecryptRC4(data []byte, key []byte) ([]byte, error) {
	// RC4 decryption is the same as encryption
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)
	
	return EncryptRC4(data, key)
}
