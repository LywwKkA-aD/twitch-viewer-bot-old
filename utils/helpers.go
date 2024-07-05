package utils

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func TryClickConsent(page *rod.Page, buttonConsent string) error {

	// Find the button using the selector
	btn, err := page.Element(buttonConsent)
	if err != nil {
		return fmt.Errorf("failed to find the accept cookies button: %w", err)
	}

	// Click on the button with the left mouse button and a click count of 1
	if err := btn.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return fmt.Errorf("failed to click the consent button: %w", err)
	}

	return nil
}

func SetLocalStorage(page *rod.Page, storage map[string]string) {
	var scriptBuilder strings.Builder

	// Start the script
	scriptBuilder.WriteString("")

	for key, value := range storage {
		// Escape single quotes in key and value to ensure JavaScript code validity
		escapedKey := strings.ReplaceAll(key, "'", "\\'")
		escapedValue := strings.ReplaceAll(value, "'", "\\'")

		// Add each local storage set item to the script
		scriptBuilder.WriteString(fmt.Sprintf("localStorage.setItem('%s', '%s');\n", escapedKey, escapedValue))
	}

	// Execute the complete script
	page.MustEvaluate(scriptBuilder.String())
}
