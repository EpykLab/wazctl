package actions

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	api "github.com/EpykLab/wasabi"
	"github.com/EpykLab/wazctl/config"
)

// Config creates and validates the Wazuh API client configuration
func Config() (*api.Configuration, error) {
	confs, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Validate configuration
	if confs.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is empty")
	}
	if confs.Protocol != "http" && confs.Protocol != "https" {
		return nil, fmt.Errorf("invalid protocol: %s", confs.Protocol)
	}
	if confs.Port == "" {
		return nil, fmt.Errorf("port is empty")
	}

	cfg := api.NewConfiguration()
	cfg.Host = confs.Endpoint
	cfg.Scheme = confs.Protocol
	cfg.Servers[0].Variables["port"] = api.ServerVariable{DefaultValue: confs.Port}
	cfg.UserAgent = "WazctlClient/1.0"
	cfg.Debug = true
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // WARNING: Insecure for testing
			},
		},
	}

	// Log configuration
	log.Printf("Config: Host=%s, Scheme=%s, Port=%s", cfg.Host, cfg.Scheme, confs.Port)

	return cfg, nil
}
