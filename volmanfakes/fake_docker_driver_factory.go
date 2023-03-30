// Code generated by counterfeiter. DO NOT EDIT.
package volmanfakes

import (
	sync "sync"

	dockerdriver "code.cloudfoundry.org/dockerdriver"
	lager "code.cloudfoundry.org/lager/v3"
	voldiscoverers "code.cloudfoundry.org/volman/voldiscoverers"
)

type FakeDockerDriverFactory struct {
	DockerDriverStub        func(lager.Logger, string, string, string) (dockerdriver.Driver, error)
	dockerDriverMutex       sync.RWMutex
	dockerDriverArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}
	dockerDriverReturns struct {
		result1 dockerdriver.Driver
		result2 error
	}
	dockerDriverReturnsOnCall map[int]struct {
		result1 dockerdriver.Driver
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDockerDriverFactory) DockerDriver(arg1 lager.Logger, arg2 string, arg3 string, arg4 string) (dockerdriver.Driver, error) {
	fake.dockerDriverMutex.Lock()
	ret, specificReturn := fake.dockerDriverReturnsOnCall[len(fake.dockerDriverArgsForCall)]
	fake.dockerDriverArgsForCall = append(fake.dockerDriverArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DockerDriver", []interface{}{arg1, arg2, arg3, arg4})
	fake.dockerDriverMutex.Unlock()
	if fake.DockerDriverStub != nil {
		return fake.DockerDriverStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dockerDriverReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDockerDriverFactory) DockerDriverCallCount() int {
	fake.dockerDriverMutex.RLock()
	defer fake.dockerDriverMutex.RUnlock()
	return len(fake.dockerDriverArgsForCall)
}

func (fake *FakeDockerDriverFactory) DockerDriverCalls(stub func(lager.Logger, string, string, string) (dockerdriver.Driver, error)) {
	fake.dockerDriverMutex.Lock()
	defer fake.dockerDriverMutex.Unlock()
	fake.DockerDriverStub = stub
}

func (fake *FakeDockerDriverFactory) DockerDriverArgsForCall(i int) (lager.Logger, string, string, string) {
	fake.dockerDriverMutex.RLock()
	defer fake.dockerDriverMutex.RUnlock()
	argsForCall := fake.dockerDriverArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDockerDriverFactory) DockerDriverReturns(result1 dockerdriver.Driver, result2 error) {
	fake.dockerDriverMutex.Lock()
	defer fake.dockerDriverMutex.Unlock()
	fake.DockerDriverStub = nil
	fake.dockerDriverReturns = struct {
		result1 dockerdriver.Driver
		result2 error
	}{result1, result2}
}

func (fake *FakeDockerDriverFactory) DockerDriverReturnsOnCall(i int, result1 dockerdriver.Driver, result2 error) {
	fake.dockerDriverMutex.Lock()
	defer fake.dockerDriverMutex.Unlock()
	fake.DockerDriverStub = nil
	if fake.dockerDriverReturnsOnCall == nil {
		fake.dockerDriverReturnsOnCall = make(map[int]struct {
			result1 dockerdriver.Driver
			result2 error
		})
	}
	fake.dockerDriverReturnsOnCall[i] = struct {
		result1 dockerdriver.Driver
		result2 error
	}{result1, result2}
}

func (fake *FakeDockerDriverFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dockerDriverMutex.RLock()
	defer fake.dockerDriverMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDockerDriverFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ voldiscoverers.DockerDriverFactory = new(FakeDockerDriverFactory)
