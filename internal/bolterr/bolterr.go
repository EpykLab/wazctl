package bolterr

// Code defines the class of error.
type Code uint8

const (
	// UnknownError is the default.
	UnknownError Code = iota
	// UserError is for errors caused by user input (e.g., bad flag value).
	// These messages are safe to show directly to the user.
	UserError
	// SystemError is for internal problems (e.g., can't write to a file).
	// The message is for logging; a generic message is shown to the user.
	SystemError
)

// Error is the standard error type for the application.
type Error struct {
	// The class of error.
	Code Code
	// The user-facing message.
	Message string
	// The underlying error for logging and debugging.
	Err error
}

func (e *Error) Error() string {
	return e.Message
}
