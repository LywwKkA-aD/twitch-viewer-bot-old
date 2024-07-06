package bot

import (
	"fmt"
	"time"
	"twitch-viewer-bot/utils"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
)

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
	page := browser.MustPage("https://www.twitch.tv/milten_uzumaki")
	page.MustWaitLoad() // Ensure the page is fully loaded

	// Interact with consent and mute button
	acceptCookies := `button[data-a-target="consent-banner-accept"]`

	if err := utils.TryClickConsent(page, acceptCookies); err != nil {
		fmt.Println("Failed to click consent button:", err)
		return
	}

	// Set the local storage value for 'mature'
	script1 := `window.localStorage.setItem('mature', 'true`
	script2 := `window.localStorage.setItem('video-muted', '{"default": "false"}');`
	script3 := `window.localStorage.setItem('volume', '0.5');`
	script4 := `window.localStorage.setItem('video-quality', '{"default": "160p30"}');`
	script5 := `window.localStorage.setItem('lowLatencyModeEnabled', 'false');`
	_, err1 := page.Evaluate(rod.Eval(script1))
	if err1 != nil {
		fmt.Printf("Failed to set localStorage key 'mature': %v\n", err1)
	}
	_, err2 := page.Evaluate(rod.Eval(script2))
	if err2 != nil {
		fmt.Printf("Failed to set localStorage key 'mature': %v\n", err2)
	}
	_, err3 := page.Evaluate(rod.Eval(script3))
	if err3 != nil {
		fmt.Printf("Failed to set localStorage key 'mature': %v\n", err3)
	}
	_, err4 := page.Evaluate(rod.Eval(script4))
	if err4 != nil {
		fmt.Printf("Failed to set localStorage key 'mature': %v\n", err4)
	}
	_, err5 := page.Evaluate(rod.Eval(script5))
	if err5 != nil {
		fmt.Printf("Failed to set localStorage key 'mature': %v\n", err5)
	}

	time.Sleep(5 * time.Second) // Wait for the page to apply the changes
	page.Reload()               // Reload the page to apply the changes

	select {} // Keep the browser open
}
