package main

import (
	"fmt"

	"github.com/brandonspitz/Go-Discord-Bot-Template/bot"
	"github.com/brandonspitz/Go-Discord-Bot-Template/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
