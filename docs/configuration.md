# ⚙️ Configuration Guide

## Overview

wazctl requires configuration to connect to your Wazuh manager and OpenSearch indexer. Configuration is stored in YAML format and can be placed in multiple locations.

## 🚀 Quick Configuration

### 1. Generate Configuration File

Create a configuration file with default values:

```bash
wazctl init config
```

This creates a `.wazctl.yaml` file in your current directory with the following structure:

```yaml
wazuh:
    endpoint: your-instance.com
    port: "55000"
    protocol: https
    wuiPassword: password
    wuiUsername: wui
    httpDebug: false
    skipTlsVerify: true
indexer:
    endpoint: your-instance.com
    port: "9200"
    protocol: https
    indexerPassword: password
    indexerUsername: wui
    httpDebug: false
    skipTlsVerify: true
local:
    repoVersion: v4.12.0
```

### 2. Edit Configuration

Replace the placeholder values with your actual Wazuh instance details:

```yaml
wazuh:
    endpoint: wazuh.company.com
    port: "55000"
    protocol: https
    wuiPassword: your_secure_password
    wuiUsername: your_username
    httpDebug: false
    skipTlsVerify: false  # Set to true only for testing
indexer:
    endpoint: indexer.company.com
    port: "9200"
    protocol: https
    indexerPassword: indexer_password
    indexerUsername: indexer_username
    httpDebug: false
    skipTlsVerify: false
local:
    repoVersion: v4.12.0  # Wazuh version for local instances
```

### 3. Test Configuration

Verify your configuration works:

```bash
wazctl test auth
```

If successful, this will display a JWT token, confirming authentication.

## 📁 Configuration File Locations

wazctl searches for configuration files in the following order:

1. **Current Directory**: `.wazctl.yaml`
2. **Home Directory**: `~/.wazctl.yaml`
3. **Config Directory**: `~/.config/wazctl.yaml`

Place your configuration file in any of these locations based on your preference.

## 🔧 Configuration Sections

### Wazuh Manager Section

Configuration for your Wazuh manager API:

| Field | Description | Default | Required |
|-------|-------------|---------|----------|
| `endpoint` | Wazuh manager hostname/IP | `your-instance.com` | ✅ |
| `port` | API port | `"55000"` | ✅ |
| `protocol` | HTTP protocol | `https` | ✅ |
| `wuiUsername` | API username | `wui` | ✅ |
| `wuiPassword` | API password | `password` | ✅ |
| `httpDebug` | Enable HTTP debug logging | `false` | ❌ |
| `skipTlsVerify` | Skip TLS certificate verification | `true` | ❌ |

### OpenSearch Indexer Section

Configuration for OpenSearch/Elasticsearch indexer:

| Field | Description | Default | Required |
|-------|-------------|---------|----------|
| `endpoint` | Indexer hostname/IP | `your-instance.com` | ✅ |
| `port` | Indexer port | `"9200"` | ✅ |
| `protocol` | HTTP protocol | `https` | ✅ |
| `indexerUsername` | Indexer username | `wui` | ✅ |
| `indexerPassword` | Indexer password | `password` | ✅ |
| `httpDebug` | Enable HTTP debug logging | `false` | ❌ |
| `skipTlsVerify` | Skip TLS certificate verification | `true` | ❌ |

### Local Development Section

Configuration for local Wazuh instances:

| Field | Description | Default | Required |
|-------|-------------|---------|----------|
| `repoVersion` | Wazuh version for local instances | `v4.12.0` | ✅ |

## 🔒 Security Best Practices

### 🛡️ Credential Management

1. **Never commit credentials to version control**
2. **Use environment variables for sensitive data** (feature planned)
3. **Restrict file permissions**:
   ```bash
   chmod 600 ~/.wazctl.yaml
   ```
4. **Use dedicated service accounts** with minimal required permissions

### 🔐 TLS Configuration

- **Production**: Always set `skipTlsVerify: false`
- **Development**: Only use `skipTlsVerify: true` for testing with self-signed certificates
- **Consider using proper SSL certificates** in your Wazuh deployment

## 🐛 Troubleshooting Configuration

### Authentication Failures

```bash
# Test your configuration
wazctl test auth
```

**Common Issues**:
- Incorrect credentials
- Wrong endpoint or port
- Network connectivity issues
- Firewall blocking connections

### Debug Mode

Enable debug logging to troubleshoot connection issues:

```yaml
wazuh:
    httpDebug: true
indexer:
    httpDebug: true
```

### Configuration Not Found

If wazctl can't find your configuration:

1. **Check file location** - Ensure the file is in one of the expected locations
2. **Check file permissions** - Ensure the file is readable
3. **Check YAML syntax** - Use a YAML validator to verify format

## ➡️ Next Steps

Once configured, try the [Quick Start Tutorial](quick-start.md) to explore wazctl's features.