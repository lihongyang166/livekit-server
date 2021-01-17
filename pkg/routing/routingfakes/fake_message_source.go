// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FakeMessageSource struct {
	ReadMessageStub        func() (protoreflect.ProtoMessage, error)
	readMessageMutex       sync.RWMutex
	readMessageArgsForCall []struct {
	}
	readMessageReturns struct {
		result1 protoreflect.ProtoMessage
		result2 error
	}
	readMessageReturnsOnCall map[int]struct {
		result1 protoreflect.ProtoMessage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMessageSource) ReadMessage() (protoreflect.ProtoMessage, error) {
	fake.readMessageMutex.Lock()
	ret, specificReturn := fake.readMessageReturnsOnCall[len(fake.readMessageArgsForCall)]
	fake.readMessageArgsForCall = append(fake.readMessageArgsForCall, struct {
	}{})
	stub := fake.ReadMessageStub
	fakeReturns := fake.readMessageReturns
	fake.recordInvocation("ReadMessage", []interface{}{})
	fake.readMessageMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMessageSource) ReadMessageCallCount() int {
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	return len(fake.readMessageArgsForCall)
}

func (fake *FakeMessageSource) ReadMessageCalls(stub func() (protoreflect.ProtoMessage, error)) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = stub
}

func (fake *FakeMessageSource) ReadMessageReturns(result1 protoreflect.ProtoMessage, result2 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	fake.readMessageReturns = struct {
		result1 protoreflect.ProtoMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeMessageSource) ReadMessageReturnsOnCall(i int, result1 protoreflect.ProtoMessage, result2 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	if fake.readMessageReturnsOnCall == nil {
		fake.readMessageReturnsOnCall = make(map[int]struct {
			result1 protoreflect.ProtoMessage
			result2 error
		})
	}
	fake.readMessageReturnsOnCall[i] = struct {
		result1 protoreflect.ProtoMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeMessageSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMessageSource) recordInvocation(key string, args []interface{}) {
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

var _ routing.MessageSource = new(FakeMessageSource)
