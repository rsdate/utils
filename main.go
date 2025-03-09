package main

import "fmt"

func main() {
	fmt.Println("============ This is the utils module ============")
	fmt.Println(`Packages: 
 - errors
 - types`)
	fmt.Println(`errors package: 
 Functions:
  - func (e *ErrChecker) CreateErr(message string) error
  - func (e *ErrChecker) CheckErr(eMindex string, testMode bool, y func() (any, error)) (any, error)
  - func (e *ErrChecker) CheckErrBool(eMindex string, testMode bool, y func() (bool, string, any)) (any, error)
 Types:
  - type ErrChecker struct {
   ErrPrefix string,
   PanicMode string,
   EM map[string]string,
  }
 Globals: 
  - var eMrpkgengine map[string]string
  - var eMItypes map[string]string
  - var eMIerrors map[string]string
  - var EM map[string]map[string]string
 Tests: 
  - TestCreateErr
  - TestCheckErr
  - TestCheckErrBool`)
}
