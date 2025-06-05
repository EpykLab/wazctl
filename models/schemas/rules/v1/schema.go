package v1

import "encoding/json"
import "fmt"
import "reflect"

// Schema for defining Wazuh rule test cases using edge cases with executable
// commands
type SchemaJson struct {
	// Description of the rule and its purpose
	Description string `json:"description" yaml:"description" mapstructure:"description"`

	// List of edge cases to test the rule
	Edges []SchemaJsonEdgesElem `json:"edges" yaml:"edges" mapstructure:"edges"`

	// Author of the rule
	RuleAuthor string `json:"rule_author,omitempty" yaml:"ruleAuthor,omitempty" mapstructure:"ruleAuthor,omitempty"`

	// Unique identifier for the Wazuh rule
	RuleId string `json:"rule_id" yaml:"ruleId" mapstructure:"ruleId"`

	// Human-readable name of the rule
	RuleName string `json:"rule_name" yaml:"ruleName" mapstructure:"ruleName"`
}

type SchemaJsonEdgesElem struct {
	// Command to execute to trigger the rule
	Command SchemaJsonEdgesElemCommand `json:"command" yaml:"command" mapstructure:"command"`

	// Description of the edge case and expected behavior
	Description string `json:"description" yaml:"description" mapstructure:"description"`

	// Expected outcome when the command is executed (e.g., rule triggered or not)
	ExpectedOutcome string `json:"expected_outcome" yaml:"expected_outcome" mapstructure:"expected_outcome"`

	// Title of the edge case
	Title string `json:"title" yaml:"title" mapstructure:"title"`
}

// Command to execute to trigger the rule
type SchemaJsonEdgesElemCommand struct {
	// Type of command
	Type SchemaJsonEdgesElemCommandType `json:"type" yaml:"type" mapstructure:"type"`

	// The command to execute
	Value string `json:"value" yaml:"value" mapstructure:"value"`
}

type SchemaJsonEdgesElemCommandType string

const SchemaJsonEdgesElemCommandTypeBash SchemaJsonEdgesElemCommandType = "bash"
const SchemaJsonEdgesElemCommandTypePowershell SchemaJsonEdgesElemCommandType = "powershell"

var enumValues_SchemaJsonEdgesElemCommandType = []interface{}{
	"bash",
	"powershell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJsonEdgesElemCommandType) UnmarshalJSON(value []byte) error {
	var v string
	if err := json.Unmarshal(value, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SchemaJsonEdgesElemCommandType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SchemaJsonEdgesElemCommandType, v)
	}
	*j = SchemaJsonEdgesElemCommandType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJsonEdgesElemCommand) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in SchemaJsonEdgesElemCommand: required")
	}
	if _, ok := raw["value"]; raw != nil && !ok {
		return fmt.Errorf("field value in SchemaJsonEdgesElemCommand: required")
	}
	type Plain SchemaJsonEdgesElemCommand
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Value) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "value", 1)
	}
	*j = SchemaJsonEdgesElemCommand(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJsonEdgesElem) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["command"]; raw != nil && !ok {
		return fmt.Errorf("field command in SchemaJsonEdgesElem: required")
	}
	if _, ok := raw["description"]; raw != nil && !ok {
		return fmt.Errorf("field description in SchemaJsonEdgesElem: required")
	}
	if _, ok := raw["expected_outcome"]; raw != nil && !ok {
		return fmt.Errorf("field expected_outcome in SchemaJsonEdgesElem: required")
	}
	if _, ok := raw["title"]; raw != nil && !ok {
		return fmt.Errorf("field title in SchemaJsonEdgesElem: required")
	}
	type Plain SchemaJsonEdgesElem
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Description) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "description", 1)
	}
	if len(plain.ExpectedOutcome) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "expected_outcome", 1)
	}
	if len(plain.Title) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "title", 1)
	}
	*j = SchemaJsonEdgesElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["description"]; raw != nil && !ok {
		return fmt.Errorf("field description in SchemaJson: required")
	}
	if _, ok := raw["edges"]; raw != nil && !ok {
		return fmt.Errorf("field edges in SchemaJson: required")
	}
	if _, ok := raw["rule_id"]; raw != nil && !ok {
		return fmt.Errorf("field rule_id in SchemaJson: required")
	}
	if _, ok := raw["rule_name"]; raw != nil && !ok {
		return fmt.Errorf("field rule_name in SchemaJson: required")
	}
	type Plain SchemaJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Description) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "description", 1)
	}
	if plain.Edges != nil && len(plain.Edges) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "edges", 1)
	}
	if len(plain.RuleAuthor) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "rule_author", 1)
	}
	if len(plain.RuleId) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "rule_id", 1)
	}
	if len(plain.RuleName) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "rule_name", 1)
	}
	*j = SchemaJson(plain)
	return nil
}
