// This file was generated by counterfeiter
package fake_id_provider

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/shed/layercake"
)

type FakeIDProvider struct {
	ProvideIDStub        func(path string) (layercake.ID, error)
	provideIDMutex       sync.RWMutex
	provideIDArgsForCall []struct {
		path string
	}
	provideIDReturns struct {
		result1 layercake.ID
		result2 error
	}
}

func (fake *FakeIDProvider) ProvideID(path string) (layercake.ID, error) {
	fake.provideIDMutex.Lock()
	fake.provideIDArgsForCall = append(fake.provideIDArgsForCall, struct {
		path string
	}{path})
	fake.provideIDMutex.Unlock()
	if fake.ProvideIDStub != nil {
		return fake.ProvideIDStub(path)
	} else {
		return fake.provideIDReturns.result1, fake.provideIDReturns.result2
	}
}

func (fake *FakeIDProvider) ProvideIDCallCount() int {
	fake.provideIDMutex.RLock()
	defer fake.provideIDMutex.RUnlock()
	return len(fake.provideIDArgsForCall)
}

func (fake *FakeIDProvider) ProvideIDArgsForCall(i int) string {
	fake.provideIDMutex.RLock()
	defer fake.provideIDMutex.RUnlock()
	return fake.provideIDArgsForCall[i].path
}

func (fake *FakeIDProvider) ProvideIDReturns(result1 layercake.ID, result2 error) {
	fake.ProvideIDStub = nil
	fake.provideIDReturns = struct {
		result1 layercake.ID
		result2 error
	}{result1, result2}
}

var _ layercake.IDProvider = new(FakeIDProvider)