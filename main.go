package main

import (
	"log"
	"os"
	"time"
)

func main() {

	log.Println("Version: 0.01")
	log.Println("Start bot...")

	time.Sleep(1 * time.Second)

	//Создаем таблицу
	if os.Getenv("CREATE_TABLE") == "yes" {
		if os.Getenv("DB_SWITCH") == "on" {
			log.Println("Create table...")
			if err := createTable(); err != nil {
				panic(err)
			}
		}
		log.Println("Table created")
	}

	time.Sleep(1 * time.Second)
	//Вызываем бота
	if os.Getenv("TELEGRAM_BOT") == "on" {
		log.Println("Telegram bot...")
		telegramBot()
	}
}
