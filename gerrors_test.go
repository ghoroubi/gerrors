package gerrors

import (
	"encoding/json"
	"errors"
	"testing"
)

var (
	ErrValidationError = errors.New("invalid data provided")
)

var (
	wrongJsonStr = `
{
	"id":"1",
	"username":"admin"
}
`
	trueJsonStr = `
{
	"id":1,
	"username":"admin"
}
`
)

// User
type user struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func TestWrapError(t *testing.T) {

	u := new(user)

	// False
	err := json.Unmarshal([]byte(wrongJsonStr), &u)
	err = WrapError(ErrValidationError, err)
	if !errors.Is(err, ErrValidationError) {
		t.FailNow()
	}

	// True
	err = json.Unmarshal([]byte(trueJsonStr), &u)
	err = WrapError(ErrValidationError, err)
	if errors.Is(err, ErrValidationError) {
		t.FailNow()
	}
}

func TestHandleGracefully(t *testing.T) {
	u := new(user)

	// False
	err := json.Unmarshal([]byte(wrongJsonStr), &u)

	err = WrapError(ErrValidationError, err)
	if !errors.Is(err, ErrValidationError) {
		t.FailNow()
	}

	// Error in production mode
	productionErr := HandleGracefully(err, false)

	// Error in development/debug mode
	developErr := HandleGracefully(err, true)

	if !errors.As(ErrValidationError, &productionErr) {
		t.FailNow()
	}

	if !errors.As(ErrValidationError, &developErr) {
		t.FailNow()
	}

	if !errors.Is(ErrValidationError, productionErr) {
		t.FailNow()
	}

	if !errors.Is(ErrValidationError, developErr) {
		t.FailNow()
	}
}
