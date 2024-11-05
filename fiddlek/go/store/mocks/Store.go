// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	store "go.skia.org/infra/fiddlek/go/store"

	types "go.skia.org/infra/fiddlek/go/types"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// DeleteName provides a mock function with given fields: name
func (_m *Store) DeleteName(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: hash
func (_m *Store) Exists(hash string) error {
	ret := _m.Called(hash)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCode provides a mock function with given fields: fiddleHash
func (_m *Store) GetCode(fiddleHash string) (string, *types.Options, error) {
	ret := _m.Called(fiddleHash)

	if len(ret) == 0 {
		panic("no return value specified for GetCode")
	}

	var r0 string
	var r1 *types.Options
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (string, *types.Options, error)); ok {
		return rf(fiddleHash)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(fiddleHash)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) *types.Options); ok {
		r1 = rf(fiddleHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Options)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(fiddleHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetHashFromName provides a mock function with given fields: name
func (_m *Store) GetHashFromName(name string) (string, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetHashFromName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMedia provides a mock function with given fields: fiddleHash, media
func (_m *Store) GetMedia(fiddleHash string, media store.Media) ([]byte, string, string, error) {
	ret := _m.Called(fiddleHash, media)

	if len(ret) == 0 {
		panic("no return value specified for GetMedia")
	}

	var r0 []byte
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(string, store.Media) ([]byte, string, string, error)); ok {
		return rf(fiddleHash, media)
	}
	if rf, ok := ret.Get(0).(func(string, store.Media) []byte); ok {
		r0 = rf(fiddleHash, media)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, store.Media) string); ok {
		r1 = rf(fiddleHash, media)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, store.Media) string); ok {
		r2 = rf(fiddleHash, media)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(string, store.Media) error); ok {
		r3 = rf(fiddleHash, media)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// ListAllNames provides a mock function with given fields:
func (_m *Store) ListAllNames() ([]store.Named, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAllNames")
	}

	var r0 []store.Named
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]store.Named, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []store.Named); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]store.Named)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: code, options, results
func (_m *Store) Put(code string, options types.Options, results *types.Result) (string, error) {
	ret := _m.Called(code, options, results)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, types.Options, *types.Result) (string, error)); ok {
		return rf(code, options, results)
	}
	if rf, ok := ret.Get(0).(func(string, types.Options, *types.Result) string); ok {
		r0 = rf(code, options, results)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, types.Options, *types.Result) error); ok {
		r1 = rf(code, options, results)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMedia provides a mock function with given fields: options, fiddleHash, results
func (_m *Store) PutMedia(options types.Options, fiddleHash string, results *types.Result) error {
	ret := _m.Called(options, fiddleHash, results)

	if len(ret) == 0 {
		panic("no return value specified for PutMedia")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Options, string, *types.Result) error); ok {
		r0 = rf(options, fiddleHash, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStatus provides a mock function with given fields: name, status
func (_m *Store) SetStatus(name string, status string) error {
	ret := _m.Called(name, status)

	if len(ret) == 0 {
		panic("no return value specified for SetStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidName provides a mock function with given fields: name
func (_m *Store) ValidName(name string) bool {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for ValidName")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// WriteName provides a mock function with given fields: name, hash, user, status
func (_m *Store) WriteName(name string, hash string, user string, status string) error {
	ret := _m.Called(name, hash, user, status)

	if len(ret) == 0 {
		panic("no return value specified for WriteName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(name, hash, user, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
