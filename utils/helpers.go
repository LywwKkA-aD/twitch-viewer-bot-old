package utils

import (
	"net/http"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func TryClickConsent(page *rod.Page, buttonConsent string) {
	btn, err := page.Element(buttonConsent)
	if err != nil {
		log.Error("Consent button not found: ", err)
		return
	}
	if err := btn.Click(proto.InputMouseButtonLeft, 1); err != nil {
		log.Error("Error clicking consent button: ", err)
	}
}

func VerifyProxy(proxyURL string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://www.twitch.tv/")
	if err != nil {
		log.Error("Proxy verification failed: ", err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
