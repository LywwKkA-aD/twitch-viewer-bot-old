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
	stopChans     map[int]chan struct{}
)

func init() {
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.Level = logrus.DebugLevel
	stopChans = make(map[int]chan struct{})
}

func adjustBots(targetCount int) {
	botCountMutex.Lock()
	defer botCountMutex.Unlock()

	currentCount := len(stopChans)
	diff := targetCount - currentCount
	if diff > 0 {
		for i := 0; i < diff; i++ {
			stopChan := make(chan struct{})
			id := len(stopChans)
			stopChans[id] = stopChan
			delay := time.Duration(rand.Intn(5000)) * time.Millisecond // Random start delay up to 5 seconds
			go func(id int) {
				time.Sleep(delay)
				proxyURL := "http://p.webshare.io:9999" // Proxy URL should be dynamic or configurable
				log.Infof("Starting bot %d with a delay of %v", id, delay)
				bot.OpenBot(id, proxyURL, stopChan)
				manageBotLifespan(id, stopChan)
			}(id)
		}
	}
	log.Infof("Adjusted bot count from %d to %d at %v", currentCount, len(stopChans), time.Now().Format("2006-01-02 15:04:05"))
}

func manageBotLifespan(id int, stopChan chan struct{}) {
	duration := time.Duration(rand.Intn(25)+5) * time.Minute // Bot operates between 5 to 30 minutes
	log.Infof("Bot %d will run for a lifespan of %v", id, duration)
	time.Sleep(duration)
	close(stopChan)

	botCountMutex.Lock()
	defer botCountMutex.Unlock()
	delete(stopChans, id)
	log.Infof("Bot %d stopped after %v", id, duration)
}

func botManager() {
	middleAmount := 7 // Set this to your desired middle amount of bots
	for {
		variation := rand.Intn(11) - 5 // Random number between -5 and +5
		newBotCount := middleAmount + variation
		log.Infof("Checking bot count adjustment at %v", time.Now().Format("2006-01-02 15:04:05"))
		adjustBots(newBotCount)
		time.Sleep(5 * time.Minute) // Adjust bot count every 5 minutes
	}
}

func main() {
	go botManager()
	select {} // Keep the main goroutine running indefinitely
}
