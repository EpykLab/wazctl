package config

import (
	"bytes"
	"github.com/EpykLab/wazctl/models/configurations"
	"text/template"
)

func ScaffoldFromTempl() bytes.Buffer {
	tmpl := `wazuh:
    endpoint: {{.WazuhInstanceConfigurations.Endpoint}}
    port: {{.WazuhInstanceConfigurations.Port}}
    protocol: {{.WazuhInstanceConfigurations.Protocol}}
    wuiPassword: {{.WazuhInstanceConfigurations.WuiPassword}}
    wuiUsername: {{.WazuhInstanceConfigurations.WuiUsername}}
    httpDebug: {{.WazuhInstanceConfigurations.HttpDebug}}
    skipTlsVerify: {{.WazuhInstanceConfigurations.SkipTlsVerify}}
indexer:
    endpoint: {{.IndexerInstanceConfiguration.Endpoint}}
    port: {{.IndexerInstanceConfiguration.Port}}
    protocol: {{.IndexerInstanceConfiguration.Protocol}}
    indexerPassword: {{.IndexerInstanceConfiguration.IndexerPassword}}
    indexerUsername: {{.IndexerInstanceConfiguration.IndexerUsername}}
    httpDebug: {{.IndexerInstanceConfiguration.HttpDebug}}
    skipTlsVerify: {{.IndexerInstanceConfiguration.SkipTlsVerify}}
local:
    repoVersion: {{.LocalInstanceConfiguration.RepoVersion}}
`

	data := configurations.WazuhCtlConfig{
		WazuhInstanceConfigurations: configurations.WazuhInstanceConfigurations{
			Endpoint:      "your-instance.com",
			Port:          "55000",
			Protocol:      "https",
			WuiPassword:   "password",
			WuiUsername:   "wui",
			SkipTlsVerify: true,
			HttpDebug:     false,
		},
		IndexerInstanceConfiguration: configurations.IndexerInstanceConfiguration{
			Endpoint:        "your-instance.com",
			Port:            "9200",
			Protocol:        "https",
			IndexerPassword: "password",
			IndexerUsername: "wui",
			SkipTlsVerify:   true,
			HttpDebug:       false,
		},
		LocalInstanceConfiguration: configurations.LocalInstanceConfiguration{
			RepoVersion: "4.12.0",
		},
	}

	var buf bytes.Buffer

	t := template.Must(template.New("wazuhConfig").Parse(tmpl))
	t.Execute(&buf, data)

	return buf
}
