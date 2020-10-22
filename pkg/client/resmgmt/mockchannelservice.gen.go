// Code generated by counterfeiter. DO NOT EDIT.
package resmgmt

import (
	reqContext "context"
	"sync"

	"github.com/studyzy/fabric-sdk-go/pkg/common/options"
	"github.com/studyzy/fabric-sdk-go/pkg/common/providers/fab"
)

type MockChannelService struct {
	ConfigStub        func() (fab.ChannelConfig, error)
	configMutex       sync.RWMutex
	configArgsForCall []struct{}
	configReturns     struct {
		result1 fab.ChannelConfig
		result2 error
	}
	configReturnsOnCall map[int]struct {
		result1 fab.ChannelConfig
		result2 error
	}
	EventServiceStub        func(opts ...options.Opt) (fab.EventService, error)
	eventServiceMutex       sync.RWMutex
	eventServiceArgsForCall []struct {
		opts []options.Opt
	}
	eventServiceReturns struct {
		result1 fab.EventService
		result2 error
	}
	eventServiceReturnsOnCall map[int]struct {
		result1 fab.EventService
		result2 error
	}
	MembershipStub        func() (fab.ChannelMembership, error)
	membershipMutex       sync.RWMutex
	membershipArgsForCall []struct{}
	membershipReturns     struct {
		result1 fab.ChannelMembership
		result2 error
	}
	membershipReturnsOnCall map[int]struct {
		result1 fab.ChannelMembership
		result2 error
	}
	ChannelConfigStub        func() (fab.ChannelCfg, error)
	channelConfigMutex       sync.RWMutex
	channelConfigArgsForCall []struct{}
	channelConfigReturns     struct {
		result1 fab.ChannelCfg
		result2 error
	}
	channelConfigReturnsOnCall map[int]struct {
		result1 fab.ChannelCfg
		result2 error
	}
	TransactorStub        func(reqCtx reqContext.Context) (fab.Transactor, error)
	transactorMutex       sync.RWMutex
	transactorArgsForCall []struct {
		reqCtx reqContext.Context
	}
	transactorReturns struct {
		result1 fab.Transactor
		result2 error
	}
	transactorReturnsOnCall map[int]struct {
		result1 fab.Transactor
		result2 error
	}
	DiscoveryStub        func() (fab.DiscoveryService, error)
	discoveryMutex       sync.RWMutex
	discoveryArgsForCall []struct{}
	discoveryReturns     struct {
		result1 fab.DiscoveryService
		result2 error
	}
	discoveryReturnsOnCall map[int]struct {
		result1 fab.DiscoveryService
		result2 error
	}
	SelectionStub        func() (fab.SelectionService, error)
	selectionMutex       sync.RWMutex
	selectionArgsForCall []struct{}
	selectionReturns     struct {
		result1 fab.SelectionService
		result2 error
	}
	selectionReturnsOnCall map[int]struct {
		result1 fab.SelectionService
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockChannelService) Config() (fab.ChannelConfig, error) {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct{}{})
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.configReturns.result1, fake.configReturns.result2
}

func (fake *MockChannelService) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *MockChannelService) ConfigReturns(result1 fab.ChannelConfig, result2 error) {
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 fab.ChannelConfig
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) ConfigReturnsOnCall(i int, result1 fab.ChannelConfig, result2 error) {
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 fab.ChannelConfig
			result2 error
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 fab.ChannelConfig
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) EventService(opts ...options.Opt) (fab.EventService, error) {
	fake.eventServiceMutex.Lock()
	ret, specificReturn := fake.eventServiceReturnsOnCall[len(fake.eventServiceArgsForCall)]
	fake.eventServiceArgsForCall = append(fake.eventServiceArgsForCall, struct {
		opts []options.Opt
	}{opts})
	fake.recordInvocation("EventService", []interface{}{opts})
	fake.eventServiceMutex.Unlock()
	if fake.EventServiceStub != nil {
		return fake.EventServiceStub(opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.eventServiceReturns.result1, fake.eventServiceReturns.result2
}

func (fake *MockChannelService) EventServiceCallCount() int {
	fake.eventServiceMutex.RLock()
	defer fake.eventServiceMutex.RUnlock()
	return len(fake.eventServiceArgsForCall)
}

func (fake *MockChannelService) EventServiceArgsForCall(i int) []options.Opt {
	fake.eventServiceMutex.RLock()
	defer fake.eventServiceMutex.RUnlock()
	return fake.eventServiceArgsForCall[i].opts
}

func (fake *MockChannelService) EventServiceReturns(result1 fab.EventService, result2 error) {
	fake.EventServiceStub = nil
	fake.eventServiceReturns = struct {
		result1 fab.EventService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) EventServiceReturnsOnCall(i int, result1 fab.EventService, result2 error) {
	fake.EventServiceStub = nil
	if fake.eventServiceReturnsOnCall == nil {
		fake.eventServiceReturnsOnCall = make(map[int]struct {
			result1 fab.EventService
			result2 error
		})
	}
	fake.eventServiceReturnsOnCall[i] = struct {
		result1 fab.EventService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) Membership() (fab.ChannelMembership, error) {
	fake.membershipMutex.Lock()
	ret, specificReturn := fake.membershipReturnsOnCall[len(fake.membershipArgsForCall)]
	fake.membershipArgsForCall = append(fake.membershipArgsForCall, struct{}{})
	fake.recordInvocation("Membership", []interface{}{})
	fake.membershipMutex.Unlock()
	if fake.MembershipStub != nil {
		return fake.MembershipStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.membershipReturns.result1, fake.membershipReturns.result2
}

func (fake *MockChannelService) MembershipCallCount() int {
	fake.membershipMutex.RLock()
	defer fake.membershipMutex.RUnlock()
	return len(fake.membershipArgsForCall)
}

func (fake *MockChannelService) MembershipReturns(result1 fab.ChannelMembership, result2 error) {
	fake.MembershipStub = nil
	fake.membershipReturns = struct {
		result1 fab.ChannelMembership
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) MembershipReturnsOnCall(i int, result1 fab.ChannelMembership, result2 error) {
	fake.MembershipStub = nil
	if fake.membershipReturnsOnCall == nil {
		fake.membershipReturnsOnCall = make(map[int]struct {
			result1 fab.ChannelMembership
			result2 error
		})
	}
	fake.membershipReturnsOnCall[i] = struct {
		result1 fab.ChannelMembership
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) ChannelConfig() (fab.ChannelCfg, error) {
	fake.channelConfigMutex.Lock()
	ret, specificReturn := fake.channelConfigReturnsOnCall[len(fake.channelConfigArgsForCall)]
	fake.channelConfigArgsForCall = append(fake.channelConfigArgsForCall, struct{}{})
	fake.recordInvocation("ChannelConfig", []interface{}{})
	fake.channelConfigMutex.Unlock()
	if fake.ChannelConfigStub != nil {
		return fake.ChannelConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.channelConfigReturns.result1, fake.channelConfigReturns.result2
}

func (fake *MockChannelService) ChannelConfigCallCount() int {
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	return len(fake.channelConfigArgsForCall)
}

func (fake *MockChannelService) ChannelConfigReturns(result1 fab.ChannelCfg, result2 error) {
	fake.ChannelConfigStub = nil
	fake.channelConfigReturns = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) ChannelConfigReturnsOnCall(i int, result1 fab.ChannelCfg, result2 error) {
	fake.ChannelConfigStub = nil
	if fake.channelConfigReturnsOnCall == nil {
		fake.channelConfigReturnsOnCall = make(map[int]struct {
			result1 fab.ChannelCfg
			result2 error
		})
	}
	fake.channelConfigReturnsOnCall[i] = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) Transactor(reqCtx reqContext.Context) (fab.Transactor, error) {
	fake.transactorMutex.Lock()
	ret, specificReturn := fake.transactorReturnsOnCall[len(fake.transactorArgsForCall)]
	fake.transactorArgsForCall = append(fake.transactorArgsForCall, struct {
		reqCtx reqContext.Context
	}{reqCtx})
	fake.recordInvocation("Transactor", []interface{}{reqCtx})
	fake.transactorMutex.Unlock()
	if fake.TransactorStub != nil {
		return fake.TransactorStub(reqCtx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.transactorReturns.result1, fake.transactorReturns.result2
}

func (fake *MockChannelService) TransactorCallCount() int {
	fake.transactorMutex.RLock()
	defer fake.transactorMutex.RUnlock()
	return len(fake.transactorArgsForCall)
}

func (fake *MockChannelService) TransactorArgsForCall(i int) reqContext.Context {
	fake.transactorMutex.RLock()
	defer fake.transactorMutex.RUnlock()
	return fake.transactorArgsForCall[i].reqCtx
}

func (fake *MockChannelService) TransactorReturns(result1 fab.Transactor, result2 error) {
	fake.TransactorStub = nil
	fake.transactorReturns = struct {
		result1 fab.Transactor
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) TransactorReturnsOnCall(i int, result1 fab.Transactor, result2 error) {
	fake.TransactorStub = nil
	if fake.transactorReturnsOnCall == nil {
		fake.transactorReturnsOnCall = make(map[int]struct {
			result1 fab.Transactor
			result2 error
		})
	}
	fake.transactorReturnsOnCall[i] = struct {
		result1 fab.Transactor
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) Discovery() (fab.DiscoveryService, error) {
	fake.discoveryMutex.Lock()
	ret, specificReturn := fake.discoveryReturnsOnCall[len(fake.discoveryArgsForCall)]
	fake.discoveryArgsForCall = append(fake.discoveryArgsForCall, struct{}{})
	fake.recordInvocation("Discovery", []interface{}{})
	fake.discoveryMutex.Unlock()
	if fake.DiscoveryStub != nil {
		return fake.DiscoveryStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.discoveryReturns.result1, fake.discoveryReturns.result2
}

func (fake *MockChannelService) DiscoveryCallCount() int {
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	return len(fake.discoveryArgsForCall)
}

func (fake *MockChannelService) DiscoveryReturns(result1 fab.DiscoveryService, result2 error) {
	fake.DiscoveryStub = nil
	fake.discoveryReturns = struct {
		result1 fab.DiscoveryService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) DiscoveryReturnsOnCall(i int, result1 fab.DiscoveryService, result2 error) {
	fake.DiscoveryStub = nil
	if fake.discoveryReturnsOnCall == nil {
		fake.discoveryReturnsOnCall = make(map[int]struct {
			result1 fab.DiscoveryService
			result2 error
		})
	}
	fake.discoveryReturnsOnCall[i] = struct {
		result1 fab.DiscoveryService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) Selection() (fab.SelectionService, error) {
	fake.selectionMutex.Lock()
	ret, specificReturn := fake.selectionReturnsOnCall[len(fake.selectionArgsForCall)]
	fake.selectionArgsForCall = append(fake.selectionArgsForCall, struct{}{})
	fake.recordInvocation("Selection", []interface{}{})
	fake.selectionMutex.Unlock()
	if fake.SelectionStub != nil {
		return fake.SelectionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.selectionReturns.result1, fake.selectionReturns.result2
}

func (fake *MockChannelService) SelectionCallCount() int {
	fake.selectionMutex.RLock()
	defer fake.selectionMutex.RUnlock()
	return len(fake.selectionArgsForCall)
}

func (fake *MockChannelService) SelectionReturns(result1 fab.SelectionService, result2 error) {
	fake.SelectionStub = nil
	fake.selectionReturns = struct {
		result1 fab.SelectionService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) SelectionReturnsOnCall(i int, result1 fab.SelectionService, result2 error) {
	fake.SelectionStub = nil
	if fake.selectionReturnsOnCall == nil {
		fake.selectionReturnsOnCall = make(map[int]struct {
			result1 fab.SelectionService
			result2 error
		})
	}
	fake.selectionReturnsOnCall[i] = struct {
		result1 fab.SelectionService
		result2 error
	}{result1, result2}
}

func (fake *MockChannelService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.eventServiceMutex.RLock()
	defer fake.eventServiceMutex.RUnlock()
	fake.membershipMutex.RLock()
	defer fake.membershipMutex.RUnlock()
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	fake.transactorMutex.RLock()
	defer fake.transactorMutex.RUnlock()
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	fake.selectionMutex.RLock()
	defer fake.selectionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockChannelService) recordInvocation(key string, args []interface{}) {
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

var _ fab.ChannelService = new(MockChannelService)
