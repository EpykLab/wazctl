package rules

import (
	"bytes"
	v1 "github.com/EpykLab/wazctl/models/schemas/rules/v1"
	"text/template"
)

func ScaffoldFromTempl() bytes.Buffer {

	tmpl := `ruleId: {{.RuleId}}
ruleName: {{.RuleName}}
ruleAuthor: {{.RuleAuthor}}
description: {{.Description}}
edges:
{{- range .Edges}}
  - title: {{.Title}}
    description: {{.Description}}
    command:
      type: {{.Command.Type}}
      value: |-
       {{.Command.Value}}
    expected_outcome: {{.ExpectedOutcome}}
{{- end}}`

	data := v1.SchemaJson{
		RuleId:      "rule_001",
		RuleName:    "Unauthorized Access",
		RuleAuthor:  "John Doe",
		Description: "Tests unauthorized access attempts",
		Edges: []v1.SchemaJsonEdgesElem{
			{
				Title:           "Invalid Login",
				Description:     "Simulate invalid login attempt",
				Command:         v1.SchemaJsonEdgesElemCommand{Type: "bash", Value: "ssh invalid@server"},
				ExpectedOutcome: "Rule triggers alert",
			},
		},
	}

	var buf bytes.Buffer

	t := template.Must(template.New("ruleTest").Parse(tmpl))
	t.Execute(&buf, data)

	return buf
}
