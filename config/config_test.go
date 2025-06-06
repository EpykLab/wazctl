package config

import (
	"reflect"
	"testing"

	"github.com/EpykLab/wazctl/models/configurations"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *configurations.WazuhCtlConfig
		wantErr bool
	}{
		{
			name: "get proper config",
			want: &configurations.WazuhCtlConfig{
				Endpoint:    "10.77.60.11",
				Port:        "55000",
				Protocol:    "https",
				WuiPassword: "wazuh-wui",
				WuiUsername: "wazuh-wui",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
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
