package actions

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/EpykLab/wazctl/config"
	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

type IdexerClientMethods interface {
	// User operations
	CreateNewUserInOSIndexer()
}

type IndexerClient struct {
	Client *opensearchapi.Client
	Ctx    context.Context
}

// Config creates and validates the Wazuh API client configuration
func IndexerConfig() (*opensearchapi.Config, error) {
	// Initialize the client with SSL/TLS enabled.
	confs, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &opensearchapi.Config{
		Client: opensearch.Config{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: confs.IndexerInstanceConfiguration.SkipTlsVerify},
			},
			Addresses: []string{fmt.Sprintf("%s://%s:%s",
				confs.IndexerInstanceConfiguration.Protocol,
				confs.IndexerInstanceConfiguration.Endpoint,
				confs.IndexerInstanceConfiguration.Port)},
			Username: confs.IndexerInstanceConfiguration.IndexerUsername,
			Password: confs.IndexerInstanceConfiguration.IndexerPassword,
		},
	}, nil
}

func IndexerClientFactory() *opensearchapi.Client {

	config, err := IndexerConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := opensearchapi.NewClient(*config)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
