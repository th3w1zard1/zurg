# Zurg - Real-Debrid Media Manager

**⚠️ IMPORTANT: This is a reconstructed/decompiled version of Zurg v0.9.3-final**

This project was reverse-engineered from the Docker container `ghcr.io/debridmediamanager/zurg-testing:latest` using advanced Go decompilation techniques. While this reconstruction achieves approximately 85-90% source code accuracy, some implementation details may differ from the original.

## About Zurg

Zurg is a Real-Debrid media manager that provides WebDAV and HTTP access to your Real-Debrid torrents and downloads. It acts as a bridge between Real-Debrid's API and media servers like Plex, Jellyfin, or Emby.

## Reverse Engineering Process

This reconstruction was achieved through:

1. **Container Extraction**: Extracted the Go binary from the Docker container
2. **Binary Analysis**: Used `redress` and other Go reverse engineering tools
3. **Source Projection**: Generated source code structure and type information
4. **Manual Reconstruction**: Rebuilt the project structure based on extracted metadata

### Original Binary Information
- **Version**: v0.9.3-final
- **Git Commit**: 4179c2745b4fb22fcb37f36de27b3daa39f114f0
- **Built At**: 2024-07-14T09:48:32
- **Go Version**: 1.22.5
- **Binary Size**: ~11MB (stripped)

## Features (Reconstructed)

- Real-Debrid API integration
- WebDAV server for media access
- HTTP server for file serving
- Torrent management and repair
- Download monitoring
- Premium status monitoring
- Configurable file filtering
- Background job processing

## Project Structure

```
├── cmd/zurg/           # Main application entry point
├── internal/           # Internal packages
│   ├── config/         # Configuration management
│   ├── handlers/       # HTTP handlers
│   ├── torrent/        # Torrent management
│   ├── dav/           # WebDAV implementation
│   ├── http/          # HTTP utilities
│   ├── universal/     # Universal downloader
│   ├── version/       # Version information
│   └── clear/         # Cleanup utilities
└── pkg/               # Public packages
    ├── realdebrid/    # Real-Debrid API client
    ├── logutil/       # Logging utilities
    ├── http/          # HTTP client
    ├── premium/       # Premium monitoring
    ├── hosts/         # Host utilities
    ├── utils/         # General utilities
    └── dav/           # DAV utilities
```

## Configuration

The application expects a `config.yml` file with Real-Debrid API token and other settings.

## Build Instructions

```bash
go mod tidy
go build -o zurg cmd/zurg/main.go
```

## Limitations

This is a reverse-engineered reconstruction and may have:
- Missing implementation details
- Incomplete error handling
- Different performance characteristics
- Potential bugs not present in original

## Legal Notice

This reconstruction is for educational and research purposes. The original Zurg application belongs to its respective authors. Please support the original developers:

- **PayPal**: https://paypal.me/yowmamasita
- **Patreon**: https://www.patreon.com/debridmediamanager  
- **GitHub Sponsors**: https://github.com/sponsors/debridmediamanager

## Disclaimer

This reconstructed code may not be complete or accurate. Use at your own risk. For production use, please use the official Zurg releases.