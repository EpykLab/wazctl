package actions

import (
	"crypto/tls"
	"net/http"

	"github.com/EpykLab/wazctl/pkg/opensearch"
)

type IndexerClient struct {
	osConfig opensearch.IndexerClientConfig
	client   http.Client
}

func IndexerClientFactory() *IndexerClient {

	return &IndexerClient{
		osConfig: *opensearch.NewClientConfig(),
		client: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}
