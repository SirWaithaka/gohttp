package gohttp

//Error defines that an exception has occurred with
// HTTPClient.
type Error struct {
	// custom message
	msg string

	// original error returned by HTTPClient
	Err error
}

func (e Error) Error() string {
	return e.msg
}

// ConnectionRefusedError
type ConnectionRefusedError struct {
	msg string

	// Err is the error that occurred during the operation
	Err error
}

func (e ConnectionRefusedError) Error() string {
	return e.msg
}
