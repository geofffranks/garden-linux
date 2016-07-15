// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/garden-linux/linux_backend"
)

type FakeContainerProvider struct {
	ProvideContainerStub        func(linux_backend.LinuxContainerSpec) linux_backend.Container
	provideContainerMutex       sync.RWMutex
	provideContainerArgsForCall []struct {
		arg1 linux_backend.LinuxContainerSpec
	}
	provideContainerReturns struct {
		result1 linux_backend.Container
	}
}

func (fake *FakeContainerProvider) ProvideContainer(arg1 linux_backend.LinuxContainerSpec) linux_backend.Container {
	fake.provideContainerMutex.Lock()
	fake.provideContainerArgsForCall = append(fake.provideContainerArgsForCall, struct {
		arg1 linux_backend.LinuxContainerSpec
	}{arg1})
	fake.provideContainerMutex.Unlock()
	if fake.ProvideContainerStub != nil {
		return fake.ProvideContainerStub(arg1)
	} else {
		return fake.provideContainerReturns.result1
	}
}

func (fake *FakeContainerProvider) ProvideContainerCallCount() int {
	fake.provideContainerMutex.RLock()
	defer fake.provideContainerMutex.RUnlock()
	return len(fake.provideContainerArgsForCall)
}

func (fake *FakeContainerProvider) ProvideContainerArgsForCall(i int) linux_backend.LinuxContainerSpec {
	fake.provideContainerMutex.RLock()
	defer fake.provideContainerMutex.RUnlock()
	return fake.provideContainerArgsForCall[i].arg1
}

func (fake *FakeContainerProvider) ProvideContainerReturns(result1 linux_backend.Container) {
	fake.ProvideContainerStub = nil
	fake.provideContainerReturns = struct {
		result1 linux_backend.Container
	}{result1}
}

var _ linux_backend.ContainerProvider = new(FakeContainerProvider)
