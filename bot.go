package bot

import (
	"fmt"
	"time"
	"twitch-viwer-bot/utils"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
)

// OpenBot launches a browser instance with the given proxy configuration.
func OpenBot(id int8) {
	proxyURL := fmt.Sprintf("http://%s:%s", "p.webshare.io", "9999")
	fmt.Printf("Function instance %d is running with proxy: %s\n", id, proxyURL)

	l := launcher.New().
		Headless(true).
		Proxy(proxyURL).
		Set("disable-web-security", "true").
		Set("autoplay-policy", "no-user-gesture-required").
		Set("enable-features", "NetworkService,NetworkServiceInProcess").
		Set("disable-features", "OutOfBlinkCors").
		Set("disable-blink-features", "AutomationControlled").
		Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe")

	url := l.MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect().DefaultDevice(devices.Clear)

	// Navigate to Twitch in a new page
	twitchPage := browser.MustPage("https://www.twitch.tv/lmckgkfj9")
	twitchPage.MustWaitLoad() // Ensure the page is fully loaded

	// Wait a bit before trying to interact with elements
	time.Sleep(5 * time.Second)

	acceptCookies := `button[data-a-target="consent-banner-accept"]`
	muteButton := `button[data-a-target="player-mute-unmute-button"]`

	if err := utils.TryClickConsent(twitchPage, acceptCookies); err != nil {
		fmt.Println("Failed to click consent button:", err)
		return
	}

	if err := utils.TryClickConsent(twitchPage, muteButton); err != nil {
		fmt.Println("Failed to click consent button:", err)
		return
	}

	select {} // Keep the browser open
}
