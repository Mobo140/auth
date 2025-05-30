// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Mobo140/auth/internal/repository.LogRepository -o log_repository_minimock.go -n LogRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Mobo140/auth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// LogRepositoryMock implements mm_repository.LogRepository
type LogRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateLogAuth          func(ctx context.Context, logEntry *model.LogEntryAuth) (err error)
	funcCreateLogAuthOrigin    string
	inspectFuncCreateLogAuth   func(ctx context.Context, logEntry *model.LogEntryAuth)
	afterCreateLogAuthCounter  uint64
	beforeCreateLogAuthCounter uint64
	CreateLogAuthMock          mLogRepositoryMockCreateLogAuth

	funcCreateLogUser          func(ctx context.Context, logEntry *model.LogEntryUser) (err error)
	funcCreateLogUserOrigin    string
	inspectFuncCreateLogUser   func(ctx context.Context, logEntry *model.LogEntryUser)
	afterCreateLogUserCounter  uint64
	beforeCreateLogUserCounter uint64
	CreateLogUserMock          mLogRepositoryMockCreateLogUser
}

// NewLogRepositoryMock returns a mock for mm_repository.LogRepository
func NewLogRepositoryMock(t minimock.Tester) *LogRepositoryMock {
	m := &LogRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateLogAuthMock = mLogRepositoryMockCreateLogAuth{mock: m}
	m.CreateLogAuthMock.callArgs = []*LogRepositoryMockCreateLogAuthParams{}

	m.CreateLogUserMock = mLogRepositoryMockCreateLogUser{mock: m}
	m.CreateLogUserMock.callArgs = []*LogRepositoryMockCreateLogUserParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLogRepositoryMockCreateLogAuth struct {
	optional           bool
	mock               *LogRepositoryMock
	defaultExpectation *LogRepositoryMockCreateLogAuthExpectation
	expectations       []*LogRepositoryMockCreateLogAuthExpectation

	callArgs []*LogRepositoryMockCreateLogAuthParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// LogRepositoryMockCreateLogAuthExpectation specifies expectation struct of the LogRepository.CreateLogAuth
type LogRepositoryMockCreateLogAuthExpectation struct {
	mock               *LogRepositoryMock
	params             *LogRepositoryMockCreateLogAuthParams
	paramPtrs          *LogRepositoryMockCreateLogAuthParamPtrs
	expectationOrigins LogRepositoryMockCreateLogAuthExpectationOrigins
	results            *LogRepositoryMockCreateLogAuthResults
	returnOrigin       string
	Counter            uint64
}

// LogRepositoryMockCreateLogAuthParams contains parameters of the LogRepository.CreateLogAuth
type LogRepositoryMockCreateLogAuthParams struct {
	ctx      context.Context
	logEntry *model.LogEntryAuth
}

// LogRepositoryMockCreateLogAuthParamPtrs contains pointers to parameters of the LogRepository.CreateLogAuth
type LogRepositoryMockCreateLogAuthParamPtrs struct {
	ctx      *context.Context
	logEntry **model.LogEntryAuth
}

// LogRepositoryMockCreateLogAuthResults contains results of the LogRepository.CreateLogAuth
type LogRepositoryMockCreateLogAuthResults struct {
	err error
}

// LogRepositoryMockCreateLogAuthOrigins contains origins of expectations of the LogRepository.CreateLogAuth
type LogRepositoryMockCreateLogAuthExpectationOrigins struct {
	origin         string
	originCtx      string
	originLogEntry string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Optional() *mLogRepositoryMockCreateLogAuth {
	mmCreateLogAuth.optional = true
	return mmCreateLogAuth
}

// Expect sets up expected params for LogRepository.CreateLogAuth
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Expect(ctx context.Context, logEntry *model.LogEntryAuth) *mLogRepositoryMockCreateLogAuth {
	if mmCreateLogAuth.mock.funcCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Set")
	}

	if mmCreateLogAuth.defaultExpectation == nil {
		mmCreateLogAuth.defaultExpectation = &LogRepositoryMockCreateLogAuthExpectation{}
	}

	if mmCreateLogAuth.defaultExpectation.paramPtrs != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by ExpectParams functions")
	}

	mmCreateLogAuth.defaultExpectation.params = &LogRepositoryMockCreateLogAuthParams{ctx, logEntry}
	mmCreateLogAuth.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateLogAuth.expectations {
		if minimock.Equal(e.params, mmCreateLogAuth.defaultExpectation.params) {
			mmCreateLogAuth.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateLogAuth.defaultExpectation.params)
		}
	}

	return mmCreateLogAuth
}

// ExpectCtxParam1 sets up expected param ctx for LogRepository.CreateLogAuth
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) ExpectCtxParam1(ctx context.Context) *mLogRepositoryMockCreateLogAuth {
	if mmCreateLogAuth.mock.funcCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Set")
	}

	if mmCreateLogAuth.defaultExpectation == nil {
		mmCreateLogAuth.defaultExpectation = &LogRepositoryMockCreateLogAuthExpectation{}
	}

	if mmCreateLogAuth.defaultExpectation.params != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Expect")
	}

	if mmCreateLogAuth.defaultExpectation.paramPtrs == nil {
		mmCreateLogAuth.defaultExpectation.paramPtrs = &LogRepositoryMockCreateLogAuthParamPtrs{}
	}
	mmCreateLogAuth.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateLogAuth.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateLogAuth
}

// ExpectLogEntryParam2 sets up expected param logEntry for LogRepository.CreateLogAuth
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) ExpectLogEntryParam2(logEntry *model.LogEntryAuth) *mLogRepositoryMockCreateLogAuth {
	if mmCreateLogAuth.mock.funcCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Set")
	}

	if mmCreateLogAuth.defaultExpectation == nil {
		mmCreateLogAuth.defaultExpectation = &LogRepositoryMockCreateLogAuthExpectation{}
	}

	if mmCreateLogAuth.defaultExpectation.params != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Expect")
	}

	if mmCreateLogAuth.defaultExpectation.paramPtrs == nil {
		mmCreateLogAuth.defaultExpectation.paramPtrs = &LogRepositoryMockCreateLogAuthParamPtrs{}
	}
	mmCreateLogAuth.defaultExpectation.paramPtrs.logEntry = &logEntry
	mmCreateLogAuth.defaultExpectation.expectationOrigins.originLogEntry = minimock.CallerInfo(1)

	return mmCreateLogAuth
}

// Inspect accepts an inspector function that has same arguments as the LogRepository.CreateLogAuth
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Inspect(f func(ctx context.Context, logEntry *model.LogEntryAuth)) *mLogRepositoryMockCreateLogAuth {
	if mmCreateLogAuth.mock.inspectFuncCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("Inspect function is already set for LogRepositoryMock.CreateLogAuth")
	}

	mmCreateLogAuth.mock.inspectFuncCreateLogAuth = f

	return mmCreateLogAuth
}

// Return sets up results that will be returned by LogRepository.CreateLogAuth
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Return(err error) *LogRepositoryMock {
	if mmCreateLogAuth.mock.funcCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Set")
	}

	if mmCreateLogAuth.defaultExpectation == nil {
		mmCreateLogAuth.defaultExpectation = &LogRepositoryMockCreateLogAuthExpectation{mock: mmCreateLogAuth.mock}
	}
	mmCreateLogAuth.defaultExpectation.results = &LogRepositoryMockCreateLogAuthResults{err}
	mmCreateLogAuth.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateLogAuth.mock
}

// Set uses given function f to mock the LogRepository.CreateLogAuth method
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Set(f func(ctx context.Context, logEntry *model.LogEntryAuth) (err error)) *LogRepositoryMock {
	if mmCreateLogAuth.defaultExpectation != nil {
		mmCreateLogAuth.mock.t.Fatalf("Default expectation is already set for the LogRepository.CreateLogAuth method")
	}

	if len(mmCreateLogAuth.expectations) > 0 {
		mmCreateLogAuth.mock.t.Fatalf("Some expectations are already set for the LogRepository.CreateLogAuth method")
	}

	mmCreateLogAuth.mock.funcCreateLogAuth = f
	mmCreateLogAuth.mock.funcCreateLogAuthOrigin = minimock.CallerInfo(1)
	return mmCreateLogAuth.mock
}

// When sets expectation for the LogRepository.CreateLogAuth which will trigger the result defined by the following
// Then helper
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) When(ctx context.Context, logEntry *model.LogEntryAuth) *LogRepositoryMockCreateLogAuthExpectation {
	if mmCreateLogAuth.mock.funcCreateLogAuth != nil {
		mmCreateLogAuth.mock.t.Fatalf("LogRepositoryMock.CreateLogAuth mock is already set by Set")
	}

	expectation := &LogRepositoryMockCreateLogAuthExpectation{
		mock:               mmCreateLogAuth.mock,
		params:             &LogRepositoryMockCreateLogAuthParams{ctx, logEntry},
		expectationOrigins: LogRepositoryMockCreateLogAuthExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateLogAuth.expectations = append(mmCreateLogAuth.expectations, expectation)
	return expectation
}

// Then sets up LogRepository.CreateLogAuth return parameters for the expectation previously defined by the When method
func (e *LogRepositoryMockCreateLogAuthExpectation) Then(err error) *LogRepositoryMock {
	e.results = &LogRepositoryMockCreateLogAuthResults{err}
	return e.mock
}

// Times sets number of times LogRepository.CreateLogAuth should be invoked
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Times(n uint64) *mLogRepositoryMockCreateLogAuth {
	if n == 0 {
		mmCreateLogAuth.mock.t.Fatalf("Times of LogRepositoryMock.CreateLogAuth mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateLogAuth.expectedInvocations, n)
	mmCreateLogAuth.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateLogAuth
}

func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) invocationsDone() bool {
	if len(mmCreateLogAuth.expectations) == 0 && mmCreateLogAuth.defaultExpectation == nil && mmCreateLogAuth.mock.funcCreateLogAuth == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateLogAuth.mock.afterCreateLogAuthCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateLogAuth.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateLogAuth implements mm_repository.LogRepository
func (mmCreateLogAuth *LogRepositoryMock) CreateLogAuth(ctx context.Context, logEntry *model.LogEntryAuth) (err error) {
	mm_atomic.AddUint64(&mmCreateLogAuth.beforeCreateLogAuthCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateLogAuth.afterCreateLogAuthCounter, 1)

	mmCreateLogAuth.t.Helper()

	if mmCreateLogAuth.inspectFuncCreateLogAuth != nil {
		mmCreateLogAuth.inspectFuncCreateLogAuth(ctx, logEntry)
	}

	mm_params := LogRepositoryMockCreateLogAuthParams{ctx, logEntry}

	// Record call args
	mmCreateLogAuth.CreateLogAuthMock.mutex.Lock()
	mmCreateLogAuth.CreateLogAuthMock.callArgs = append(mmCreateLogAuth.CreateLogAuthMock.callArgs, &mm_params)
	mmCreateLogAuth.CreateLogAuthMock.mutex.Unlock()

	for _, e := range mmCreateLogAuth.CreateLogAuthMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateLogAuth.CreateLogAuthMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.params
		mm_want_ptrs := mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.paramPtrs

		mm_got := LogRepositoryMockCreateLogAuthParams{ctx, logEntry}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateLogAuth.t.Errorf("LogRepositoryMock.CreateLogAuth got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.logEntry != nil && !minimock.Equal(*mm_want_ptrs.logEntry, mm_got.logEntry) {
				mmCreateLogAuth.t.Errorf("LogRepositoryMock.CreateLogAuth got unexpected parameter logEntry, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.expectationOrigins.originLogEntry, *mm_want_ptrs.logEntry, mm_got.logEntry, minimock.Diff(*mm_want_ptrs.logEntry, mm_got.logEntry))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateLogAuth.t.Errorf("LogRepositoryMock.CreateLogAuth got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateLogAuth.CreateLogAuthMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateLogAuth.t.Fatal("No results are set for the LogRepositoryMock.CreateLogAuth")
		}
		return (*mm_results).err
	}
	if mmCreateLogAuth.funcCreateLogAuth != nil {
		return mmCreateLogAuth.funcCreateLogAuth(ctx, logEntry)
	}
	mmCreateLogAuth.t.Fatalf("Unexpected call to LogRepositoryMock.CreateLogAuth. %v %v", ctx, logEntry)
	return
}

// CreateLogAuthAfterCounter returns a count of finished LogRepositoryMock.CreateLogAuth invocations
func (mmCreateLogAuth *LogRepositoryMock) CreateLogAuthAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLogAuth.afterCreateLogAuthCounter)
}

// CreateLogAuthBeforeCounter returns a count of LogRepositoryMock.CreateLogAuth invocations
func (mmCreateLogAuth *LogRepositoryMock) CreateLogAuthBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLogAuth.beforeCreateLogAuthCounter)
}

// Calls returns a list of arguments used in each call to LogRepositoryMock.CreateLogAuth.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateLogAuth *mLogRepositoryMockCreateLogAuth) Calls() []*LogRepositoryMockCreateLogAuthParams {
	mmCreateLogAuth.mutex.RLock()

	argCopy := make([]*LogRepositoryMockCreateLogAuthParams, len(mmCreateLogAuth.callArgs))
	copy(argCopy, mmCreateLogAuth.callArgs)

	mmCreateLogAuth.mutex.RUnlock()

	return argCopy
}

// MinimockCreateLogAuthDone returns true if the count of the CreateLogAuth invocations corresponds
// the number of defined expectations
func (m *LogRepositoryMock) MinimockCreateLogAuthDone() bool {
	if m.CreateLogAuthMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateLogAuthMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateLogAuthMock.invocationsDone()
}

// MinimockCreateLogAuthInspect logs each unmet expectation
func (m *LogRepositoryMock) MinimockCreateLogAuthInspect() {
	for _, e := range m.CreateLogAuthMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogAuth at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateLogAuthCounter := mm_atomic.LoadUint64(&m.afterCreateLogAuthCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateLogAuthMock.defaultExpectation != nil && afterCreateLogAuthCounter < 1 {
		if m.CreateLogAuthMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogAuth at\n%s", m.CreateLogAuthMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogAuth at\n%s with params: %#v", m.CreateLogAuthMock.defaultExpectation.expectationOrigins.origin, *m.CreateLogAuthMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateLogAuth != nil && afterCreateLogAuthCounter < 1 {
		m.t.Errorf("Expected call to LogRepositoryMock.CreateLogAuth at\n%s", m.funcCreateLogAuthOrigin)
	}

	if !m.CreateLogAuthMock.invocationsDone() && afterCreateLogAuthCounter > 0 {
		m.t.Errorf("Expected %d calls to LogRepositoryMock.CreateLogAuth at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateLogAuthMock.expectedInvocations), m.CreateLogAuthMock.expectedInvocationsOrigin, afterCreateLogAuthCounter)
	}
}

type mLogRepositoryMockCreateLogUser struct {
	optional           bool
	mock               *LogRepositoryMock
	defaultExpectation *LogRepositoryMockCreateLogUserExpectation
	expectations       []*LogRepositoryMockCreateLogUserExpectation

	callArgs []*LogRepositoryMockCreateLogUserParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// LogRepositoryMockCreateLogUserExpectation specifies expectation struct of the LogRepository.CreateLogUser
type LogRepositoryMockCreateLogUserExpectation struct {
	mock               *LogRepositoryMock
	params             *LogRepositoryMockCreateLogUserParams
	paramPtrs          *LogRepositoryMockCreateLogUserParamPtrs
	expectationOrigins LogRepositoryMockCreateLogUserExpectationOrigins
	results            *LogRepositoryMockCreateLogUserResults
	returnOrigin       string
	Counter            uint64
}

// LogRepositoryMockCreateLogUserParams contains parameters of the LogRepository.CreateLogUser
type LogRepositoryMockCreateLogUserParams struct {
	ctx      context.Context
	logEntry *model.LogEntryUser
}

// LogRepositoryMockCreateLogUserParamPtrs contains pointers to parameters of the LogRepository.CreateLogUser
type LogRepositoryMockCreateLogUserParamPtrs struct {
	ctx      *context.Context
	logEntry **model.LogEntryUser
}

// LogRepositoryMockCreateLogUserResults contains results of the LogRepository.CreateLogUser
type LogRepositoryMockCreateLogUserResults struct {
	err error
}

// LogRepositoryMockCreateLogUserOrigins contains origins of expectations of the LogRepository.CreateLogUser
type LogRepositoryMockCreateLogUserExpectationOrigins struct {
	origin         string
	originCtx      string
	originLogEntry string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Optional() *mLogRepositoryMockCreateLogUser {
	mmCreateLogUser.optional = true
	return mmCreateLogUser
}

// Expect sets up expected params for LogRepository.CreateLogUser
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Expect(ctx context.Context, logEntry *model.LogEntryUser) *mLogRepositoryMockCreateLogUser {
	if mmCreateLogUser.mock.funcCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Set")
	}

	if mmCreateLogUser.defaultExpectation == nil {
		mmCreateLogUser.defaultExpectation = &LogRepositoryMockCreateLogUserExpectation{}
	}

	if mmCreateLogUser.defaultExpectation.paramPtrs != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by ExpectParams functions")
	}

	mmCreateLogUser.defaultExpectation.params = &LogRepositoryMockCreateLogUserParams{ctx, logEntry}
	mmCreateLogUser.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateLogUser.expectations {
		if minimock.Equal(e.params, mmCreateLogUser.defaultExpectation.params) {
			mmCreateLogUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateLogUser.defaultExpectation.params)
		}
	}

	return mmCreateLogUser
}

// ExpectCtxParam1 sets up expected param ctx for LogRepository.CreateLogUser
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) ExpectCtxParam1(ctx context.Context) *mLogRepositoryMockCreateLogUser {
	if mmCreateLogUser.mock.funcCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Set")
	}

	if mmCreateLogUser.defaultExpectation == nil {
		mmCreateLogUser.defaultExpectation = &LogRepositoryMockCreateLogUserExpectation{}
	}

	if mmCreateLogUser.defaultExpectation.params != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Expect")
	}

	if mmCreateLogUser.defaultExpectation.paramPtrs == nil {
		mmCreateLogUser.defaultExpectation.paramPtrs = &LogRepositoryMockCreateLogUserParamPtrs{}
	}
	mmCreateLogUser.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateLogUser.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateLogUser
}

// ExpectLogEntryParam2 sets up expected param logEntry for LogRepository.CreateLogUser
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) ExpectLogEntryParam2(logEntry *model.LogEntryUser) *mLogRepositoryMockCreateLogUser {
	if mmCreateLogUser.mock.funcCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Set")
	}

	if mmCreateLogUser.defaultExpectation == nil {
		mmCreateLogUser.defaultExpectation = &LogRepositoryMockCreateLogUserExpectation{}
	}

	if mmCreateLogUser.defaultExpectation.params != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Expect")
	}

	if mmCreateLogUser.defaultExpectation.paramPtrs == nil {
		mmCreateLogUser.defaultExpectation.paramPtrs = &LogRepositoryMockCreateLogUserParamPtrs{}
	}
	mmCreateLogUser.defaultExpectation.paramPtrs.logEntry = &logEntry
	mmCreateLogUser.defaultExpectation.expectationOrigins.originLogEntry = minimock.CallerInfo(1)

	return mmCreateLogUser
}

// Inspect accepts an inspector function that has same arguments as the LogRepository.CreateLogUser
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Inspect(f func(ctx context.Context, logEntry *model.LogEntryUser)) *mLogRepositoryMockCreateLogUser {
	if mmCreateLogUser.mock.inspectFuncCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("Inspect function is already set for LogRepositoryMock.CreateLogUser")
	}

	mmCreateLogUser.mock.inspectFuncCreateLogUser = f

	return mmCreateLogUser
}

// Return sets up results that will be returned by LogRepository.CreateLogUser
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Return(err error) *LogRepositoryMock {
	if mmCreateLogUser.mock.funcCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Set")
	}

	if mmCreateLogUser.defaultExpectation == nil {
		mmCreateLogUser.defaultExpectation = &LogRepositoryMockCreateLogUserExpectation{mock: mmCreateLogUser.mock}
	}
	mmCreateLogUser.defaultExpectation.results = &LogRepositoryMockCreateLogUserResults{err}
	mmCreateLogUser.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateLogUser.mock
}

// Set uses given function f to mock the LogRepository.CreateLogUser method
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Set(f func(ctx context.Context, logEntry *model.LogEntryUser) (err error)) *LogRepositoryMock {
	if mmCreateLogUser.defaultExpectation != nil {
		mmCreateLogUser.mock.t.Fatalf("Default expectation is already set for the LogRepository.CreateLogUser method")
	}

	if len(mmCreateLogUser.expectations) > 0 {
		mmCreateLogUser.mock.t.Fatalf("Some expectations are already set for the LogRepository.CreateLogUser method")
	}

	mmCreateLogUser.mock.funcCreateLogUser = f
	mmCreateLogUser.mock.funcCreateLogUserOrigin = minimock.CallerInfo(1)
	return mmCreateLogUser.mock
}

// When sets expectation for the LogRepository.CreateLogUser which will trigger the result defined by the following
// Then helper
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) When(ctx context.Context, logEntry *model.LogEntryUser) *LogRepositoryMockCreateLogUserExpectation {
	if mmCreateLogUser.mock.funcCreateLogUser != nil {
		mmCreateLogUser.mock.t.Fatalf("LogRepositoryMock.CreateLogUser mock is already set by Set")
	}

	expectation := &LogRepositoryMockCreateLogUserExpectation{
		mock:               mmCreateLogUser.mock,
		params:             &LogRepositoryMockCreateLogUserParams{ctx, logEntry},
		expectationOrigins: LogRepositoryMockCreateLogUserExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateLogUser.expectations = append(mmCreateLogUser.expectations, expectation)
	return expectation
}

// Then sets up LogRepository.CreateLogUser return parameters for the expectation previously defined by the When method
func (e *LogRepositoryMockCreateLogUserExpectation) Then(err error) *LogRepositoryMock {
	e.results = &LogRepositoryMockCreateLogUserResults{err}
	return e.mock
}

// Times sets number of times LogRepository.CreateLogUser should be invoked
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Times(n uint64) *mLogRepositoryMockCreateLogUser {
	if n == 0 {
		mmCreateLogUser.mock.t.Fatalf("Times of LogRepositoryMock.CreateLogUser mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateLogUser.expectedInvocations, n)
	mmCreateLogUser.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateLogUser
}

func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) invocationsDone() bool {
	if len(mmCreateLogUser.expectations) == 0 && mmCreateLogUser.defaultExpectation == nil && mmCreateLogUser.mock.funcCreateLogUser == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateLogUser.mock.afterCreateLogUserCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateLogUser.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateLogUser implements mm_repository.LogRepository
func (mmCreateLogUser *LogRepositoryMock) CreateLogUser(ctx context.Context, logEntry *model.LogEntryUser) (err error) {
	mm_atomic.AddUint64(&mmCreateLogUser.beforeCreateLogUserCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateLogUser.afterCreateLogUserCounter, 1)

	mmCreateLogUser.t.Helper()

	if mmCreateLogUser.inspectFuncCreateLogUser != nil {
		mmCreateLogUser.inspectFuncCreateLogUser(ctx, logEntry)
	}

	mm_params := LogRepositoryMockCreateLogUserParams{ctx, logEntry}

	// Record call args
	mmCreateLogUser.CreateLogUserMock.mutex.Lock()
	mmCreateLogUser.CreateLogUserMock.callArgs = append(mmCreateLogUser.CreateLogUserMock.callArgs, &mm_params)
	mmCreateLogUser.CreateLogUserMock.mutex.Unlock()

	for _, e := range mmCreateLogUser.CreateLogUserMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateLogUser.CreateLogUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateLogUser.CreateLogUserMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateLogUser.CreateLogUserMock.defaultExpectation.params
		mm_want_ptrs := mmCreateLogUser.CreateLogUserMock.defaultExpectation.paramPtrs

		mm_got := LogRepositoryMockCreateLogUserParams{ctx, logEntry}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateLogUser.t.Errorf("LogRepositoryMock.CreateLogUser got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateLogUser.CreateLogUserMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.logEntry != nil && !minimock.Equal(*mm_want_ptrs.logEntry, mm_got.logEntry) {
				mmCreateLogUser.t.Errorf("LogRepositoryMock.CreateLogUser got unexpected parameter logEntry, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateLogUser.CreateLogUserMock.defaultExpectation.expectationOrigins.originLogEntry, *mm_want_ptrs.logEntry, mm_got.logEntry, minimock.Diff(*mm_want_ptrs.logEntry, mm_got.logEntry))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateLogUser.t.Errorf("LogRepositoryMock.CreateLogUser got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateLogUser.CreateLogUserMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateLogUser.CreateLogUserMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateLogUser.t.Fatal("No results are set for the LogRepositoryMock.CreateLogUser")
		}
		return (*mm_results).err
	}
	if mmCreateLogUser.funcCreateLogUser != nil {
		return mmCreateLogUser.funcCreateLogUser(ctx, logEntry)
	}
	mmCreateLogUser.t.Fatalf("Unexpected call to LogRepositoryMock.CreateLogUser. %v %v", ctx, logEntry)
	return
}

// CreateLogUserAfterCounter returns a count of finished LogRepositoryMock.CreateLogUser invocations
func (mmCreateLogUser *LogRepositoryMock) CreateLogUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLogUser.afterCreateLogUserCounter)
}

// CreateLogUserBeforeCounter returns a count of LogRepositoryMock.CreateLogUser invocations
func (mmCreateLogUser *LogRepositoryMock) CreateLogUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLogUser.beforeCreateLogUserCounter)
}

// Calls returns a list of arguments used in each call to LogRepositoryMock.CreateLogUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateLogUser *mLogRepositoryMockCreateLogUser) Calls() []*LogRepositoryMockCreateLogUserParams {
	mmCreateLogUser.mutex.RLock()

	argCopy := make([]*LogRepositoryMockCreateLogUserParams, len(mmCreateLogUser.callArgs))
	copy(argCopy, mmCreateLogUser.callArgs)

	mmCreateLogUser.mutex.RUnlock()

	return argCopy
}

// MinimockCreateLogUserDone returns true if the count of the CreateLogUser invocations corresponds
// the number of defined expectations
func (m *LogRepositoryMock) MinimockCreateLogUserDone() bool {
	if m.CreateLogUserMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateLogUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateLogUserMock.invocationsDone()
}

// MinimockCreateLogUserInspect logs each unmet expectation
func (m *LogRepositoryMock) MinimockCreateLogUserInspect() {
	for _, e := range m.CreateLogUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogUser at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateLogUserCounter := mm_atomic.LoadUint64(&m.afterCreateLogUserCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateLogUserMock.defaultExpectation != nil && afterCreateLogUserCounter < 1 {
		if m.CreateLogUserMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogUser at\n%s", m.CreateLogUserMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLogUser at\n%s with params: %#v", m.CreateLogUserMock.defaultExpectation.expectationOrigins.origin, *m.CreateLogUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateLogUser != nil && afterCreateLogUserCounter < 1 {
		m.t.Errorf("Expected call to LogRepositoryMock.CreateLogUser at\n%s", m.funcCreateLogUserOrigin)
	}

	if !m.CreateLogUserMock.invocationsDone() && afterCreateLogUserCounter > 0 {
		m.t.Errorf("Expected %d calls to LogRepositoryMock.CreateLogUser at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateLogUserMock.expectedInvocations), m.CreateLogUserMock.expectedInvocationsOrigin, afterCreateLogUserCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LogRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateLogAuthInspect()

			m.MinimockCreateLogUserInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LogRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *LogRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateLogAuthDone() &&
		m.MinimockCreateLogUserDone()
}
