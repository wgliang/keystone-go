/*
Copyright 2018 The keystone-go Authors.

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

package options

import (
	"github.com/spf13/pflag"
)

// ServerOption is the main context object for the controller manager.
type ServerOption struct {
	UserName           string
	Password           string
	TenantName         string
	TenantId           string
	RegionName         string
	IdentityApiVersion string
	Token              string
	Endpoint           string
	ForceNewToken      string
	Cacert             string
	Key                string
}

// NewServerOption creates a new CMServer with a default config.
func NewServerOption() *ServerOption {
	s := ServerOption{}
	return &s
}

// AddFlags adds flags for a specific CMServer to the specified FlagSet
func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.UserName, "os-username", s.UserName, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.TenantName, "os-tenant-name", s.TenantName, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.TenantId, "os-tenant-id", s.TenantId, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.RegionName, "os-region-name", s.RegionName, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.IdentityApiVersion, "os-identity-api-version", s.IdentityApiVersion, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Token, "os-token", s.Token, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Endpoint, "os-endpoint", s.Endpoint, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.ForceNewToken, "force-new-token", s.ForceNewToken, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Cacert, "os-cacert", s.Cacert, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Key, "os-key", s.Key, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
}

func (s *ServerOption) CheckOptionOrDie() {

}
