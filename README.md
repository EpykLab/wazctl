<div align="center">
  <h1>Under Heavy Construction</h1>
  <p><strong>This project is a work in progress and still very infant. Expect bugs, and breaking changes!</strong></p>
</div>

# wazctl - Your Wazuh Command-Line Companion

[](https://golang.org)
[](https://opensource.org/licenses/MIT)
[](https://www.google.com/search?q=https://github.com/EpykLab/wazctl/blob/main/CONTRIBUTING.md)

`wazctl` is a powerful, intuitive command-line interface (CLI) designed to streamline your interactions with the Wazuh Security Platform. Whether you're managing agents, testing new rules, or automating security tasks, `wazctl` is the tool you need to get the job done efficiently.

wazctl is built using [wasabi](https://github.com/EpykLab/wasabi), our Wazuh API SDK, auto-generated from the Wazuh OpenAPI specification.


## Intended Functionality

The vision for `wazctl` is to provide a comprehensive toolkit for Wazuh administrators and security engineers.

  * **Simplified API Interaction:** Authenticate and interact with the Wazuh API using simple commands, abstracting away the complexities of direct API calls.
  * **Agent Management:** Manage the lifecycle of your Wazuh agents directly from your terminal. The tool currently supports listing agents, with plans to expand to other management functions.
  * **Rule Testing Framework:** A core feature of `wazctl` is its ability to scaffold and (eventually) run test cases for your Wazuh rules. Define edge cases in simple YAML files to ensure your rules work as expected.
  * **Effortless Configuration:** Quickly generate the configuration files needed to connect `wazctl` to your Wazuh manager.

## Installation

`wazctl` is built with Go and can be installed using `go install`:

```bash
go install github.com/EpykLab/wazctl@latest
```

Make sure your `$(go env GOPATH)/bin` directory is in your system's `PATH`.

## Getting Started

Getting up and running with `wazctl` is easy.

### 1. Create a Configuration File

First, you need to tell `wazctl` how to connect to your Wazuh manager. Generate a configuration file with the `init config` command.

```bash
wazctl init config
```

This will create a `.wazctl.yaml` file in your current directory with the following structure:

```yaml
wazuh:
  endpoint: your-instance.com
  port: "55000"
  protocol: https
  wuiPassword: password
  wuiUsername: wui
  httpDebug: false
  skipTlsVerify: true
indexer:
  endpoint: your-instance.com
  port: "9200"
  protocol: https
  indexerPassword: password
  indexerUsername: wui
  httpDebug: false
  skipTlsVerify: true
local:
  repoVersion: "v4.12.0"   # include the leading 'v' for the version
```

Edit this file with your Wazuh API credentials and endpoint details. `wazctl` looks for config in (first found wins): `.wazctl.yaml`, `~/.wazctl.yaml`, `~/.config/wazctl.yaml`.

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

## Local environment (Docker) setup

You can run a full Wazuh single-node stack in Docker for development or testing. Config is **optional** for starting the local env: you can run `wazctl localenv docker --start` with no config file; wazctl will use default values (e.g. Wazuh Docker repo version `v4.12.0`).

### Prerequisites

- **Docker** and **Docker Compose** installed and running
- **Git** (used to clone the [wazuh-docker](https://github.com/wazuh/wazuh-docker) repository)

### Config for localenv

| Scenario | Config required? | What to set |
|----------|------------------|-------------|
| Start/stop/clean local stack | No | Omit config or use defaults. If you have a config file, the `local` section is used. |
| Use wazctl against the local stack (e.g. `test auth`, `api agents list`) | Yes | Create and edit `.wazctl.yaml` so the **wazuh** section points at your local instance. |

**Optional: pin Wazuh Docker version**

If you create a config file (e.g. `wazctl init config`), you can set the Wazuh Docker repo version under `local`:

```yaml
local:
  repoVersion: "v4.12.0"   # must include the leading 'v'; used by localenv docker --start
```

If you omit config or leave `local.repoVersion` unset, wazctl uses `v4.12.0`.

**After the stack is running: point wazctl at it**

To run `wazctl test auth`, `wazctl api agents list`, etc. against your local stack, set the **wazuh** section to the manager API (typically HTTPS on localhost, port 55000). The single-node Docker deployment uses default credentials; use the same in your config or the credentials you configured in the stack:

```yaml
wazuh:
  endpoint: localhost
  port: "55000"
  protocol: https
  wuiUsername: admin
  wuiPassword: SecretPassword
  skipTlsVerify: true
  httpDebug: false
```

Place this in `.wazctl.yaml` in your project dir, or in `~/.wazctl.yaml` / `~/.config/wazctl.yaml`. Then run `wazctl test auth` to confirm connectivity.

### Where the deployment lives

- **Working directory:** `~/.wazuh-docker`
- **Single-node Compose:** `~/.wazuh-docker/single-node`

The first time you run `--start`, wazctl clones the wazuh-docker repo into `~/.wazuh-docker`, generates certificates, and brings the stack up. Later `--start` runs reuse that clone.

### Lifecycle commands

| Command | Effect |
|---------|--------|
| `wazctl localenv docker --start` | Clone repo (if needed), generate certs, start the stack. Dashboard is at **https://localhost** (user `admin`, password `SecretPassword`). |
| `wazctl localenv docker --stop` | Stop containers; data and `~/.wazuh-docker` are left in place. |
| `wazctl localenv docker --clean` | Stop containers, remove volumes, and delete `~/.wazuh-docker`. Use this for a full reset. |

### Suggested workflow: try local stack first, then add config

1. Start the local stack (no config needed):  
   `wazctl localenv docker --start`
2. Wait until the dashboard is up at https://localhost (admin / SecretPassword).
3. Create a config so wazctl can talk to it:  
   `wazctl init config`  
   Then edit `.wazctl.yaml`: set `wazuh.endpoint` to `localhost`, `wazuh.port` to `"55000"`, `wazuh.wuiUsername` / `wazuh.wuiPassword` to the dashboard credentials, and `wazuh.skipTlsVerify: true`.
4. Test:  
   `wazctl test auth`  
   `wazctl api agents list`

## CLI Reference

All commands and flags. Use `wazctl [command] --help` for details.

| Command | Description | Flags |
|---------|-------------|--------|
| `wazctl` | Base CLI (no default action) | `-t, --toggle` (misc), `-h, --help` |
| **init** | Scaffold config or rule files | `-h, --help` |
| `wazctl init config` | Create `.wazctl.yaml` in current directory | (none) |
| `wazctl init rule` | Create a new rule test YAML file | `-n, --name` (required): base name for the file (e.g. `my_test` → `my_test.yaml`) |
| **config** | Same as `init config` | (none) |
| **rule** | Same as `init rule` | `-n, --name` (required) |
| **localenv** | Launch or manage a local Wazuh instance | `-h, --help` |
| `wazctl localenv docker` | Run Wazuh in Docker (clone repo, compose) | `--start`: start instance, `--stop`: stop instance, `--clean`: remove instance (volumes) |
| **api** | Wazuh API commands | `-h, --help` |
| `wazctl api agents list` | List agents enrolled in the manager (JSON) | (none) |
| **agents** | Same as `api agents` | `-h, --help` |
| `wazctl agents list` | Same as `api agents list` | (none) |
| **test** | Test connectivity and auth | `-h, --help` |
| `wazctl test auth` | Authenticate and print JWT | (none) |
| **user** | Manage users (Wazuh or Indexer) | `-h, --help` |
| `wazctl user add` | Create a new user | `-u, --username` (required), `-p, --password` (required), `-c, --component` (required): `wazuh` or `indexer`, `-r, --role`: indexer role (required when `component=indexer`) |
| **help** | Help for any command | `wazctl help [command]` |
| **completion** | Shell completion (Cobra) | `wazctl completion [bash\|zsh\|fish\|powershell]` |

Config file search order: `.wazctl.yaml` (current dir), then `~/.wazctl.yaml`, then `~/.config/wazctl.yaml`.

## Example Workflows

### First-time setup and verify connection

```bash
wazctl init config
# Edit .wazctl.yaml with your endpoint and credentials
wazctl test auth
# Expect a JWT printed on success
wazctl api agents list
```

### Local Wazuh in Docker

See **Local environment (Docker) setup** above for prerequisites, config, and directory layout.

```bash
wazctl localenv docker --start
# When done:
wazctl localenv docker --stop
# To remove containers and volumes:
wazctl localenv docker --clean
```

### Add a Wazuh manager user

```bash
wazctl user add --username alice --password secret --component wazuh
# Or short: wazctl user add -u alice -p secret -c wazuh
```

### Add an Indexer user with role

```bash
wazctl user add -u bob -p secret -c indexer -r all_access
# -r/--role is required when component is indexer
```

### Scaffold and customize a rule test

```bash
wazctl init rule -n ssh_bruteforce
# Edit ssh_bruteforce.yaml (ruleId, edges, commands, expected_outcome)
```

## Project Roadmap

This project is under active development. Here is a look at what's done and what's planned.

  * [x] **Initial Setup Commands** (`init config`, `init rule` or `config`, `rule`)
  * [x] **Authentication** (`test auth`)
  * [x] **List Wazuh Agents** (`api agents list` or `agents list`)
  * [x] **Local Docker environment** (`localenv docker --start/--stop/--clean`)
  * [x] **User management** (`user add` for Wazuh and Indexer)
  * [ ] **Rule Test Execution Engine** (e.g., `wazctl test rule <file.yaml>`)
  * [ ] **Expanded Agent Management** (e.g., `restart`, `update`, `remove` agents)
  * [ ] **Enhanced Output Formatting** (Tables, JSON, etc.)
  * [ ] **Broader API Support** (Managing rules, decoders, CDB lists, etc.)
  * [ ] **Pre-compiled Binaries** for multiple platforms.
  ...and much more.

## How to Contribute

Contributions are what make the open-source community such an amazing place to
learn, inspire, and create. Any contributions you make are **greatly
appreciated**. This Wazuh API is huge, so this is a massive undertaking.

1.  **Fork the Project**
2.  **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4.  **Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **Open a Pull Request**

Please feel free to open an issue with the tag "bug" or "enhancement" as well!

## License

Distributed under the MIT License. See the `LICENSE` file in the original
repository for more information. The copyright notice in the source files
indicates it is available under a permissive license.
