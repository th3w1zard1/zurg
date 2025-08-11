package realdebrid

import (
	"encoding/json"
	"time"
)

// User represents Real-Debrid user information
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Points    int    `json:"points"`
	Locale    string `json:"locale"`
	Avatar    string `json:"avatar"`
	Type      string `json:"type"`
	Premium   int    `json:"premium"`
	Expiration string `json:"expiration"`
}

// Torrent represents a Real-Debrid torrent
type Torrent struct {
	ID       string    `json:"id"`
	Filename string    `json:"filename"`
	Hash     string    `json:"hash"`
	Bytes    int64     `json:"bytes"`
	Host     string    `json:"host"`
	Split    int       `json:"split"`
	Progress float64   `json:"progress"`
	Status   string    `json:"status"`
	Added    time.Time `json:"added"`
	Files    []File    `json:"files"`
	Links    []string  `json:"links"`
	Ended    *time.Time `json:"ended,omitempty"`
}

// IsDone checks if torrent is completed
func (t *Torrent) IsDone() bool {
	return t.Status == "downloaded" || t.Status == "seeding"
}

// UnmarshalJSON custom unmarshaling for Torrent
func (t *Torrent) UnmarshalJSON(data []byte) error {
	type Alias Torrent
	aux := &struct {
		Added string `json:"added"`
		Ended string `json:"ended"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Added != "" {
		if addedTime, err := time.Parse("2006-01-02T15:04:05.000Z", aux.Added); err == nil {
			t.Added = addedTime
		}
	}

	if aux.Ended != "" {
		if endedTime, err := time.Parse("2006-01-02T15:04:05.000Z", aux.Ended); err == nil {
			t.Ended = &endedTime
		}
	}

	return nil
}

// File represents a file within a torrent
type File struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Bytes    int64  `json:"bytes"`
	Selected int    `json:"selected"`
}

// TorrentInfo represents detailed torrent information
type TorrentInfo struct {
	ID           string    `json:"id"`
	Filename     string    `json:"filename"`
	OriginalFilename string `json:"original_filename"`
	Hash         string    `json:"hash"`
	Bytes        int64     `json:"bytes"`
	OriginalBytes int64    `json:"original_bytes"`
	Host         string    `json:"host"`
	Split        int       `json:"split"`
	Progress     float64   `json:"progress"`
	Status       string    `json:"status"`
	Added        time.Time `json:"added"`
	Files        []File    `json:"files"`
	Links        []string  `json:"links"`
	Ended        *time.Time `json:"ended,omitempty"`
}

// IsComplete checks if torrent info is complete
func (ti *TorrentInfo) IsComplete() bool {
	if ti.Status != "downloaded" && ti.Status != "seeding" {
		return false
	}
	
	for _, file := range ti.Files {
		if file.Selected == 1 && file.Bytes == 0 {
			return false
		}
	}
	
	return true
}

// MarshalJSON custom marshaling for TorrentInfo
func (ti *TorrentInfo) MarshalJSON() ([]byte, error) {
	type Alias TorrentInfo
	aux := &struct {
		Added string `json:"added"`
		Ended string `json:"ended,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(ti),
		Added: ti.Added.Format("2006-01-02T15:04:05.000Z"),
	}
	
	if ti.Ended != nil {
		aux.Ended = ti.Ended.Format("2006-01-02T15:04:05.000Z")
	}
	
	return json.Marshal(aux)
}

// UnmarshalJSON custom unmarshaling for TorrentInfo
func (ti *TorrentInfo) UnmarshalJSON(data []byte) error {
	type Alias TorrentInfo
	aux := &struct {
		Added string `json:"added"`
		Ended string `json:"ended"`
		*Alias
	}{
		Alias: (*Alias)(ti),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Added != "" {
		if addedTime, err := time.Parse("2006-01-02T15:04:05.000Z", aux.Added); err == nil {
			ti.Added = addedTime
		}
	}

	if aux.Ended != "" {
		if endedTime, err := time.Parse("2006-01-02T15:04:05.000Z", aux.Ended); err == nil {
			ti.Ended = &endedTime
		}
	}

	return nil
}

// Download represents a Real-Debrid download
type Download struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	MimeType string `json:"mimeType"`
	FileSize int64  `json:"filesize"`
	Link     string `json:"link"`
	Host     string `json:"host"`
	Chunks   int    `json:"chunks"`
	Download string `json:"download"`
	Generated time.Time `json:"generated"`
}

// UnmarshalJSON custom unmarshaling for Download
func (d *Download) UnmarshalJSON(data []byte) error {
	type Alias Download
	aux := &struct {
		Generated string `json:"generated"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Generated != "" {
		if genTime, err := time.Parse("2006-01-02T15:04:05.000Z", aux.Generated); err == nil {
			d.Generated = genTime
		}
	}

	return nil
}

// HostInfo represents Real-Debrid host information
type HostInfo struct {
	Host   string `json:"host"`
	Status string `json:"status"`
}

// Alias represents a Real-Debrid alias/unrestrict response
type Alias struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	MimeType string `json:"mimeType"`
	FileSize int64  `json:"filesize"`
	Link     string `json:"link"`
	Host     string `json:"host"`
	Chunks   int    `json:"chunks"`
	Download string `json:"download"`
}