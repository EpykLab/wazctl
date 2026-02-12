# 👤 User Management

wazctl provides comprehensive user management capabilities for both Wazuh managers and OpenSearch indexers, allowing you to create and manage users across your security infrastructure.

## 🎯 Overview

wazctl can manage users in:
- **🛡️ Wazuh Manager**: API users for Wazuh operations
- **🔍 OpenSearch Indexer**: Users for log analysis and visualization
- **🔗 Role Management**: Assign appropriate roles and permissions

## 👥 Creating Users

### Wazuh Manager Users

Create users in your Wazuh manager:

```bash
wazctl user add --username analyst1 --password SecurePass123! --component wazuh
```

**Parameters**:
- `--username` (`-u`): Username for the new user
- `--password` (`-p`): Password for the new user
- `--component` (`-c`): Set to `wazuh` for manager users

### OpenSearch Indexer Users

Create users in your OpenSearch indexer with specific roles:

```bash
wazctl user add --username analyst2 --password SecurePass123! --component indexer --role kibana_user
```

**Parameters**:
- `--username` (`-u`): Username for the new user
- `--password` (`-p`): Password for the new user  
- `--component` (`-c`): Set to `indexer` for OpenSearch users
- `--role` (`-r`): Role to assign (required for indexer users)

## 🔐 User Roles and Permissions

### Wazuh Manager Roles

Common Wazuh user roles and their capabilities:

| Role | Permissions | Use Case |
|------|-------------|----------|
| `administrator` | Full system access | System administrators |
| `agent_manager` | Agent management | Agent operations teams |
| `readonly` | Read-only access | Auditors, compliance |
| `normal` | Basic user access | Limited operations |

### OpenSearch Indexer Roles

Common OpenSearch roles for security analysis:

| Role | Permissions | Use Case |
|------|-------------|----------|
| `kibana_user` | Kibana dashboard access | SOC analysts |
| `logstash_writer` | Write log data | Log ingestion |
| `readonly` | Read-only log access | Compliance, reporting |
| `all_access` | Full indexer access | Administrators |
| `security_analyst` | Security-focused access | Security operations |

## 📋 User Management Examples

### SOC Team Setup

Create a complete SOC team with appropriate access:

```bash
#!/bin/bash
echo "🔧 Setting up SOC team users..."

# SOC Manager - Full access
wazctl user add --username soc_manager --password Manager123! --component wazuh
wazctl user add --username soc_manager --password Manager123! --component indexer --role all_access

# Senior Analysts - Agent management + analysis
wazctl user add --username senior_analyst1 --password Analyst123! --component wazuh
wazctl user add --username senior_analyst1 --password Analyst123! --component indexer --role security_analyst

wazctl user add --username senior_analyst2 --password Analyst123! --component wazuh  
wazctl user add --username senior_analyst2 --password Analyst123! --component indexer --role security_analyst

# Junior Analysts - Read-only access
wazctl user add --username junior_analyst1 --password Junior123! --component wazuh
wazctl user add --username junior_analyst1 --password Junior123! --component indexer --role kibana_user

wazctl user add --username junior_analyst2 --password Junior123! --component wazuh
wazctl user add --username junior_analyst2 --password Junior123! --component indexer --role kibana_user

# Compliance Officer - Read-only everywhere
wazctl user add --username compliance_officer --password Compliance123! --component wazuh
wazctl user add --username compliance_officer --password Compliance123! --component indexer --role readonly

echo "✅ SOC team setup complete!"
```

### Development Team Setup

Create development-focused users:

```bash
#!/bin/bash
echo "🛠️ Setting up development team users..."

# Security Developer - Rule development access
wazctl user add --username sec_developer --password DevSecure123! --component wazuh
wazctl user add --username sec_developer --password DevSecure123! --component indexer --role kibana_user

# Test Engineer - Agent management for testing
wazctl user add --username test_engineer --password TestSecure123! --component wazuh
wazctl user add --username test_engineer --password TestSecure123! --component indexer --role kibana_user

echo "✅ Development team setup complete!"
```

### Service Account Setup

Create service accounts for automation:

```bash
#!/bin/bash
echo "🤖 Setting up service accounts..."

# Monitoring Service - Read-only for metrics
wazctl user add --username monitoring_svc --password MonitorSvc123! --component wazuh
wazctl user add --username monitoring_svc --password MonitorSvc123! --component indexer --role readonly

# Backup Service - Read access for backups
wazctl user add --username backup_svc --password BackupSvc123! --component indexer --role readonly

# Integration Service - API access for integrations
wazctl user add --username integration_svc --password IntegrationSvc123! --component wazuh

echo "✅ Service accounts setup complete!"
```

## 🔒 Security Best Practices

### Password Requirements

Ensure strong passwords for all users:

```bash
#!/bin/bash
# Password generation helper
generate_password() {
    openssl rand -base64 32 | tr -d "=+/" | cut -c1-16
}

USERNAME="new_analyst"
PASSWORD=$(generate_password)

echo "Creating user: $USERNAME"
echo "Generated password: $PASSWORD"

wazctl user add --username "$USERNAME" --password "$PASSWORD" --component wazuh
wazctl user add --username "$USERNAME" --password "$PASSWORD" --component indexer --role kibana_user

echo "⚠️  Please securely share password with user"
```

### Account Lifecycle Management

```bash
#!/bin/bash
# User onboarding script

FIRSTNAME=$1
LASTNAME=$2
ROLE=$3

if [ $# -ne 3 ]; then
    echo "Usage: $0 <firstname> <lastname> <role>"
    echo "Roles: analyst, manager, readonly"
    exit 1
fi

USERNAME="${FIRSTNAME,,}.${LASTNAME,,}"  # Convert to lowercase
PASSWORD=$(openssl rand -base64 32 | tr -d "=+/" | cut -c1-16)

echo "👤 Creating user: $USERNAME"
echo "🎭 Role: $ROLE"

case $ROLE in
    "analyst")
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component wazuh
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component indexer --role security_analyst
        ;;
    "manager")
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component wazuh
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component indexer --role all_access
        ;;
    "readonly")
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component wazuh
        wazctl user add --username "$USERNAME" --password "$PASSWORD" --component indexer --role readonly
        ;;
    *)
        echo "❌ Invalid role: $ROLE"
        exit 1
        ;;
esac

echo "✅ User created successfully!"
echo "📧 Send these credentials securely to the user:"
echo "   Username: $USERNAME"
echo "   Password: $PASSWORD"
echo "   Wazuh URL: https://your-wazuh-manager.com"
echo "   Kibana URL: https://your-indexer.com:5601"
```

### Role-Based Access Control

```yaml
# rbac_mapping.yaml - Document your role mappings
roles:
  soc_manager:
    wazuh_permissions:
      - full_access
    indexer_role: all_access
    description: "Full access to all systems"
    
  senior_analyst:
    wazuh_permissions:
      - agent_management
      - rule_management
      - read_access
    indexer_role: security_analyst
    description: "Agent and rule management with analysis access"
    
  junior_analyst:
    wazuh_permissions:
      - read_access
    indexer_role: kibana_user
    description: "Read-only access with dashboard privileges"
    
  compliance_officer:
    wazuh_permissions:
      - read_access
    indexer_role: readonly
    description: "Read-only access for compliance reporting"
```

## 📊 User Management Workflows

### Daily User Operations

```bash
#!/bin/bash
# Daily user management tasks

echo "📊 Daily User Management Report - $(date)"
echo "=========================================="

# Check authentication for service accounts
echo "🔐 Testing service account authentication..."
SERVICE_ACCOUNTS=("monitoring_svc" "backup_svc" "integration_svc")

for account in "${SERVICE_ACCOUNTS[@]}"; do
    # Note: This would require temporarily updating config
    echo "Testing $account..."
    # wazctl test auth (would need config switch)
done

# TODO: Add user activity monitoring when API supports it
echo "📈 User activity monitoring not yet available"

# Check for expired passwords (manual process currently)
echo "⚠️  Manual tasks:"
echo "   - Review user access logs"
echo "   - Check for inactive accounts"
echo "   - Verify role assignments"
```

### User Audit Script

```bash
#!/bin/bash
# User audit and compliance check

echo "🔍 Wazuh User Audit Report"
echo "=========================="
echo "Date: $(date)"
echo ""

echo "📋 Current Users (Manual Review Required):"
echo "   - Use Wazuh Web UI to review current users"
echo "   - Check user last login times"
echo "   - Verify role assignments"
echo ""

echo "🔐 Password Policy Compliance:"
echo "   - Ensure passwords meet complexity requirements"
echo "   - Check password age and rotation"
echo "   - Verify MFA where applicable"
echo ""

echo "👥 Role Assignment Review:"
echo "   - Verify principle of least privilege"
echo "   - Check for unnecessary permissions"
echo "   - Review service account access"
echo ""

echo "📊 Recommendations:"
echo "   - Regular password rotation (90 days)"
echo "   - Annual access review"
echo "   - Monitor user activity logs"
echo "   - Implement MFA for privileged accounts"
```

## 🔧 Troubleshooting User Issues

### Common User Creation Errors

#### 1. Authentication Failure
```bash
Error: authentication failed
```

**Solution**: Verify wazctl configuration and API credentials:
```bash
wazctl test auth
```

#### 2. Insufficient Permissions
```bash
Error: insufficient permissions to create user
```

**Solution**: Ensure your wazctl user has administrative privileges:
- Check user role in Wazuh manager
- Verify API permissions
- Use an administrator account

#### 3. Invalid Role for Indexer
```bash
Error: role required for indexer component
```

**Solution**: Always specify role for indexer users:
```bash
wazctl user add --username user1 --password pass123 --component indexer --role kibana_user
```

#### 4. User Already Exists
```bash
Error: user already exists
```

**Solution**: Choose a different username or remove existing user first.

### User Creation Validation

```bash
#!/bin/bash
# Validate user creation

USERNAME=$1
COMPONENT=$2

if [ -z "$USERNAME" ] || [ -z "$COMPONENT" ]; then
    echo "Usage: $0 <username> <wazuh|indexer>"
    exit 1
fi

echo "🔍 Validating user creation for: $USERNAME"

case $COMPONENT in
    "wazuh")
        echo "📝 Checking Wazuh user creation..."
        # Test authentication with new user would require config change
        echo "   Manual verification needed in Wazuh Web UI"
        ;;
    "indexer")
        echo "📝 Checking Indexer user creation..."
        # Test indexer access would require indexer API calls
        echo "   Manual verification needed in Kibana"
        ;;
    *)
        echo "❌ Invalid component: $COMPONENT"
        exit 1
        ;;
esac

echo "✅ Please verify user $USERNAME can:"
echo "   - Log into the web interface"
echo "   - Access appropriate dashboards"
echo "   - Perform role-specific functions"
```

## 🚧 Planned Features

Future user management enhancements:

- **📋 User Listing**: List existing users
- **🗑️ User Deletion**: Remove users safely
- **🔄 Password Reset**: Reset user passwords
- **👥 Role Management**: Modify user roles
- **📊 User Activity**: Monitor user actions
- **🔐 MFA Support**: Multi-factor authentication
- **📅 Account Expiration**: Automatic account lifecycle
- **🔗 LDAP Integration**: Enterprise directory integration

## 📚 Integration Examples

### Slack User Notifications

```bash
#!/bin/bash
# Notify team of new user creation

USERNAME=$1
ROLE=$2
SLACK_WEBHOOK="your-slack-webhook-url"

MESSAGE="👤 New user created: \`$USERNAME\` with role \`$ROLE\`"

curl -X POST -H 'Content-type: application/json' \
     --data "{\"text\":\"$MESSAGE\"}" \
     "$SLACK_WEBHOOK"
```

### CSV User Export

```bash
#!/bin/bash
# Export user list to CSV (template for future implementation)

echo "Username,Component,Role,Created_Date,Last_Login" > users.csv
echo "Note: Manual data entry required until API supports user listing" >> users.csv
```

## ➡️ Next Steps

- Set up [Docker Environment](docker.md)
- Explore [API Commands](api.md)
- Review [Troubleshooting](troubleshooting.md)