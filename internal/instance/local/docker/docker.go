package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// WazuhDockerManager manages the Wazuh Docker deployment.
type WazuhDockerManager struct {
	RepoURL       string // Git repository URL for Wazuh Docker
	RepoVersion   string // Version tag (e.g., "v4.12.0")
	WorkDir       string // Directory to clone the repository into
	SingleNodeDir string // Path to the single-node directory
}

// NewWazuhDockerManager initializes a new WazuhDockerManager.
func NewWazuhDockerManager() (*WazuhDockerManager, error) {
	// Use a directory in the user's home for persistence
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get user home directory: %w", err)
	}
	workDir := filepath.Join(home, ".wazuh-docker")
	return &WazuhDockerManager{
		RepoURL:       "https://github.com/wazuh/wazuh-docker.git",
		RepoVersion:   "v4.12.0",
		WorkDir:       workDir,
		SingleNodeDir: filepath.Join(workDir, "single-node"),
	}, nil
}

// ensureGit ensures the git command is available.
func (m *WazuhDockerManager) ensureGit() error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("git not found in PATH; please install Git")
	}
	fmt.Println("✅ Git is installed.")
	return nil
}

// cloneRepository clones the Wazuh Docker repository if it doesn't exist.
func (m *WazuhDockerManager) cloneRepository() error {
	if _, err := os.Stat(m.WorkDir); !os.IsNotExist(err) {
		fmt.Printf("✅ Wazuh Docker repository already cloned at %s\n", m.WorkDir)
		return nil
	}

	fmt.Printf("📥 Cloning Wazuh Docker repository (%s) to %s...\n", m.RepoVersion, m.WorkDir)
	cmd := exec.Command("git", "clone", "-b", m.RepoVersion, "--single-branch", m.RepoURL, m.WorkDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone Wazuh Docker repository: %w", err)
	}
	fmt.Println("✅ Repository cloned successfully.")
	return nil
}

// generateCertificates runs the certificate generation step.
func (m *WazuhDockerManager) generateCertificates() error {
	certDir := filepath.Join(m.SingleNodeDir, "config", "wazuh_indexer_ssl_certs")
	if _, err := os.Stat(certDir); !os.IsNotExist(err) {
		// Check if key certificate files exist
		adminCert := filepath.Join(certDir, "admin.pem")
		if _, err := os.Stat(adminCert); err == nil {
			fmt.Println("✅ Certificates already generated.")
			return nil
		}
	}

	fmt.Println("🔒 Generating SSL certificates for Wazuh...")
	cmd := exec.Command("/bin/sh", "-c", "docker compose -f generate-indexer-certs.yml run --no-TTY --rm generator")
	cmd.Dir = m.SingleNodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate certificates: %w", err)
	}
	fmt.Println("✅ Certificates generated successfully.")
	return nil
}

// Start deploys the Wazuh stack using Docker Compose.
func (m *WazuhDockerManager) Start() error {
	// Ensure prerequisites
	if err := m.ensureGit(); err != nil {
		return err
	}

	// Clone the repository
	if err := m.cloneRepository(); err != nil {
		return err
	}

	// Generate certificates
	if err := m.generateCertificates(); err != nil {
		return err
	}

	// Start the Wazuh stack
	fmt.Println("🚀 Starting Wazuh Docker stack...")
	cmd := exec.Command("/bin/bash", "-c", "docker compose up -d")
	cmd.Dir = m.SingleNodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start Wazuh stack: %w", err)
	}

	// Wait for the dashboard to be ready (up to 2 minutes)
	fmt.Println("⏳ Waiting for Wazuh dashboard to be ready (this may take up to 2 minutes)...")
	if err := m.waitForDashboard(); err != nil {
		return fmt.Errorf("Wazuh dashboard failed to start: %w", err)
	}

	fmt.Println("✅ Wazuh stack started successfully.")
	fmt.Println("🌐 Access the Wazuh dashboard at: https://localhost")
	fmt.Println("   Username: admin")
	fmt.Println("   Password: SecretPassword")
	return nil
}

// waitForDashboard polls the Wazuh dashboard container logs to check if it's ready.
func (m *WazuhDockerManager) waitForDashboard() error {
	timeout := 2 * time.Minute
	start := time.Now()
	containerName := "single-node-wazuh.dashboard-1"

	for time.Since(start) < timeout {
		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("docker logs %s", containerName))
		output, err := cmd.CombinedOutput()
		if err != nil {
			// Container might not be up yet, wait and retry
			time.Sleep(5 * time.Second)
			continue
		}

		logs := string(output)
		// Check for the actual ready indicator in Wazuh dashboard logs
		if strings.Contains(logs, "http server running at https://0.0.0.0:5601") {
			return nil
		}

		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("timeout waiting for Wazuh dashboard to be ready")
}

// Stop stops the Wazuh stack.
func (m *WazuhDockerManager) Stop() error {
	if _, err := os.Stat(m.SingleNodeDir); os.IsNotExist(err) {
		fmt.Println("⚠️ Wazuh Docker directory not found; nothing to stop.")
		return nil
	}

	fmt.Println("🛑 Stopping Wazuh Docker stack...")
	cmd := exec.Command("/bin/bash", "-c", "docker compose down")
	cmd.Dir = m.SingleNodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop Wazuh stack: %w", err)
	}

	fmt.Println("✅ Wazuh stack stopped successfully.")
	return nil
}

// Clean removes the Wazuh Docker directory and volumes.
func (m *WazuhDockerManager) Clean() error {
	// Stop the stack first
	if err := m.Stop(); err != nil {
		return err
	}

	// Remove volumes
	fmt.Println("🧹 Removing Wazuh Docker volumes...")
	cmd := exec.Command("/bin/bash", "-c", "docker compose down -v")
	cmd.Dir = m.SingleNodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove Wazuh volumes: %w", err)
	}

	// Remove the working directory
	if err := os.RemoveAll(m.WorkDir); err != nil {
		return fmt.Errorf("failed to remove Wazuh Docker directory: %w", err)
	}

	fmt.Println("✅ Cleaned up Wazuh Docker deployment.")
	return nil
}
