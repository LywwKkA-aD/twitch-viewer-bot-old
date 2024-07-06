package bot

import (
	"time"
	"twitch-viewer-bot/utils"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func OpenBot(id int8, stopChan chan bool) {
	proxyURL := "http://p.webshare.io:9999"
	if !utils.VerifyProxy(proxyURL) {
		log.Error("Proxy failed for bot: ", id)
		return
	}
	log.Infof("Function instance %d is running with proxy: %s\n", id, proxyURL)

	l := launcher.New().
		Headless(true).
		Proxy(proxyURL).
		Set("disable-web-security", "true").
		Set("autoplay-policy", "user-gesture-required").
		Set("enable-features", "NetworkService,NetworkServiceInProcess").
		Set("disable-features", "OutOfBlinkCors").
		Set("disable-blink-features", "AutomationControlled").
		Set("mute-audio", "true").
		Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe").
		NoSandbox(true)

	url := l.MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect().DefaultDevice(devices.Clear)
	page := browser.MustPage("https://www.twitch.tv/billieblaze")
	page.MustWaitLoad()

	utils.TryClickConsent(page, `button[data-a-target="consent-banner-accept"]`)

	scripts := []string{
		`window.localStorage.setItem('mature', 'true');`,
		`window.localStorage.setItem('video-muted', '{"default": "true"}');`,
		`window.localStorage.setItem('volume', '0.0');`,
		`window.localStorage.setItem('video-quality', '{"default": "160p30"}');`,
		`window.localStorage.setItem('lowLatencyModeEnabled', 'false');`,
	}

	for _, script := range scripts {
		page.Evaluate(rod.Eval(script))
	}

	time.Sleep(2 * time.Second)
	page.Reload()

	select {
	case <-stopChan:
		log.Info("Stopping bot: ", id)
		browser.MustClose()
		return
	default:
		// continue running
	}
}
