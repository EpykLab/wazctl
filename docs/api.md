# 🔌 API Commands

wazctl provides direct access to Wazuh API endpoints through structured commands. This allows you to interact with the Wazuh manager programmatically while maintaining the convenience of the CLI interface.

## 🎯 Overview

The API command structure provides:
- **🔗 Direct API Access**: Raw access to Wazuh API endpoints
- **📊 Structured Output**: Consistent JSON formatting
- **🛡️ Authentication**: Automatic JWT token management
- **🔄 Error Handling**: Comprehensive error reporting

## 📋 Available API Commands

### `wazctl api agents`

Agent-related API operations.

#### `wazctl api agents list`

List all agents enrolled in the Wazuh manager.

```bash
wazctl api agents list
```

**Equivalent to**: Direct call to `/agents` endpoint

**Output Format**:
```json
{
  "data": {
    "affected_items": [
      {
        "id": "000",
        "name": "wazuh-manager",
        "ip": "127.0.0.1",
        "status": "active",
        "node_name": "node01",
        "dateAdd": "2024-01-15T10:30:00Z",
        "version": "4.12.0",
        "manager": "wazuh-manager",
        "os": {
          "arch": "x86_64",
          "major": "20",
          "minor": "04",
          "name": "Ubuntu",
          "platform": "ubuntu",
          "version": "20.04.6 LTS"
        },
        "lastKeepAlive": "2024-01-15T15:45:00Z",
        "group": ["default"]
      }
    ],
    "total_affected_items": 1,
    "total_failed_items": 0,
    "failed_items": []
  }
}
```

## 🔧 API Command Patterns

### Basic Usage

```bash
# List agents via API
wazctl api agents list

# Process output with jq
wazctl api agents list | jq '.data.affected_items[].name'
```

### Error Handling

API commands return structured error information:

```json
{
  "data": {
    "affected_items": [],
    "total_affected_items": 0,
    "total_failed_items": 1,
    "failed_items": [
      {
        "error": {
          "code": 1701,
          "message": "Invalid token"
        }
      }
    ]
  }
}
```

## 🚧 Planned API Commands

Future releases will include additional API endpoints:

### Agent Management (Planned)
```bash
# Agent details
wazctl api agents get <agent_id>

# Agent restart
wazctl api agents restart <agent_id>

# Agent update
wazctl api agents update <agent_id>

# Agent deletion
wazctl api agents delete <agent_id>
```

### Rule Management (Planned)
```bash
# List rules
wazctl api rules list

# Get specific rule
wazctl api rules get <rule_id>

# Create rule
wazctl api rules create --file <rule_file>

# Update rule
wazctl api rules update <rule_id> --file <rule_file>

# Delete rule
wazctl api rules delete <rule_id>
```

### Configuration Management (Planned)
```bash
# Get configuration
wazctl api config get

# Update configuration
wazctl api config update --file <config_file>

# Restart manager
wazctl api manager restart
```

### User Management (Planned)
```bash
# List users
wazctl api users list

# Get user details
wazctl api users get <username>

# Update user
wazctl api users update <username> --role <new_role>

# Delete user
wazctl api users delete <username>
```

## 📊 Integration Examples

### Monitoring Script

```bash
#!/bin/bash
# API-based monitoring script

echo "🔍 Wazuh System Status"
echo "====================="

# Test API connectivity
echo "📡 Testing API connectivity..."
if wazctl test auth > /dev/null; then
    echo "✅ API connection successful"
else
    echo "❌ API connection failed"
    exit 1
fi

# Get agent statistics
echo "👥 Agent Status Summary:"
AGENTS_DATA=$(wazctl api agents list)

if [ $? -eq 0 ]; then
    TOTAL=$(echo "$AGENTS_DATA" | jq '.data.total_affected_items')
    ACTIVE=$(echo "$AGENTS_DATA" | jq '.data.affected_items[] | select(.status=="active")' | wc -l)
    DISCONNECTED=$(echo "$AGENTS_DATA" | jq '.data.affected_items[] | select(.status=="disconnected")' | wc -l)
    NEVER_CONNECTED=$(echo "$AGENTS_DATA" | jq '.data.affected_items[] | select(.status=="never_connected")' | wc -l)
    
    echo "  Total Agents: $TOTAL"
    echo "  Active: $ACTIVE"
    echo "  Disconnected: $DISCONNECTED"
    echo "  Never Connected: $NEVER_CONNECTED"
    
    # Alert on disconnected agents
    if [ "$DISCONNECTED" -gt 0 ]; then
        echo "⚠️  Disconnected Agents:"
        echo "$AGENTS_DATA" | jq -r '.data.affected_items[] | select(.status=="disconnected") | "  - \(.name) (\(.ip))"'
    fi
else
    echo "❌ Failed to retrieve agent data"
fi
```

### Data Export

```bash
#!/bin/bash
# Export agent data via API

echo "📤 Exporting agent data..."

# Get agent data
AGENTS_DATA=$(wazctl api agents list)

if [ $? -eq 0 ]; then
    # Export to CSV
    echo "ID,Name,IP,Status,Version,OS,Last_Keep_Alive" > agents_export.csv
    echo "$AGENTS_DATA" | jq -r '.data.affected_items[] | [.id, .name, .ip, .status, .version, (.os.name + " " + .os.version), .lastKeepAlive] | @csv' >> agents_export.csv
    
    # Export to JSON
    echo "$AGENTS_DATA" | jq '.data.affected_items' > agents_export.json
    
    echo "✅ Data exported to:"
    echo "  - agents_export.csv"
    echo "  - agents_export.json"
else
    echo "❌ Failed to export agent data"
    exit 1
fi
```

### Health Check API

```bash
#!/bin/bash
# Comprehensive health check using API

perform_health_check() {
    local component=$1
    local endpoint=$2
    
    echo "🔍 Checking $component..."
    
    if curl -k -s "$endpoint" > /dev/null; then
        echo "✅ $component is responsive"
        return 0
    else
        echo "❌ $component is not responding"
        return 1
    fi
}

echo "🏥 Wazuh Health Check"
echo "===================="

# Check Wazuh API
if wazctl test auth > /dev/null; then
    echo "✅ Wazuh API authentication successful"
    
    # Check agent endpoint
    if wazctl api agents list > /dev/null; then
        echo "✅ Agent API endpoint functional"
    else
        echo "❌ Agent API endpoint failed"
    fi
else
    echo "❌ Wazuh API authentication failed"
fi

# Check other endpoints (when available)
perform_health_check "Wazuh Manager" "https://localhost:55000"
perform_health_check "OpenSearch" "https://localhost:9200"
perform_health_check "Wazuh Dashboard" "https://localhost:443"
```

## 🔄 API Response Processing

### Standard Response Format

All API responses follow the Wazuh API standard:

```json
{
  "data": {
    "affected_items": [...],      // Successful items
    "total_affected_items": 10,   // Count of successful items
    "total_failed_items": 0,      // Count of failed items
    "failed_items": []            // Failed items with error details
  }
}
```

### Processing Examples

```bash
# Extract all agent names
wazctl api agents list | jq -r '.data.affected_items[].name'

# Filter by status
wazctl api agents list | jq '.data.affected_items[] | select(.status=="active")'

# Count by status
wazctl api agents list | jq '.data.affected_items | group_by(.status) | map({status: .[0].status, count: length})'

# Get specific fields
wazctl api agents list | jq '.data.affected_items[] | {name: .name, ip: .ip, status: .status}'

# Check for errors
wazctl api agents list | jq '.data.failed_items[]'
```

## 🛠️ Advanced API Usage

### Scripting Best Practices

```bash
#!/bin/bash
# Best practices for API scripting

# Function to check API response
check_api_response() {
    local response="$1"
    local failed_count=$(echo "$response" | jq '.data.total_failed_items')
    
    if [ "$failed_count" -gt 0 ]; then
        echo "⚠️  API operation had $failed_count failed items:"
        echo "$response" | jq '.data.failed_items'
        return 1
    fi
    
    return 0
}

# Function to retry API calls
retry_api_call() {
    local max_attempts=3
    local attempt=1
    
    while [ $attempt -le $max_attempts ]; do
        echo "🔄 Attempt $attempt/$max_attempts..."
        
        if result=$(wazctl api agents list 2>/dev/null); then
            if check_api_response "$result"; then
                echo "$result"
                return 0
            fi
        fi
        
        attempt=$((attempt + 1))
        sleep 2
    done
    
    echo "❌ API call failed after $max_attempts attempts"
    return 1
}

# Usage
if agents_data=$(retry_api_call); then
    echo "✅ Successfully retrieved agent data"
    # Process data...
else
    echo "❌ Failed to retrieve agent data"
    exit 1
fi
```

### Rate Limiting Awareness

```bash
#!/bin/bash
# Handle API rate limiting

api_call_with_rate_limit() {
    local delay=1  # Delay between calls
    
    # Make API call
    result=$(wazctl api agents list)
    
    # Check for rate limit error (if API returns it)
    if echo "$result" | grep -q "rate limit"; then
        echo "⚠️  Rate limit reached, waiting..."
        sleep 5
        result=$(wazctl api agents list)
    fi
    
    echo "$result"
    sleep $delay  # Prevent rate limiting
}
```

## 🔐 Security Considerations

### Token Management

```bash
# API commands automatically handle JWT tokens
# No manual token management required

# For debugging, you can see the token:
TOKEN=$(wazctl test auth)
echo "Current token: $TOKEN"

# Use token with curl for custom API calls:
curl -k -H "Authorization: Bearer $TOKEN" https://your-wazuh-manager:55000/agents
```

### Secure Scripting

```bash
#!/bin/bash
# Secure API scripting practices

# Don't log sensitive data
set +x  # Disable command echoing

# Check authentication before proceeding
if ! wazctl test auth > /dev/null 2>&1; then
    echo "❌ Authentication failed"
    exit 1
fi

# Use secure temporary files
TEMP_FILE=$(mktemp)
trap "rm -f $TEMP_FILE" EXIT

# Store API response securely
wazctl api agents list > "$TEMP_FILE"

# Process data
if [ -s "$TEMP_FILE" ]; then
    # Process the file...
    echo "✅ Data processed successfully"
else
    echo "❌ No data received"
    exit 1
fi
```

## ➡️ Next Steps

- Explore [Agent Management](agents.md) for detailed agent operations
- Check [User Management](users.md) for user operations
- Review [Command Reference](command-reference.md) for all available commands