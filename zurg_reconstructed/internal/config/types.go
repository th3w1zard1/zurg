package config

// ZurgConfig represents the main configuration
type ZurgConfig struct {
	Token                           string                 `yaml:"token"`
	Host                           string                 `yaml:"host"`
	Port                           int                    `yaml:"port"`
	Username                       string                 `yaml:"username"`
	Password                       string                 `yaml:"password"`
	Proxy                          string                 `yaml:"proxy"`
	NumOfWorkers                   int                    `yaml:"num_of_workers"`
	RefreshEverySecs               int                    `yaml:"refresh_every_secs"`
	RepairEveryMins                int                    `yaml:"repair_every_mins"`
	DownloadsEveryMins             int                    `yaml:"downloads_every_mins"`
	EnableRepair                   bool                   `yaml:"enable_repair"`
	OnLibraryUpdate                string                 `yaml:"on_library_update"`
	NetworkBufferSize              int                    `yaml:"network_buffer_size"`
	EnableRetainFolderNameExtension bool                   `yaml:"enable_retain_folder_name_extension"`
	EnableRetainRDTorrentName       bool                   `yaml:"enable_retain_rd_torrent_name"`
	ShouldIgnoreRenames            bool                   `yaml:"should_ignore_renames"`
	ShouldServeFromRclone          bool                   `yaml:"should_serve_from_rclone"`
	ShouldForceIPv6                bool                   `yaml:"should_force_ipv6"`
	RetriesUntilFailed             int                    `yaml:"retries_until_failed"`
	EnableDownloadMount            bool                   `yaml:"enable_download_mount"`
	ApiTimeoutSecs                 int                    `yaml:"api_timeout_secs"`
	DownloadTimeoutSecs            int                    `yaml:"download_timeout_secs"`
	RateLimitSleepSecs             int                    `yaml:"rate_limit_sleep_secs"`
	ShouldDeleteRarFiles           bool                   `yaml:"should_delete_rar_files"`
	PlayableExtensions             []string               `yaml:"playable_extensions"`
	TorrentsCount                  int                    `yaml:"torrents_count"`
	APIRateLimitPerMinute          int                    `yaml:"api_rate_limit_per_minute"`
	TorrentsRateLimitPerMinute     int                    `yaml:"torrents_rate_limit_per_minute"`
	Directories                    map[string]interface{} `yaml:"directories"`
}

// GetConfig returns the config itself (for interface compatibility)
func (zc *ZurgConfig) GetConfig() *ZurgConfig {
	return zc
}

// GetToken returns the API token
func (zc *ZurgConfig) GetToken() string {
	return zc.Token
}

// GetHost returns the host
func (zc *ZurgConfig) GetHost() string {
	if zc.Host == "" {
		return "0.0.0.0"
	}
	return zc.Host
}

// GetPort returns the port
func (zc *ZurgConfig) GetPort() int {
	if zc.Port == 0 {
		return 9999
	}
	return zc.Port
}

// GetUsername returns the username
func (zc *ZurgConfig) GetUsername() string {
	return zc.Username
}

// GetPassword returns the password
func (zc *ZurgConfig) GetPassword() string {
	return zc.Password
}

// GetProxy returns the proxy setting
func (zc *ZurgConfig) GetProxy() string {
	return zc.Proxy
}

// GetNumOfWorkers returns number of workers
func (zc *ZurgConfig) GetNumOfWorkers() int {
	if zc.NumOfWorkers == 0 {
		return 10
	}
	return zc.NumOfWorkers
}

// GetRefreshEverySecs returns refresh interval
func (zc *ZurgConfig) GetRefreshEverySecs() int {
	if zc.RefreshEverySecs == 0 {
		return 120
	}
	return zc.RefreshEverySecs
}

// GetRepairEveryMins returns repair interval
func (zc *ZurgConfig) GetRepairEveryMins() int {
	if zc.RepairEveryMins == 0 {
		return 30
	}
	return zc.RepairEveryMins
}

// GetDownloadsEveryMins returns downloads refresh interval
func (zc *ZurgConfig) GetDownloadsEveryMins() int {
	if zc.DownloadsEveryMins == 0 {
		return 5
	}
	return zc.DownloadsEveryMins
}

// EnableRepair returns repair enabled status
func (zc *ZurgConfig) EnableRepair() bool {
	return zc.EnableRepair
}

// GetOnLibraryUpdate returns library update hook
func (zc *ZurgConfig) GetOnLibraryUpdate() string {
	return zc.OnLibraryUpdate
}

// GetNetworkBufferSize returns network buffer size
func (zc *ZurgConfig) GetNetworkBufferSize() int {
	if zc.NetworkBufferSize == 0 {
		return 1024 * 1024 // 1MB
	}
	return zc.NetworkBufferSize
}

// EnableRetainFolderNameExtension returns folder name extension retention setting
func (zc *ZurgConfig) EnableRetainFolderNameExtension() bool {
	return zc.EnableRetainFolderNameExtension
}

// EnableRetainRDTorrentName returns RD torrent name retention setting
func (zc *ZurgConfig) EnableRetainRDTorrentName() bool {
	return zc.EnableRetainRDTorrentName
}

// ShouldIgnoreRenames returns ignore renames setting
func (zc *ZurgConfig) ShouldIgnoreRenames() bool {
	return zc.ShouldIgnoreRenames
}

// ShouldServeFromRclone returns serve from rclone setting
func (zc *ZurgConfig) ShouldServeFromRclone() bool {
	return zc.ShouldServeFromRclone
}

// ShouldForceIPv6 returns force IPv6 setting
func (zc *ZurgConfig) ShouldForceIPv6() bool {
	return zc.ShouldForceIPv6
}

// GetRetriesUntilFailed returns retry count
func (zc *ZurgConfig) GetRetriesUntilFailed() int {
	if zc.RetriesUntilFailed == 0 {
		return 3
	}
	return zc.RetriesUntilFailed
}

// EnableDownloadMount returns download mount setting
func (zc *ZurgConfig) EnableDownloadMount() bool {
	return zc.EnableDownloadMount
}

// GetApiTimeoutSecs returns API timeout
func (zc *ZurgConfig) GetApiTimeoutSecs() int {
	if zc.ApiTimeoutSecs == 0 {
		return 30
	}
	return zc.ApiTimeoutSecs
}

// GetDownloadTimeoutSecs returns download timeout
func (zc *ZurgConfig) GetDownloadTimeoutSecs() int {
	if zc.DownloadTimeoutSecs == 0 {
		return 300
	}
	return zc.DownloadTimeoutSecs
}

// GetRateLimitSleepSecs returns rate limit sleep duration
func (zc *ZurgConfig) GetRateLimitSleepSecs() int {
	if zc.RateLimitSleepSecs == 0 {
		return 5
	}
	return zc.RateLimitSleepSecs
}

// ShouldDeleteRarFiles returns delete RAR files setting
func (zc *ZurgConfig) ShouldDeleteRarFiles() bool {
	return zc.ShouldDeleteRarFiles
}

// GetPlayableExtensions returns playable file extensions
func (zc *ZurgConfig) GetPlayableExtensions() []string {
	if len(zc.PlayableExtensions) == 0 {
		return []string{".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm", ".m4v", ".3gp", ".ts", ".m2ts"}
	}
	return zc.PlayableExtensions
}

// GetTorrentsCount returns torrents count
func (zc *ZurgConfig) GetTorrentsCount() int {
	if zc.TorrentsCount == 0 {
		return 2500
	}
	return zc.TorrentsCount
}

// GetAPIRateLimitPerMinute returns API rate limit
func (zc *ZurgConfig) GetAPIRateLimitPerMinute() int {
	if zc.APIRateLimitPerMinute == 0 {
		return 60
	}
	return zc.APIRateLimitPerMinute
}

// GetTorrentsRateLimitPerMinute returns torrents rate limit
func (zc *ZurgConfig) GetTorrentsRateLimitPerMinute() int {
	if zc.TorrentsRateLimitPerMinute == 0 {
		return 1
	}
	return zc.TorrentsRateLimitPerMinute
}