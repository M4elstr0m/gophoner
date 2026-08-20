package google

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/modules"
	"github.com/charmbracelet/log"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const (
	identifierInputSelector = `#identifierId`
	identifierNextSelector  = `#identifierNext`
	identifierCheckRPCQuery = "rpcids=MI613e"

	registeredAccountSignal   = "FIRST_AUTH_FACTOR"
	unregisteredAccountSignal = `"MI613e","[null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,[3]]"`
)

func chromeAllocatorOptions(browserProfile *browser_profiles.BrowserProfile) []chromedp.ExecAllocatorOption {
	return append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-infobars", true),
		chromedp.UserAgent(browserProfile.UserAgent),
	)
}

func newChromeContext(browserProfile *browser_profiles.BrowserProfile, timeout time.Duration) (context.Context, context.CancelFunc) {
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), chromeAllocatorOptions(browserProfile)...)
	ctx, cancelBrowser := chromedp.NewContext(allocCtx)
	ctx, cancelTimeout := context.WithTimeout(ctx, timeout)

	return ctx, func() {
		cancelTimeout()
		cancelBrowser()
		cancelAlloc()
	}
}

func checkIdentifierRegistration(ctx context.Context, phoneNumber string) (bool, error) {
	var (
		mu        sync.Mutex
		requestID network.RequestID
		found     bool
	)
	done := make(chan struct{})
	var closeOnce sync.Once

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventResponseReceived:
			if strings.Contains(e.Response.URL, "batchexecute") && strings.Contains(e.Response.URL, identifierCheckRPCQuery) {
				mu.Lock()
				requestID = e.RequestID
				found = true
				mu.Unlock()
			}
		case *network.EventLoadingFinished:
			mu.Lock()
			match := found && e.RequestID == requestID
			mu.Unlock()
			if match {
				closeOnce.Do(func() { close(done) })
			}
		}
	})

	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(LOGIN_PAGE_URL),
		chromedp.WaitVisible(identifierInputSelector, chromedp.ByID),
		chromedp.SendKeys(identifierInputSelector, phoneNumber, chromedp.ByID),
		chromedp.Click(identifierNextSelector, chromedp.ByID),
	)
	if err != nil {
		log.Warn("Failed to submit identifier", "error", err)
		return false, err
	}

	select {
	case <-done:
	case <-ctx.Done():
		log.Warn("Timed out waiting for identifier-check response", "error", ctx.Err())
		return false, ctx.Err()
	}

	mu.Lock()
	rid := requestID
	mu.Unlock()

	var body []byte
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		body, err = network.GetResponseBody(rid).Do(ctx)
		return err
	}))
	if err != nil {
		log.Warn("Failed to read identifier-check response", "error", err)
		return false, err
	}

	bodyStr := string(body)

	switch {
	case strings.Contains(bodyStr, registeredAccountSignal):
		return true, nil
	case strings.Contains(bodyStr, unregisteredAccountSignal):
		return false, nil
	default:
		err := fmt.Errorf("unexpected identifier-check response shape: %w", modules.ErrorLimitReached)
		log.Warn("Unexpected identifier-check response", "body", bodyStr, "error", err)
		return false, err
	}
}
