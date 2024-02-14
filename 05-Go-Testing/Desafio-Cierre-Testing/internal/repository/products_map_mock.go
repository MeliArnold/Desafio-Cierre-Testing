package repository

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryProducts struct {
	mock.Mock
}

func (_m *MockRepositoryProducts) SearchProducts(_param0 internal.ProductQuery) (map[int]internal.Product, error) {
	ret := _m.Called(_param0)

	var r0 map[int]internal.Product
	if rf, ok := ret.Get(0).(func(internal.ProductQuery) map[int]internal.Product); ok {
		r0 = rf(_param0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]internal.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ProductQuery) error); ok {
		r1 = rf(_param0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
