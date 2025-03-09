package errors

// These are types that can be used to standardize error handling

type ErrChecker struct {
	ErrPrefix string
	PanicMode string
	EM        map[string]string
	TestMode  bool
}
