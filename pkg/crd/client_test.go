/*
Copyright 2016 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crd

import (
	"testing"

	"k8s.io/client-go/dynamic"
)

func TestGetDynamicClient(t *testing.T) {
	tests := []struct {
		name    string
		want    dynamic.Interface
		wantErr bool
	}{
		{
			name:    "TestGetDynamicClient",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDynamicClient()
			if err != nil {
				t.Errorf("GetDynamicClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetDynamicClient() client is nil")
				return
			}
		})
	}
}

func TestGetKubeconfig(t *testing.T) {
	tests := []struct {
		name           string
		wantKubeconfig string
		wantErr        bool
	}{
		{
			name:           "TestGetKubeconfig with env",
			wantKubeconfig: "~/.kube/config",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKubeconfig, err := GetKubeconfig()
			if err != nil {
				t.Errorf("GetKubeconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotKubeconfig) == 0 {
				t.Errorf("GetKubeconfig() gotKubeconfig = %v, want %v", gotKubeconfig, tt.wantKubeconfig)
			}
		})
	}
}
