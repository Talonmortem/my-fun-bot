package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	gigachat "github.com/saintbyte/gigachat_api"
)

func gchat(input string) string {
	chat := gigachat.NewGigachat()
	aswer, err := chat.Ask(input)
	if err != nil {
		log.Fatal(err)
	}
	return aswer
}

func gchatImage(input string) string {

	url := "https://gigachat.devices.sberbank.ru/api/v1/files/:file_id/content"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Accept", "image/jpg")
	gigatocken := os.Getenv("GIGACHAT_AUTH_DATA")
	req.Header.Add("Authorization", "Bearer "+gigatocken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return (string(body))
}
