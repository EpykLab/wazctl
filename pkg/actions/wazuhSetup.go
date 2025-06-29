package actions

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	api "github.com/EpykLab/wasabi"
	"github.com/EpykLab/wazctl/config"
)

type WazuhClientMethods interface {
	// Agent operations
	GetAllAgentsFromWazuhManager()

	// User operations
	CreateNewUserInWazuh()
	GetUsersInWazuh()
}

type WazctlClient struct {
	Client *api.APIClient
	Ctx    context.Context
}

// WazuhConfig creates and validates the Wazuh API client configuration
func WazuhConfig() (*api.Configuration, error) {
	confs, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Validate configuration
	if confs.WazuhInstanceConfigurations.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is empty")
	}
	if confs.WazuhInstanceConfigurations.Protocol != "http" && confs.WazuhInstanceConfigurations.Protocol != "https" {
		return nil, fmt.Errorf("invalid protocol: %s", confs.WazuhInstanceConfigurations.Protocol)
	}
	if confs.WazuhInstanceConfigurations.Port == "" {
		return nil, fmt.Errorf("port is empty")
	}

	cfg := api.NewConfiguration()

	// ✅ **THE FIX IS HERE** ✅
	// Instead of setting Host, Scheme, and Port variables separately,
	// we construct the full server URL directly. This ensures the
	// Host header includes the port.
	serverURL := fmt.Sprintf("%s://%s:%s", confs.WazuhInstanceConfigurations.Protocol,
		confs.WazuhInstanceConfigurations.Endpoint,
		confs.WazuhInstanceConfigurations.Port)
	cfg.Servers = api.ServerConfigurations{
		{
			URL:         serverURL,
			Description: "Wazuh API Server",
		},
	}

	cfg.UserAgent = "WazctlClient/1.0"
	cfg.Debug = confs.WazuhInstanceConfigurations.HttpDebug
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: confs.WazuhInstanceConfigurations.SkipTlsVerify,
			},
		},
	}

	// Log configuration
	log.Printf("Config: ServerURL=%s", serverURL)

	return cfg, nil
}

func WazctlClientFactory() *WazctlClient {

	conf, err := config.New()
	if err != nil {
		log.Println(err)
	}

	token := AuthWithUsernameAndPassword(*conf).JWT().String()

	config, err := WazuhConfig()
	if err != nil {
		log.Println(err)
	}

	return &WazctlClient{
		Client: api.NewAPIClient(config),
		Ctx: context.WithValue(context.Background(),
			api.ContextAccessToken,
			strings.TrimSpace(token)),
	}
}
