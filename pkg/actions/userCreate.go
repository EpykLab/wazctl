package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	wasabi "github.com/EpykLab/wasabi"
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
		return nil, errors.New(fmt.Sprintf("Error when calling `SecurityAPI.ApiControllersSecurityControllerCreateUser``: %v\n", err))
	}

	return resp.MarshalJSON()
}

func (ctl *IndexerClient) CreateNewUserInOSIndexer(opts *CreateNewIndexerUserOptions) ([]byte, error) {

	// TODO: create query for new users
	query := strings.NewReader(`{}`)

	createUserRequest, err := http.NewRequest(http.MethodPost, "", query)
	if err != nil {
		return nil, err
	}

	createUserRequest.Header["Content-Type"] = []string{"application/json"}

	resp, err := ctl.Client.Client.Perform(createUserRequest)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

// TODO: Create ability to map roles in wazuh
