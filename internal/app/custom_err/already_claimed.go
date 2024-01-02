package custom_err

type ClaimedError struct {
	Message string `json:"message"`
}

func NewClaimedError() *ClaimedError {
	return &ClaimedError{
		Message: "You have already claimed your daily reward today.",
	}
}

func (e *ClaimedError) Error() string {
	return e.Message
}
