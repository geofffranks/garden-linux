// This file was generated by counterfeiter
package fake_capabilities

import (
	"sync"

	"code.cloudfoundry.org/garden-linux/containerizer"
)

type FakeCapabilities struct {
	LimitStub        func(bool) error
	limitMutex       sync.RWMutex
	limitArgsForCall []struct {
		arg1 bool
	}
	limitReturns struct {
		result1 error
	}
}

func (fake *FakeCapabilities) Limit(arg1 bool) error {
	fake.limitMutex.Lock()
	fake.limitArgsForCall = append(fake.limitArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.limitMutex.Unlock()
	if fake.LimitStub != nil {
		return fake.LimitStub(arg1)
	} else {
		return fake.limitReturns.result1
	}
}

func (fake *FakeCapabilities) LimitCallCount() int {
	fake.limitMutex.RLock()
	defer fake.limitMutex.RUnlock()
	return len(fake.limitArgsForCall)
}

func (fake *FakeCapabilities) LimitArgsForCall(i int) bool {
	fake.limitMutex.RLock()
	defer fake.limitMutex.RUnlock()
	return fake.limitArgsForCall[i].arg1
}

func (fake *FakeCapabilities) LimitReturns(result1 error) {
	fake.LimitStub = nil
	fake.limitReturns = struct {
		result1 error
	}{result1}
}

var _ containerizer.Capabilities = new(FakeCapabilities)
