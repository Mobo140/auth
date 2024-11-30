// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Mobo140/auth/internal/service.AccessService -o access_service_minimock.go -n AccessServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// AccessServiceMock implements mm_service.AccessService
type AccessServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCheck          func(ctx context.Context, accessToken string, endpoint string) (err error)
	funcCheckOrigin    string
	inspectFuncCheck   func(ctx context.Context, accessToken string, endpoint string)
	afterCheckCounter  uint64
	beforeCheckCounter uint64
	CheckMock          mAccessServiceMockCheck
}

// NewAccessServiceMock returns a mock for mm_service.AccessService
func NewAccessServiceMock(t minimock.Tester) *AccessServiceMock {
	m := &AccessServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckMock = mAccessServiceMockCheck{mock: m}
	m.CheckMock.callArgs = []*AccessServiceMockCheckParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mAccessServiceMockCheck struct {
	optional           bool
	mock               *AccessServiceMock
	defaultExpectation *AccessServiceMockCheckExpectation
	expectations       []*AccessServiceMockCheckExpectation

	callArgs []*AccessServiceMockCheckParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// AccessServiceMockCheckExpectation specifies expectation struct of the AccessService.Check
type AccessServiceMockCheckExpectation struct {
	mock               *AccessServiceMock
	params             *AccessServiceMockCheckParams
	paramPtrs          *AccessServiceMockCheckParamPtrs
	expectationOrigins AccessServiceMockCheckExpectationOrigins
	results            *AccessServiceMockCheckResults
	returnOrigin       string
	Counter            uint64
}

// AccessServiceMockCheckParams contains parameters of the AccessService.Check
type AccessServiceMockCheckParams struct {
	ctx         context.Context
	accessToken string
	endpoint    string
}

// AccessServiceMockCheckParamPtrs contains pointers to parameters of the AccessService.Check
type AccessServiceMockCheckParamPtrs struct {
	ctx         *context.Context
	accessToken *string
	endpoint    *string
}

// AccessServiceMockCheckResults contains results of the AccessService.Check
type AccessServiceMockCheckResults struct {
	err error
}

// AccessServiceMockCheckOrigins contains origins of expectations of the AccessService.Check
type AccessServiceMockCheckExpectationOrigins struct {
	origin            string
	originCtx         string
	originAccessToken string
	originEndpoint    string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCheck *mAccessServiceMockCheck) Optional() *mAccessServiceMockCheck {
	mmCheck.optional = true
	return mmCheck
}

// Expect sets up expected params for AccessService.Check
func (mmCheck *mAccessServiceMockCheck) Expect(ctx context.Context, accessToken string, endpoint string) *mAccessServiceMockCheck {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &AccessServiceMockCheckExpectation{}
	}

	if mmCheck.defaultExpectation.paramPtrs != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by ExpectParams functions")
	}

	mmCheck.defaultExpectation.params = &AccessServiceMockCheckParams{ctx, accessToken, endpoint}
	mmCheck.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCheck.expectations {
		if minimock.Equal(e.params, mmCheck.defaultExpectation.params) {
			mmCheck.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheck.defaultExpectation.params)
		}
	}

	return mmCheck
}

// ExpectCtxParam1 sets up expected param ctx for AccessService.Check
func (mmCheck *mAccessServiceMockCheck) ExpectCtxParam1(ctx context.Context) *mAccessServiceMockCheck {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &AccessServiceMockCheckExpectation{}
	}

	if mmCheck.defaultExpectation.params != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Expect")
	}

	if mmCheck.defaultExpectation.paramPtrs == nil {
		mmCheck.defaultExpectation.paramPtrs = &AccessServiceMockCheckParamPtrs{}
	}
	mmCheck.defaultExpectation.paramPtrs.ctx = &ctx
	mmCheck.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCheck
}

// ExpectAccessTokenParam2 sets up expected param accessToken for AccessService.Check
func (mmCheck *mAccessServiceMockCheck) ExpectAccessTokenParam2(accessToken string) *mAccessServiceMockCheck {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &AccessServiceMockCheckExpectation{}
	}

	if mmCheck.defaultExpectation.params != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Expect")
	}

	if mmCheck.defaultExpectation.paramPtrs == nil {
		mmCheck.defaultExpectation.paramPtrs = &AccessServiceMockCheckParamPtrs{}
	}
	mmCheck.defaultExpectation.paramPtrs.accessToken = &accessToken
	mmCheck.defaultExpectation.expectationOrigins.originAccessToken = minimock.CallerInfo(1)

	return mmCheck
}

// ExpectEndpointParam3 sets up expected param endpoint for AccessService.Check
func (mmCheck *mAccessServiceMockCheck) ExpectEndpointParam3(endpoint string) *mAccessServiceMockCheck {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &AccessServiceMockCheckExpectation{}
	}

	if mmCheck.defaultExpectation.params != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Expect")
	}

	if mmCheck.defaultExpectation.paramPtrs == nil {
		mmCheck.defaultExpectation.paramPtrs = &AccessServiceMockCheckParamPtrs{}
	}
	mmCheck.defaultExpectation.paramPtrs.endpoint = &endpoint
	mmCheck.defaultExpectation.expectationOrigins.originEndpoint = minimock.CallerInfo(1)

	return mmCheck
}

// Inspect accepts an inspector function that has same arguments as the AccessService.Check
func (mmCheck *mAccessServiceMockCheck) Inspect(f func(ctx context.Context, accessToken string, endpoint string)) *mAccessServiceMockCheck {
	if mmCheck.mock.inspectFuncCheck != nil {
		mmCheck.mock.t.Fatalf("Inspect function is already set for AccessServiceMock.Check")
	}

	mmCheck.mock.inspectFuncCheck = f

	return mmCheck
}

// Return sets up results that will be returned by AccessService.Check
func (mmCheck *mAccessServiceMockCheck) Return(err error) *AccessServiceMock {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &AccessServiceMockCheckExpectation{mock: mmCheck.mock}
	}
	mmCheck.defaultExpectation.results = &AccessServiceMockCheckResults{err}
	mmCheck.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCheck.mock
}

// Set uses given function f to mock the AccessService.Check method
func (mmCheck *mAccessServiceMockCheck) Set(f func(ctx context.Context, accessToken string, endpoint string) (err error)) *AccessServiceMock {
	if mmCheck.defaultExpectation != nil {
		mmCheck.mock.t.Fatalf("Default expectation is already set for the AccessService.Check method")
	}

	if len(mmCheck.expectations) > 0 {
		mmCheck.mock.t.Fatalf("Some expectations are already set for the AccessService.Check method")
	}

	mmCheck.mock.funcCheck = f
	mmCheck.mock.funcCheckOrigin = minimock.CallerInfo(1)
	return mmCheck.mock
}

// When sets expectation for the AccessService.Check which will trigger the result defined by the following
// Then helper
func (mmCheck *mAccessServiceMockCheck) When(ctx context.Context, accessToken string, endpoint string) *AccessServiceMockCheckExpectation {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("AccessServiceMock.Check mock is already set by Set")
	}

	expectation := &AccessServiceMockCheckExpectation{
		mock:               mmCheck.mock,
		params:             &AccessServiceMockCheckParams{ctx, accessToken, endpoint},
		expectationOrigins: AccessServiceMockCheckExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCheck.expectations = append(mmCheck.expectations, expectation)
	return expectation
}

// Then sets up AccessService.Check return parameters for the expectation previously defined by the When method
func (e *AccessServiceMockCheckExpectation) Then(err error) *AccessServiceMock {
	e.results = &AccessServiceMockCheckResults{err}
	return e.mock
}

// Times sets number of times AccessService.Check should be invoked
func (mmCheck *mAccessServiceMockCheck) Times(n uint64) *mAccessServiceMockCheck {
	if n == 0 {
		mmCheck.mock.t.Fatalf("Times of AccessServiceMock.Check mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCheck.expectedInvocations, n)
	mmCheck.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCheck
}

func (mmCheck *mAccessServiceMockCheck) invocationsDone() bool {
	if len(mmCheck.expectations) == 0 && mmCheck.defaultExpectation == nil && mmCheck.mock.funcCheck == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCheck.mock.afterCheckCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCheck.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Check implements mm_service.AccessService
func (mmCheck *AccessServiceMock) Check(ctx context.Context, accessToken string, endpoint string) (err error) {
	mm_atomic.AddUint64(&mmCheck.beforeCheckCounter, 1)
	defer mm_atomic.AddUint64(&mmCheck.afterCheckCounter, 1)

	mmCheck.t.Helper()

	if mmCheck.inspectFuncCheck != nil {
		mmCheck.inspectFuncCheck(ctx, accessToken, endpoint)
	}

	mm_params := AccessServiceMockCheckParams{ctx, accessToken, endpoint}

	// Record call args
	mmCheck.CheckMock.mutex.Lock()
	mmCheck.CheckMock.callArgs = append(mmCheck.CheckMock.callArgs, &mm_params)
	mmCheck.CheckMock.mutex.Unlock()

	for _, e := range mmCheck.CheckMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCheck.CheckMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheck.CheckMock.defaultExpectation.Counter, 1)
		mm_want := mmCheck.CheckMock.defaultExpectation.params
		mm_want_ptrs := mmCheck.CheckMock.defaultExpectation.paramPtrs

		mm_got := AccessServiceMockCheckParams{ctx, accessToken, endpoint}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCheck.t.Errorf("AccessServiceMock.Check got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCheck.CheckMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.accessToken != nil && !minimock.Equal(*mm_want_ptrs.accessToken, mm_got.accessToken) {
				mmCheck.t.Errorf("AccessServiceMock.Check got unexpected parameter accessToken, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCheck.CheckMock.defaultExpectation.expectationOrigins.originAccessToken, *mm_want_ptrs.accessToken, mm_got.accessToken, minimock.Diff(*mm_want_ptrs.accessToken, mm_got.accessToken))
			}

			if mm_want_ptrs.endpoint != nil && !minimock.Equal(*mm_want_ptrs.endpoint, mm_got.endpoint) {
				mmCheck.t.Errorf("AccessServiceMock.Check got unexpected parameter endpoint, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCheck.CheckMock.defaultExpectation.expectationOrigins.originEndpoint, *mm_want_ptrs.endpoint, mm_got.endpoint, minimock.Diff(*mm_want_ptrs.endpoint, mm_got.endpoint))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheck.t.Errorf("AccessServiceMock.Check got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCheck.CheckMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheck.CheckMock.defaultExpectation.results
		if mm_results == nil {
			mmCheck.t.Fatal("No results are set for the AccessServiceMock.Check")
		}
		return (*mm_results).err
	}
	if mmCheck.funcCheck != nil {
		return mmCheck.funcCheck(ctx, accessToken, endpoint)
	}
	mmCheck.t.Fatalf("Unexpected call to AccessServiceMock.Check. %v %v %v", ctx, accessToken, endpoint)
	return
}

// CheckAfterCounter returns a count of finished AccessServiceMock.Check invocations
func (mmCheck *AccessServiceMock) CheckAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.afterCheckCounter)
}

// CheckBeforeCounter returns a count of AccessServiceMock.Check invocations
func (mmCheck *AccessServiceMock) CheckBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.beforeCheckCounter)
}

// Calls returns a list of arguments used in each call to AccessServiceMock.Check.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheck *mAccessServiceMockCheck) Calls() []*AccessServiceMockCheckParams {
	mmCheck.mutex.RLock()

	argCopy := make([]*AccessServiceMockCheckParams, len(mmCheck.callArgs))
	copy(argCopy, mmCheck.callArgs)

	mmCheck.mutex.RUnlock()

	return argCopy
}

// MinimockCheckDone returns true if the count of the Check invocations corresponds
// the number of defined expectations
func (m *AccessServiceMock) MinimockCheckDone() bool {
	if m.CheckMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CheckMock.invocationsDone()
}

// MinimockCheckInspect logs each unmet expectation
func (m *AccessServiceMock) MinimockCheckInspect() {
	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AccessServiceMock.Check at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCheckCounter := mm_atomic.LoadUint64(&m.afterCheckCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CheckMock.defaultExpectation != nil && afterCheckCounter < 1 {
		if m.CheckMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to AccessServiceMock.Check at\n%s", m.CheckMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to AccessServiceMock.Check at\n%s with params: %#v", m.CheckMock.defaultExpectation.expectationOrigins.origin, *m.CheckMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheck != nil && afterCheckCounter < 1 {
		m.t.Errorf("Expected call to AccessServiceMock.Check at\n%s", m.funcCheckOrigin)
	}

	if !m.CheckMock.invocationsDone() && afterCheckCounter > 0 {
		m.t.Errorf("Expected %d calls to AccessServiceMock.Check at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CheckMock.expectedInvocations), m.CheckMock.expectedInvocationsOrigin, afterCheckCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AccessServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCheckInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AccessServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *AccessServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCheckDone()
}
