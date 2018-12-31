package event

import (
	"reflect"
	"testing"
)

func TestGetEventManager(t *testing.T) {
	if reflect.TypeOf(Manager()) != reflect.TypeOf(&manager{}) || Manager() == nil {
		t.Error("Unexpected value for event manager")
	}
}

func TestManager_Dispatch(t *testing.T) {

}

func TestManager_Listen(t *testing.T) {

}

func TestManager_RunConcurrent(t *testing.T) {
	sut := Manager()
	sut.RunConcurrent()
	if sut.runAsync != true {
		t.Error("failed to start event manager routine")
	}
}

func TestManager_Unlisten(t *testing.T) {

}

func TestManager_Unregister(t *testing.T) {
	sut := Manager()
	sut.RunConcurrent()
	if sut.runAsync != true {
		t.Error("failed to start event manager routine")
	}
	sut.Unregister()
	if sut.runAsync != false {
		t.Error("failed to stop event manager routine")
	}
}
