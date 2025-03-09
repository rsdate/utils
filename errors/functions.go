package errors

import (
	"fmt"
	"os"
)

// These are error helper functions that can be used to standardize error handling in all my code.

func (e *ErrChecker) CheckErr(eMindex string, y func() (any, error)) (any, error) {
	rvalue, err := y()
	if !e.TestMode {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s. error message: %s\n", []any{e.ErrPrefix, e.EM[eMindex], err.Error()}...)
			if e.PanicMode == "true" {
				panic(fmt.Errorf("%s: %s. error message: %s", []any{e.ErrPrefix, e.EM[eMindex], err.Error()}...))
			} else if e.PanicMode == "false" {
				os.Exit(1)
			}
		}
		return rvalue, nil
	} else {
		if err != nil {
			return nil, fmt.Errorf("%s: %s. error message: %s", []any{e.ErrPrefix, e.EM[eMindex], err.Error()}...)
		} else {
			return rvalue, nil
		}
	}
}

func (e *ErrChecker) CheckErrBool(eMindex string, y func() (bool, string, any)) (any, error) {
	ok, message, rvalue := y()
	if !e.TestMode {
		if !ok {
			fmt.Fprintf(os.Stderr, "%s: %s. error message: %s\n", []any{e.ErrPrefix, e.EM[eMindex], message}...)
			if e.PanicMode == "true" {
				panic(fmt.Errorf("%s: %s. error message: %s", []any{e.ErrPrefix, e.EM[eMindex], message}...))
			} else if e.PanicMode == "false" {
				os.Exit(1)
			}
		}
		return rvalue, nil
	} else {
		if !ok {
			return nil, fmt.Errorf("%s: %s. error message: %s", []any{e.ErrPrefix, e.EM[eMindex], message}...)
		} else {
			return rvalue, nil
		}
	}
}
