package main

import (
	"math/rand"
	"sync"
	"time"
	"twitch-viewer-bot/bot"

	"github.com/sirupsen/logrus"
)

var (
	log           = logrus.New()
	botCountMutex sync.Mutex
	stopChans     map[int]chan bool
)

func init() {
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	log.Level = logrus.DebugLevel
	stopChans = make(map[int]chan bool)
}

func adjustBots(targetCount int) {
	botCountMutex.Lock()
	defer botCountMutex.Unlock()

	currentCount := len(stopChans)
	diff := targetCount - currentCount
	if diff > 0 {
		for i := 0; i < diff; i++ {
			stopChan := make(chan bool)
			id := len(stopChans)
			stopChans[id] = stopChan
			delay := time.Duration(rand.Intn(5000)) * time.Millisecond // Random start delay up to 5 seconds
			go func(id int) {
				time.Sleep(delay)
				bot.OpenBot(int8(id), stopChan)
				manageBotLifespan(id, stopChan) // Call to manage the bot's lifespan
			}(id)
		}
	}
	log.Infof("Adjusted bot count from %d to %d at %v", currentCount, len(stopChans), time.Now())
}

func manageBotLifespan(id int, stopChan chan bool) {
	duration := time.Duration(rand.Intn(25)+5) * time.Minute // Bot operates between 5 to 30 minutes
	log.Infof("Bot %d starting for a lifespan of %v", id, duration)
	time.Sleep(duration)
	stopChan <- true
	close(stopChan)
	botCountMutex.Lock()
	delete(stopChans, id)
	botCountMutex.Unlock()
	log.Infof("Bot %d stopped after %v", id, duration)
}

func botManager() {
	middleAmount := 10 // Set this to your desired middle amount of bots
	for {
		variation := rand.Intn(11) - 5 // Random number between -5 and +5
		newBotCount := middleAmount + variation
		log.Infof("Checking bot count adjustment at %v", time.Now())
		adjustBots(newBotCount)
		time.Sleep(5 * time.Minute) // Adjust bot count every 5 minutes
	}
}

func main() {
	go botManager()
	select {} // Keep the main goroutine running indefinitely
}
