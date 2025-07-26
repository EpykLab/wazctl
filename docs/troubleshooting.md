# 🔧 Troubleshooting

This guide helps you diagnose and resolve common issues with wazctl. Most problems fall into a few categories: authentication, network connectivity, configuration, or environment setup.

## 🎯 Quick Diagnostics

Before diving into specific issues, run these quick diagnostic commands:

```bash
# Test basic functionality
wazctl --help

# Check configuration
cat .wazctl.yaml

# Test authentication
wazctl test auth

# Check network connectivity
curl -k https://your-wazuh-manager:55000

# Check Docker (if using local environment)
docker ps | grep wazuh
```

## 🔐 Authentication Issues

### Problem: "authentication failed"

```
Error: authentication failed
```

**Possible Causes**:
- Incorrect username or password
- Wrong endpoint configuration
- Account locked or disabled
- API service not running

**Solutions**:

1. **Verify Credentials**:
   ```bash
   # Check configuration
   cat .wazctl.yaml
   
   # Verify credentials in Wazuh web interface
   # https://your-wazuh-manager/
   ```

2. **Test API Endpoint**:
   ```bash
   # Test basic connectivity
   curl -k https://your-wazuh-manager:55000
   
   # Should return Wazuh API information
   ```

3. **Check Service Status**:
   ```bash
   # On Wazuh manager server
   systemctl status wazuh-manager
   systemctl status wazuh-api
   ```

4. **Debug Authentication**:
   ```bash
   # Enable debug mode in .wazctl.yaml
   wazuh:
       httpDebug: true
   
   # Run auth test with debug output
   wazctl test auth
   ```

### Problem: "connection refused"

```
Error: connection refused
```

**Possible Causes**:
- Wrong endpoint or port
- Firewall blocking connection
- Service not running
- Network issues

**Solutions**:

1. **Verify Network Connectivity**:
   ```bash
   # Test port connectivity
   telnet your-wazuh-manager 55000
   
   # Or using nc
   nc -zv your-wazuh-manager 55000
   ```

2. **Check Firewall**:
   ```bash
   # On Wazuh manager, check if port is open
   netstat -tulpn | grep :55000
   
   # Check firewall rules
   iptables -L | grep 55000
   ```

3. **Verify Configuration**:
   ```bash
   # Check endpoint and port in config
   grep -E "(endpoint|port)" .wazctl.yaml
   ```

### Problem: "certificate signed by unknown authority"

```
Error: x509: certificate signed by unknown authority
```

**Solutions**:

1. **For Development/Testing**:
   ```yaml
   # In .wazctl.yaml
   wazuh:
       skipTlsVerify: true
   ```

2. **For Production**:
   ```bash
   # Install proper SSL certificates on Wazuh manager
   # Or add CA certificate to system trust store
   
   # On Linux:
   sudo cp wazuh-ca.crt /usr/local/share/ca-certificates/
   sudo update-ca-certificates
   ```

## 🌐 Network and Connectivity Issues

### Problem: Slow or timeout responses

**Solutions**:

1. **Check Network Latency**:
   ```bash
   # Test round-trip time
   ping your-wazuh-manager
   
   # Test HTTP response time
   time curl -k https://your-wazuh-manager:55000
   ```

2. **Increase Timeout** (future feature):
   ```yaml
   # In .wazctl.yaml (planned)
   wazuh:
       timeout: 30s
   ```

3. **Check DNS Resolution**:
   ```bash
   # Verify hostname resolves correctly
   nslookup your-wazuh-manager
   
   # Or use IP address directly
   ```

### Problem: Connection works sometimes, fails others

**Possible Causes**:
- Load balancer issues
- Intermittent network problems
- API rate limiting

**Solutions**:

1. **Check API Rate Limits**:
   ```bash
   # Wazuh API default: 300 requests/minute
   # Add delays between requests in scripts
   
   sleep 1  # Add 1-second delay between commands
   ```

2. **Monitor API Logs**:
   ```bash
   # On Wazuh manager
   tail -f /var/ossec/logs/api.log
   ```

## ⚙️ Configuration Issues

### Problem: "no valid config file found"

```
Error: no valid config file found in locations: [.wazctl.yaml ~/.wazctl.yaml ~/.config/wazctl.yaml]
```

**Solutions**:

1. **Create Configuration**:
   ```bash
   wazctl init config
   ```

2. **Check File Locations**:
   ```bash
   # Look for existing config files
   ls -la .wazctl.yaml
   ls -la ~/.wazctl.yaml
   ls -la ~/.config/wazctl.yaml
   ```

3. **Verify File Permissions**:
   ```bash
   # Ensure file is readable
   chmod 600 .wazctl.yaml
   ```

### Problem: Configuration syntax errors

**Solutions**:

1. **Validate YAML Syntax**:
   ```bash
   # Using Python
   python -c "import yaml; yaml.safe_load(open('.wazctl.yaml'))"
   
   # Using yq (if installed)
   yq eval '.wazuh.endpoint' .wazctl.yaml
   ```

2. **Check Indentation**:
   ```yaml
   # Correct indentation (spaces, not tabs)
   wazuh:
       endpoint: your-instance.com
       port: "55000"
   ```

3. **Reset Configuration**:
   ```bash
   # Start fresh
   mv .wazctl.yaml .wazctl.yaml.backup
   wazctl init config
   # Manually copy your settings
   ```

## 👥 Agent Management Issues

### Problem: No agents listed

```json
{
  "data": {
    "affected_items": [],
    "total_affected_items": 0
  }
}
```

**Possible Causes**:
- No agents enrolled
- API user lacks permissions
- Agents not properly configured

**Solutions**:

1. **Check Agent Enrollment**:
   ```bash
   # On Wazuh manager
   /var/ossec/bin/agent_control -l
   ```

2. **Verify API Permissions**:
   ```bash
   # Check user permissions in Wazuh web interface
   # User should have agent read permissions
   ```

3. **Check Agent Status**:
   ```bash
   # On agent machine
   systemctl status wazuh-agent
   
   # Check agent logs
   tail -f /var/ossec/logs/ossec.log
   ```

### Problem: Agents show as "never_connected"

**Solutions**:

1. **Check Agent Configuration**:
   ```bash
   # On agent machine, verify manager IP
   grep MANAGER_IP /var/ossec/etc/ossec.conf
   ```

2. **Test Network Connectivity**:
   ```bash
   # From agent to manager on port 1514
   telnet wazuh-manager 1514
   ```

3. **Check Firewall Rules**:
   ```bash
   # Ensure port 1514 is open between agent and manager
   ```

## 🐳 Docker Environment Issues

### Problem: Docker environment won't start

**Solutions**:

1. **Check Docker Status**:
   ```bash
   # Verify Docker is running
   docker info
   
   # Check available resources
   docker system df
   ```

2. **Free Up Resources**:
   ```bash
   # Clean up Docker resources
   docker system prune
   
   # Remove old containers
   docker container prune
   ```

3. **Check Port Conflicts**:
   ```bash
   # Check if ports are already in use
   netstat -tulpn | grep -E ':(443|5601|9200|55000)'
   
   # Stop conflicting services
   sudo systemctl stop apache2  # If using port 443
   ```

### Problem: Services not responding after start

**Solutions**:

1. **Wait Longer**:
   ```bash
   # Services need time to initialize
   sleep 120
   wazctl test auth
   ```

2. **Check Container Status**:
   ```bash
   # View running containers
   docker ps
   
   # Check container health
   docker inspect <container_id> | grep -i health
   ```

3. **View Container Logs**:
   ```bash
   # Check logs for errors
   docker logs <container_name>
   
   # Follow logs in real-time
   docker logs -f <container_name>
   ```

### Problem: Docker cleanup fails

**Solutions**:

1. **Force Container Removal**:
   ```bash
   # Stop all containers
   docker stop $(docker ps -q)
   
   # Force remove containers
   docker rm -f $(docker ps -aq)
   ```

2. **Remove Volumes**:
   ```bash
   # List volumes
   docker volume ls
   
   # Remove specific volumes
   docker volume rm <volume_name>
   
   # Remove all unused volumes
   docker volume prune
   ```

## 👤 User Management Issues

### Problem: User creation fails

**Solutions**:

1. **Check Required Parameters**:
   ```bash
   # Ensure all required flags are provided
   wazctl user add --username test --password pass123 --component wazuh
   
   # For indexer, role is required
   wazctl user add --username test --password pass123 --component indexer --role kibana_user
   ```

2. **Verify Permissions**:
   ```bash
   # Ensure your API user can create users
   # Check permissions in Wazuh web interface
   ```

### Problem: "user already exists"

**Solutions**:

1. **Choose Different Username**:
   ```bash
   # Use unique username
   wazctl user add --username analyst_new --password pass123 --component wazuh
   ```

2. **Check Existing Users**:
   ```bash
   # Manually check users in Wazuh web interface
   # Plan: wazctl user list (future feature)
   ```

## 📊 Performance Issues

### Problem: Slow command execution

**Solutions**:

1. **Check System Resources**:
   ```bash
   # Monitor system load
   top
   htop
   
   # Check memory usage
   free -h
   ```

2. **Optimize Configuration**:
   ```yaml
   # Disable debug mode if enabled
   wazuh:
       httpDebug: false
   ```

3. **Network Optimization**:
   ```bash
   # Use IP instead of hostname
   # Ensure good network connection
   ```

## 🔍 Installation Issues

### Problem: "wazctl: command not found"

**Solutions**:

1. **Check Installation**:
   ```bash
   # Verify Go installation
   go version
   
   # Check GOPATH
   echo $GOPATH
   ```

2. **Update PATH**:
   ```bash
   # Add Go bin to PATH
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

3. **Reinstall**:
   ```bash
   go install github.com/EpykLab/wazctl@latest
   ```

### Problem: Go installation issues

**Solutions**:

1. **Install Go**:
   ```bash
   # Download and install Go from https://golang.org/dl/
   
   # Or using package manager
   sudo apt install golang-go  # Ubuntu/Debian
   sudo yum install golang      # CentOS/RHEL
   brew install go             # macOS
   ```

2. **Set Environment Variables**:
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

## 🛠️ Advanced Troubleshooting

### Debug Mode

Enable detailed logging for troubleshooting:

```yaml
# In .wazctl.yaml
wazuh:
    httpDebug: true
indexer:
    httpDebug: true
```

### Network Analysis

```bash
# Capture network traffic
sudo tcpdump -i any host your-wazuh-manager

# Analyze DNS resolution
dig your-wazuh-manager

# Test SSL connection
openssl s_client -connect your-wazuh-manager:55000
```

### System Information Collection

```bash
#!/bin/bash
# Diagnostic information script

echo "=== wazctl Diagnostic Information ==="
echo "Date: $(date)"
echo ""

echo "=== Environment ==="
echo "OS: $(uname -a)"
echo "Go version: $(go version 2>/dev/null || echo 'Go not installed')"
echo "Docker version: $(docker --version 2>/dev/null || echo 'Docker not installed')"
echo ""

echo "=== Configuration ==="
if [ -f .wazctl.yaml ]; then
    echo "Configuration found:"
    cat .wazctl.yaml | sed 's/password:.*/password: ***HIDDEN***/'
else
    echo "No configuration file found"
fi
echo ""

echo "=== Network Connectivity ==="
if command -v curl >/dev/null; then
    echo "Testing connectivity..."
    curl -k -I https://localhost:55000 2>&1 | head -5
else
    echo "curl not available"
fi
echo ""

echo "=== Docker Status ==="
if command -v docker >/dev/null; then
    echo "Docker containers:"
    docker ps | grep wazuh
else
    echo "Docker not available"
fi
```

## 📞 Getting Help

### Self-Help Resources

1. **Built-in Help**:
   ```bash
   wazctl --help
   wazctl <command> --help
   ```

2. **Documentation**:
   - [Quick Start Guide](quick-start.md)
   - [Configuration Guide](configuration.md)
   - [Command Reference](command-reference.md)

3. **Community Resources**:
   - GitHub Issues: Report bugs and get help
   - GitHub Discussions: Community support
   - Wazuh Community: General Wazuh help

### When to Seek Help

Contact the community when:
- You've tried the solutions in this guide
- You have a reproducible bug
- You need a feature that doesn't exist
- You want to contribute to the project

### Information to Include

When reporting issues, include:
- wazctl version (when available)
- Operating system and version
- Configuration file (with passwords redacted)
- Complete error messages
- Steps to reproduce the issue
- Output from diagnostic commands

### Bug Report Template

```
## Issue Description
Brief description of the problem

## Environment
- OS: [e.g., Ubuntu 20.04]
- wazctl version: [e.g., v0.4.0]
- Go version: [e.g., 1.24.3]

## Configuration
```yaml
# Paste configuration (redact passwords)
```

## Steps to Reproduce
1. Step one
2. Step two
3. Step three

## Expected Behavior
What you expected to happen

## Actual Behavior
What actually happened

## Error Messages
```
Paste complete error messages here
```

## Additional Context
Any other relevant information
```

---

> 💡 **Remember**: Most issues are configuration or network related. Check authentication and connectivity first before diving into complex troubleshooting.