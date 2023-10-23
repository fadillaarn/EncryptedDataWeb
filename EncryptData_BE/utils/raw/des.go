package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

type UserData struct {
	Name     string
	Email    string
	Phone    string
	Password string
	Address  string
	Role     string
}

func main() {
	// key := []byte("mysecretPasswordkeySiz24")
	key, err := generateRandomKey(24)
	if err != nil {
		fmt.Errorf("Error generating random key: %s", err)
		panic(err)
	}
	userData := UserData{
		Name:     "John Doe",
		Email:    "john@example.com",
		Phone:    "1234567890",
		Password: "supersecret",
		Address:  "123 Main St",
		Role:     "user",
	}

	// Encrypt individual fields with time measurement
	// encryptedName, timeEncryptName := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Name)
	// })

	// encryptedEmail, timeEncryptEmail := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Email)
	// })

	// encryptedPhone, timeEncryptPhone := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Phone)
	// })

	// encryptedPassword, timeEncryptPassword := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Password)
	// })

	// encryptedAddress, timeEncryptAddress := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Address)
	// })

	// encryptedRole, timeEncryptRole := MeasureTime(func() string {
	// 	return EncryptField(key, userData.Role)
	// })
	encryptedName := EncryptField(key, userData.Name)
	encryptedEmail := EncryptField(key, userData.Email)
	encryptedPhone := EncryptField(key, userData.Phone)
	encryptedPassword := EncryptField(key, userData.Password)
	encryptedAddress := EncryptField(key, userData.Address)
	encryptedRole := EncryptField(key, userData.Role)

	// Repeat the same for other fields...

	fmt.Println("Encrypted User Data:")
	fmt.Printf("Encrypted Name: %s\n", encryptedName)
	fmt.Printf("Encrypted Email: %s\n", encryptedEmail)
	fmt.Printf("Encrypted Phone: %s\n", encryptedPhone)
	fmt.Printf("Encrypted Password: %s\n", encryptedPassword)
	fmt.Printf("Encrypted Address: %s\n", encryptedAddress)
	fmt.Printf("Encrypted Role: %s\n", encryptedRole)
	
	// Print time elapsed for encryption
	// fmt.Printf("Time elapsed for Name encryption: %s\n", timeEncryptName.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Email encryption: %s\n", timeEncryptEmail.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Phone encryption: %s\n", timeEncryptPhone.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Password encryption: %s\n", timeEncryptPassword.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Address encryption: %s\n", timeEncryptAddress.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Role encryption: %s\n", timeEncryptRole.Round(time.Microsecond))

	// Decrypt individual fields with time measurement
	// decryptedName, timeDecryptName := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedName)
	// })

	// decryptedEmail, timeDecryptEmail := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedEmail)
	// })
	// decryptedPhone, timeDecryptPhone := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedPhone)
	// })
	// decryptedPassword, timeDecryptPassword := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedPassword)
	// })
	// decryptedAddress, timeDecryptAddress := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedAddress)
	// })
	// decryptedRole, timeDecryptRole := MeasureTime(func() string {
	// 	return DecryptField(key, encryptedRole)
	// })

	// Decrypt individual fields
	decryptedName := DecryptField(key, encryptedName)
	decryptedEmail := DecryptField(key, encryptedEmail)
	decryptedPhone := DecryptField(key, encryptedPhone)
	decryptedPassword := DecryptField(key, encryptedPassword)
	decryptedAddress := DecryptField(key, encryptedAddress)
	decryptedRole := DecryptField(key, encryptedRole)

	// Repeat the same for other fields...

	fmt.Println("Decrypted User Data:")
	fmt.Printf("Decrypted Name: %s\n", decryptedName)
	fmt.Printf("Decrypted Email: %s\n", decryptedEmail)
	fmt.Printf("Decrypted Phone: %s\n", decryptedPhone)
	fmt.Printf("Decrypted Password: %s\n", decryptedPassword)
	fmt.Printf("Decrypted Address: %s\n", decryptedAddress)
	fmt.Printf("Decrypted Role: %s\n", decryptedRole)
	// Print time elapsed for decryption
	// fmt.Printf("Time elapsed for Name decryption: %s\n", timeDecryptName.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Email decryption: %s\n", timeDecryptEmail.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Phone decryption: %s\n", timeDecryptPhone.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Password decryption: %s\n", timeDecryptPassword.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Address decryption: %s\n", timeDecryptAddress.Round(time.Microsecond))
	// fmt.Printf("Time elapsed for Role decryption: %s\n", timeDecryptRole.Round(time.Microsecond))

}

func generateRandomKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func EncryptField(key []byte, data string) string {
	// Generate a random IV
	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Errorf("Error generating IV: %s", err)
		panic(err)
	}

	c, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Errorf("NewTripleDESCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}

	// Calculate the required padding
	padding := 8 - len(data)%8 // DES block size is 8 bytes

	// Append the padding
	for i := 0; i < padding; i++ {
		data += string(byte(padding))
	}

	// Create a CBC mode block cipher with the IV
	mode := cipher.NewCBCEncrypter(c, iv)

	out := make([]byte, len(data))
	mode.CryptBlocks(out, []byte(data))

	// Prepend the IV to the ciphertext
	ciphertext := append(iv, out...)

	return base64.StdEncoding.EncodeToString(ciphertext)
	// cipher.NewCBCEncrypter(c, key[:8]).CryptBlocks(out, []byte(data))
	// return base64.StdEncoding.EncodeToString(out)
}

func DecryptField(key []byte, ct string) string {
	ctBytes, err := base64.StdEncoding.DecodeString(ct)
	if err != nil {
		fmt.Errorf("Base64 decode error: %s", err)
		panic(err)
	}

	c, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Errorf("NewTripleDESCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}

	// Extract the IV from the first block
	iv := ctBytes[:des.BlockSize]
	ctBytes = ctBytes[des.BlockSize:]

	// Create a CBC mode block cipher with the IV
	mode := cipher.NewCBCDecrypter(c, iv)

	plain := make([]byte, len(ctBytes))
	// cipher.NewCBCDecrypter(c, key[:8]).CryptBlocks(plain, ctBytes)
	mode.CryptBlocks(plain, ctBytes)
	decryptedData := string(plain[:])
	// Remove padding
	padding := int(decryptedData[len(decryptedData)-1])
	return decryptedData[:len(decryptedData)-padding]
}

// MeasureTime is a function that measures the time taken to execute a function and returns the result and time elapsed.
func MeasureTime(fn func() string) (result string, elapsed time.Duration) {
	start := time.Now()
	result = fn()
	elapsed = time.Since(start)
	return result, elapsed
}