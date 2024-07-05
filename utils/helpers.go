package utils

import (
	"fmt"

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
