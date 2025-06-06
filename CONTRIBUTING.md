# 🤝 Contributing to wazctl

First off, thank you for considering contributing to `wazctl`! It's people like
you that make the open-source community such a fantastic place. We welcome any
and all contributions, from bug reports to new features.

Following these guidelines helps to communicate that you respect the time of
the developers managing and developing this open-source project. In return,
they should reciprocate that respect in addressing your issue, assessing
changes, and helping you finalize your pull requests.

## 🤔 How Can I Contribute?

There are many ways to contribute, and all of them are valuable.

### ❤️ Sponsoring

If you rely on this project, please consider sponsoring its development. It
helps us dedicate more time to maintenance and new features. You can sponsor
the project through [GitHub
Sponsors](https://github.com/sponsors/YOUR_USERNAME). Any amount is greatly
appreciated!

### 🐛 Reporting Bugs

If you find a bug, please make sure it hasn't already been reported by
searching the [Issues](https://github.com/EpykLab/wazctl/issues) on GitHub.

If you can't find an open issue addressing the problem, please [open a new
one](https://github.com/EpykLab/wazctl/issues/new). Be sure to include:
* **A clear and descriptive title**.
* **A detailed description of the problem**, including steps to reproduce it.
* **The version of `wazctl`** you are using.
* **Your operating system**.
* **Any relevant logs or error messages**.

### ✨ Suggesting Enhancements

If you have an idea for a new feature or an improvement to an existing one:
1.  **Check the project roadmap** in the `README.md` to see if your idea is
    already planned.
2.  **Search the [Issues](https://github.com/EpykLab/wazctl/issues)** to see if
    the enhancement has already been suggested.
3.  If not, [open a new issue](https://github.com/EpykLab/wazctl/issues/new)
    with the "enhancement" label, providing a clear and detailed description of
your suggestion and why it would be valuable.

### 📝 Submitting a Pull Request

If you're ready to contribute code, that's fantastic! Here’s how to get started.

## 🛠️ Development Setup

To get your local development environment up and running, follow these steps:

1.  **Prerequisites**:
    * Go (version 1.18 or later is recommended)
    * Git

2.  **Fork and Clone the Repository**:
    * Fork the repository on GitHub.
    * Clone your fork locally:
        ```bash
        git clone [https://github.com/YOUR_USERNAME/wazctl.git](https://github.com/YOUR_USERNAME/wazctl.git)
        cd wazctl
        ```

3.  **Build the Project**:
    You can build the `wazctl` binary using the standard Go command:
    ```bash
    go build .
    ```
    This will create a `wazctl` executable in the root directory. You can run it with `./wazctl`.

4.  **Run Tests**:
    Before making any changes, make sure the existing tests pass:
    ```bash
    go test ./...
    ```

## ✅ Pull Request Process

1.  **Create a new branch** for your feature or bugfix. Please use a descriptive branch name (e.g., `feature/add-agent-restart` or `fix/config-parsing-error`).
    ```bash
    git checkout -b feature/my-amazing-feature
    ```

2.  **Make your changes**. Write clean, readable code and add comments where necessary.

3.  **Add or update tests**. If you're adding a new feature, please include tests for it. If you're fixing a bug, add a test that catches the bug to prevent regressions.

4.  **Ensure all tests pass**:
    ```bash
    go test ./...
    ```

5.  **Format your code**:
    Make sure your Go code is formatted correctly using `gofmt`.
    ```bash
    gofmt -w .
    ```

6.  **Commit your changes**. Use a clear and descriptive commit message. We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. For example:
    * `feat: add restart functionality for agents`
    * `fix: resolve panic when config file is empty`
    * `docs: update agent management section in README`

7.  **Push your branch** to your fork on GitHub:
    ```bash
    git push origin feature/my-amazing-feature
    ```

8.  **Open a Pull Request** to the `main` branch of the original `wazctl`
    repository. Provide a clear title and a detailed description of your
changes. Link to any relevant issues.

Thank you again for your contribution!
