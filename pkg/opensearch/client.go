package opensearch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EpykLab/wazctl/config"
)

type IndexerClientConfig struct {
	SkipTLSVerify bool
	Address       string
	Username      string
	Password      string
	// TODO: added in field for http debug
}

type indexerMethods interface {
	IndexerApiRequest(payload any, uri endpoints) (*http.Request, error)
}

func NewClientConfig() *IndexerClientConfig {

	confs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	return &IndexerClientConfig{
		SkipTLSVerify: confs.IndexerInstanceConfiguration.SkipTlsVerify,
		Address: fmt.Sprintf("%s://%s:%s",
			confs.IndexerInstanceConfiguration.Protocol,
			confs.IndexerInstanceConfiguration.Endpoint,
			confs.IndexerInstanceConfiguration.Port),
		Username: confs.IndexerInstanceConfiguration.IndexerUsername,
		Password: confs.IndexerInstanceConfiguration.IndexerPassword,
	}
}

// IndexerApiRequest()
//
//	:uri of type string should make use of endpoints helpers
func (c *IndexerClientConfig) IndexerApiRequest(payload any, uri string, method string) (*http.Request, error) {

	Authententication := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.Username, c.Password)))

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, reqErr := http.NewRequest(method, uri, bytes.NewBuffer(jsonData))
	if reqErr != nil {
		return nil, reqErr
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", Authententication))
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

