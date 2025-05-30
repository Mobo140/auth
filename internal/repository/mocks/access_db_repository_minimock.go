// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Mobo140/auth/internal/repository.AccessDBRepository -o access_db_repository_minimock.go -n AccessDBRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Mobo140/auth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// AccessDBRepositoryMock implements mm_repository.AccessDBRepository
type AccessDBRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetEndpointsAccess          func(ctx context.Context) (apa1 []*model.AccessEndpoint, err error)
	funcGetEndpointsAccessOrigin    string
	inspectFuncGetEndpointsAccess   func(ctx context.Context)
	afterGetEndpointsAccessCounter  uint64
	beforeGetEndpointsAccessCounter uint64
	GetEndpointsAccessMock          mAccessDBRepositoryMockGetEndpointsAccess
}

// NewAccessDBRepositoryMock returns a mock for mm_repository.AccessDBRepository
func NewAccessDBRepositoryMock(t minimock.Tester) *AccessDBRepositoryMock {
	m := &AccessDBRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetEndpointsAccessMock = mAccessDBRepositoryMockGetEndpointsAccess{mock: m}
	m.GetEndpointsAccessMock.callArgs = []*AccessDBRepositoryMockGetEndpointsAccessParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mAccessDBRepositoryMockGetEndpointsAccess struct {
	optional           bool
	mock               *AccessDBRepositoryMock
	defaultExpectation *AccessDBRepositoryMockGetEndpointsAccessExpectation
	expectations       []*AccessDBRepositoryMockGetEndpointsAccessExpectation

	callArgs []*AccessDBRepositoryMockGetEndpointsAccessParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// AccessDBRepositoryMockGetEndpointsAccessExpectation specifies expectation struct of the AccessDBRepository.GetEndpointsAccess
type AccessDBRepositoryMockGetEndpointsAccessExpectation struct {
	mock               *AccessDBRepositoryMock
	params             *AccessDBRepositoryMockGetEndpointsAccessParams
	paramPtrs          *AccessDBRepositoryMockGetEndpointsAccessParamPtrs
	expectationOrigins AccessDBRepositoryMockGetEndpointsAccessExpectationOrigins
	results            *AccessDBRepositoryMockGetEndpointsAccessResults
	returnOrigin       string
	Counter            uint64
}

// AccessDBRepositoryMockGetEndpointsAccessParams contains parameters of the AccessDBRepository.GetEndpointsAccess
type AccessDBRepositoryMockGetEndpointsAccessParams struct {
	ctx context.Context
}

// AccessDBRepositoryMockGetEndpointsAccessParamPtrs contains pointers to parameters of the AccessDBRepository.GetEndpointsAccess
type AccessDBRepositoryMockGetEndpointsAccessParamPtrs struct {
	ctx *context.Context
}

// AccessDBRepositoryMockGetEndpointsAccessResults contains results of the AccessDBRepository.GetEndpointsAccess
type AccessDBRepositoryMockGetEndpointsAccessResults struct {
	apa1 []*model.AccessEndpoint
	err  error
}

// AccessDBRepositoryMockGetEndpointsAccessOrigins contains origins of expectations of the AccessDBRepository.GetEndpointsAccess
type AccessDBRepositoryMockGetEndpointsAccessExpectationOrigins struct {
	origin    string
	originCtx string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Optional() *mAccessDBRepositoryMockGetEndpointsAccess {
	mmGetEndpointsAccess.optional = true
	return mmGetEndpointsAccess
}

// Expect sets up expected params for AccessDBRepository.GetEndpointsAccess
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Expect(ctx context.Context) *mAccessDBRepositoryMockGetEndpointsAccess {
	if mmGetEndpointsAccess.mock.funcGetEndpointsAccess != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by Set")
	}

	if mmGetEndpointsAccess.defaultExpectation == nil {
		mmGetEndpointsAccess.defaultExpectation = &AccessDBRepositoryMockGetEndpointsAccessExpectation{}
	}

	if mmGetEndpointsAccess.defaultExpectation.paramPtrs != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by ExpectParams functions")
	}

	mmGetEndpointsAccess.defaultExpectation.params = &AccessDBRepositoryMockGetEndpointsAccessParams{ctx}
	mmGetEndpointsAccess.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetEndpointsAccess.expectations {
		if minimock.Equal(e.params, mmGetEndpointsAccess.defaultExpectation.params) {
			mmGetEndpointsAccess.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetEndpointsAccess.defaultExpectation.params)
		}
	}

	return mmGetEndpointsAccess
}

// ExpectCtxParam1 sets up expected param ctx for AccessDBRepository.GetEndpointsAccess
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) ExpectCtxParam1(ctx context.Context) *mAccessDBRepositoryMockGetEndpointsAccess {
	if mmGetEndpointsAccess.mock.funcGetEndpointsAccess != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by Set")
	}

	if mmGetEndpointsAccess.defaultExpectation == nil {
		mmGetEndpointsAccess.defaultExpectation = &AccessDBRepositoryMockGetEndpointsAccessExpectation{}
	}

	if mmGetEndpointsAccess.defaultExpectation.params != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by Expect")
	}

	if mmGetEndpointsAccess.defaultExpectation.paramPtrs == nil {
		mmGetEndpointsAccess.defaultExpectation.paramPtrs = &AccessDBRepositoryMockGetEndpointsAccessParamPtrs{}
	}
	mmGetEndpointsAccess.defaultExpectation.paramPtrs.ctx = &ctx
	mmGetEndpointsAccess.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGetEndpointsAccess
}

// Inspect accepts an inspector function that has same arguments as the AccessDBRepository.GetEndpointsAccess
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Inspect(f func(ctx context.Context)) *mAccessDBRepositoryMockGetEndpointsAccess {
	if mmGetEndpointsAccess.mock.inspectFuncGetEndpointsAccess != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("Inspect function is already set for AccessDBRepositoryMock.GetEndpointsAccess")
	}

	mmGetEndpointsAccess.mock.inspectFuncGetEndpointsAccess = f

	return mmGetEndpointsAccess
}

// Return sets up results that will be returned by AccessDBRepository.GetEndpointsAccess
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Return(apa1 []*model.AccessEndpoint, err error) *AccessDBRepositoryMock {
	if mmGetEndpointsAccess.mock.funcGetEndpointsAccess != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by Set")
	}

	if mmGetEndpointsAccess.defaultExpectation == nil {
		mmGetEndpointsAccess.defaultExpectation = &AccessDBRepositoryMockGetEndpointsAccessExpectation{mock: mmGetEndpointsAccess.mock}
	}
	mmGetEndpointsAccess.defaultExpectation.results = &AccessDBRepositoryMockGetEndpointsAccessResults{apa1, err}
	mmGetEndpointsAccess.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetEndpointsAccess.mock
}

// Set uses given function f to mock the AccessDBRepository.GetEndpointsAccess method
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Set(f func(ctx context.Context) (apa1 []*model.AccessEndpoint, err error)) *AccessDBRepositoryMock {
	if mmGetEndpointsAccess.defaultExpectation != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("Default expectation is already set for the AccessDBRepository.GetEndpointsAccess method")
	}

	if len(mmGetEndpointsAccess.expectations) > 0 {
		mmGetEndpointsAccess.mock.t.Fatalf("Some expectations are already set for the AccessDBRepository.GetEndpointsAccess method")
	}

	mmGetEndpointsAccess.mock.funcGetEndpointsAccess = f
	mmGetEndpointsAccess.mock.funcGetEndpointsAccessOrigin = minimock.CallerInfo(1)
	return mmGetEndpointsAccess.mock
}

// When sets expectation for the AccessDBRepository.GetEndpointsAccess which will trigger the result defined by the following
// Then helper
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) When(ctx context.Context) *AccessDBRepositoryMockGetEndpointsAccessExpectation {
	if mmGetEndpointsAccess.mock.funcGetEndpointsAccess != nil {
		mmGetEndpointsAccess.mock.t.Fatalf("AccessDBRepositoryMock.GetEndpointsAccess mock is already set by Set")
	}

	expectation := &AccessDBRepositoryMockGetEndpointsAccessExpectation{
		mock:               mmGetEndpointsAccess.mock,
		params:             &AccessDBRepositoryMockGetEndpointsAccessParams{ctx},
		expectationOrigins: AccessDBRepositoryMockGetEndpointsAccessExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetEndpointsAccess.expectations = append(mmGetEndpointsAccess.expectations, expectation)
	return expectation
}

// Then sets up AccessDBRepository.GetEndpointsAccess return parameters for the expectation previously defined by the When method
func (e *AccessDBRepositoryMockGetEndpointsAccessExpectation) Then(apa1 []*model.AccessEndpoint, err error) *AccessDBRepositoryMock {
	e.results = &AccessDBRepositoryMockGetEndpointsAccessResults{apa1, err}
	return e.mock
}

// Times sets number of times AccessDBRepository.GetEndpointsAccess should be invoked
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Times(n uint64) *mAccessDBRepositoryMockGetEndpointsAccess {
	if n == 0 {
		mmGetEndpointsAccess.mock.t.Fatalf("Times of AccessDBRepositoryMock.GetEndpointsAccess mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetEndpointsAccess.expectedInvocations, n)
	mmGetEndpointsAccess.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetEndpointsAccess
}

func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) invocationsDone() bool {
	if len(mmGetEndpointsAccess.expectations) == 0 && mmGetEndpointsAccess.defaultExpectation == nil && mmGetEndpointsAccess.mock.funcGetEndpointsAccess == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetEndpointsAccess.mock.afterGetEndpointsAccessCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetEndpointsAccess.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetEndpointsAccess implements mm_repository.AccessDBRepository
func (mmGetEndpointsAccess *AccessDBRepositoryMock) GetEndpointsAccess(ctx context.Context) (apa1 []*model.AccessEndpoint, err error) {
	mm_atomic.AddUint64(&mmGetEndpointsAccess.beforeGetEndpointsAccessCounter, 1)
	defer mm_atomic.AddUint64(&mmGetEndpointsAccess.afterGetEndpointsAccessCounter, 1)

	mmGetEndpointsAccess.t.Helper()

	if mmGetEndpointsAccess.inspectFuncGetEndpointsAccess != nil {
		mmGetEndpointsAccess.inspectFuncGetEndpointsAccess(ctx)
	}

	mm_params := AccessDBRepositoryMockGetEndpointsAccessParams{ctx}

	// Record call args
	mmGetEndpointsAccess.GetEndpointsAccessMock.mutex.Lock()
	mmGetEndpointsAccess.GetEndpointsAccessMock.callArgs = append(mmGetEndpointsAccess.GetEndpointsAccessMock.callArgs, &mm_params)
	mmGetEndpointsAccess.GetEndpointsAccessMock.mutex.Unlock()

	for _, e := range mmGetEndpointsAccess.GetEndpointsAccessMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.apa1, e.results.err
		}
	}

	if mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.Counter, 1)
		mm_want := mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.params
		mm_want_ptrs := mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.paramPtrs

		mm_got := AccessDBRepositoryMockGetEndpointsAccessParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetEndpointsAccess.t.Errorf("AccessDBRepositoryMock.GetEndpointsAccess got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetEndpointsAccess.t.Errorf("AccessDBRepositoryMock.GetEndpointsAccess got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetEndpointsAccess.GetEndpointsAccessMock.defaultExpectation.results
		if mm_results == nil {
			mmGetEndpointsAccess.t.Fatal("No results are set for the AccessDBRepositoryMock.GetEndpointsAccess")
		}
		return (*mm_results).apa1, (*mm_results).err
	}
	if mmGetEndpointsAccess.funcGetEndpointsAccess != nil {
		return mmGetEndpointsAccess.funcGetEndpointsAccess(ctx)
	}
	mmGetEndpointsAccess.t.Fatalf("Unexpected call to AccessDBRepositoryMock.GetEndpointsAccess. %v", ctx)
	return
}

// GetEndpointsAccessAfterCounter returns a count of finished AccessDBRepositoryMock.GetEndpointsAccess invocations
func (mmGetEndpointsAccess *AccessDBRepositoryMock) GetEndpointsAccessAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetEndpointsAccess.afterGetEndpointsAccessCounter)
}

// GetEndpointsAccessBeforeCounter returns a count of AccessDBRepositoryMock.GetEndpointsAccess invocations
func (mmGetEndpointsAccess *AccessDBRepositoryMock) GetEndpointsAccessBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetEndpointsAccess.beforeGetEndpointsAccessCounter)
}

// Calls returns a list of arguments used in each call to AccessDBRepositoryMock.GetEndpointsAccess.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetEndpointsAccess *mAccessDBRepositoryMockGetEndpointsAccess) Calls() []*AccessDBRepositoryMockGetEndpointsAccessParams {
	mmGetEndpointsAccess.mutex.RLock()

	argCopy := make([]*AccessDBRepositoryMockGetEndpointsAccessParams, len(mmGetEndpointsAccess.callArgs))
	copy(argCopy, mmGetEndpointsAccess.callArgs)

	mmGetEndpointsAccess.mutex.RUnlock()

	return argCopy
}

// MinimockGetEndpointsAccessDone returns true if the count of the GetEndpointsAccess invocations corresponds
// the number of defined expectations
func (m *AccessDBRepositoryMock) MinimockGetEndpointsAccessDone() bool {
	if m.GetEndpointsAccessMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetEndpointsAccessMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetEndpointsAccessMock.invocationsDone()
}

// MinimockGetEndpointsAccessInspect logs each unmet expectation
func (m *AccessDBRepositoryMock) MinimockGetEndpointsAccessInspect() {
	for _, e := range m.GetEndpointsAccessMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AccessDBRepositoryMock.GetEndpointsAccess at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetEndpointsAccessCounter := mm_atomic.LoadUint64(&m.afterGetEndpointsAccessCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetEndpointsAccessMock.defaultExpectation != nil && afterGetEndpointsAccessCounter < 1 {
		if m.GetEndpointsAccessMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to AccessDBRepositoryMock.GetEndpointsAccess at\n%s", m.GetEndpointsAccessMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to AccessDBRepositoryMock.GetEndpointsAccess at\n%s with params: %#v", m.GetEndpointsAccessMock.defaultExpectation.expectationOrigins.origin, *m.GetEndpointsAccessMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetEndpointsAccess != nil && afterGetEndpointsAccessCounter < 1 {
		m.t.Errorf("Expected call to AccessDBRepositoryMock.GetEndpointsAccess at\n%s", m.funcGetEndpointsAccessOrigin)
	}

	if !m.GetEndpointsAccessMock.invocationsDone() && afterGetEndpointsAccessCounter > 0 {
		m.t.Errorf("Expected %d calls to AccessDBRepositoryMock.GetEndpointsAccess at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetEndpointsAccessMock.expectedInvocations), m.GetEndpointsAccessMock.expectedInvocationsOrigin, afterGetEndpointsAccessCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AccessDBRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetEndpointsAccessInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AccessDBRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *AccessDBRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetEndpointsAccessDone()
}
