# 🗺️ wazctl Roadmap

This roadmap outlines the development progress and planned features for wazctl. The project follows a phased approach, building core functionality first and expanding to advanced features.

## 📊 Current Status (v0.4.0)

> **⚠️ Status**: Under Heavy Development  
> **🚧 Stability**: Expect bugs and breaking changes

### ✅ Completed Features

#### Core Infrastructure
- ✅ **CLI Framework**: Cobra-based command structure
- ✅ **Configuration Management**: YAML-based configuration
- ✅ **API Authentication**: JWT-based Wazuh API authentication
- ✅ **Error Handling**: Structured error reporting
- ✅ **JSON Output**: Consistent JSON formatting

#### Agent Management
- ✅ **Agent Listing**: List all enrolled agents
- ✅ **Agent Status**: View connection status and details
- ✅ **JSON Processing**: API response formatting

#### Rule Testing Framework
- ✅ **Test Templates**: YAML-based rule test scaffolding
- ✅ **Test Scenarios**: Edge case definition structure
- ✅ **Command Types**: Support for bash, PowerShell, manual tests

#### User Management
- ✅ **Wazuh Users**: Create users in Wazuh manager
- ✅ **Indexer Users**: Create users in OpenSearch indexer
- ✅ **Role Assignment**: Assign roles to indexer users

#### Local Development
- ✅ **Docker Environment**: Local Wazuh instance management
- ✅ **Environment Control**: Start, stop, clean operations

#### Development Tools
- ✅ **Template Generation**: Configuration and rule templates
- ✅ **Go Module**: Proper dependency management

## 🔄 Phase 1: Foundation Strengthening (v0.5.0 - Q2 2024)

### 🎯 Primary Goals
- Stabilize core functionality
- Improve error handling and user experience
- Add essential missing features

### 📋 Planned Features

#### Enhanced Agent Management
- 🔄 **Agent Restart**: Remotely restart agents
- 🔄 **Agent Update**: Update agent versions
- 🔄 **Agent Status Detail**: Detailed agent information
- 🔄 **Agent Filtering**: Advanced filtering and search

#### Improved Configuration
- 🔄 **Environment Variables**: Support for env var configuration
- 🔄 **Multiple Profiles**: Switch between different environments
- 🔄 **Configuration Validation**: Validate config before use
- 🔄 **Auto-discovery**: Detect local Wazuh instances

#### Enhanced User Management
- 🔄 **User Listing**: List existing users
- 🔄 **User Deletion**: Safely remove users
- 🔄 **Password Reset**: Reset user passwords
- 🔄 **Role Modification**: Change user roles

#### Developer Experience
- 🔄 **Better Error Messages**: More helpful error descriptions
- 🔄 **Command Completion**: Shell completion support
- 🔄 **Verbose Mode**: Detailed operation logging
- 🔄 **Help Improvements**: Better documentation in CLI

## 🚀 Phase 2: Advanced Features (v0.6.0 - Q3 2024)

### 🎯 Primary Goals
- Rule testing execution engine
- Advanced API functionality
- Performance optimization

### 📋 Planned Features

#### Rule Testing Engine
- 📈 **Test Execution**: Automated rule test running
- 📈 **Test Reports**: Generate test result reports
- 📈 **CI/CD Integration**: Support for automated testing
- 📈 **Test Coverage**: Track rule test coverage

#### Expanded API Support
- 📈 **Rule Management**: CRUD operations for rules
- 📈 **Decoder Management**: Manage log decoders
- 📈 **CDB Lists**: Manage custom database lists
- 📈 **Configuration API**: Manage Wazuh configuration

#### Enhanced Docker Support
- 📈 **Custom Configurations**: Deploy with custom settings
- 📈 **Multi-node Setup**: Clustered environments
- 📈 **Performance Tuning**: Optimized resource allocation
- 📈 **Scenario Templates**: Pre-configured test scenarios

#### Monitoring and Observability
- 📈 **Agent Monitoring**: Real-time agent health tracking
- 📈 **Performance Metrics**: System performance monitoring
- 📈 **Log Analysis**: Built-in log analysis tools
- 📈 **Alerting**: Built-in alerting mechanisms

## 🎯 Phase 3: Enterprise Features (v0.7.0 - Q4 2024)

### 🎯 Primary Goals
- Enterprise-grade features
- Integration capabilities
- Advanced security features

### 📋 Planned Features

#### Security Enhancements
- 🌐 **MFA Support**: Multi-factor authentication
- 🌐 **RBAC Enhancement**: Advanced role-based access control
- 🌐 **Audit Logging**: Comprehensive audit trails
- 🌐 **Certificate Management**: TLS certificate handling

#### Enterprise Integration
- 🌐 **LDAP Integration**: Enterprise directory integration
- 🌐 **SIEM Integration**: Connect with external SIEM systems
- 🌐 **API Gateway**: Advanced API management
- 🌐 **SSO Support**: Single sign-on integration

#### Advanced Rule Features
- 🌐 **Rule Validation**: Syntax validation and testing
- 🌐 **Rule Analytics**: Performance analysis and optimization
- 🌐 **Rule Templates**: Pre-built rule templates
- 🌐 **Rule Versioning**: Version control for rules

#### Compliance and Reporting
- 🌐 **Compliance Reports**: Automated compliance reporting
- 🌐 **Custom Dashboards**: Build custom dashboards
- 🌐 **Data Export**: Advanced data export capabilities
- 🌐 **Scheduled Reports**: Automated report generation

## 🔮 Phase 4: AI and Automation (v0.8.0 - Q1 2025)

### 🎯 Primary Goals
- AI-powered features
- Advanced automation
- Predictive capabilities

### 📋 Planned Features

#### AI-Powered Analysis
- 🤖 **Threat Detection**: AI-powered threat analysis
- 🤖 **Anomaly Detection**: Behavioral anomaly detection
- 🤖 **Rule Suggestions**: AI-suggested rule improvements
- 🤖 **False Positive Reduction**: AI-driven tuning

#### Advanced Automation
- 🤖 **Response Automation**: Automated incident response
- 🤖 **Playbooks**: Security playbook execution
- 🤖 **Workflow Engine**: Custom workflow automation
- 🤖 **Auto-remediation**: Automated threat remediation

#### Predictive Features
- 🤖 **Capacity Planning**: Predictive capacity planning
- 🤖 **Threat Prediction**: Predictive threat modeling
- 🤖 **Performance Optimization**: AI-driven optimization
- 🤖 **Risk Assessment**: Automated risk assessment

## 🏗️ Infrastructure Improvements (Ongoing)

### Quality and Reliability
- **Testing**: Comprehensive test suite
- **Documentation**: Complete API documentation
- **Performance**: Optimization and benchmarking
- **Security**: Security auditing and hardening

### Distribution and Packaging
- **Binary Releases**: Pre-compiled binaries for all platforms
- **Package Managers**: Support for Homebrew, APT, YUM
- **Container Images**: Official Docker images
- **Installation Scripts**: One-line installation scripts

### Community and Ecosystem
- **Plugin System**: Third-party plugin support
- **Community Rules**: Shared rule repository
- **Integration Library**: Pre-built integrations
- **Training Materials**: Comprehensive training resources

## 📅 Release Schedule

| Version | Target Date | Focus Area |
|---------|------------|------------|
| v0.5.0 | Q2 2024 | Foundation Strengthening |
| v0.6.0 | Q3 2024 | Advanced Features |
| v0.7.0 | Q4 2024 | Enterprise Features |
| v0.8.0 | Q1 2025 | AI and Automation |
| v1.0.0 | Q2 2025 | Stable Release |

## 🎯 Feature Requests and Priorities

### High Priority
1. **Rule Testing Engine**: Automated test execution
2. **Agent Management**: Complete lifecycle management
3. **User Management**: Full CRUD operations
4. **Performance**: Optimization and caching
5. **Documentation**: Complete user guides

### Medium Priority
1. **Multi-environment Support**: Better environment switching
2. **Advanced Filtering**: Powerful query capabilities
3. **Export/Import**: Configuration and data migration
4. **Monitoring**: Real-time monitoring capabilities
5. **Integration**: Third-party tool integration

### Future Considerations
1. **Web Interface**: Optional web UI
2. **Mobile App**: Mobile management capabilities
3. **Cloud Integration**: Cloud provider integration
4. **Machine Learning**: Advanced ML capabilities
5. **Blockchain**: Immutable audit trails

## 🤝 Community Involvement

### How to Contribute
- **🐛 Bug Reports**: Report issues and bugs
- **💡 Feature Requests**: Suggest new features
- **💻 Code Contributions**: Submit pull requests
- **📖 Documentation**: Improve documentation
- **🧪 Testing**: Test pre-release versions

### Current Priorities for Contributors
1. **Testing**: Help test current features
2. **Documentation**: Improve and expand docs
3. **Examples**: Create usage examples
4. **Bug Fixes**: Fix existing issues
5. **Feature Implementation**: Implement planned features

## 📊 Success Metrics

### v1.0 Goals
- **⭐ 1000+ GitHub Stars**: Community adoption
- **📦 10+ Integrations**: Third-party integrations
- **🏢 100+ Organizations**: Enterprise adoption
- **📈 95%+ Uptime**: Reliability metrics
- **⚡ <1s Response Time**: Performance metrics

### Long-term Vision
wazctl aims to become the de facto standard for Wazuh management, providing:
- **Comprehensive Management**: Complete Wazuh lifecycle management
- **Developer-Friendly**: Excellent developer experience
- **Enterprise-Ready**: Production-grade reliability and security
- **Community-Driven**: Strong open-source community
- **Innovation**: Cutting-edge security features

## 🔄 Feedback and Updates

This roadmap is a living document and will be updated regularly based on:
- **Community Feedback**: User input and suggestions
- **Market Needs**: Industry requirements and trends
- **Technical Constraints**: Development limitations and opportunities
- **Resource Availability**: Development team capacity

### How to Provide Feedback
- **GitHub Issues**: Feature requests and bug reports
- **Discussions**: Community discussions on GitHub
- **Pull Requests**: Direct code contributions
- **Documentation**: Improvements to this roadmap

---

**Last Updated**: January 2024  
**Next Review**: February 2024

> 💡 **Note**: This roadmap represents current intentions and may change based on community feedback, technical discoveries, and resource availability.