package config

import (
	"log"
	"os"

	"github.com/EpykLab/wazctl/internal/files"
	"github.com/EpykLab/wazctl/models/configurations"
	"gopkg.in/yaml.v3"
)

var (
	// lists the defined locations where the config will be located by default
	configLocs = []string{"~/.wazctl.yaml", "~/.config/wazctl.yaml", ".wazctl.yaml"}
)

func New() *configurations.WazuhCtlConfig {

	var config configurations.WazuhCtlConfig
	var content []byte

	for _, x := range configLocs {
		_, err := os.Stat(x)
		if err != nil {
			continue
		} else {
			content, err = files.ReadFileFromSpecifiedPath(x)
			if err != nil {
				log.Println("encountered error finding and opening config file ", err)
			}
		}
	}

	err := yaml.Unmarshal(content, config)
	if err != nil {
		log.Println(err)
	}

	return &configurations.WazuhCtlConfig{
		Endpoint:    config.Endpoint,
		WuiPassword: config.WuiPassword,
		WuiUsername: config.WuiUsername,
	}
}
