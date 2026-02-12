# ⚙️ Configuration Reference

This comprehensive reference covers all wazctl configuration options, file formats, and advanced configuration scenarios.

## 📋 Configuration File Format

wazctl uses YAML configuration files with three main sections:

```yaml
wazuh:          # Wazuh manager configuration
  # ... manager settings
indexer:        # OpenSearch indexer configuration  
  # ... indexer settings
local:          # Local development configuration
  # ... local settings
```

## 🏠 Configuration File Locations

wazctl searches for configuration files in order:

1. **Current Directory**: `./.wazctl.yaml`
2. **User Home**: `~/.wazctl.yaml`  
3. **Config Directory**: `~/.config/wazctl.yaml`

The first file found is used. Create configuration files in your preferred location.

## 🛡️ Wazuh Manager Configuration

### Required Settings

```yaml
wazuh:
  endpoint: "wazuh-manager.company.com"    # Wazuh manager hostname/IP
  port: "55000"                           # API port (usually 55000)
  protocol: "https"                       # Protocol (http/https)
  wuiUsername: "admin"                    # API username
  wuiPassword: "secure_password"          # API password
```

### Optional Settings

```yaml
wazuh:
  # Security settings
  skipTlsVerify: false                    # Skip TLS certificate verification
  httpDebug: false                        # Enable HTTP debug logging
  
  # Future settings (planned)
  timeout: "30s"                          # Request timeout
  retries: 3                              # Number of retries
  rateLimit: 300                          # Requests per minute
```

### Complete Wazuh Section Example

```yaml
wazuh:
  # Connection settings
  endpoint: "wazuh-prod.company.com"
  port: "55000"
  protocol: "https"
  
  # Authentication
  wuiUsername: "wazctl_service"
  wuiPassword: "SuperSecurePassword123!"
  
  # Security (production)
  skipTlsVerify: false
  httpDebug: false
  
  # Development override
  # skipTlsVerify: true
  # httpDebug: true
```

## 🔍 OpenSearch Indexer Configuration

### Required Settings

```yaml
indexer:
  endpoint: "opensearch.company.com"      # OpenSearch hostname/IP
  port: "9200"                           # OpenSearch port
  protocol: "https"                      # Protocol (http/https)
  indexerUsername: "admin"               # OpenSearch username
  indexerPassword: "secure_password"     # OpenSearch password
```

### Optional Settings

```yaml
indexer:
  # Security settings
  skipTlsVerify: false                   # Skip TLS certificate verification
  httpDebug: false                       # Enable HTTP debug logging
  
  # Future settings (planned)
  timeout: "30s"                         # Request timeout
  index: "wazuh-alerts-*"               # Default index pattern
```

### Complete Indexer Section Example

```yaml
indexer:
  # Connection settings
  endpoint: "opensearch-cluster.company.com"
  port: "9200"
  protocol: "https"
  
  # Authentication
  indexerUsername: "wazctl_indexer"
  indexerPassword: "IndexerSecurePass456!"
  
  # Security (production)
  skipTlsVerify: false
  httpDebug: false
```

## 🏗️ Local Development Configuration

### Required Settings

```yaml
local:
  repoVersion: "v4.12.0"                 # Wazuh version for Docker environment
```

### Optional Settings

```yaml
local:
  # Docker settings (planned)
  dockerRegistry: "docker.io"           # Docker registry
  namespace: "wazuh"                     # Docker namespace
  pullPolicy: "IfNotPresent"            # Image pull policy
  
  # Resource settings (planned)
  memory: "4Gi"                         # Memory allocation
  cpu: "2000m"                          # CPU allocation
```

### Complete Local Section Example

```yaml
local:
  # Wazuh version
  repoVersion: "v4.12.0"
  
  # Future Docker customization
  # dockerRegistry: "private-registry.company.com"
  # namespace: "security-tools" 
  # memory: "8Gi"
  # cpu: "4000m"
```

## 📄 Complete Configuration Examples

### Production Environment

```yaml
# Production configuration example
wazuh:
  endpoint: "wazuh-prod.company.com"
  port: "55000" 
  protocol: "https"
  wuiUsername: "wazctl_prod"
  wuiPassword: "ProductionSecurePass123!"
  skipTlsVerify: false
  httpDebug: false

indexer:
  endpoint: "opensearch-prod.company.com"
  port: "9200"
  protocol: "https" 
  indexerUsername: "wazctl_prod"
  indexerPassword: "ProdIndexerPass456!"
  skipTlsVerify: false
  httpDebug: false

local:
  repoVersion: "v4.12.0"
```

### Development Environment

```yaml
# Development configuration example
wazuh:
  endpoint: "wazuh-dev.company.com"
  port: "55000"
  protocol: "https"
  wuiUsername: "developer"
  wuiPassword: "DevPassword123!"
  skipTlsVerify: true      # Allow self-signed certificates
  httpDebug: true          # Enable debugging

indexer:
  endpoint: "opensearch-dev.company.com"
  port: "9200"
  protocol: "https"
  indexerUsername: "developer"  
  indexerPassword: "DevIndexerPass456!"
  skipTlsVerify: true      # Allow self-signed certificates
  httpDebug: true          # Enable debugging

local:
  repoVersion: "v4.11.0"   # Use older version for compatibility testing
```

### Local Docker Environment

```yaml
# Local Docker environment configuration
wazuh:
  endpoint: "localhost"
  port: "55000"
  protocol: "https"
  wuiUsername: "admin"
  wuiPassword: "admin"      # Default Docker credentials
  skipTlsVerify: true       # Required for local development
  httpDebug: false

indexer:
  endpoint: "localhost"
  port: "9200"
  protocol: "https"
  indexerUsername: "admin"
  indexerPassword: "admin"  # Default Docker credentials
  skipTlsVerify: true       # Required for local development
  httpDebug: false

local:
  repoVersion: "v4.12.0"
```

## 🔧 Environment-Specific Configurations

### Multi-Environment Setup

```bash
# Create environment-specific configs
cp .wazctl.yaml .wazctl.prod.yaml
cp .wazctl.yaml .wazctl.dev.yaml
cp .wazctl.yaml .wazctl.local.yaml

# Switch environments
cp .wazctl.prod.yaml .wazctl.yaml   # Use production
cp .wazctl.dev.yaml .wazctl.yaml    # Use development
cp .wazctl.local.yaml .wazctl.yaml  # Use local Docker
```

### Environment Switching Script

```bash
#!/bin/bash
# wazctl-env.sh - Environment switching script

switch_environment() {
    local env=$1
    local config_file=".wazctl.${env}.yaml"
    
    if [ ! -f "$config_file" ]; then
        echo "❌ Configuration file not found: $config_file"
        exit 1
    fi
    
    # Backup current config
    if [ -f ".wazctl.yaml" ]; then
        cp .wazctl.yaml .wazctl.backup.yaml
        echo "📄 Current config backed up to .wazctl.backup.yaml"
    fi
    
    # Switch to new environment
    cp "$config_file" .wazctl.yaml
    echo "✅ Switched to $env environment"
    
    # Test connection
    echo "🔍 Testing connection..."
    if wazctl test auth > /dev/null; then
        echo "✅ Connection successful"
    else
        echo "❌ Connection failed - check configuration"
    fi
}

case $1 in
    "prod"|"production")
        switch_environment "prod"
        ;;
    "dev"|"development")
        switch_environment "dev"
        ;;
    "local")
        switch_environment "local"
        ;;
    "list")
        echo "Available environments:"
        ls .wazctl.*.yaml 2>/dev/null | sed 's/.wazctl\.\(.*\)\.yaml/  - \1/'
        ;;
    *)
        echo "Usage: $0 {prod|dev|local|list}"
        echo "Example: $0 prod"
        exit 1
        ;;
esac
```

## 🔒 Security Best Practices

### Credential Management

```yaml
# ❌ BAD: Hardcoded passwords
wazuh:
  wuiPassword: "admin123"

# ✅ GOOD: Use strong passwords
wazuh:
  wuiPassword: "Wz#K9$mP@3nQ7&vL2sX8"
```

### File Permissions

```bash
# Set secure permissions
chmod 600 .wazctl.yaml
chmod 600 ~/.wazctl.yaml

# Verify permissions
ls -la .wazctl.yaml
# Should show: -rw------- (600)
```

### Environment Variables (Planned)

Future versions will support environment variables:

```yaml
# Future: Environment variable support
wazuh:
  endpoint: "${WAZUH_ENDPOINT}"
  wuiUsername: "${WAZUH_USERNAME}"
  wuiPassword: "${WAZUH_PASSWORD}"
```

```bash
# Set environment variables
export WAZUH_ENDPOINT="wazuh-prod.company.com"
export WAZUH_USERNAME="wazctl_service"
export WAZUH_PASSWORD="$(cat /secure/wazuh-password)"
```

## 🐛 Configuration Validation

### Manual Validation

```bash
# Test YAML syntax
python -c "import yaml; yaml.safe_load(open('.wazctl.yaml'))"

# Test specific fields
yq eval '.wazuh.endpoint' .wazctl.yaml
yq eval '.indexer.port' .wazctl.yaml

# Validate configuration
wazctl test auth
```

### Validation Script

```bash
#!/bin/bash
# validate-config.sh - Configuration validation script

validate_config() {
    local config_file="${1:-.wazctl.yaml}"
    
    echo "🔍 Validating configuration: $config_file"
    
    # Check file exists
    if [ ! -f "$config_file" ]; then
        echo "❌ Configuration file not found: $config_file"
        return 1
    fi
    
    # Check file permissions
    local permissions=$(stat -c "%a" "$config_file" 2>/dev/null || stat -f "%A" "$config_file" 2>/dev/null)
    if [ "$permissions" != "600" ]; then
        echo "⚠️  Warning: File permissions are $permissions, should be 600"
        echo "   Run: chmod 600 $config_file"
    fi
    
    # Validate YAML syntax
    if python -c "import yaml; yaml.safe_load(open('$config_file'))" 2>/dev/null; then
        echo "✅ YAML syntax is valid"
    else
        echo "❌ YAML syntax error"
        return 1
    fi
    
    # Check required fields
    local required_fields=(
        ".wazuh.endpoint"
        ".wazuh.port" 
        ".wazuh.protocol"
        ".wazuh.wuiUsername"
        ".wazuh.wuiPassword"
        ".local.repoVersion"
    )
    
    for field in "${required_fields[@]}"; do
        if yq eval "$field" "$config_file" | grep -q "null"; then
            echo "❌ Missing required field: $field"
            return 1
        else
            echo "✅ Required field present: $field"
        fi
    done
    
    echo "✅ Configuration validation passed"
    return 0
}

# Validate current config
validate_config
```

## 📚 Configuration Templates

### Minimal Configuration

```yaml
# Minimal working configuration
wazuh:
  endpoint: "your-wazuh-manager.com"
  port: "55000"
  protocol: "https"
  wuiUsername: "admin"
  wuiPassword: "your-password"
  skipTlsVerify: true

local:
  repoVersion: "v4.12.0"
```

### Complete Configuration Template

```yaml
# Complete configuration template with all options
wazuh:
  # Required connection settings
  endpoint: "wazuh-manager.company.com"
  port: "55000"
  protocol: "https"
  
  # Required authentication
  wuiUsername: "wazctl_user"
  wuiPassword: "secure_password_here"
  
  # Security settings
  skipTlsVerify: false              # true for dev, false for prod
  httpDebug: false                  # true for debugging only

indexer:
  # Required connection settings
  endpoint: "opensearch.company.com"
  port: "9200"
  protocol: "https"
  
  # Required authentication
  indexerUsername: "wazctl_indexer"
  indexerPassword: "secure_indexer_password"
  
  # Security settings
  skipTlsVerify: false              # true for dev, false for prod
  httpDebug: false                  # true for debugging only

local:
  # Required for Docker environment
  repoVersion: "v4.12.0"           # Wazuh version to deploy
  
  # Future options (commented out)
  # dockerRegistry: "docker.io"
  # namespace: "wazuh"
  # memory: "4Gi"
  # cpu: "2000m"
```

## 🔧 Troubleshooting Configuration

### Common Issues

1. **YAML Syntax Errors**:
   ```bash
   # Check indentation (spaces, not tabs)
   # Validate quotes and colons
   python -c "import yaml; yaml.safe_load(open('.wazctl.yaml'))"
   ```

2. **File Not Found**:
   ```bash
   # Check file locations
   ls -la .wazctl.yaml ~/.wazctl.yaml ~/.config/wazctl.yaml
   ```

3. **Permission Issues**:
   ```bash
   # Fix permissions
   chmod 600 .wazctl.yaml
   ```

4. **Authentication Failures**:
   ```bash
   # Test authentication
   wazctl test auth
   ```

### Debug Configuration

```yaml
# Enable debugging for troubleshooting
wazuh:
  httpDebug: true                   # Shows HTTP requests/responses
  
indexer:
  httpDebug: true                   # Shows HTTP requests/responses
```

## ➡️ Next Steps

- Review [Authentication Guide](authentication.md) for connection setup
- Check [Troubleshooting](troubleshooting.md) for common issues
- Explore [Command Reference](command-reference.md) for usage examples