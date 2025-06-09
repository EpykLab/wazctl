package rules

import (
	"bytes"
	v1 "github.com/EpykLab/wazctl/models/schemas/rules/v1"
	"text/template"
)

const exampleRule = `<rule id="100234" level="3">
    <if_sid>230</if_sid>
    <field name="alert_type">normal</field>
    <description>The file limit set for this agent is $(file_limit). Now, $(file_count) files are being monitored.</description>
   <group>syscheck,fim_db_state,</group>
 </rule>`

func ScaffoldFromTempl() bytes.Buffer {

	tmpl := `ruleId: {{.RuleId}}
ruleName: {{.RuleName}}
ruleAuthor: {{.RuleAuthor}}
ruleContent: |-
 {{.RuleContent}}
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
		RuleContent: exampleRule,
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
