package main

import (
	"twitch-viewer-bot/bot"
)

func main() {
	var i int8
	for i = 0; i < 1; i++ {
		go bot.OpenBot(i)
	}
	select {} // This will block forever without consuming CPU
}
