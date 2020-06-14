package gerrors

import (
	"encoding/json"
	"errors"
	"testing"
)

var (
	ErrValidationError = errors.New("invalid data provided")
	wrongJSONStr       = `
{
	"id":"1",
	"username":"admin"
}
`
	trueJSONStr = `
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
	err := json.Unmarshal([]byte(wrongJSONStr), &u)
	err = WrapError(ErrValidationError, err)
	if !errors.Is(err, ErrValidationError) {
		t.FailNow()
	}

	// True
	err = json.Unmarshal([]byte(trueJSONStr), &u)
	err = WrapError(ErrValidationError, err)
	if errors.Is(err, ErrValidationError) {
		t.FailNow()
	}
}

func TestHandleGracefully(t *testing.T) {
	u := new(user)

	// False
	err := json.Unmarshal([]byte(wrongJSONStr), &u)

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
