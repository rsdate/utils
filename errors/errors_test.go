package errors_test

// This is a test file for the errors package.

import (
	"fmt"
	"testing"

	"github.com/rsdate/utils/errors"
)

var errChecker = errors.ErrChecker{
	ErrPrefix: "test error",
	PanicMode: "false",
	EM:        errors.EM["eMIer"],
	TestMode:  true,
}

func TestCheckErr(t *testing.T) {
	const message string = "test error message"
	_, err := errChecker.CheckErr("tsterr1", func() (any, error) {
		return nil, fmt.Errorf("%s", []any{message}...)
	})
	if err.Error() != "test error: test error 1. error message: test error message" {
		t.Errorf("CheckErr() failed, expected error message %s, got %s", []any{"test error: test error 1. error message: test error message", err.Error()}...)
	} else if err == nil {
		t.Errorf("CheckErr() failed, expected an error, got nil")
	}
}

func TestCheckErrBool(t *testing.T) {
	var message string = "test error message"
	_, err := errChecker.CheckErrBool("tsterr1", func() (bool, string, any) {
		return false, message, nil
	})
	var errMessage string = "test error: test error 1. error message: test error message"
	if err.Error() != errMessage {
		t.Errorf("CheckErrBool() failed, expected error message %s, got %s", []any{errMessage, err.Error()}...)
	} else if err == nil {
		t.Errorf("CheckErrBool() failed, expected an error, got nil")
	}
}
