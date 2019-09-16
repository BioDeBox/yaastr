// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ast "github.com/biodebox/yaastr/ast"
import mock "github.com/stretchr/testify/mock"

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// GetParent provides a mock function with given fields:
func (_m *Node) GetParent() ast.ParentNode {
	ret := _m.Called()

	var r0 ast.ParentNode
	if rf, ok := ret.Get(0).(func() ast.ParentNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ast.ParentNode)
		}
	}

	return r0
}

// SetParent provides a mock function with given fields: node
func (_m *Node) SetParent(node ast.ParentNode) {
	_m.Called(node)
}
