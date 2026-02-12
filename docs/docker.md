# 🐳 Docker Environment

wazctl provides built-in Docker support for local Wazuh development and testing. This feature allows you to quickly spin up complete Wazuh environments for rule development, testing, and learning.

## 🎯 Overview

The Docker environment includes:
- **🛡️ Wazuh Manager**: Complete security platform
- **🔍 OpenSearch**: Log indexing and storage
- **📊 OpenSearch Dashboards**: Data visualization (Kibana replacement)
- **🌐 Wazuh Dashboard**: Security operations interface
- **🔗 Pre-configured Integration**: Ready-to-use setup

## 🚀 Quick Start

### Starting Your Environment

Launch a complete Wazuh stack:

```bash
wazctl localenv docker --start
```

This command:
1. Downloads required Docker images
2. Creates necessary networks and volumes
3. Starts all Wazuh components
4. Configures component communication
5. Waits for services to be ready

**Expected output**:
```
🐳 Starting Wazuh Docker environment...
📥 Pulling Docker images...
🔧 Creating network and volumes...
🚀 Starting containers...
✅ Wazuh environment ready!

📊 Access URLs:
   Wazuh Dashboard: https://localhost:443
   OpenSearch Dashboards: https://localhost:5601
   
🔐 Default credentials:
   Username: admin
   Password: admin
```

### Stopping Your Environment

Stop all services while preserving data:

```bash
wazctl localenv docker --stop
```

### Cleaning Up

Remove all containers, networks, and volumes:

```bash
wazctl localenv docker --clean
```

⚠️ **Warning**: This permanently deletes all data in your local environment.

## 🔧 Environment Management

### Complete Lifecycle Management

```bash
#!/bin/bash
# Complete environment lifecycle script

echo "🐳 Wazuh Docker Environment Manager"
echo "==================================="

case $1 in
    "start")
        echo "🚀 Starting environment..."
        wazctl localenv docker --start
        
        echo "⏳ Waiting for services to be ready..."
        sleep 60
        
        echo "🔐 Testing authentication..."
        wazctl test auth
        
        echo "✅ Environment ready for use!"
        ;;
        
    "stop")
        echo "⏸️ Stopping environment..."
        wazctl localenv docker --stop
        echo "✅ Environment stopped"
        ;;
        
    "restart")
        echo "🔄 Restarting environment..."
        wazctl localenv docker --stop
        sleep 10
        wazctl localenv docker --start
        echo "✅ Environment restarted"
        ;;
        
    "clean")
        echo "🗑️ Cleaning environment..."
        read -p "This will delete all data. Continue? (y/N): " confirm
        if [[ $confirm == [yY] ]]; then
            wazctl localenv docker --clean
            echo "✅ Environment cleaned"
        else
            echo "❌ Cleanup cancelled"
        fi
        ;;
        
    "status")
        echo "📊 Environment status:"
        docker ps | grep wazuh
        ;;
        
    *)
        echo "Usage: $0 {start|stop|restart|clean|status}"
        exit 1
        ;;
esac
```

### Environment Health Check

```bash
#!/bin/bash
# Health check script for Docker environment

echo "🏥 Wazuh Docker Environment Health Check"
echo "========================================"

# Check Docker daemon
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker daemon not running"
    exit 1
fi

# Check container status
echo "📦 Container Status:"
containers=$(docker ps --format "table {{.Names}}\t{{.Status}}" | grep wazuh)
if [ -z "$containers" ]; then
    echo "❌ No Wazuh containers running"
    echo "💡 Start environment with: wazctl localenv docker --start"
    exit 1
else
    echo "$containers"
fi

# Check service endpoints
echo ""
echo "🌐 Service Availability:"

# Wazuh API
if curl -k -s https://localhost:55000 > /dev/null; then
    echo "✅ Wazuh API (port 55000)"
else
    echo "❌ Wazuh API (port 55000)"
fi

# Wazuh Dashboard
if curl -k -s https://localhost:443 > /dev/null; then
    echo "✅ Wazuh Dashboard (port 443)"
else
    echo "❌ Wazuh Dashboard (port 443)"
fi

# OpenSearch
if curl -k -s https://localhost:9200 > /dev/null; then
    echo "✅ OpenSearch (port 9200)"
else
    echo "❌ OpenSearch (port 9200)"
fi

# OpenSearch Dashboards
if curl -k -s https://localhost:5601 > /dev/null; then
    echo "✅ OpenSearch Dashboards (port 5601)"
else
    echo "❌ OpenSearch Dashboards (port 5601)"
fi

# Test wazctl authentication
echo ""
echo "🔐 Authentication Test:"
if wazctl test auth > /dev/null 2>&1; then
    echo "✅ wazctl authentication successful"
else
    echo "❌ wazctl authentication failed"
    echo "💡 Check configuration and wait for services to fully start"
fi
```

## 🔧 Configuration and Customization

### Environment Variables

The Docker environment can be customized using your `.wazctl.yaml` configuration:

```yaml
local:
    repoVersion: v4.12.0  # Wazuh version to deploy
wazuh:
    endpoint: localhost   # Will be set automatically for Docker
    port: "55000"
    protocol: https
    skipTlsVerify: true   # Required for local development
```

### Version Management

Deploy specific Wazuh versions:

```bash
# Edit .wazctl.yaml to change version
nano .wazctl.yaml

# Update local.repoVersion to desired version:
# local:
#     repoVersion: v4.11.0

# Clean and restart with new version
wazctl localenv docker --clean
wazctl localenv docker --start
```

### Persistent Data

Docker volumes preserve data across container restarts:

```bash
# View Docker volumes
docker volume ls | grep wazuh

# Backup volume data
docker run --rm -v wazuh_data:/data -v $(pwd):/backup alpine tar czf /backup/wazuh-backup.tar.gz /data

# Restore volume data
docker run --rm -v wazuh_data:/data -v $(pwd):/backup alpine tar xzf /backup/wazuh-backup.tar.gz -C /
```

## 🛠️ Development Workflows

### Rule Development Workflow

```bash
#!/bin/bash
# Rule development with Docker environment

RULE_NAME=$1
if [ -z "$RULE_NAME" ]; then
    echo "Usage: $0 <rule_name>"
    exit 1
fi

echo "🎯 Starting rule development for: $RULE_NAME"

# Start environment if not running
if ! docker ps | grep wazuh > /dev/null; then
    echo "🐳 Starting Docker environment..."
    wazctl localenv docker --start
    
    echo "⏳ Waiting for services..."
    sleep 60
fi

# Create rule test
echo "📝 Creating rule test template..."
wazctl init rule --name "${RULE_NAME}_test"

# Test authentication
echo "🔐 Testing authentication..."
wazctl test auth

# Check agents
echo "👥 Checking agents..."
wazctl agents list

echo "✅ Environment ready for rule development!"
echo "📋 Next steps:"
echo "   1. Edit ${RULE_NAME}_test.yaml"
echo "   2. Access Wazuh Dashboard: https://localhost:443"
echo "   3. Test your rules"
echo "   4. Monitor alerts in real-time"
```

### Agent Testing Setup

```bash
#!/bin/bash
# Set up agent testing in Docker environment

echo "🔧 Setting up agent testing environment..."

# Start main environment
wazctl localenv docker --start

# Wait for services
sleep 60

# Check environment
wazctl test auth
wazctl agents list

# Create test files for agent simulation
mkdir -p test-data
cat > test-data/auth-test.sh << 'EOF'
#!/bin/bash
# Simulate authentication events
for i in {1..5}; do
    logger "sshd: Failed password for testuser from 192.168.1.100"
    sleep 2
done
EOF

cat > test-data/web-test.sh << 'EOF'
#!/bin/bash
# Simulate web attacks
logger "apache: GET /admin.php 404"
logger "apache: GET /wp-admin/ 200"
logger "apache: POST /login.php 401"
EOF

chmod +x test-data/*.sh

echo "✅ Agent testing environment ready!"
echo "📋 Available test scripts:"
echo "   - test-data/auth-test.sh (authentication events)"
echo "   - test-data/web-test.sh (web application events)"
```

## 📊 Monitoring and Debugging

### Container Log Monitoring

```bash
#!/bin/bash
# Monitor container logs

echo "📊 Wazuh Container Logs"
echo "======================="

# Get Wazuh container names
containers=$(docker ps --format "{{.Names}}" | grep wazuh)

if [ -z "$containers" ]; then
    echo "❌ No Wazuh containers running"
    exit 1
fi

echo "Available containers:"
echo "$containers"
echo ""

# Follow logs for all containers
echo "📝 Following logs (Ctrl+C to stop)..."
for container in $containers; do
    echo "--- $container ---"
    docker logs --tail=10 -f $container &
done

wait
```

### Performance Monitoring

```bash
#!/bin/bash
# Monitor Docker environment performance

echo "📈 Docker Environment Performance"
echo "================================="

# Container resource usage
echo "💾 Container Resource Usage:"
docker stats --no-stream --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}"

# Disk usage
echo ""
echo "💽 Disk Usage:"
docker system df

# Volume usage
echo ""
echo "📦 Volume Usage:"
docker volume ls --format "table {{.Name}}\t{{.Driver}}"
```

### Troubleshooting Common Issues

#### Environment Won't Start

```bash
# Check Docker resources
docker system df
docker system prune  # Free up space if needed

# Check port conflicts
netstat -tulpn | grep -E ':(443|5601|9200|55000)'

# View startup logs
docker logs $(docker ps -q | head -1)
```

#### Services Not Responding

```bash
# Restart individual containers
docker restart $(docker ps --format "{{.Names}}" | grep wazuh)

# Check container health
docker inspect $(docker ps -q) | grep -i health

# Force clean restart
wazctl localenv docker --clean
wazctl localenv docker --start
```

#### Authentication Issues

```bash
# Wait longer for services
sleep 120
wazctl test auth

# Check Wazuh API logs
docker logs $(docker ps | grep wazuh-manager | awk '{print $1}')

# Verify configuration
cat .wazctl.yaml
```

## 🔗 Integration with External Tools

### VS Code Development

```json
// .vscode/tasks.json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "Start Wazuh Environment",
            "type": "shell",
            "command": "wazctl localenv docker --start",
            "group": "build",
            "presentation": {
                "echo": true,
                "reveal": "always",
                "focus": false,
                "panel": "shared"
            }
        },
        {
            "label": "Stop Wazuh Environment",
            "type": "shell", 
            "command": "wazctl localenv docker --stop",
            "group": "build"
        },
        {
            "label": "Clean Wazuh Environment",
            "type": "shell",
            "command": "wazctl localenv docker --clean",
            "group": "build"
        }
    ]
}
```

### Docker Compose Integration

For advanced users who want to customize the Docker setup:

```yaml
# docker-compose.override.yml
version: '3.8'
services:
  wazuh-manager:
    ports:
      - "1515:1515"  # Agent enrollment
      - "514:514/udp"  # Syslog
    volumes:
      - ./custom-rules:/var/ossec/etc/rules/local_rules.xml
```

## 🚧 Planned Features

Future Docker environment enhancements:

- **⚙️ Custom Configurations**: Deploy with custom settings
- **📊 Multi-node Setup**: Clustered environments
- **🔧 Plugin Support**: Additional Wazuh integrations  
- **📈 Performance Tuning**: Optimized resource allocation
- **🎯 Scenario Templates**: Pre-configured test scenarios
- **🔄 Auto-update**: Automatic version management

## ➡️ Next Steps

- Create [Rule Tests](rule-testing.md) in your environment
- Explore [API Commands](api.md)
- Check [Command Reference](command-reference.md)