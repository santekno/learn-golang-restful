// Code generated by mockery v2.42.1. DO NOT EDIT.

package usecase

import (
	context "context"

	models "github.com/santekno/learn-golang-restful/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockArticleUsecase is an autogenerated mock type for the ArticleUsecase type
type MockArticleUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockArticleUsecase) Delete(ctx context.Context, id int64) (bool, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (bool, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockArticleUsecase) GetAll(ctx context.Context) ([]models.ArticleResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.ArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.ArticleResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.ArticleResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ArticleResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockArticleUsecase) GetByID(ctx context.Context, id int64) (models.ArticleResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.ArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (models.ArticleResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.ArticleResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.ArticleResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, article
func (_m *MockArticleUsecase) Store(ctx context.Context, article models.ArticleCreateRequest) (models.ArticleResponse, error) {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 models.ArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ArticleCreateRequest) (models.ArticleResponse, error)); ok {
		return rf(ctx, article)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ArticleCreateRequest) models.ArticleResponse); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Get(0).(models.ArticleResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ArticleCreateRequest) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, article
func (_m *MockArticleUsecase) Update(ctx context.Context, article models.ArticleUpdateRequest) (models.ArticleResponse, error) {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.ArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ArticleUpdateRequest) (models.ArticleResponse, error)); ok {
		return rf(ctx, article)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ArticleUpdateRequest) models.ArticleResponse); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Get(0).(models.ArticleResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ArticleUpdateRequest) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockArticleUsecase creates a new instance of MockArticleUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockArticleUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockArticleUsecase {
	mock := &MockArticleUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
