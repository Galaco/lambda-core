package logger

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"testing"
)

func TestNotice(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)

	expected := "hello world"
	Notice(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestWarn(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)
	expected := "hello world"
	Warn(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestError(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)

	expected := "hello world"
	Error(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}

	Error(errors.New("foo"))
}

func TestFatal(t *testing.T) {
	t.Skip("can't test a wrapper for log.Panic()")
}

func TestEnablePretty(t *testing.T) {
	EnablePretty()
	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)

	colour := aurora.NewAurora(true)
	expected := "hello world"
	Error(expected)
	actual := writer.String()
	expected = fmt.Sprint(colour.Red(expected))
	if expected != actual {
		t.Error(actual)
		t.Error("log info doesn't match expected")
	}
}

func TestDisablePretty(t *testing.T) {
	EnablePretty()
	DisablePretty()

	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)

	expected := "hello world"
	Error(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestSetDestination(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 0))
	SetWriter(writer)

	expected := "hello world"
	Notice(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}
