package main

import (
	"log"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

const appKey = "estaesunaclavede32bytes123456"
const otpFileName = "otp.key"
const qrFileName = "qrcode.png"
const appName = "Bot Muestra"
const accountName = "Laboratorio de Algoritmos y Estructuras de Datos"

func generateQRCode(text string) error {
	qrcode, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return err
	}

	return qrcode.WriteFile(256, qrFileName)
}

/*func InitOtp() {
	if archivoExiste(otpFileName) && archivoExiste(otpFileName) {
		return
	} else {
		GenerateOtpKey(appName, accountName)
	}
}*/

func GenerateOtpKey(appname string, accountname string) string {
	// Crear una nueva clave secreta OTP
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appname,
		AccountName: accountname,
	})
	if err != nil {
		log.Fatal("Error al generar la clave secreta OTP:", err)
	}

	// Generar y mostrar el código QR
	err = generateQRCode(key.URL())
	if err != nil {
		log.Fatal("Error al generar el código QR:", err)
	}
	fmt.Println("Escanee el QR utilizando Google Authenticator")
	//guardarTextoCifradoEnArchivo(otpFileName, key.Secret())
	return key.Secret()
}

func ValidateOtp(inputOTP string, secretKey string) bool {
	//secretKey, _ := leerTextoCifradoDesdeArchivo(otpFileName)
	return totp.Validate(inputOTP, secretKey)
}

func encrypt(texto string, clave []byte) ([]byte, error) {
	block, err := aes.NewCipher(clave)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(texto))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(texto))

	return ciphertext, nil
}

func decrypt(ciphertext []byte, clave []byte) (string, error) {
	block, err := aes.NewCipher(clave)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("El texto cifrado es demasiado corto")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func guardarTextoCifradoEnArchivo(encryptedFileName string, texto string) error {
	ciphertext, err := encrypt(texto, []byte(appKey))
	if err != nil {
		return err
	}

	archivo, err := os.Create(encryptedFileName)
	if err != nil {
		return err
	}
	defer archivo.Close()

	encoder := base64.NewEncoder(base64.StdEncoding, archivo)
	_, err = encoder.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}

func leerTextoCifradoDesdeArchivo(encryptedFileName string) (string, error) {
	archivo, err := os.Open(encryptedFileName)
	if err != nil {
		return "", err
	}
	defer archivo.Close()

	decoder := base64.NewDecoder(base64.StdEncoding, archivo)
	ciphertext, err := io.ReadAll(decoder)
	if err != nil {
		return "", err
	}

	texto, err := decrypt(ciphertext, []byte(appKey))
	if err != nil {
		return "", err
	}

	return texto, nil
}

func archivoExiste(nombreArchivo string) bool {
	_, err := os.Stat(nombreArchivo)
	if err == nil {
		// El archivo existe
		return true
	}
	if os.IsNotExist(err) {
		// El archivo no existe
		return false
	}
	// Ocurrió un error al verificar la existencia del archivo
	return false
}
