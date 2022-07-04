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

// Code generated by lister-gen. DO NOT EDIT.

package v2beta1

import (
	v2beta1 "github.com/ovrclk/provider-services/pkg/apis/akash.network/v2beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InventoryRequestLister helps list InventoryRequests.
// All objects returned here must be treated as read-only.
type InventoryRequestLister interface {
	// List lists all InventoryRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta1.InventoryRequest, err error)
	// Get retrieves the InventoryRequest from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta1.InventoryRequest, error)
	InventoryRequestListerExpansion
}

// inventoryRequestLister implements the InventoryRequestLister interface.
type inventoryRequestLister struct {
	indexer cache.Indexer
}

// NewInventoryRequestLister returns a new InventoryRequestLister.
func NewInventoryRequestLister(indexer cache.Indexer) InventoryRequestLister {
	return &inventoryRequestLister{indexer: indexer}
}

// List lists all InventoryRequests in the indexer.
func (s *inventoryRequestLister) List(selector labels.Selector) (ret []*v2beta1.InventoryRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta1.InventoryRequest))
	})
	return ret, err
}

// Get retrieves the InventoryRequest from the index for a given name.
func (s *inventoryRequestLister) Get(name string) (*v2beta1.InventoryRequest, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta1.Resource("inventoryrequest"), name)
	}
	return obj.(*v2beta1.InventoryRequest), nil
}
