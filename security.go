package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const keySize = 32

var hmacKey []byte
var readData BotConfiguration

const (
	encryptedFile = "data.json.enc"
)

type BotConfiguration struct {
	TelegramKey      string `json:"telegramkey"`
	GcbaClientId     string `json:"gcbaclientid"`
	GcbaClientSecret string `json:"gcbaclientsecret"`
	OtpSecret        string `json:"otpsecret"`
}

func InitSecurity(adminKey string) *BotConfiguration {
	// Verificar si el archivo ya existe
	if !verifyKeyStore() {
		// Si el archivo no existe, lo creamos
		fmt.Println("Creating key store")
		if err := createKeyStore(adminKey); err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return nil
		}
	}

	jsonData := ReadDecrypt(adminKey, encryptedFile)
	err := json.Unmarshal(jsonData, &readData)
	if err != nil {
		log.Fatal(err)
	}
	return &readData

}

func createKeyStore(encryptionKey string) error {
	var telegramToken string
	var clientId string
	var clientSecret string

	fmt.Println("Ingrese el token de Telegram:")
	fmt.Scanln(&telegramToken)

	fmt.Println("Ingrese el client ID de GCBA:")
	fmt.Scanln(&clientId)

	fmt.Println("Ingrese el client secret de GCBA:")
	fmt.Scanln(&clientSecret)

	otpKey := GenerateOtpKey("Laboratorio de Algoritmos y Estructuras de Datos", "Muestra 2023")

	data := BotConfiguration{
		TelegramKey:      telegramToken,
		GcbaClientId:     clientId,
		GcbaClientSecret: clientSecret,
		OtpSecret:        otpKey,
	}

	// Serializar la estructura a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Guardar el JSON en un archivo
	/*	err = os.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	*/
	encryptSave(jsonData, encryptionKey, encryptedFile)
	return nil
}

func verifyKeyStore() bool {
	// Verificar si el archivo existe
	_, err := os.Stat(encryptedFile)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

/*
package main

import (

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

)

	func main() {
		userKey := "mySecretKey"
		plainData := []byte("Hello, World!")
		filename := "encrypted_data.bin"

		// Encriptar y guardar los datos en un archivo
		encrypt(plainData, userKey, filename)

		// Desencriptar los datos del archivo
		decryptedData := decrypt(userKey, filename)
		fmt.Println("Decrypted Data:", string(decryptedData))
	}
*/
func encryptSave(plainData []byte, userKey string, filename string) {
	// Calcular el hash SHA-256 de la clave del usuario
	userKeyHash := sha256.Sum256([]byte(userKey))
	aesKey := userKeyHash[:]

	// Crear un nuevo archivo para escribir los datos encriptados
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Crear un cifrador AES con la clave
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		log.Fatal(err)
	}

	// Generar un IV (Vector de Inicialización) aleatorio
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	// Escribir el IV en el archivo
	if _, err := file.Write(iv); err != nil {
		log.Fatal(err)
	}

	// Crear un modo de cifrado en bloque con el cifrador y el IV
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encriptar y escribir los datos en el archivo
	stream.XORKeyStream(plainData, plainData)
	if _, err := file.Write(plainData); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data encrypted and saved to", filename)
}

func ReadDecrypt(userKey string, filename string) []byte {
	// Calcular el hash SHA-256 de la clave del usuario
	userKeyHash := sha256.Sum256([]byte(userKey))
	aesKey := userKeyHash[:]

	// Leer el archivo cifrado
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Extraer el IV del archivo
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Crear un cifrador AES con la clave
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		log.Fatal(err)
	}

	// Crear un modo de cifrado en bloque con el cifrador y el IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// Desencriptar los datos
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}
