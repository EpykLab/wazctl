package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/EpykLab/wazctl/models/configurations"
)

func TestNew(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, ".wazctl.yaml")
	content := `wazuh:
  endpoint: "10.77.60.11"
  port: "55000"
  protocol: https
  wuiPassword: wazuh-wui
  wuiUsername: wazuh-wui
  skipTlsVerify: false
  httpDebug: false
indexer:
  endpoint: ""
  port: "9200"
  protocol: https
  indexerPassword: ""
  indexerUsername: ""
  skipTlsVerify: false
  httpDebug: false
local:
  repoVersion: "v4.12.0"
`
	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		t.Fatalf("write test config: %v", err)
	}
	prev, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}
	defer func() { _ = os.Chdir(prev) }()

	want := &configurations.WazuhCtlConfig{
		WazuhInstanceConfigurations: configurations.WazuhInstanceConfigurations{
			Endpoint:      "10.77.60.11",
			Port:          "55000",
			Protocol:      "https",
			WuiPassword:   "wazuh-wui",
			WuiUsername:   "wazuh-wui",
			SkipTlsVerify: false,
			HttpDebug:     false,
		},
		IndexerInstanceConfiguration: configurations.IndexerInstanceConfiguration{
			Port:            "9200",
			Protocol:        "https",
			SkipTlsVerify:   false,
			HttpDebug:       false,
		},
		LocalInstanceConfiguration: configurations.LocalInstanceConfiguration{
			RepoVersion: "v4.12.0",
		},
	}

	got, err := New()
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}

func TestLoadOptional(t *testing.T) {
	t.Run("no config returns nil nil", func(t *testing.T) {
		dir := t.TempDir()
		prev, _ := os.Getwd()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("chdir: %v", err)
		}
		defer func() { _ = os.Chdir(prev) }()

		got, err := LoadOptional()
		if err != nil {
			t.Errorf("LoadOptional() error = %v", err)
			return
		}
		if got != nil {
			t.Errorf("LoadOptional() = %v, want nil when no config", got)
		}
	})

	t.Run("with config returns parsed config", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, ".wazctl.yaml")
		content := `wazuh:
  endpoint: "10.77.60.11"
  port: "55000"
  protocol: https
local:
  repoVersion: "v4.11.0"
`
		if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
			t.Fatalf("write test config: %v", err)
		}
		prev, _ := os.Getwd()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("chdir: %v", err)
		}
		defer func() { _ = os.Chdir(prev) }()

		got, err := LoadOptional()
		if err != nil {
			t.Errorf("LoadOptional() error = %v", err)
			return
		}
		if got == nil {
			t.Fatal("LoadOptional() = nil, want config")
		}
		if got.WazuhInstanceConfigurations.Endpoint != "10.77.60.11" || got.LocalInstanceConfiguration.RepoVersion != "v4.11.0" {
			t.Errorf("LoadOptional() = %+v", got)
		}
	})
}

func Test_expandHomeDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandHomeDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandHomeDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("expandHomeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
