package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//	loggedIn := false
	loggedUsers := make(map[int64]bool)
	var adminKey string
	//factoryKey := "mySecretKey"

	fmt.Println("Ingrese la clave de administrador:")
	fmt.Scanln(&adminKey)

	fmt.Println("Inicializando...")
	config := InitSecurity(adminKey)

	fmt.Println("Leyendo claves...")
	fmt.Println("TelegramKey: ", config.TelegramKey)
	fmt.Println("GcbaClientId: ", config.GcbaClientId)
	fmt.Println("GcbaClientSecret: ", config.GcbaClientSecret)
	fmt.Println("OtpSecret: ", config.OtpSecret)

	bot, err := tgbotapi.NewBotAPI(config.TelegramKey)

	//secretKey := os.Getenv("SECRET_KEY")

	bot.Debug = true

	log.Printf("Autorizado para la cuenta %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//InitOtp()

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignorar actualizaciones que no sean mensajes
			continue
		}

		if !update.Message.IsCommand() { // ignorar mensajes que no sean comandos
			continue
		}

		// Crear un nuevo MessageConfig vacio
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extraer el comando del mensaje
		switch update.Message.Command() {

		case "screenshot":
			msg.Text = ""
			photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("img/screenshot.jpeg"))
			if _, err = bot.Send(photo); err != nil {
				log.Fatalln(err)
			}

		case "bitcoin":
			//if loggedIn {
			if value, exists := loggedUsers[update.Message.From.ID]; exists && value {
				msg.Text = GetBitcoinPrice()
			} else {
				msg.Text = "Acceso no autorizado"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		case "subway":
			/*if value, exists := loggedUsers[update.Message.From.ID]; exists && value {*/
			msg.Text = SubwayAlerts(config.GcbaClientId, config.GcbaClientSecret)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			/*}*/

		case "bus":
			/*if value, exists := loggedUsers[update.Message.From.ID]; exists && value {*/
			msg.Text = BusAlerts(config.GcbaClientId, config.GcbaClientSecret)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			/*}*/

		case "bikes":
			//if loggedIn {
			/*if value, exists := loggedUsers[update.Message.From.ID]; exists && value {*/
			msg.Text = StationInfo(config.GcbaClientId, config.GcbaClientSecret, 21)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			/*}*/

		case "weather":
			//if loggedIn {
			/*if value, exists := loggedUsers[update.Message.From.ID]; exists && value {*/
			msg.Text = WeatherReport("Buenos Aires")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			/*}*/

		default:
			inputOtp := update.Message.Command()
			if ValidateOtp(inputOtp, config.OtpSecret) {
				loggedUsers[update.Message.From.ID] = true
				msg.Text = "Bienvenido " + update.Message.From.FirstName + "(ID " + strconv.FormatInt(update.Message.From.ID, 10) + ")"
				fmt.Println("OTP valido")
			} else {
				loggedUsers[update.Message.From.ID] = false
				msg.Text = "OTP no valido"
				fmt.Println("OTP no valido")
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		}

	}

}
