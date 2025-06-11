package actions

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/EpykLab/wazctl/models/configurations"
)

type Response struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Error int `json:"error"`
}

type Auth struct {
	b []byte
}

func AuthWithUsernameAndPassword(config configurations.WazuhCtlConfig) *Auth {
	// Replace these with your actual values

	// Create the request
	url := fmt.Sprintf("%s://%s:%s/security/user/authenticate",
		config.WazuhInstanceConfigurations.Protocol,
		config.WazuhInstanceConfigurations.Endpoint,
		config.WazuhInstanceConfigurations.Port)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil
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
		return nil
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return nil
	}

	return &Auth{b: body}
}

func (a *Auth) JWT() *Auth {

	var response Response

	err := json.Unmarshal(a.b, &response)
	if err != nil {
		log.Println(err)
	}

	return &Auth{
		b: []byte(response.Data.Token),
	}
}

func (a *Auth) String() string {
	return string(a.b)
}
