// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// VolumeAttachmentsGetter is an autogenerated mock type for the VolumeAttachmentsGetter type
type VolumeAttachmentsGetter struct {
	mock.Mock
}

// VolumeAttachments provides a mock function with given fields:
func (_m *VolumeAttachmentsGetter) VolumeAttachments() v1.VolumeAttachmentInterface {
	ret := _m.Called()

	var r0 v1.VolumeAttachmentInterface
	if rf, ok := ret.Get(0).(func() v1.VolumeAttachmentInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.VolumeAttachmentInterface)
		}
	}

	return r0
}

// NewVolumeAttachmentsGetter creates a new instance of VolumeAttachmentsGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewVolumeAttachmentsGetter(t testing.TB) *VolumeAttachmentsGetter {
	mock := &VolumeAttachmentsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
