# [rsdate/utils](github.com/rsdate/utils)

### Summary
`rsdate/utils` is a collection of utilities for different projects of mine that use go as the main language. Currentlt it only contains 2 packages.

## Contents
### [1 - Included Packages](#included-packages)
- #### [1a - `errors`](#errors)
    - ##### [1aa - `ErrChecker` struct](#errchecker-struct)
    - ##### [1ab - `CheckErr` function](#checkerr-function-errchecker-struct)
- #### [1b - `types`](#types)
### [2 - Testing](#testing)

## Included Packages
### Errors
The `errors` package is a package that contains utilities to handle errors.
##### `ErrChecker` struct
The `ErrChecker` struct is the main way to use the `errors` package for error handling. This struct contains fields that include the error prefix, the panic mode, and the list of error messages. Errors are in the form of 
`ErrPrefix: example error. error message: example error from another function.` \
Example:
```go
...
var errChecker = errors.ErrChecker{
    ErrPrefix: "example",
    PanicMode: "true",
    EM: errors.EM[eMIer],
}
```
where `ErrPrefix` is the prefix that will be in the error, `PanicMode` is whether the `CheckErr*` functions will panic or not, and `EM` is the map of error messages that all the `CheckErr*` functions will use.
##### Explanation
This struct is designed to simplify the parameters of the `CheckErr*` functions. That means if you use those functions a lot in your program, you won't have to be repetitive. If for instance you use the functions like this:
```go
...
val1 := errChecker.CheckErr(somevalue, ..., func() (any, error) {
    ...
    foo
    ...
})
...
val2 := errChecker.CheckErr(othervalue, ..., func() (any, error) {
    ...
    bar
    ...
})
...
```
then you don't have to specify the same parameters multiple times. Instead, the variable values change from `somevalue` to `othervalue`. The functions changed from doing operation `foo` to operation `bar`. All other data, such as the error prefix, test mode, etc. is persisted in your local `ErrChecker` instance.
##### `CheckErr` function (`ErrChecker` struct)
The `CheckErr` function, as it says, checks the error. It does this using a function parameter `y`, which is supposed to return values of the type `any` and `error`. It returns a value of type `any`. \
Example:
```go
...
val := errChecker.CheckErr("tsterr1", false, func() (any, error) {
    // whatever you want to do in here
    return return_value, fmt.Errorf("...") // or nil if there is no error
})
...
```
##### Explanation
The first parameter is the error message index. In case you don't remember, when you initialized the `ErrChecker` struct, you passed in a list of error messages. This is so that you don't have to write long error messages every time you call this function. If you don't want to, then you don't have to, but this makes it easier (recomended format for EM field below). Next, the `func` parameter. This parameter takes in an anonymous function and then prints the error or panics if and error is returned, or it just returns the return value if no error is returned. This means instead of writing this:
```go
...
61 some_val, err := // some operation
62 if err != nil {
63  fmt.Print(errMessage)
64     if somePanicModeFlag == "true" {
65         panic(fmt.Errorf("%s", errMessage))
66     } else {
67         os.Exit(1)
68     }
69 }
...
```
which is 9 lines of code, you can do this:
```go
...
61 some_val, _ := errChecker.CheckErr(EMindex, func() (any, error) {
62     some_vaL, err := // some operation
63     return some_vaL, err // or nil if no error
64 })
...
```
which is 4 lines of code (~56% reduction of overall size). Just remember that this is best when you want this style of error checking (printing the error message, then either panicing or exiting).
# Rest coming soon...

