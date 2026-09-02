package _chrome

import (
	"context"
	"time"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/chromedp/chromedp"
)

func allocatorOptions(browserProfile *browser_profiles.BrowserProfile, headless bool) []chromedp.ExecAllocatorOption {
	return append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-infobars", true),
		chromedp.UserAgent(browserProfile.UserAgent),
	)
}

func NewContext(browserProfile *browser_profiles.BrowserProfile, timeout time.Duration, headless bool) (context.Context, context.CancelFunc) {
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), allocatorOptions(browserProfile, headless)...)
	ctx, cancelBrowser := chromedp.NewContext(allocCtx)
	ctx, cancelTimeout := context.WithTimeout(ctx, timeout)

	return ctx, func() {
		cancelTimeout()
		cancelBrowser()
		cancelAlloc()
	}
}
