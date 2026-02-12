# 📦 Installation Guide

## Prerequisites

- **Go 1.24.3+** - wazctl is built with Go and requires a recent version
- **Network Access** - Connection to your Wazuh manager API (typically port 55000)

## Installation Methods

### 🎯 Go Install (Recommended)

The fastest way to install wazctl is using Go's built-in package manager:

```bash
go install github.com/EpykLab/wazctl@latest
```

**Important**: Ensure your `$(go env GOPATH)/bin` directory is in your system's `PATH`.

### 📥 Pre-compiled Binaries

> ⚠️ **Coming Soon**: Pre-compiled binaries for multiple platforms are planned for future releases.

### 🔨 Build from Source

1. **Clone the repository**:
   ```bash
   git clone https://github.com/EpykLab/wazctl.git
   cd wazctl
   ```

2. **Build the binary**:
   ```bash
   go build -o wazctl main.go
   ```

3. **Install to your PATH** (optional):
   ```bash
   sudo mv wazctl /usr/local/bin/
   ```

## ✅ Verify Installation

Confirm wazctl is installed correctly:

```bash
wazctl --help
```

You should see the main help output with available commands.

## 🔄 Updating

To update to the latest version:

```bash
go install github.com/EpykLab/wazctl@latest
```

## 🗑️ Uninstalling

To remove wazctl installed via `go install`:

```bash
rm $(go env GOPATH)/bin/wazctl
```

## 🆘 Troubleshooting Installation

### Go Not Found
If you get a "go: command not found" error, install Go from [golang.org](https://golang.org/dl/).

### PATH Issues
If wazctl command isn't found after installation, add Go's bin directory to your PATH:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### Network Issues
If installation fails due to network issues, try:
- Using a VPN or different network
- Setting up Go module proxy: `export GOPROXY=https://proxy.golang.org,direct`

## ➡️ Next Steps

Once installed, proceed to the [Configuration Guide](configuration.md) to set up your Wazuh connection.