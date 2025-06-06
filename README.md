<div align="center">
  <h1>🚧 Under Heavy Construction 🚧</h1>
  <p><strong>This project is a work in progress and still very infant. Expect bugs, and breaking changes!</strong></p>
</div>

# 🛡️ wazctl - Your Wazuh Command-Line Companion 🦉

[](https://golang.org)
[](https://opensource.org/licenses/MIT)
[](https://www.google.com/search?q=https://github.com/EpykLab/wazctl/blob/main/CONTRIBUTING.md)

`wazctl` is a powerful, intuitive command-line interface (CLI) designed to streamline your interactions with the Wazuh Security Platform. Whether you're managing agents, testing new rules, or automating security tasks, `wazctl` is the tool you need to get the job done efficiently.

wazclt in built using [wasabi](https://github.com/EpykLab/wasabi), our wazuh api sdk, auto generated from wazuh openapi specification.


## ✨ Intended Functionality

The vision for `wazctl` is to provide a comprehensive toolkit for Wazuh administrators and security engineers.

  * **⚡️ Simplified API Interaction:** Authenticate and interact with the Wazuh API using simple commands, abstracting away the complexities of direct API calls.
  * **👤 Agent Management:** Manage the lifecycle of your Wazuh agents directly from your terminal. The tool currently supports listing agents, with plans to expand to other management functions.
  * **📐 Rule Testing Framework:** A core feature of `wazctl` is its ability to scaffold and (eventually) run test cases for your Wazuh rules. Define edge cases in simple YAML files to ensure your rules work as expected.
  * **⚙️ Effortless Configuration:** Quickly generate the configuration files needed to connect `wazctl` to your Wazuh manager.

## 🚀 Installation

`wazctl` is built with Go and can be installed using `go install`:

```bash
go install github.com/EpykLab/wazctl@latest
```

Make sure your `$(go env GOPATH)/bin` directory is in your system's `PATH`.

## ▶️ Getting Started

Getting up and running with `wazctl` is easy.

### 1. Create a Configuration File

First, you need to tell `wazctl` how to connect to your Wazuh manager. Generate a configuration file with the `init config` command.

```bash
wazctl init config
```

This will create a `.wazctl.yaml` file in your current directory with the following content:

```yaml
endpoint: your-instance.com
port: "55000"
protocol: https
wuiPassword: password
wuiUsername: wui
httpDebug: false
skipTlsVerify: true
```

Edit this file with your Wazuh API credentials and endpoint details. `wazctl`
also checks for this file in `~/.wazctl.yaml` and `~/.config/wazctl.yaml`.

### 2. Test Your Connection

Verify that your credentials are correct by running the `test auth` command.

```bash
wazctl test auth
```

If successful, this will print a JWT token to your console, confirming that
`wazctl` can authenticate with your Wazuh manager.

### 3. Interact with the API

You can now use `wazctl` to interact with the Wazuh API. For example, to list
your connected agents:

```bash
wazctl api agents list
```

This command fetches and displays all agents enrolled in the manager.

### 4. Scaffold a Rule Test

To create a new rule test file, use the `init rule` command. This is perfect
for building a library of test cases for your custom rules.

```bash
wazctl init rule --name "my_suspicious_login_test"
```

This generates a YAML file named `my_suspicious_login_test.yaml` with a
pre-defined structure, ready for you to customize:

```yaml
ruleId: rule_001
ruleName: Unauthorized Access
ruleAuthor: John Doe
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

## 🗺️ Project Roadmap

This project is under active development. Here is a look at what's done and what's planned.

  * [✅] **Initial Setup Commands** (`init config`, `init rule`)
  * [✅] **Authentication** (`test auth`)
  * [✅] **List Wazuh Agents** (`api agents list`)
  * [🚧] **Rule Test Execution Engine** (e.g., `wazctl test rule <file.yaml>`)
  * [🔄] **Expanded Agent Management** (e.g., `restart`, `update`, `remove` agents)
  * [📈] **Enhanced Output Formatting** (Tables, JSON, etc.)
  * [🌐] **Broader API Support** (Managing rules, decoders, CDB lists, etc.)
  * [📦] **Pre-compiled Binaries** for multiple platforms.
  ...and much more.

## 🤝 How to Contribute

Contributions are what make the open-source community such an amazing place to
learn, inspire, and create. Any contributions you make are **greatly
appreciated**. This Wazuh API is huge, so this is a massive undertaking.

1.  **Fork the Project**
2.  **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4.  **Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **Open a Pull Request**

Please feel free to open an issue with the tag "bug" or "enhancement" as well!

## 📜 License

Distributed under the MIT License. See the `LICENSE` file in the original
repository for more information. The copyright notice in the source files
indicates it is available under a permissive license.
