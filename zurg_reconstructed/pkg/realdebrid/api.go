package realdebrid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/debridmediamanager/zurg/pkg/http"
)

const (
	BaseURL = "https://api.real-debrid.com/rest/1.0"
)

// RealDebrid represents the Real-Debrid API client
type RealDebrid struct {
	client *http.HTTPClient
	token  string
}

// NewRealDebrid creates a new Real-Debrid client
func NewRealDebrid(token string, httpClient *http.HTTPClient) *RealDebrid {
	return &RealDebrid{
		client: httpClient,
		token:  token,
	}
}

// UnrestrictLink unrestricts a download link
func (rd *RealDebrid) UnrestrictLink(ctx context.Context, link string) (*Alias, error) {
	data := url.Values{}
	data.Set("link", link)

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/unrestrict/link", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := rd.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var alias Alias
	if err := json.NewDecoder(resp.Body).Decode(&alias); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &alias, nil
}

// GetTorrents gets all torrents
func (rd *RealDebrid) GetTorrents(ctx context.Context, offset, page int) ([]Torrent, error) {
	reqURL := fmt.Sprintf("%s/torrents?offset=%d&page=%d", BaseURL, offset, page)
	
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var torrents []Torrent
	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return torrents, nil
}

// GetTorrentInfo gets detailed information about a torrent
func (rd *RealDebrid) GetTorrentInfo(ctx context.Context, id string) (*TorrentInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", BaseURL+"/torrents/info/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var torrentInfo TorrentInfo
	if err := json.NewDecoder(resp.Body).Decode(&torrentInfo); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &torrentInfo, nil
}

// SelectTorrentFiles selects files from a torrent
func (rd *RealDebrid) SelectTorrentFiles(ctx context.Context, id string, files string) error {
	data := url.Values{}
	data.Set("files", files)

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/torrents/selectFiles/"+id, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := rd.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// DeleteTorrent deletes a torrent
func (rd *RealDebrid) DeleteTorrent(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", BaseURL+"/torrents/delete/"+id, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// AddMagnetHash adds a magnet/hash to torrents
func (rd *RealDebrid) AddMagnetHash(ctx context.Context, magnet string) (string, error) {
	data := url.Values{}
	data.Set("magnet", magnet)

	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+"/torrents/addMagnet", strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := rd.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		ID  string `json:"id"`
		URI string `json:"uri"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return result.ID, nil
}

// GetActiveTorrentCount gets the count of active torrents
func (rd *RealDebrid) GetActiveTorrentCount(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", BaseURL+"/torrents/activeCount", nil)
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Count int `json:"nb"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("error decoding response: %w", err)
	}

	return result.Count, nil
}

// GetDownloads gets downloads with pagination
func (rd *RealDebrid) GetDownloads(ctx context.Context, offset, page int) ([]Download, error) {
	reqURL := fmt.Sprintf("%s/downloads?offset=%d&page=%d", BaseURL, offset, page)
	
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var downloads []Download
	if err := json.NewDecoder(resp.Body).Decode(&downloads); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return downloads, nil
}

// GetUserInformation gets user information
func (rd *RealDebrid) GetUserInformation(ctx context.Context) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", BaseURL+"/user", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+rd.token)

	resp, err := rd.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &user, nil
}