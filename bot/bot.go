package bot

import (
	"fmt"
	"twitch-viewer-bot/utils"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
)

// LocalStorageConfig represents the default local storage settings for Twitch.
var LocalStorageConfig = map[string]string{
	"mature":                "true",
	"video-muted":           `{"default": "false"}`,
	"volume":                "0.5",
	"video-quality":         `{"default": "160p30"}`,
	"lowLatencyModeEnabled": "false",
}

func OpenBot(id int8) {
	proxyURL := fmt.Sprintf("http://%s:%s", "p.webshare.io", "9999")
	fmt.Printf("Function instance %d is running with proxy: %s\n", id, proxyURL)

	l := launcher.New().
		Headless(false).
		Set("disable-web-security", "true").
		Set("autoplay-policy", "no-user-gesture-required").
		Set("enable-features", "NetworkService,NetworkServiceInProcess").
		Set("disable-features", "OutOfBlinkCors").
		Set("disable-blink-features", "AutomationControlled").
		Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe")

	url := l.MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect().DefaultDevice(devices.Clear)
	page := browser.MustPage("https://www.twitch.tv/lmckgkfj9")
	page.MustWaitLoad() // Ensure the page is fully loaded

	// Interact with consent and mute button
	acceptCookies := `button[data-a-target="consent-banner-accept"]`
	muteButton := `button[data-a-target="player-mute-unmute-button"]`

	if err := utils.TryClickConsent(page, acceptCookies); err != nil {
		fmt.Println("Failed to click consent button:", err)
		return
	}

	if err := utils.TryClickConsent(page, muteButton); err != nil {
		fmt.Println("Failed to click mute button:", err)
		return
	}

	// Now set local storage
	if err := utils.SetLocalStorage(page, LocalStorageConfig); err != nil {
		fmt.Println("Failed to set local storage:", err)
		return
	}

	select {} // Keep the browser open
}
