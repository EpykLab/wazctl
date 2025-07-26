# 🔐 Authentication

wazctl uses the Wazuh API authentication system to interact with your Wazuh manager. This guide covers authentication setup, testing, and troubleshooting.

## 🎯 Overview

wazctl authenticates using:
- **Username/Password**: Standard Wazuh API credentials
- **JWT Tokens**: Session-based authentication (handled automatically)
- **TLS**: HTTPS connections (configurable TLS verification)

## 🧪 Testing Authentication

### Basic Authentication Test

The most important command for verifying your setup:

```bash
wazctl test auth
```

**Successful Output**:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3YXp1aCIsImF1ZCI6IldhenVoIEFQSSBSRVNUIiwiaWF0IjoxNzA2NzE0NDI4LCJleHAiOjE3MDY3MTUzMjgsInN1YiI6IndlbGNvbWUifQ.jwt_token_here
```

This JWT token confirms successful authentication.

**Failed Output**:
```
Error: authentication failed
```

### Understanding JWT Tokens

The returned JWT token contains:
- **Issuer**: Wazuh API
- **Subject**: Your username
- **Expiration**: Token validity period
- **Permissions**: Based on your user role

## ⚙️ Authentication Configuration

### Required Credentials

In your `.wazctl.yaml` file:

```yaml
wazuh:
    endpoint: your-wazuh-manager.com
    port: "55000"
    protocol: https
    wuiUsername: your_username
    wuiPassword: your_password
    skipTlsVerify: false  # Production: false, Testing: true
```

### Credential Requirements

| Field | Description | Example |
|-------|-------------|---------|
| `wuiUsername` | Wazuh API username | `admin`, `analyst`, `readonly` |
| `wuiPassword` | Wazuh API password | `SecurePassword123!` |
| `endpoint` | Wazuh manager hostname | `wazuh.company.com` |
| `port` | API port (usually 55000) | `"55000"` |
| `protocol` | Connection protocol | `https` (recommended) |

## 🔒 Security Configuration

### TLS Certificate Verification

**Production Environment** (Recommended):
```yaml
wazuh:
    protocol: https
    skipTlsVerify: false
```

**Development/Testing** (Self-signed certificates):
```yaml
wazuh:
    protocol: https
    skipTlsVerify: true
```

### HTTP Debug Mode

Enable detailed HTTP logging for troubleshooting:

```yaml
wazuh:
    httpDebug: true
```

⚠️ **Warning**: Debug mode exposes sensitive information in logs. Use only for troubleshooting.

## 👤 User Roles and Permissions

### Common Wazuh API User Roles

| Role | Permissions | wazctl Compatibility |
|------|-------------|---------------------|
| `administrator` | Full API access | ✅ All features |
| `agent_manager` | Agent management | ✅ Agent commands |
| `readonly` | Read-only access | ⚠️ Limited features |
| `normal` | Basic access | ⚠️ Very limited |

### Required Permissions for wazctl Features

| Feature | Required Permissions |
|---------|---------------------|
| `test auth` | Basic authentication |
| `agents list` | Agent read permissions |
| `user add` | User management permissions |
| API commands | Varies by endpoint |

## 🐛 Troubleshooting Authentication

### Common Authentication Errors

#### 1. Connection Refused
```bash
Error: connection refused
```

**Causes & Solutions**:
- **Wrong endpoint**: Verify hostname/IP in configuration
- **Wrong port**: Ensure port 55000 is correct
- **Firewall**: Check network connectivity
- **Service down**: Verify Wazuh manager is running

#### 2. Authentication Failed
```bash
Error: authentication failed
```

**Causes & Solutions**:
- **Wrong credentials**: Verify username/password
- **Account locked**: Check Wazuh user status
- **Insufficient permissions**: Verify user has API access

#### 3. TLS Certificate Errors
```bash
Error: x509: certificate signed by unknown authority
```

**Solutions**:
- **Development**: Set `skipTlsVerify: true`
- **Production**: Install proper SSL certificates
- **Self-signed**: Add certificate to system trust store

#### 4. Token Expired
```bash
Error: token expired
```

**Solution**: wazctl automatically handles token refresh. If this persists:
- Check system clock synchronization
- Verify Wazuh manager time settings

### Debug Authentication Issues

1. **Enable debug logging**:
   ```yaml
   wazuh:
       httpDebug: true
   ```

2. **Test basic connectivity**:
   ```bash
   curl -k https://your-wazuh-manager:55000
   ```

3. **Verify API endpoint**:
   ```bash
   curl -k -X GET "https://your-wazuh-manager:55000/" \
        -H "Authorization: Bearer $(wazctl test auth)"
   ```

4. **Check Wazuh manager logs**:
   ```bash
   # On Wazuh manager
   tail -f /var/ossec/logs/api.log
   ```

## 🔧 Advanced Authentication

### Multiple Environment Setup

Manage different environments with separate configurations:

```bash
# Development
cp .wazctl.yaml .wazctl.dev.yaml

# Production
cp .wazctl.yaml .wazctl.prod.yaml

# Switch environments
cp .wazctl.dev.yaml .wazctl.yaml
wazctl test auth
```

### Credential Security

**Best Practices**:
1. **Never commit credentials** to version control
2. **Use dedicated service accounts** with minimal permissions
3. **Rotate passwords regularly**
4. **Restrict file permissions**:
   ```bash
   chmod 600 .wazctl.yaml
   ```

### API Rate Limiting

Be aware of Wazuh API rate limits:
- **Default**: 300 requests per minute per IP
- **Recommendation**: Add delays between bulk operations
- **Monitor**: Check API logs for rate limit warnings

## 🔄 Token Management

### Automatic Token Refresh

wazctl automatically:
- Obtains JWT tokens on first request
- Refreshes expired tokens
- Handles token renewal transparently

### Manual Token Operations

```bash
# Get current token
TOKEN=$(wazctl test auth)

# Use token with curl
curl -k -X GET "https://wazuh-manager:55000/agents" \
     -H "Authorization: Bearer $TOKEN"
```

## ➡️ Next Steps

Once authentication is working:
- Explore [Agent Management](agents.md)
- Try [API Commands](api.md)  
- Set up [User Management](users.md)