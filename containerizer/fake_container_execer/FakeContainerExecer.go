// This file was generated by counterfeiter
package fake_container_execer

import (
	"sync"

	"code.cloudfoundry.org/garden-linux/containerizer"
)

type FakeContainerExecer struct {
	ExecStub        func(binPath string, args ...string) (int, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		binPath string
		args    []string
	}
	execReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeContainerExecer) Exec(binPath string, args ...string) (int, error) {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		binPath string
		args    []string
	}{binPath, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(binPath, args...)
	} else {
		return fake.execReturns.result1, fake.execReturns.result2
	}
}

func (fake *FakeContainerExecer) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeContainerExecer) ExecArgsForCall(i int) (string, []string) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].binPath, fake.execArgsForCall[i].args
}

func (fake *FakeContainerExecer) ExecReturns(result1 int, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ containerizer.ContainerExecer = new(FakeContainerExecer)
