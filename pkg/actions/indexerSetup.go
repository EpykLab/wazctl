package actions

import (
	"github.com/EpykLab/wazctl/pkg/opensearch"
)

type IndexerClient struct {
	oSConfig opensearch.IndexerClientConfig
}

func IndexerClientFactory() *IndexerClient {

	return &IndexerClient{
		oSConfig: *opensearch.NewClientConfig(),
	}
}
