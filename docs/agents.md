# 👥 Agent Management

Agent management is one of wazctl's core features, allowing you to monitor and manage Wazuh agents directly from the command line.

## 🎯 Overview

wazctl provides commands to:
- 📋 List all enrolled agents
- 📊 View agent status and details  
- 🔍 Filter and search agents
- 📈 Monitor agent health

> 🚧 **Note**: Additional agent management features (restart, update, remove) are planned for future releases.

## 📋 Listing Agents

### Basic Agent Listing

List all agents enrolled in your Wazuh manager:

```bash
wazctl agents list
```

**Alternative command**:
```bash
wazctl api agents list
```

Both commands provide identical functionality through different command paths.

### Sample Output

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
      },
      {
        "id": "001",
        "name": "web-server-01", 
        "ip": "192.168.1.100",
        "status": "active",
        "node_name": "node01",
        "dateAdd": "2024-01-10T09:15:00Z",
        "version": "4.12.0",
        "manager": "wazuh-manager",
        "os": {
          "arch": "x86_64",
          "major": "8",
          "minor": "0",
          "name": "CentOS Linux",
          "platform": "centos", 
          "version": "8.0"
        },
        "lastKeepAlive": "2024-01-15T15:44:30Z",
        "group": ["webservers", "production"]
      }
    ],
    "total_affected_items": 2,
    "total_failed_items": 0,
    "failed_items": []
  }
}
```

## 📊 Understanding Agent Data

### Agent Status Values

| Status | Description | Action Needed |
|--------|-------------|---------------|
| `active` | Agent is connected and sending data | ✅ None |
| `disconnected` | Agent is enrolled but not connected | 🔍 Investigate connectivity |
| `never_connected` | Agent enrolled but never connected | 🔧 Check agent configuration |
| `pending` | Agent waiting for approval | ✋ Approve agent enrollment |

### Key Agent Fields

| Field | Description | Use Case |
|-------|-------------|----------|
| `id` | Unique agent identifier | Agent-specific operations |
| `name` | Agent hostname/name | Human-readable identification |
| `ip` | Agent IP address | Network troubleshooting |
| `status` | Current connection status | Health monitoring |
| `lastKeepAlive` | Last communication time | Connectivity tracking |
| `version` | Wazuh agent version | Version management |
| `os` | Operating system details | Platform-specific rules |
| `group` | Agent groups | Group-based configuration |

## 🔍 Filtering and Processing Agent Data

### Using jq for Advanced Filtering

wazctl outputs JSON, making it perfect for processing with `jq`:

#### List Only Agent Names
```bash
wazctl agents list | jq -r '.data.affected_items[].name'
```

#### Show Disconnected Agents
```bash
wazctl agents list | jq '.data.affected_items[] | select(.status=="disconnected")'
```

#### Count Agents by Status
```bash
wazctl agents list | jq '.data.affected_items | group_by(.status) | map({status: .[0].status, count: length})'
```

#### Find Agents by IP Range
```bash
wazctl agents list | jq '.data.affected_items[] | select(.ip | startswith("192.168.1"))'
```

#### List Agents with Old Versions
```bash
wazctl agents list | jq '.data.affected_items[] | select(.version != "4.12.0")'
```

### Using grep for Simple Filtering

#### Find Specific Agent
```bash
wazctl agents list | grep -i "web-server"
```

#### Show Only Active Agents
```bash
wazctl agents list | grep '"status": "active"'
```

#### Find Never Connected Agents
```bash
wazctl agents list | grep '"status": "never_connected"'
```

## 📈 Monitoring Workflows

### Daily Health Check

```bash
#!/bin/bash
echo "🔍 Wazuh Agent Health Check - $(date)"
echo "======================================"

# Test authentication
echo "🔐 Testing authentication..."
wazctl test auth > /dev/null && echo "✅ Authentication successful" || echo "❌ Authentication failed"

# Get agent statistics
echo "📊 Agent Statistics:"
TOTAL=$(wazctl agents list | jq '.data.total_affected_items')
ACTIVE=$(wazctl agents list | jq '.data.affected_items[] | select(.status=="active")' | wc -l)
DISCONNECTED=$(wazctl agents list | jq '.data.affected_items[] | select(.status=="disconnected")' | wc -l)
NEVER_CONNECTED=$(wazctl agents list | jq '.data.affected_items[] | select(.status=="never_connected")' | wc -l)

echo "  Total Agents: $TOTAL"
echo "  Active: $ACTIVE"
echo "  Disconnected: $DISCONNECTED" 
echo "  Never Connected: $NEVER_CONNECTED"

# Show problematic agents
if [ "$DISCONNECTED" -gt 0 ] || [ "$NEVER_CONNECTED" -gt 0 ]; then
    echo "⚠️  Agents Needing Attention:"
    wazctl agents list | jq -r '.data.affected_items[] | select(.status=="disconnected" or .status=="never_connected") | "  - \(.name) (\(.ip)) - \(.status)"'
fi
```

### Agent Inventory Report

```bash
#!/bin/bash
echo "📋 Agent Inventory Report"
echo "========================"

# Operating System Distribution
echo "🖥️  Operating Systems:"
wazctl agents list | jq -r '.data.affected_items[].os | "\(.name) \(.version)"' | sort | uniq -c | sort -rn

# Agent Version Distribution  
echo -e "\n📦 Agent Versions:"
wazctl agents list | jq -r '.data.affected_items[].version' | sort | uniq -c | sort -rn

# Group Membership
echo -e "\n👥 Group Membership:"
wazctl agents list | jq -r '.data.affected_items[] | "\(.name): \(.group | join(", "))"'
```

## 🔧 Troubleshooting Agent Issues

### Common Agent Problems

#### 1. Agent Never Connected

**Symptoms**:
```bash
wazctl agents list | jq '.data.affected_items[] | select(.status=="never_connected")'
```

**Possible Causes**:
- Incorrect manager IP in agent configuration
- Firewall blocking port 1514
- Agent service not started
- Authentication key issues

**Investigation Steps**:
1. Check agent configuration on the endpoint
2. Verify network connectivity to port 1514
3. Review agent logs on the endpoint
4. Confirm enrollment process completed

#### 2. Agent Disconnected

**Symptoms**:
```bash
wazctl agents list | jq '.data.affected_items[] | select(.status=="disconnected")'
```

**Possible Causes**:
- Network connectivity issues
- Agent service stopped
- Manager restart
- High network latency

**Investigation Steps**:
1. Check last keep-alive time
2. Verify agent service status
3. Test network connectivity
4. Review manager logs

#### 3. Agent Version Mismatch

**Symptoms**:
```bash
wazctl agents list | jq '.data.affected_items[] | select(.version != "4.12.0")'
```

**Actions**:
- Plan agent updates
- Test compatibility
- Schedule maintenance windows

## 📚 Integration Examples

### Prometheus Metrics Export

```bash
#!/bin/bash
# Export agent metrics for Prometheus
echo "# HELP wazuh_agents_total Total number of Wazuh agents"
echo "# TYPE wazuh_agents_total gauge"
wazctl agents list | jq -r '
  .data.affected_items | 
  group_by(.status) | 
  map("wazuh_agents_total{status=\"\(.[0].status)\"} \(length)")[]
'
```

### Slack Notifications

```bash
#!/bin/bash
WEBHOOK_URL="your-slack-webhook-url"

DISCONNECTED=$(wazctl agents list | jq '.data.affected_items[] | select(.status=="disconnected")' | wc -l)

if [ "$DISCONNECTED" -gt 0 ]; then
    MESSAGE="⚠️ Wazuh Alert: $DISCONNECTED agents are disconnected"
    curl -X POST -H 'Content-type: application/json' \
         --data "{\"text\":\"$MESSAGE\"}" \
         "$WEBHOOK_URL"
fi
```

### CSV Export

```bash
#!/bin/bash
# Export agent data to CSV
echo "ID,Name,IP,Status,Version,OS,LastKeepAlive" > agents.csv
wazctl agents list | jq -r '
  .data.affected_items[] | 
  [.id, .name, .ip, .status, .version, (.os.name + " " + .os.version), .lastKeepAlive] | 
  @csv
' >> agents.csv
```

## 🚧 Planned Features

Future agent management capabilities:

- **🔄 Agent Restart**: Remotely restart agents
- **📦 Agent Update**: Update agent versions  
- **🗑️ Agent Removal**: Remove/deregister agents
- **⚙️ Configuration Push**: Deploy agent configurations
- **📊 Enhanced Filtering**: Advanced query capabilities
- **📈 Historical Data**: Agent status history

## ➡️ Next Steps

- Learn about [Rule Testing](rule-testing.md)
- Explore [API Commands](api.md)
- Set up [User Management](users.md)