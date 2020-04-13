// Code generated by mockery v1.0.0. DO NOT EDIT.

package scanner

import context "context"
import mock "github.com/stretchr/testify/mock"
import types "github.com/aquasecurity/fanal/types"

// MockAnalyzer is an autogenerated mock type for the Analyzer type
type MockAnalyzer struct {
	mock.Mock
}

type AnalyzerAnalyzeArgs struct {
	Ctx         context.Context
	CtxAnything bool
}

type AnalyzerAnalyzeReturns struct {
	Info types.ImageReference
	Err  error
}

type AnalyzerAnalyzeExpectation struct {
	Args    AnalyzerAnalyzeArgs
	Returns AnalyzerAnalyzeReturns
}

func (_m *MockAnalyzer) ApplyAnalyzeExpectation(e AnalyzerAnalyzeExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	_m.On("Analyze", args...).Return(e.Returns.Info, e.Returns.Err)
}

func (_m *MockAnalyzer) ApplyAnalyzeExpectations(expectations []AnalyzerAnalyzeExpectation) {
	for _, e := range expectations {
		_m.ApplyAnalyzeExpectation(e)
	}
}

// Analyze provides a mock function with given fields: ctx
func (_m *MockAnalyzer) Analyze(ctx context.Context) (types.ImageReference, error) {
	ret := _m.Called(ctx)

	var r0 types.ImageReference
	if rf, ok := ret.Get(0).(func(context.Context) types.ImageReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.ImageReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}