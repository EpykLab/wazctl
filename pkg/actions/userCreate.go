package actions

import (
	"errors"
	"fmt"
	"os"

	wasabi "github.com/EpykLab/wasabi"
)

type CreateNewUserInWazuhManagerOptions struct {
	Username string
	Password string
}

// Creates a new user in the wazuh instance
func (ctl *WazctlClient) CreateNewUserInWazuhManager(opts *CreateNewUserInWazuhManagerOptions) ([]byte, error) {

	pretty := true
	waitForComplete := true
	newUser := wasabi.NewApiControllersSecurityControllerCreateUserRequest(opts.Username, opts.Password)

	resp, r, err := ctl.Client.SecurityAPI.ApiControllersSecurityControllerCreateUser(ctl.Ctx).
		Pretty(pretty).
		WaitForComplete(waitForComplete).
		ApiControllersSecurityControllerCreateUserRequest(*newUser).
		Execute()
	if err != nil {

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, errors.New(fmt.Sprintf("Error when calling `SecurityAPI.ApiControllersSecurityControllerCreateUser``: %v\n", err))
	}

	return resp.MarshalJSON()
}
