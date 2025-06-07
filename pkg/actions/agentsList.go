package actions

import (
	"encoding/json"
	"io"
	"log"
)

func (ctl *WazctlClient) GetAllAgentsFromWazuhManager() ([]byte, error) {

	resp, httpResp, err := ctl.Client.AgentsAPI.ApiControllersAgentControllerGetAgents(ctl.Ctx).
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
		}
		log.Fatalf("Get agents failed: %v", err)
	}

	return resp.Data.MarshalJSON()

}
