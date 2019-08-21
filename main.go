package main

import (
	"github.com/JellalLeonhardt/flowerss-bot/bot"
	"github.com/JellalLeonhardt/flowerss-bot/task"
)

func main() {
	go task.Update()
	bot.Start()
}
