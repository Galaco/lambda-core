package logger

import (
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"testing"
)

func TestNotice(t *testing.T) {
	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	expected := "hello world"
	Notice(expected)
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestWarn(t *testing.T) {
	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	expected := "hello world"
	Warn(expected)
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestError(t *testing.T) {
	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	expected := "hello world"
	Error(expected)
	if expected != actual {
		t.Error("log info doesn't match expected")
	}

	Error(errors.New("foo"))
}

func TestFatal(t *testing.T) {
	t.Skip("can't test a wrapper for log.Fatal()")
}

func TestEnablePretty(t *testing.T) {
	EnablePretty()

	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	colour := aurora.NewAurora(true)
	expected := "hello world"
	Error(expected)
	expected = fmt.Sprint(colour.Red(expected))
	if expected != actual {
		t.Error(actual)
		t.Error("log info doesn't match expected")
	}
}

func TestDisablePretty(t *testing.T) {
	EnablePretty()
	DisablePretty()

	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	expected := "hello world"
	Error(expected)
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestSetDestination(t *testing.T) {
	actual := ""
	SetOutputPipeFunc(func(val string) {
		actual = val
	})

	expected := "hello world"
	Notice(expected)
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}