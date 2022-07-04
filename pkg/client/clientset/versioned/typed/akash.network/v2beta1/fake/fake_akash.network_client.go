/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2beta1 "github.com/ovrclk/provider-services/pkg/client/clientset/versioned/typed/akash.network/v2beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAkashV2beta1 struct {
	*testing.Fake
}

func (c *FakeAkashV2beta1) Inventories() v2beta1.InventoryInterface {
	return &FakeInventories{c}
}

func (c *FakeAkashV2beta1) InventoryRequests() v2beta1.InventoryRequestInterface {
	return &FakeInventoryRequests{c}
}

func (c *FakeAkashV2beta1) Manifests(namespace string) v2beta1.ManifestInterface {
	return &FakeManifests{c, namespace}
}

func (c *FakeAkashV2beta1) ProviderHosts(namespace string) v2beta1.ProviderHostInterface {
	return &FakeProviderHosts{c, namespace}
}

func (c *FakeAkashV2beta1) ProviderLeasedIPs(namespace string) v2beta1.ProviderLeasedIPInterface {
	return &FakeProviderLeasedIPs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAkashV2beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
