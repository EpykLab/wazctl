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

type Client interface {
	// Agent operations
	GetAllAgentsFromWazuhManager()
}

type WazctlClient struct {
	Client *api.APIClient
	Ctx    context.Context
}

// Config creates and validates the Wazuh API client configuration
func Config() (*api.Configuration, error) {
	confs, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	log.Println(confs.SkipTlsVerify)

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

	// ✅ **THE FIX IS HERE** ✅
	// Instead of setting Host, Scheme, and Port variables separately,
	// we construct the full server URL directly. This ensures the
	// Host header includes the port.
	serverURL := fmt.Sprintf("%s://%s:%s", confs.Protocol, confs.Endpoint, confs.Port)
	cfg.Servers = api.ServerConfigurations{
		{
			URL:         serverURL,
			Description: "Wazuh API Server",
		},
	}

	cfg.UserAgent = "WazctlClient/1.0"
	cfg.Debug = confs.HttpDebug
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: confs.SkipTlsVerify,
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

	config, err := Config()
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
