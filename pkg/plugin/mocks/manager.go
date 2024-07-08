// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	restoreitemactionv1 "github.com/vmware-tanzu/velero/pkg/plugin/velero/restoreitemaction/v1"

	restoreitemactionv2 "github.com/vmware-tanzu/velero/pkg/plugin/velero/restoreitemaction/v2"

	v1 "github.com/vmware-tanzu/velero/pkg/plugin/velero/backupitemaction/v1"

	v2 "github.com/vmware-tanzu/velero/pkg/plugin/velero/backupitemaction/v2"

	velero "github.com/vmware-tanzu/velero/pkg/plugin/velero"

	volumesnapshotterv1 "github.com/vmware-tanzu/velero/pkg/plugin/velero/volumesnapshotter/v1"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CleanupClients provides a mock function with given fields:
func (_m *Manager) CleanupClients() {
	_m.Called()
}

// GetBackupItemAction provides a mock function with given fields: name
func (_m *Manager) GetBackupItemAction(name string) (v1.BackupItemAction, error) {
	ret := _m.Called(name)

	var r0 v1.BackupItemAction
	if rf, ok := ret.Get(0).(func(string) v1.BackupItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupItemActionV2 provides a mock function with given fields: name
func (_m *Manager) GetBackupItemActionV2(name string) (v2.BackupItemAction, error) {
	ret := _m.Called(name)

	var r0 v2.BackupItemAction
	if rf, ok := ret.Get(0).(func(string) v2.BackupItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupItemActions provides a mock function with given fields:
func (_m *Manager) GetBackupItemActions() ([]v1.BackupItemAction, error) {
	ret := _m.Called()

	var r0 []v1.BackupItemAction
	if rf, ok := ret.Get(0).(func() []v1.BackupItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupItemActionsV2 provides a mock function with given fields:
func (_m *Manager) GetBackupItemActionsV2() ([]v2.BackupItemAction, error) {
	ret := _m.Called()

	var r0 []v2.BackupItemAction
	if rf, ok := ret.Get(0).(func() []v2.BackupItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v2.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeleteItemAction provides a mock function with given fields: name
func (_m *Manager) GetDeleteItemAction(name string) (velero.DeleteItemAction, error) {
	ret := _m.Called(name)

	var r0 velero.DeleteItemAction
	if rf, ok := ret.Get(0).(func(string) velero.DeleteItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.DeleteItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeleteItemActions provides a mock function with given fields:
func (_m *Manager) GetDeleteItemActions() ([]velero.DeleteItemAction, error) {
	ret := _m.Called()

	var r0 []velero.DeleteItemAction
	if rf, ok := ret.Get(0).(func() []velero.DeleteItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]velero.DeleteItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectStore provides a mock function with given fields: name
func (_m *Manager) GetObjectStore(name string) (velero.ObjectStore, error) {
	ret := _m.Called(name)

	var r0 velero.ObjectStore
	if rf, ok := ret.Get(0).(func(string) velero.ObjectStore); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.ObjectStore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemAction provides a mock function with given fields: name
func (_m *Manager) GetRestoreItemAction(name string) (restoreitemactionv1.RestoreItemAction, error) {
	ret := _m.Called(name)

	var r0 restoreitemactionv1.RestoreItemAction
	if rf, ok := ret.Get(0).(func(string) restoreitemactionv1.RestoreItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restoreitemactionv1.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemActionV2 provides a mock function with given fields: name
func (_m *Manager) GetRestoreItemActionV2(name string) (restoreitemactionv2.RestoreItemAction, error) {
	ret := _m.Called(name)

	var r0 restoreitemactionv2.RestoreItemAction
	if rf, ok := ret.Get(0).(func(string) restoreitemactionv2.RestoreItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restoreitemactionv2.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemActions provides a mock function with given fields:
func (_m *Manager) GetRestoreItemActions() ([]restoreitemactionv1.RestoreItemAction, error) {
	ret := _m.Called()

	var r0 []restoreitemactionv1.RestoreItemAction
	if rf, ok := ret.Get(0).(func() []restoreitemactionv1.RestoreItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]restoreitemactionv1.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemActionsV2 provides a mock function with given fields:
func (_m *Manager) GetRestoreItemActionsV2() ([]restoreitemactionv2.RestoreItemAction, error) {
	ret := _m.Called()

	var r0 []restoreitemactionv2.RestoreItemAction
	if rf, ok := ret.Get(0).(func() []restoreitemactionv2.RestoreItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]restoreitemactionv2.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolumeSnapshotter provides a mock function with given fields: name
func (_m *Manager) GetVolumeSnapshotter(name string) (volumesnapshotterv1.VolumeSnapshotter, error) {
	ret := _m.Called(name)

	var r0 volumesnapshotterv1.VolumeSnapshotter
	if rf, ok := ret.Get(0).(func(string) volumesnapshotterv1.VolumeSnapshotter); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(volumesnapshotterv1.VolumeSnapshotter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}