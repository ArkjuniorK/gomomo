package helper

import "errors"

// This file contain all the custom errors that would be used across the app.
// The errors could be scoped to package or the entire app.

var (
	// ErrEmptyString This error had the 'stringconv' prefix which mean that the error
	// would be specified from the stringconv package, so it would be
	// easier for devs to track where the error came from
	ErrEmptyString = errors.New("stringconv: empty string data")
)
