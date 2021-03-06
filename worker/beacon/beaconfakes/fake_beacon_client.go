// Code generated by counterfeiter. DO NOT EDIT.
package beaconfakes

import (
	os "os"
	sync "sync"

	garden "code.cloudfoundry.org/garden"
	beacon "github.com/concourse/concourse/worker/beacon"
)

type FakeBeaconClient struct {
	DeleteWorkerStub        func() error
	deleteWorkerMutex       sync.RWMutex
	deleteWorkerArgsForCall []struct {
	}
	deleteWorkerReturns struct {
		result1 error
	}
	deleteWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	DisableKeepAliveStub        func()
	disableKeepAliveMutex       sync.RWMutex
	disableKeepAliveArgsForCall []struct {
	}
	LandWorkerStub        func() error
	landWorkerMutex       sync.RWMutex
	landWorkerArgsForCall []struct {
	}
	landWorkerReturns struct {
		result1 error
	}
	landWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	RegisterStub        func(<-chan os.Signal, chan<- struct{}) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 <-chan os.Signal
		arg2 chan<- struct{}
	}
	registerReturns struct {
		result1 error
	}
	registerReturnsOnCall map[int]struct {
		result1 error
	}
	ReportContainersStub        func(garden.Client) error
	reportContainersMutex       sync.RWMutex
	reportContainersArgsForCall []struct {
		arg1 garden.Client
	}
	reportContainersReturns struct {
		result1 error
	}
	reportContainersReturnsOnCall map[int]struct {
		result1 error
	}
	ReportVolumesStub        func() error
	reportVolumesMutex       sync.RWMutex
	reportVolumesArgsForCall []struct {
	}
	reportVolumesReturns struct {
		result1 error
	}
	reportVolumesReturnsOnCall map[int]struct {
		result1 error
	}
	RetireWorkerStub        func() error
	retireWorkerMutex       sync.RWMutex
	retireWorkerArgsForCall []struct {
	}
	retireWorkerReturns struct {
		result1 error
	}
	retireWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	SweepContainersStub        func(garden.Client) error
	sweepContainersMutex       sync.RWMutex
	sweepContainersArgsForCall []struct {
		arg1 garden.Client
	}
	sweepContainersReturns struct {
		result1 error
	}
	sweepContainersReturnsOnCall map[int]struct {
		result1 error
	}
	SweepVolumesStub        func() error
	sweepVolumesMutex       sync.RWMutex
	sweepVolumesArgsForCall []struct {
	}
	sweepVolumesReturns struct {
		result1 error
	}
	sweepVolumesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBeaconClient) DeleteWorker() error {
	fake.deleteWorkerMutex.Lock()
	ret, specificReturn := fake.deleteWorkerReturnsOnCall[len(fake.deleteWorkerArgsForCall)]
	fake.deleteWorkerArgsForCall = append(fake.deleteWorkerArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteWorker", []interface{}{})
	fake.deleteWorkerMutex.Unlock()
	if fake.DeleteWorkerStub != nil {
		return fake.DeleteWorkerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteWorkerReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) DeleteWorkerCallCount() int {
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	return len(fake.deleteWorkerArgsForCall)
}

func (fake *FakeBeaconClient) DeleteWorkerReturns(result1 error) {
	fake.DeleteWorkerStub = nil
	fake.deleteWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) DeleteWorkerReturnsOnCall(i int, result1 error) {
	fake.DeleteWorkerStub = nil
	if fake.deleteWorkerReturnsOnCall == nil {
		fake.deleteWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) DisableKeepAlive() {
	fake.disableKeepAliveMutex.Lock()
	fake.disableKeepAliveArgsForCall = append(fake.disableKeepAliveArgsForCall, struct {
	}{})
	fake.recordInvocation("DisableKeepAlive", []interface{}{})
	fake.disableKeepAliveMutex.Unlock()
	if fake.DisableKeepAliveStub != nil {
		fake.DisableKeepAliveStub()
	}
}

func (fake *FakeBeaconClient) DisableKeepAliveCallCount() int {
	fake.disableKeepAliveMutex.RLock()
	defer fake.disableKeepAliveMutex.RUnlock()
	return len(fake.disableKeepAliveArgsForCall)
}

func (fake *FakeBeaconClient) LandWorker() error {
	fake.landWorkerMutex.Lock()
	ret, specificReturn := fake.landWorkerReturnsOnCall[len(fake.landWorkerArgsForCall)]
	fake.landWorkerArgsForCall = append(fake.landWorkerArgsForCall, struct {
	}{})
	fake.recordInvocation("LandWorker", []interface{}{})
	fake.landWorkerMutex.Unlock()
	if fake.LandWorkerStub != nil {
		return fake.LandWorkerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.landWorkerReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) LandWorkerCallCount() int {
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	return len(fake.landWorkerArgsForCall)
}

func (fake *FakeBeaconClient) LandWorkerReturns(result1 error) {
	fake.LandWorkerStub = nil
	fake.landWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) LandWorkerReturnsOnCall(i int, result1 error) {
	fake.LandWorkerStub = nil
	if fake.landWorkerReturnsOnCall == nil {
		fake.landWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.landWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) Register(arg1 <-chan os.Signal, arg2 chan<- struct{}) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 <-chan os.Signal
		arg2 chan<- struct{}
	}{arg1, arg2})
	fake.recordInvocation("Register", []interface{}{arg1, arg2})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.registerReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeBeaconClient) RegisterArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBeaconClient) RegisterReturns(result1 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RegisterReturnsOnCall(i int, result1 error) {
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) ReportContainers(arg1 garden.Client) error {
	fake.reportContainersMutex.Lock()
	ret, specificReturn := fake.reportContainersReturnsOnCall[len(fake.reportContainersArgsForCall)]
	fake.reportContainersArgsForCall = append(fake.reportContainersArgsForCall, struct {
		arg1 garden.Client
	}{arg1})
	fake.recordInvocation("ReportContainers", []interface{}{arg1})
	fake.reportContainersMutex.Unlock()
	if fake.ReportContainersStub != nil {
		return fake.ReportContainersStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.reportContainersReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) ReportContainersCallCount() int {
	fake.reportContainersMutex.RLock()
	defer fake.reportContainersMutex.RUnlock()
	return len(fake.reportContainersArgsForCall)
}

func (fake *FakeBeaconClient) ReportContainersArgsForCall(i int) garden.Client {
	fake.reportContainersMutex.RLock()
	defer fake.reportContainersMutex.RUnlock()
	argsForCall := fake.reportContainersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBeaconClient) ReportContainersReturns(result1 error) {
	fake.ReportContainersStub = nil
	fake.reportContainersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) ReportContainersReturnsOnCall(i int, result1 error) {
	fake.ReportContainersStub = nil
	if fake.reportContainersReturnsOnCall == nil {
		fake.reportContainersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reportContainersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) ReportVolumes() error {
	fake.reportVolumesMutex.Lock()
	ret, specificReturn := fake.reportVolumesReturnsOnCall[len(fake.reportVolumesArgsForCall)]
	fake.reportVolumesArgsForCall = append(fake.reportVolumesArgsForCall, struct {
	}{})
	fake.recordInvocation("ReportVolumes", []interface{}{})
	fake.reportVolumesMutex.Unlock()
	if fake.ReportVolumesStub != nil {
		return fake.ReportVolumesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.reportVolumesReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) ReportVolumesCallCount() int {
	fake.reportVolumesMutex.RLock()
	defer fake.reportVolumesMutex.RUnlock()
	return len(fake.reportVolumesArgsForCall)
}

func (fake *FakeBeaconClient) ReportVolumesReturns(result1 error) {
	fake.ReportVolumesStub = nil
	fake.reportVolumesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) ReportVolumesReturnsOnCall(i int, result1 error) {
	fake.ReportVolumesStub = nil
	if fake.reportVolumesReturnsOnCall == nil {
		fake.reportVolumesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reportVolumesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RetireWorker() error {
	fake.retireWorkerMutex.Lock()
	ret, specificReturn := fake.retireWorkerReturnsOnCall[len(fake.retireWorkerArgsForCall)]
	fake.retireWorkerArgsForCall = append(fake.retireWorkerArgsForCall, struct {
	}{})
	fake.recordInvocation("RetireWorker", []interface{}{})
	fake.retireWorkerMutex.Unlock()
	if fake.RetireWorkerStub != nil {
		return fake.RetireWorkerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.retireWorkerReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) RetireWorkerCallCount() int {
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	return len(fake.retireWorkerArgsForCall)
}

func (fake *FakeBeaconClient) RetireWorkerReturns(result1 error) {
	fake.RetireWorkerStub = nil
	fake.retireWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RetireWorkerReturnsOnCall(i int, result1 error) {
	fake.RetireWorkerStub = nil
	if fake.retireWorkerReturnsOnCall == nil {
		fake.retireWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.retireWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) SweepContainers(arg1 garden.Client) error {
	fake.sweepContainersMutex.Lock()
	ret, specificReturn := fake.sweepContainersReturnsOnCall[len(fake.sweepContainersArgsForCall)]
	fake.sweepContainersArgsForCall = append(fake.sweepContainersArgsForCall, struct {
		arg1 garden.Client
	}{arg1})
	fake.recordInvocation("SweepContainers", []interface{}{arg1})
	fake.sweepContainersMutex.Unlock()
	if fake.SweepContainersStub != nil {
		return fake.SweepContainersStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sweepContainersReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) SweepContainersCallCount() int {
	fake.sweepContainersMutex.RLock()
	defer fake.sweepContainersMutex.RUnlock()
	return len(fake.sweepContainersArgsForCall)
}

func (fake *FakeBeaconClient) SweepContainersArgsForCall(i int) garden.Client {
	fake.sweepContainersMutex.RLock()
	defer fake.sweepContainersMutex.RUnlock()
	argsForCall := fake.sweepContainersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBeaconClient) SweepContainersReturns(result1 error) {
	fake.SweepContainersStub = nil
	fake.sweepContainersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) SweepContainersReturnsOnCall(i int, result1 error) {
	fake.SweepContainersStub = nil
	if fake.sweepContainersReturnsOnCall == nil {
		fake.sweepContainersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sweepContainersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) SweepVolumes() error {
	fake.sweepVolumesMutex.Lock()
	ret, specificReturn := fake.sweepVolumesReturnsOnCall[len(fake.sweepVolumesArgsForCall)]
	fake.sweepVolumesArgsForCall = append(fake.sweepVolumesArgsForCall, struct {
	}{})
	fake.recordInvocation("SweepVolumes", []interface{}{})
	fake.sweepVolumesMutex.Unlock()
	if fake.SweepVolumesStub != nil {
		return fake.SweepVolumesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sweepVolumesReturns
	return fakeReturns.result1
}

func (fake *FakeBeaconClient) SweepVolumesCallCount() int {
	fake.sweepVolumesMutex.RLock()
	defer fake.sweepVolumesMutex.RUnlock()
	return len(fake.sweepVolumesArgsForCall)
}

func (fake *FakeBeaconClient) SweepVolumesReturns(result1 error) {
	fake.SweepVolumesStub = nil
	fake.sweepVolumesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) SweepVolumesReturnsOnCall(i int, result1 error) {
	fake.SweepVolumesStub = nil
	if fake.sweepVolumesReturnsOnCall == nil {
		fake.sweepVolumesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sweepVolumesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	fake.disableKeepAliveMutex.RLock()
	defer fake.disableKeepAliveMutex.RUnlock()
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.reportContainersMutex.RLock()
	defer fake.reportContainersMutex.RUnlock()
	fake.reportVolumesMutex.RLock()
	defer fake.reportVolumesMutex.RUnlock()
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	fake.sweepContainersMutex.RLock()
	defer fake.sweepContainersMutex.RUnlock()
	fake.sweepVolumesMutex.RLock()
	defer fake.sweepVolumesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBeaconClient) recordInvocation(key string, args []interface{}) {
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

var _ beacon.BeaconClient = new(FakeBeaconClient)
