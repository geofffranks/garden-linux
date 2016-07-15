// This file was generated by counterfeiter
package fake_commander

import (
	"os/exec"
	"sync"

	"code.cloudfoundry.org/garden-linux/container_daemon"
)

type FakeCommander struct {
	CommandStub        func(args ...string) *exec.Cmd
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		args []string
	}
	commandReturns struct {
		result1 *exec.Cmd
	}
}

func (fake *FakeCommander) Command(args ...string) *exec.Cmd {
	fake.commandMutex.Lock()
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		args []string
	}{args})
	fake.commandMutex.Unlock()
	if fake.CommandStub != nil {
		return fake.CommandStub(args...)
	} else {
		return fake.commandReturns.result1
	}
}

func (fake *FakeCommander) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

func (fake *FakeCommander) CommandArgsForCall(i int) []string {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return fake.commandArgsForCall[i].args
}

func (fake *FakeCommander) CommandReturns(result1 *exec.Cmd) {
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

var _ container_daemon.Commander = new(FakeCommander)
