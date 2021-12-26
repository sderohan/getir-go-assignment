package error

import "errors"

var (
	ErrInvalid_QUERY_PARAMS = errors.New("invalid query params")
	ErrMarshalData          = errors.New("error occured while reading data")
)
