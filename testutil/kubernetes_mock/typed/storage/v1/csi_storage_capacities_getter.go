// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// CSIStorageCapacitiesGetter is an autogenerated mock type for the CSIStorageCapacitiesGetter type
type CSIStorageCapacitiesGetter struct {
	mock.Mock
}

// CSIStorageCapacities provides a mock function with given fields: namespace
func (_m *CSIStorageCapacitiesGetter) CSIStorageCapacities(namespace string) v1.CSIStorageCapacityInterface {
	ret := _m.Called(namespace)

	var r0 v1.CSIStorageCapacityInterface
	if rf, ok := ret.Get(0).(func(string) v1.CSIStorageCapacityInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.CSIStorageCapacityInterface)
		}
	}

	return r0
}

// NewCSIStorageCapacitiesGetter creates a new instance of CSIStorageCapacitiesGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSIStorageCapacitiesGetter(t testing.TB) *CSIStorageCapacitiesGetter {
	mock := &CSIStorageCapacitiesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
