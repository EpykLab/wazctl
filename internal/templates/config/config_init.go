package config

import (
	"bytes"
	"github.com/EpykLab/wazctl/models/configurations"
	"text/template"
)

func ScaffoldFromTempl() bytes.Buffer {

	tmpl := `endpoint: {{.Endpoint}}
port: {{.Port}}
protocol: {{.Protocol}}
wuiPassword: {{.WuiPassword}}
wuiUsername: {{.WuiUsername}}`

	data := configurations.WazuhCtlConfig{
		Endpoint:    "your-instance.com",
		Port:        "55000",
		Protocol:    "https",
		WuiPassword: "password",
		WuiUsername: "wui",
	}

	var buf bytes.Buffer

	t := template.Must(template.New("ruleTest").Parse(tmpl))
	t.Execute(&buf, data)

	return buf
}
