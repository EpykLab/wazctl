package agents

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	api "github.com/EpykLab/wasabi"
	"github.com/EpykLab/wazctl/config"
	"github.com/EpykLab/wazctl/pkg/actions"
	"github.com/EpykLab/wazctl/pkg/actions/auth"
)

func GetAllAgentsFromWazuhManager() {

	conf, err := config.New()
	if err != nil {
		log.Println(err)
	}

	token := auth.AuthWithUsernameAndPassword(*conf).JWT().String()

	config, err := actions.Config()
	if err != nil {
		log.Println(err)
	}

	client := api.NewAPIClient(config)
	ctx := context.WithValue(context.Background(),
		api.ContextAccessToken,
		strings.TrimSpace(token))

	_, httpResp, err := client.AgentsAPI.ApiControllersAgentControllerGetAgents(ctx).
		Pretty(true).
		Execute()
	if err != nil {
		// Log raw response
		if httpResp != nil {
			body, _ := io.ReadAll(httpResp.Body)
			var wazuhErr map[string]any
			err := json.Unmarshal(body, &wazuhErr)
			if err != nil {

			}
			httpResp.Body.Close()
			fmt.Sprint(wazuhErr)
		}
		log.Fatalf("Get agents failed: %v", err)
	}
}
