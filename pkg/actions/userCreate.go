package actions

import (
	"fmt"
	"io"
	"net/http"

	wasabi "github.com/EpykLab/wasabi"
	"github.com/EpykLab/wazctl/pkg/opensearch"
)

type CreateNewWazuhUserOptions struct {
	Username string
	Password string
}

type CreateIndexerUserPayload struct {
	Password string `json:"password"`
}

type IndexerRoleMappingPayload struct {
	Users        []string `json:"users,omitempty"`
	BackendRoles []string `json:"backend_roles,omitempty"`
}
type CreateNewIndexerUserOptions struct {
	CreateIndexerUserPayload
	IndexerRoleMappingPayload
}

// Creates a new user in the wazuh instance
func (ctl *WazctlClient) CreateNewUserInWazuhManager(opts *CreateNewWazuhUserOptions) ([]byte, error) {

	pretty := true
	waitForComplete := true
	newUser := wasabi.NewApiControllersSecurityControllerCreateUserRequest(opts.Username, opts.Password)

	resp, _, err := ctl.Client.SecurityAPI.ApiControllersSecurityControllerCreateUser(ctl.Ctx).
		Pretty(pretty).
		WaitForComplete(waitForComplete).
		ApiControllersSecurityControllerCreateUserRequest(*newUser).
		Execute()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error when calling `SecurityAPI.ApiControllersSecurityControllerCreateUser``: %v\n", err))
	}

	return resp.MarshalJSON()
}

func (ctl *IndexerClient) CreateNewUserInOSIndexer(opts *CreateNewIndexerUserOptions) ([]byte, error) {

	userPayload := opensearch.NewUserPayload{
		Attributes: struct {
			Attribute1 string "json:\"attribute1,omitzero\""
			Attribute2 string "json:\"attribute2,omitzero\""
		}{},
		BackendRoles: opts.BackendRoles,
		Password:     opts.Password,
	}

	uri := fmt.Sprintf("%s/%s/%s",
		ctl.osConfig.Address,
		opensearch.CreateNewIdexerUserURI,
		opts.Users[0])

	request, err := ctl.osConfig.IndexerApiRequest(userPayload, uri, http.MethodPut)

	resp, err := ctl.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
