package errs

type ClaimedError struct {
	Message string `json:"message"`
}

func NewClaimedError() *ClaimedError {
	return &ClaimedError{
		Message: "you have already claimed your daily reward today.",
	}
}

func (e *ClaimedError) Error() string {
	return e.Message
}
