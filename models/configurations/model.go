package configurations

type WazuhCtlConfig struct {
	Endpoint    string `json:"endpoint,omitempty" yaml:"endpoint"`
	WuiPassword string `json:"wui_password,omitempty" yaml:"wuiPassword"`
	WuiUsername string `json:"wui_username,omitempty" yaml:"wuiUsername"`
}
