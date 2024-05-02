// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package example

import (
	"context"
	"sync"
)

// Ensure, that PersonStoreMock does implement PersonStore.
// If this is not the case, regenerate this file with moq.
var _ PersonStore = &PersonStoreMock{}

// PersonStoreMock is a mock implementation of PersonStore.
//
//	func TestSomethingThatUsesPersonStore(t *testing.T) {
//
//		// make and configure a mocked PersonStore
//		mockedPersonStore := &PersonStoreMock{
//			ClearCacheFunc: func(id string)  {
//				panic("mock out the ClearCache method")
//			},
//			CreateFunc: func(ctx context.Context, person *Person, confirm bool) error {
//				panic("mock out the Create method")
//			},
//			GetFunc: func(ctx context.Context, id string) (*Person, error) {
//				panic("mock out the Get method")
//			},
//		}
//
//		// use mockedPersonStore in code that requires PersonStore
//		// and then make assertions.
//
//	}

type PersonStoreMock struct {
	// ClearCacheFunc mocks the ClearCache method.
	ClearCacheFunc func(id string)

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, person *Person, confirm bool) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id string) (*Person, error)

	// calls tracks calls to the methods.
	calls struct {
		// ClearCache holds details about calls to the ClearCache method.
		ClearCache []PersonStoreMockClearCacheCalls
		// Create holds details about calls to the Create method.
		Create []PersonStoreMockCreateCalls
		// Get holds details about calls to the Get method.
		Get []PersonStoreMockGetCalls
	}
	lockClearCache sync.RWMutex
	lockCreate     sync.RWMutex
	lockGet        sync.RWMutex
}

// PersonStoreMockClearCacheCalls holds details about calls to the ClearCache method.
type PersonStoreMockClearCacheCalls struct {
	// ID is the id argument value.
	ID string
}

// PersonStoreMockCreateCalls holds details about calls to the Create method.
type PersonStoreMockCreateCalls struct {
	// Ctx is the ctx argument value.
	Ctx context.Context
	// Person is the person argument value.
	Person *Person
	// Confirm is the confirm argument value.
	Confirm bool
}

// PersonStoreMockGetCalls holds details about calls to the Get method.
type PersonStoreMockGetCalls struct {
	// Ctx is the ctx argument value.
	Ctx context.Context
	// ID is the id argument value.
	ID string
}

// ClearCache calls ClearCacheFunc.
func (mock *PersonStoreMock) ClearCache(id string) {
	if mock.ClearCacheFunc == nil {
		panic("PersonStoreMock.ClearCacheFunc: method is nil but PersonStore.ClearCache was just called")
	}
	callInfo := PersonStoreMockClearCacheCalls{
		ID: id,
	}
	mock.lockClearCache.Lock()
	mock.calls.ClearCache = append(mock.calls.ClearCache, callInfo)
	mock.lockClearCache.Unlock()
	mock.ClearCacheFunc(id)
}

// ClearCacheCalls gets all the calls that were made to ClearCache.
// Check the length with:
//
//	len(mockedPersonStore.ClearCacheCalls())
func (mock *PersonStoreMock) ClearCacheCalls() []PersonStoreMockClearCacheCalls {
	var calls []PersonStoreMockClearCacheCalls
	mock.lockClearCache.RLock()
	calls = mock.calls.ClearCache
	mock.lockClearCache.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error {
	if mock.CreateFunc == nil {
		panic("PersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called")
	}
	callInfo := PersonStoreMockCreateCalls{
		Ctx:     ctx,
		Person:  person,
		Confirm: confirm,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, person, confirm)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedPersonStore.CreateCalls())
func (mock *PersonStoreMock) CreateCalls() []PersonStoreMockCreateCalls {
	var calls []PersonStoreMockCreateCalls
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error) {
	if mock.GetFunc == nil {
		panic("PersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called")
	}
	callInfo := PersonStoreMockGetCalls{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedPersonStore.GetCalls())
func (mock *PersonStoreMock) GetCalls() []PersonStoreMockGetCalls {
	var calls []PersonStoreMockGetCalls
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}
