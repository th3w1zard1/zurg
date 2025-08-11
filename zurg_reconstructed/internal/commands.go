package internal

import (
	"fmt"
	"github.com/debridmediamanager/zurg/internal/version"
	"github.com/debridmediamanager/zurg/internal/clear"
)

// ShowVersion displays version information
func ShowVersion() {
	fmt.Println("Zurg - Real-Debrid Media Manager")
	fmt.Println(version.GetVersion())
	fmt.Println("")
	fmt.Println("Support:")
	fmt.Println("  PayPal: https://paypal.me/yowmamasita")
	fmt.Println("  Patreon: https://www.patreon.com/debridmediamanager")
	fmt.Println("  GitHub Sponsors: https://github.com/sponsors/debridmediamanager")
}

// ClearDownloads clears all downloads from Real-Debrid
func ClearDownloads() {
	clear.ClearDownloads()
}

// ClearTorrents clears all torrents from Real-Debrid
func ClearTorrents() {
	clear.ClearTorrents()
}