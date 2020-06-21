package gerrors

import (
	"errors"
	"fmt"
)

// WrapError ...
// This method takes an error which is your service-level error
// And some external errors as interfaces and wraps your error
// With external ones.
// View the examples to deep understanding
func WrapError(appLevelError error, externalErrors ...interface{}) error {
	// Checking for nil and just wrap only if err isn't nil
	/*for _, e := range externalErrors {
		if e == nil {
			return nil
		}
	}
*/
	return fmt.Errorf("%w - %v", appLevelError, externalErrors)
}

// HandleGracefully ...
// This method takes the error and debug mode of your service
// And decides according to that mode.
// If debug mode enabled, it will return full wrapped error
// Otherwise it will return only your service error
func HandleGracefully(err error, debugMode bool) error {
	if err != nil {
		if !debugMode {
			return errors.Unwrap(err)
		}
	}

	return err
}
