package amh_go_utils

// TODO AH: Testing for these methods
type (
	Unwrappable interface {
		error
		Unwrap() error
	}

	unwrappable struct {
		newErr error
		msg    *string
		reason error
	}
)

func (u *unwrappable) Error() string {
	// Return the message if we have it, otherwise return from the new error
	if u.msg != nil {
		return *u.msg
	} else {
		return u.newErr.Error()
	}
}

func (u *unwrappable) Unwrap() error {
	return u.reason
}

//CreateWrapper just wraps up an error message into a new type
func CreateWrapper(newInfo string, cause error) Unwrappable {
	return &unwrappable{
		msg:    &newInfo,
		reason: cause,
	}
}

//CreateErrWrapper takes a new error (can't be an Unwrappable) and rolls it into a wrapped error
func CreateErrWrapper(newErr, cause error) Unwrappable {
	return &unwrappable{
		newErr: newErr,
		reason: cause,
	}
}

func TypedDeadSimpleFormatter(input Unwrappable) string {
	return DeadSimpleFormatter(input)
}
