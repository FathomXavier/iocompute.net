// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

// CSINodesGetter is an autogenerated mock type for the CSINodesGetter type
type CSINodesGetter struct {
	mock.Mock
}

// CSINodes provides a mock function with given fields:
func (_m *CSINodesGetter) CSINodes() v1beta1.CSINodeInterface {
	ret := _m.Called()

	var r0 v1beta1.CSINodeInterface
	if rf, ok := ret.Get(0).(func() v1beta1.CSINodeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.CSINodeInterface)
		}
	}

	return r0
}

// NewCSINodesGetter creates a new instance of CSINodesGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSINodesGetter(t testing.TB) *CSINodesGetter {
	mock := &CSINodesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
