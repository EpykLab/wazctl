package auth

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/EpykLab/wazctl/models/configurations"
)

func AuthWithUsernameAndPassword(config configurations.WazuhCtlConfig) (string, error) {
	// Replace these with your actual values

	// Create the request
	url := fmt.Sprintf("%s/security/user/authenticate", config.Endpoint)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", err
	}

	// Set basic authentication
	req.SetBasicAuth(config.WuiUsername, config.WuiPassword)

	// Create HTTP client with InsecureSkipVerify to mimic -k (insecure)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", err
	}

	return string(body), nil
}
