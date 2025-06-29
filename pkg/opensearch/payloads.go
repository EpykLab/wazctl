package opensearch

type UserAttributes struct {
}

// New user payload as expected by the API
type NewUserPayload struct {
	Attributes struct {
		Attribute1 string `json:"attribute1,omitzero"`
		Attribute2 string `json:"attribute2,omitzero"`
	} `json:"attributes"`
	BackendRoles            []string `json:"backend_roles"`
	OpendistroSecurityRoles []string `json:"opendistro_security_roles,omitempty"`
	Password                string   `json:"password,omitzero"`
}

