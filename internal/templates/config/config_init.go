package config

import (
	"bytes"
	"github.com/EpykLab/wazctl/models/configurations"
	"text/template"
)

func ScaffoldFromTempl() bytes.Buffer {

	tmpl := `endpoint: {{.Endpoint}}
wuiPassword: {{.WuiPassword}}
wuiUsername: {{.WuiUsername}}`

	data := configurations.WazuhCtlConfig{
		Endpoint:    "https://your-instance:5000",
		WuiPassword: "password",
		WuiUsername: "wui",
	}

	var buf bytes.Buffer

	t := template.Must(template.New("ruleTest").Parse(tmpl))
	t.Execute(&buf, data)

	return buf
}
