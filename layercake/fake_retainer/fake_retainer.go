// This file was generated by counterfeiter
package fake_retainer

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/layercake"
)

type FakeRetainer struct {
	RetainStub        func(id layercake.ID)
	retainMutex       sync.RWMutex
	retainArgsForCall []struct {
		id layercake.ID
	}
	ReleaseStub        func(id layercake.ID)
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		id layercake.ID
	}
	IsHeldStub        func(id layercake.ID) bool
	isHeldMutex       sync.RWMutex
	isHeldArgsForCall []struct {
		id layercake.ID
	}
	isHeldReturns struct {
		result1 bool
	}
}

func (fake *FakeRetainer) Retain(id layercake.ID) {
	fake.retainMutex.Lock()
	fake.retainArgsForCall = append(fake.retainArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.retainMutex.Unlock()
	if fake.RetainStub != nil {
		fake.RetainStub(id)
	}
}

func (fake *FakeRetainer) RetainCallCount() int {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return len(fake.retainArgsForCall)
}

func (fake *FakeRetainer) RetainArgsForCall(i int) layercake.ID {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return fake.retainArgsForCall[i].id
}

func (fake *FakeRetainer) Release(id layercake.ID) {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		fake.ReleaseStub(id)
	}
}

func (fake *FakeRetainer) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeRetainer) ReleaseArgsForCall(i int) layercake.ID {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].id
}

func (fake *FakeRetainer) IsHeld(id layercake.ID) bool {
	fake.isHeldMutex.Lock()
	fake.isHeldArgsForCall = append(fake.isHeldArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.isHeldMutex.Unlock()
	if fake.IsHeldStub != nil {
		return fake.IsHeldStub(id)
	} else {
		return fake.isHeldReturns.result1
	}
}

func (fake *FakeRetainer) IsHeldCallCount() int {
	fake.isHeldMutex.RLock()
	defer fake.isHeldMutex.RUnlock()
	return len(fake.isHeldArgsForCall)
}

func (fake *FakeRetainer) IsHeldArgsForCall(i int) layercake.ID {
	fake.isHeldMutex.RLock()
	defer fake.isHeldMutex.RUnlock()
	return fake.isHeldArgsForCall[i].id
}

func (fake *FakeRetainer) IsHeldReturns(result1 bool) {
	fake.IsHeldStub = nil
	fake.isHeldReturns = struct {
		result1 bool
	}{result1}
}

var _ layercake.Retainer = new(FakeRetainer)
