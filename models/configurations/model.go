package configurations

type WazuhCtlConfig struct {
	WazuhInstanceConfigurations  `json:"wazuh_instance_configurations" yaml:"wazuh"`
	IndexerInstanceConfiguration `json:"indexer_instance_configuration" yaml:"indexer"`
	LocalInstanceConfiguration   `json:"local_instance_configuration" yaml:"local"`
}

type WazuhInstanceConfigurations struct {
	Endpoint      string `json:"endpoint,omitempty" yaml:"endpoint"`
	Protocol      string `json:"protocol,omitempty" yaml:"protocol"`
	Port          string `json:"port,omitempty" yaml:"port"`
	SkipTlsVerify bool   `json:"skip_tls_verify,omitempty" yaml:"skipTlsVerify"`
	HttpDebug     bool   `json:"http_debug,omitempty" yaml:"httpDebug"`
	WuiPassword   string `json:"wui_password,omitempty" yaml:"wuiPassword"`
	WuiUsername   string `json:"wui_username,omitempty" yaml:"wuiUsername"`
}

type IndexerInstanceConfiguration struct {
	Endpoint        string `json:"endpoint,omitempty" yaml:"endpoint"`
	Protocol        string `json:"protocol,omitempty" yaml:"protocol"`
	Port            string `json:"port,omitempty" yaml:"port"`
	SkipTlsVerify   bool   `json:"skip_tls_verify,omitempty" yaml:"skipTlsVerify"`
	HttpDebug       bool   `json:"http_debug,omitempty" yaml:"httpDebug"`
	IndexerPassword string `json:"wui_password,omitempty" yaml:"indexerPassword"`
	IndexerUsername string `json:"wui_username,omitempty" yaml:"indexerUsername"`
}

type LocalInstanceConfiguration struct {
	RepoVersion string `json:"repo_version,omitempty" yaml:"repoVersion"`
}
