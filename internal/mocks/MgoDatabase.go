// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mgo "gopkg.in/mgo.v2"
import mock "github.com/stretchr/testify/mock"

// MgoDatabase is an autogenerated mock type for the MgoDatabase type
type MgoDatabase struct {
	mock.Mock
}

// C provides a mock function with given fields: _a0
func (_m *MgoDatabase) C(_a0 string) *mgo.Collection {
	ret := _m.Called(_a0)

	var r0 *mgo.Collection
	if rf, ok := ret.Get(0).(func(string) *mgo.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mgo.Collection)
		}
	}

	return r0
}
