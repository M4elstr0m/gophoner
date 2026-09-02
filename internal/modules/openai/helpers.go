package openai

import (
	"context"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	sessionEndedLoginLinkSelector = `a[href*="/auth/login_with"]`

	continueWithPhoneSelector = `button[data-dd-action-name="Continue with phone"]`
	phoneInputSelector        = `#tel`
	continueSubmitSelector    = `button[data-dd-action-name="Continue"]`

	passwordPagePathToken      = "/log-in/password"
	createAccountPagePathToken = "/create-account"

	registeredSignalToken = "usual sign-in method"

	outcomePollInterval = 300 * time.Millisecond

	passwordPageGraceWindow = 2 * time.Second
)

func checkPhoneRegistration(ctx context.Context, phoneNumber string) (bool, error) {
	if err := chromedp.Run(ctx,
		chromedp.Navigate(LOGIN_PAGE_URL),
		chromedp.WaitReady("body", chromedp.ByQuery),
	); err != nil {
		log.Warn("Failed to load login page", "error", err)
		return false, err
	}

	if err := bypassExpiredSessionInterstitial(ctx); err != nil {
		log.Warn("Failed to recover from expired-session interstitial", "error", err)
		return false, err
	}

	err := chromedp.Run(ctx,
		chromedp.WaitVisible(continueWithPhoneSelector, chromedp.ByQuery),
		chromedp.Click(continueWithPhoneSelector, chromedp.ByQuery),
		chromedp.WaitVisible(phoneInputSelector, chromedp.ByID),
		chromedp.SendKeys(phoneInputSelector, phoneNumber, chromedp.ByID),
		chromedp.Click(continueSubmitSelector, chromedp.ByQuery),
	)
	if err != nil {
		log.Warn("Failed to submit phone number", "error", err)
		return false, err
	}

	return pollRegistrationOutcome(ctx)
}

func bypassExpiredSessionInterstitial(ctx context.Context) error {
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(sessionEndedLoginLinkSelector, &nodes, chromedp.ByQuery, chromedp.AtLeast(0))); err != nil {
		return err
	}
	if len(nodes) == 0 {
		return nil
	}

	log.Debug("Session-ended interstitial encountered, following its login link")
	return chromedp.Run(ctx, chromedp.Click(sessionEndedLoginLinkSelector, chromedp.ByQuery))
}

func pollRegistrationOutcome(ctx context.Context) (bool, error) {
	ticker := time.NewTicker(outcomePollInterval)
	defer ticker.Stop()

	var onPasswordPageSince time.Time

	for {
		select {
		case <-ctx.Done():
			log.Warn("Timed out waiting for phone-registration outcome", "error", ctx.Err())
			return false, ctx.Err()
		case <-ticker.C:
			var bodyText string
			if err := chromedp.Run(ctx, chromedp.Text("body", &bodyText, chromedp.ByQuery, chromedp.NodeVisible)); err != nil {
				continue
			}
			if strings.Contains(bodyText, registeredSignalToken) {
				return true, nil
			}

			var currentURL string
			if err := chromedp.Run(ctx, chromedp.Location(&currentURL)); err != nil {
				continue
			}
			if strings.Contains(currentURL, createAccountPagePathToken) {
				return false, nil
			}
			if !strings.Contains(currentURL, passwordPagePathToken) {
				onPasswordPageSince = time.Time{}
				continue
			}
			if onPasswordPageSince.IsZero() {
				onPasswordPageSince = time.Now()
				continue
			}
			if time.Since(onPasswordPageSince) >= passwordPageGraceWindow {
				return false, nil
			}
		}
	}
}
