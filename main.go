package main

import (
	"fmt"
	"math/rand"
	"time"
	"twitch-viewer-bot/bot"
)

func main() {
	var i int8
	for i = 0; i < 20; i++ {
		go bot.OpenBot(i)
		// Set the seed
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(8) + 7 // generate a random number between 7 and 15
		fmt.Printf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
	}
	select {} // This will block forever without consuming CPU
}
