package configurations

type WazuhCtlConfig struct {
	Endpoint      string `json:"endpoint,omitempty" yaml:"endpoint"`
	Protocol      string `json:"protocol,omitempty" yaml:"protocol"`
	Port          string `json:"port,omitempty" yaml:"port"`
	SkipTlsVerify bool   `json:"skip_tls_verify,omitempty" yaml:"skipTlsVerify"`
	HttpDebug     bool   `json:"http_debug,omitempty" yaml:"httpDebug"`
	WuiPassword   string `json:"wui_password,omitempty" yaml:"wuiPassword"`
	WuiUsername   string `json:"wui_username,omitempty" yaml:"wuiUsername"`
}
