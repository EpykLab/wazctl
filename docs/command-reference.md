# 📋 Command Reference

This comprehensive reference covers all wazctl commands, options, and usage patterns.

## 🎯 Command Structure

wazctl follows a hierarchical command structure:

```
wazctl <command> <subcommand> [flags] [arguments]
```

## 🏠 Root Command

### `wazctl`

Base command for all wazctl operations.

```bash
wazctl [command]
```

**Global Flags**:
- `--help`, `-h`: Show help for any command
- `--toggle`, `-t`: Help message for toggle (legacy)

**Available Commands**:
- `agents`: Agent management operations
- `api`: Direct API interactions  
- `init`: Initialize configurations and templates
- `localenv`: Local environment management
- `test`: Testing and validation commands
- `user`: User management operations

## 🛠️ Initialization Commands

### `wazctl init`

Initialize and create templates.

```bash
wazctl init <subcommand> [flags]
```

#### `wazctl init config`

Create a wazctl configuration file.

```bash
wazctl init config
```

**Description**: Creates a `.wazctl.yaml` configuration file in the current directory with default values.

**Output**: Creates `.wazctl.yaml` with template configuration including:
- Wazuh manager settings
- OpenSearch indexer settings
- Local environment settings

**Example**:
```bash
wazctl init config
# Creates .wazctl.yaml with default configuration
```

#### `wazctl init rule`

Create a rule test template.

```bash
wazctl init rule --name <rule_name>
```

**Required Flags**:
- `--name`, `-n`: Name for the new rule test file

**Description**: Creates a YAML file with structured rule test template including metadata, rule content, and test scenarios.

**Examples**:
```bash
# Create basic rule test
wazctl init rule --name "ssh_brute_force_test"

# Create with short flag
wazctl init rule -n "web_attack_test"
```

**Output**: Creates `<rule_name>.yaml` with:
- Rule metadata (ID, name, author)
- Example rule content
- Test scenarios template

## 🔐 Authentication Commands

### `wazctl test`

Testing and validation operations.

```bash
wazctl test <subcommand>
```

#### `wazctl test auth`

Test authentication with Wazuh manager.

```bash
wazctl test auth
```

**Description**: Validates configuration and tests authentication with the Wazuh API. Returns JWT token on success.

**Prerequisites**: Valid `.wazctl.yaml` configuration file

**Success Output**: JWT token string
**Failure Output**: Error message with details

**Examples**:
```bash
# Test authentication
wazctl test auth

# Use in scripts
TOKEN=$(wazctl test auth)
```

## 👥 Agent Management Commands

### `wazctl agents`

Agent management operations.

```bash
wazctl agents <subcommand>
```

#### `wazctl agents list`

List all agents enrolled in Wazuh manager.

```bash
wazctl agents list
```

**Description**: Retrieves and displays all agents with their status, IP addresses, versions, and other metadata.

**Output Format**: JSON with agent details including:
- Agent ID and name
- Connection status
- IP address and OS information
- Last keep-alive time
- Version information

**Examples**:
```bash
# List all agents
wazctl agents list

# Filter with jq
wazctl agents list | jq '.data.affected_items[].name'

# Show only disconnected agents
wazctl agents list | jq '.data.affected_items[] | select(.status=="disconnected")'
```

## 🔌 API Commands

### `wazctl api`

Direct Wazuh API interactions.

```bash
wazctl api <subcommand>
```

#### `wazctl api agents`

API-based agent operations.

```bash
wazctl api agents <subcommand>
```

##### `wazctl api agents list`

List agents via direct API call.

```bash
wazctl api agents list
```

**Description**: Alternative path to list agents. Functionally identical to `wazctl agents list`.

**Output**: Same JSON format as `wazctl agents list`

## 👤 User Management Commands

### `wazctl user`

User management operations.

```bash
wazctl user <subcommand>
```

#### `wazctl user add`

Create new users in Wazuh or OpenSearch components.

```bash
wazctl user add --username <username> --password <password> --component <component> [--role <role>]
```

**Required Flags**:
- `--username`, `-u`: Username for the new user
- `--password`, `-p`: Password for the new user
- `--component`, `-c`: Component to create user in (`wazuh` or `indexer`)

**Conditional Flags**:
- `--role`, `-r`: Role to assign (required for `indexer` component)

**Components**:
- `wazuh`: Create user in Wazuh manager
- `indexer`: Create user in OpenSearch indexer

**Common Indexer Roles**:
- `kibana_user`: Dashboard access
- `security_analyst`: Security analysis access
- `readonly`: Read-only access
- `all_access`: Full administrative access

**Examples**:
```bash
# Create Wazuh manager user
wazctl user add --username analyst1 --password SecurePass123! --component wazuh

# Create indexer user with role
wazctl user add --username analyst2 --password SecurePass123! --component indexer --role kibana_user

# Using short flags
wazctl user add -u admin -p AdminPass123! -c wazuh
wazctl user add -u viewer -p ViewPass123! -c indexer -r readonly
```

## 🐳 Local Environment Commands

### `wazctl localenv`

Local development environment management.

```bash
wazctl localenv <subcommand>
```

#### `wazctl localenv docker`

Docker-based local Wazuh environment.

```bash
wazctl localenv docker [flags]
```

**Flags** (use one at a time):
- `--start`: Start the Wazuh Docker environment
- `--stop`: Stop the running environment
- `--clean`: Remove all containers, networks, and volumes

**Description**: Manages a complete local Wazuh stack including:
- Wazuh Manager
- OpenSearch indexer
- OpenSearch Dashboards
- Wazuh Dashboard

**Examples**:
```bash
# Start environment
wazctl localenv docker --start

# Stop environment (preserves data)
wazctl localenv docker --stop

# Clean environment (removes all data)
wazctl localenv docker --clean
```

**Default Access Points** (when started):
- Wazuh Dashboard: https://localhost:443
- OpenSearch Dashboards: https://localhost:5601
- Wazuh API: https://localhost:55000
- OpenSearch API: https://localhost:9200

## 🔍 Command Usage Patterns

### Basic Workflow

```bash
# 1. Initialize configuration
wazctl init config

# 2. Edit configuration with your details
nano .wazctl.yaml

# 3. Test authentication
wazctl test auth

# 4. List agents
wazctl agents list

# 5. Create rule test
wazctl init rule --name "my_test"
```

### Local Development Workflow

```bash
# 1. Start local environment
wazctl localenv docker --start

# 2. Wait for services
sleep 60

# 3. Test connection
wazctl test auth

# 4. Develop and test rules
wazctl init rule --name "development_test"

# 5. Stop when done
wazctl localenv docker --stop
```

### User Management Workflow

```bash
# Create Wazuh user
wazctl user add -u analyst -p SecurePass123! -c wazuh

# Create corresponding indexer user
wazctl user add -u analyst -p SecurePass123! -c indexer -r security_analyst

# Test authentication (update config first)
wazctl test auth
```

## 📊 Output Formats and Processing

### JSON Output

All data commands output JSON for easy processing:

```bash
# Pretty print
wazctl agents list | jq '.'

# Extract specific fields
wazctl agents list | jq '.data.affected_items[].name'

# Filter and count
wazctl agents list | jq '.data.affected_items | length'

# Complex filtering
wazctl agents list | jq '.data.affected_items[] | select(.status=="active") | {name: .name, ip: .ip}'
```

### Error Handling

Commands return appropriate exit codes:
- `0`: Success
- `1`: General error
- Authentication errors, network issues, etc.

### Text Processing

```bash
# Using grep
wazctl agents list | grep '"status": "active"'

# Using awk
wazctl agents list | awk '/status.*active/'

# Counting with wc
wazctl agents list | grep '"status": "disconnected"' | wc -l
```

## 🛠️ Advanced Usage

### Configuration Management

```bash
# Multiple environments
cp .wazctl.yaml .wazctl.prod.yaml
cp .wazctl.yaml .wazctl.dev.yaml

# Switch configurations
cp .wazctl.prod.yaml .wazctl.yaml
wazctl test auth
```

### Scripting and Automation

```bash
#!/bin/bash
# Health check script

echo "Testing authentication..."
if wazctl test auth > /dev/null; then
    echo "✅ Authentication successful"
else
    echo "❌ Authentication failed"
    exit 1
fi

echo "Checking agents..."
TOTAL=$(wazctl agents list | jq '.data.total_affected_items')
ACTIVE=$(wazctl agents list | jq '.data.affected_items[] | select(.status=="active")' | wc -l)

echo "Total agents: $TOTAL"
echo "Active agents: $ACTIVE"
```

### Batch Operations

```bash
# Bulk user creation
users=("analyst1" "analyst2" "analyst3")
for user in "${users[@]}"; do
    wazctl user add -u "$user" -p "TempPass123!" -c wazuh
    wazctl user add -u "$user" -p "TempPass123!" -c indexer -r kibana_user
done
```

## 🐛 Troubleshooting Commands

### Debug Information

```bash
# Test connectivity
curl -k https://your-wazuh-manager:55000

# Check configuration
cat .wazctl.yaml

# Verify Docker environment
docker ps | grep wazuh
```

### Common Issues

```bash
# Authentication fails
wazctl test auth
# Check credentials, network, and service status

# No agents listed
wazctl agents list
# Verify agent enrollment and API permissions

# Docker issues
docker logs $(docker ps | grep wazuh | awk '{print $1}')
# Check container logs for errors
```

## 📚 Help and Documentation

### Built-in Help

```bash
# General help
wazctl --help

# Command-specific help
wazctl agents --help
wazctl user add --help
wazctl init rule --help

# Subcommand help
wazctl localenv docker --help
```

### Version Information

```bash
# Check version (when available)
wazctl version

# Check build information
wazctl --version
```

## 🔄 Future Commands (Planned)

Commands planned for future releases:

```bash
# Agent management
wazctl agents restart <agent_id>
wazctl agents update <agent_id>
wazctl agents remove <agent_id>

# User management
wazctl user list
wazctl user delete <username>
wazctl user modify <username> --role <new_role>

# Rule management
wazctl rules list
wazctl rules test <rule_file>
wazctl rules deploy <rule_file>

# Configuration management
wazctl config validate
wazctl config switch <profile>
wazctl config list

# Monitoring
wazctl status
wazctl logs --follow
wazctl metrics

# Backup and restore
wazctl backup create
wazctl backup restore <backup_file>
```

---

> 💡 **Note**: This reference covers wazctl v0.4.0. Some commands may change in future versions. Check the built-in help (`--help`) for the most current information.