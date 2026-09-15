package update

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/M4elstr0m/gophoner/internal/version"
	"github.com/adrg/xdg"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"golang.org/x/mod/semver"
)

const (
	releasesURL   = "https://api.github.com/repos/" + version.AUTHOR + "/" + version.APP_NAME + "/releases/latest"
	cacheFileName = "update-check.json"
	checkInterval = 24 * time.Hour
)

var updateBadgeStyle = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#f9f9f9")).
	Background(lipgloss.Color("#00b2bb"))

type githubRelease struct {
	TagName string `json:"tag_name"`
}

type cacheData struct {
	LastChecked   time.Time `json:"last_checked"`
	LatestVersion string    `json:"latest_version"`
}

func cachePath() (string, error) {
	return xdg.StateFile(version.APP_NAME + "/" + cacheFileName)
}

func readCache() (cacheData, bool) {
	path, err := cachePath()
	if err != nil {
		return cacheData{}, false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cacheData{}, false
	}

	var c cacheData
	if err := json.Unmarshal(data, &c); err != nil {
		return cacheData{}, false
	}

	return c, true
}

func writeCache(c cacheData) {
	path, err := cachePath()
	if err != nil {
		return
	}

	data, err := json.Marshal(c)
	if err != nil {
		return
	}

	_ = os.WriteFile(path, data, 0o644)
}

func fetchLatestVersion() (string, bool) {
	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Get(releasesURL)
	if err != nil {
		log.Warn("Failed to check for updates", "error", err)
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn("Failed to check for updates", "status", resp.Status)
		return "", false
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		log.Warn("Failed to parse update check response", "error", err)
		return "", false
	}

	return release.TagName, true
}

func latestVersion() string {
	cache, hasCache := readCache()
	if hasCache && time.Since(cache.LastChecked) < checkInterval {
		return cache.LatestVersion
	}

	latest, fetched := fetchLatestVersion()
	if !fetched {
		latest = cache.LatestVersion
	}

	writeCache(cacheData{
		LastChecked:   time.Now(),
		LatestVersion: latest,
	})

	return latest
}

// Returns true if it printed something just in case
func Notify(noUpdate bool) bool {
	if noUpdate {
		return false
	}

	latest := latestVersion()
	if latest == "" {
		return false
	}

	current := "v" + version.Version
	if !semver.IsValid(latest) || !semver.IsValid(current) {
		return false
	}

	if semver.Compare(latest, current) <= 0 {
		return false
	}

	log.Info("A new version of gophoner is available",
		"current", version.Version,
		"latest", latest,
	)
	fmt.Fprintf(os.Stderr, "%s A new version of gophoner is available: %s (current: v%s)\n%s Get it here: %s\n\n",
		updateBadgeStyle.Render("NEW UPDATE"),
		latest,
		version.Version,
		updateBadgeStyle.Render(" "),
		"https://github.com/"+version.AUTHOR+"/"+version.APP_NAME+"/releases/latest",
	)

	return true
}
