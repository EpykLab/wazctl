# 🎯 Quick Start Tutorial

This tutorial will get you up and running with wazctl in just a few minutes, covering the most common workflows.

## 📋 Prerequisites

- ✅ wazctl installed ([Installation Guide](installation.md))
- ✅ Access to a Wazuh manager
- ✅ Basic understanding of Wazuh concepts

## 🚀 Step-by-Step Walkthrough

### Step 1: Initial Setup

1. **Create configuration file**:
   ```bash
   wazctl init config
   ```

2. **Edit the configuration** with your Wazuh details:
   ```bash
   # Edit the generated .wazctl.yaml file
   nano .wazctl.yaml
   ```

3. **Test authentication**:
   ```bash
   wazctl test auth
   ```
   
   ✅ **Success**: You'll see a JWT token  
   ❌ **Failure**: Check your credentials and network connectivity

### Step 2: Explore Your Wazuh Environment

1. **List all agents**:
   ```bash
   wazctl api agents list
   ```
   
   This shows all agents enrolled in your Wazuh manager with their status, IP addresses, and versions.

2. **Alternative agent listing**:
   ```bash
   wazctl agents list
   ```
   
   Both commands provide the same functionality through different command paths.

### Step 3: Create Your First Rule Test

1. **Generate a rule test template**:
   ```bash
   wazctl init rule --name "suspicious_login_test"
   ```

2. **View the generated file**:
   ```bash
   cat suspicious_login_test.yaml
   ```

   You'll see a structured test file with:
   - Rule metadata (ID, name, author)
   - Example Wazuh rule content
   - Test scenarios ("edges")
   - Expected outcomes

3. **Customize the test** for your specific rule requirements.

### Step 4: User Management

Create users in your Wazuh ecosystem:

1. **Create a Wazuh manager user**:
   ```bash
   wazctl user add --username alice --password secure123 --component wazuh
   ```

2. **Create an indexer user with role**:
   ```bash
   wazctl user add --username bob --password secure456 --component indexer --role analyst
   ```

### Step 5: Local Development Environment

Set up a local Wazuh instance for testing:

1. **Start a local Wazuh instance**:
   ```bash
   wazctl localenv docker --start
   ```

2. **Stop the instance when done**:
   ```bash
   wazctl localenv docker --stop
   ```

3. **Clean up resources**:
   ```bash
   wazctl localenv docker --clean
   ```

## 🎯 Common Workflows

### 📊 Monitoring Workflow

```bash
# Check authentication
wazctl test auth

# Review agent status
wazctl api agents list

# Identify inactive agents for investigation
wazctl api agents list | grep -i "never_connected\|disconnected"
```

### 🔍 Rule Development Workflow

```bash
# Create rule test template
wazctl init rule --name "my_custom_rule_test"

# Edit the test file with your rule logic
nano my_custom_rule_test.yaml

# Test against local environment
wazctl localenv docker --start
# (Rule execution testing planned for future releases)
```

### 👥 User Administration Workflow

```bash
# Create analyst user in Wazuh
wazctl user add --username analyst1 --password temp123 --component wazuh

# Create corresponding indexer user
wazctl user add --username analyst1 --password temp123 --component indexer --role kibana_user

# Test authentication with new credentials
# (Update .wazctl.yaml with new credentials and test)
```

## 💡 Pro Tips

### 🔍 Output Formatting

All commands output JSON for easy processing:

```bash
# Pretty print JSON output
wazctl api agents list | jq '.'

# Filter specific fields
wazctl api agents list | jq '.data.affected_items[].name'

# Count total agents
wazctl api agents list | jq '.data.total_affected_items'
```

### 📁 Configuration Management

Keep multiple configurations for different environments:

```bash
# Development environment
cp .wazctl.yaml .wazctl.dev.yaml

# Production environment  
cp .wazctl.yaml .wazctl.prod.yaml

# Switch configurations as needed
cp .wazctl.prod.yaml .wazctl.yaml
```

### 🐛 Debugging

Enable debug mode for troubleshooting:

```yaml
# In .wazctl.yaml
wazuh:
    httpDebug: true
indexer:
    httpDebug: true
```

## 🚨 Common Issues & Quick Fixes

### Authentication Fails
```bash
# Check connection
curl -k https://your-wazuh-manager:55000

# Verify credentials
wazctl test auth
```

### No Agents Listed
- Ensure agents are properly enrolled
- Check agent status in Wazuh web interface
- Verify API user has agent read permissions

### Docker Issues
```bash
# Check Docker status
docker ps

# View wazctl Docker logs
docker logs $(docker ps | grep wazuh | awk '{print $1}')
```

## ➡️ Next Steps

Now that you're familiar with the basics:

- 📖 Explore detailed feature guides:
  - [Agent Management](agents.md)
  - [Rule Testing](rule-testing.md)
  - [User Management](users.md)
- 🔧 Check the [Command Reference](command-reference.md) for all available options
- 🐛 Need help? See [Troubleshooting](troubleshooting.md)