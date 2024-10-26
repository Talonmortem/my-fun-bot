package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func telegramBot() {
	/*file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)*/
	var isdebug bool
	if os.Getenv("DEBUG") == "true" {
		isdebug = true
	} else {
		isdebug = false
	}

	log.Println("Telegram bot started")
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			// Check if the message is a command /giga
			if strings.HasPrefix(update.Message.Text, "/giga") {
				// Extract the chatid from the command
				gigaask := strings.TrimSpace(strings.Replace(update.Message.Text, "/giga", "", 1))
				// Check if the chatid is not empty
				if gigaask != "" {
					log.Printf("Asking ` %s `", gigaask)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, gchat(gigaask))
					bot.Send(msg)
				} else {
					log.Println("Invalid command: /giga запрос requires a запрос")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда.")
					bot.Send(msg)
				}
			}

			// Check if the message is a command /giga
			if strings.HasPrefix(update.Message.Text, "/image") {
				// Extract the chatid from the command
				gigaask := strings.TrimSpace(strings.Replace(update.Message.Text, "/image", "", 1))
				// Check if the chatid is not empty
				if gigaask != "" {
					log.Printf("Asking ` %s `", gigaask)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, gchatImage(gigaask))
					bot.Send(msg)
				} else {
					log.Println("Invalid command: /image запрос requires a запрос")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда.")
					bot.Send(msg)
				}
			}

			switch update.Message.Text {
			case "/gigatest":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, gchat("Кто ты? Ответь коротко."))
				bot.Send(msg)
			case "/start":
				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветствую! Я бот Evrohand, я могу искать информацию в базе данных Evrohand. Пожалуйста, предоставьте мне номер лота, которого вы хотите найти в Evrohand.")
				bot.Send(msg)

			case "/help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветствую! Я Evrohand, я могу искать информацию в базе данных Evrohand. Пожалуйста, предоставьте мне номер лота, которого вы хотите найти в Evrohand.")
				bot.Send(msg)
			case "/ping":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
				bot.Send(msg)
			case "/debug":
				if isdebug {
					isdebug = false
					log.Println("Debug mode off")
				} else {
					isdebug = true
					log.Println("Debug mode on")
				}
			case "/number_of_users":
				if os.Getenv("DB_SWITCH") == "on" {
					//Присваиваем количество пользоватьелей использовавших бота в num переменную
					num, err := getNumberOfUsers()
					if err != nil {
						//Отправлем сообщение
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error.")
						bot.Send(msg)
					}

					//Создаем строку которая содержит колличество пользователей использовавших бота
					ans := fmt.Sprintf("%d колличество пользователей использовавших бота для поиска.", num)

					//Отправлем сообщение
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
					bot.Send(msg)
				} else {

					//Отправлем сообщение
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database not connected, so i can't say you how many peoples used me.")
					bot.Send(msg)
				}
			default:
				log.Println("User: " + update.Message.From.UserName + "Usertext: " + update.Message.Text)
			}
		}
	}
}
