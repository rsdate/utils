/*
This package is a package that contains code to help with type operations.
This package is a part of the github.com/rsdate/utils

EM identifier: eMIty

EM errMessages subset: csterr,
*/
package types

import (
	"fmt"
	"strconv"

	"github.com/rsdate/utils/errors"
)

// These are helper functions that are used for different type operations.

var (
	// This is the ErrChecker variable that contains the error prefix, panic mode, and error messages.
	ErrChecker = errors.ErrChecker{
		ErrPrefix: "casting error",
		PanicMode: "false",
		EM:        errors.EM["eMIty"],
		TestMode:  false,
	}
	// Shorthand for ErrChecker.ErrPrefix
	Ep = ErrChecker.ErrPrefix
	// Shorthand for ErrChecker.EM
	Em = ErrChecker.EM
)

/*
Description: This function casts a value to a specific type. (with error handling)

Param: value (the value that will be casted) type: any

Param: typ (the type you want to cast to) type: string

Returns: val (the casted value) or an error message that says "casting error: unknown type" type: any
*/
func Cast(value any, testMode bool, typ string) any {
	switch typ {
	case "string":
		// Check if the wanted type is string
		val, _ := ErrChecker.CheckErrBool("csterr1", func() (bool, string, any) {
			strVal, ok := value.(string)
			if !ok {
				strVal = fmt.Sprintf("%v", value)
			}
			return true, "", strVal
		})
		return val
	case "int":
		// Check if the wanted type is int
		val, _ := ErrChecker.CheckErrBool("csterr2", func() (bool, string, any) {
			intVal, ok := value.(int)
			if !ok {
				floatVal, ok := value.(float64)
				if ok {
					intVal = int(floatVal)
				} else {
					strVal, ok := value.(string)
					if ok {
						intVal, _ = strconv.Atoi(strVal)
					} else {
						return true, "could not cast to int", 0
					}
				}
			}
			return true, "", intVal
		})
		return val
	case "float32":
		// Check if the wanted type is float32
		val, _ := ErrChecker.CheckErrBool("csterr3", func() (bool, string, any) {
			floatVal, ok := value.(float32)
			if !ok {
				intVal, ok := value.(int)
				if ok {
					floatVal = float32(intVal)
				} else {
					strVal, ok := value.(string)
					if ok {
						FloatVal, _ := strconv.ParseFloat(strVal, 32)
						return true, "", FloatVal
					} else {
						return true, "could not cast to float64", 0.0
					}
				}
			}
			return true, "", floatVal
		})
		return val
	case "float64":
		// Check if the wanted type is float64
		val, _ := ErrChecker.CheckErrBool("csterr4", func() (bool, string, any) {
			floatVal, ok := value.(float64)
			if !ok {
				intVal, ok := value.(int)
				if ok {
					floatVal = float64(intVal)
				} else {
					strVal, ok := value.(string)
					if ok {
						floatVal, _ = strconv.ParseFloat(strVal, 64)
					} else {
						return true, "could not cast to float64", 0.0
					}
				}
			}
			return true, "", floatVal
		})
		return val
	}
	return Ep + ": " + Em["csterr5"]
}
