/*
Copyright 2021 The Fission Authors.

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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCacheKey(t *testing.T) {
	type args struct {
		metadata *metav1.ObjectMeta
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CacheKey",
			args: args{metadata: &metav1.ObjectMeta{UID: "uid", ResourceVersion: "v1"}},
			want: "uid_v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CacheKey(tt.args.metadata); got != tt.want {
				t.Errorf("CacheKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
