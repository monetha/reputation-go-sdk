package facts

import "errors"

var (
	// ErrExchangeMustBeClosedOrAccepted returned when private data exchange is not in Closed or Accepted state
	ErrExchangeMustBeClosedOrAccepted = errors.New("private data exchange must be in Closed or Accepted state")
	// ErrInvalidExchangeKey returned when hash of exchange key does not match exchange key hash stored in private data exchange
	ErrInvalidExchangeKey = errors.New("invalid exchange key")
)
