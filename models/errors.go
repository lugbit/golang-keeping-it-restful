package models

// Application specific error codes
const (
	// User related error codes
	//
	// Email is empty
	AppCodeEmailEmpty = 3400
	// Email is invalid
	AppCodeEmailInvalid = 3499
	// Password is empty
	AppCodePasswordEmpty = 3401
	// Email doesn't exist
	AppCodeEmailNotFound = 3402
	// Password incorrect
	AppCodePasswordIncorrect = 3403

	// Notes related error codes
	//
	// Note title is empty
	AppCodeNoteTitleEmpty = 6800
	// Note text is empty
	AppCodeNoteTextEmpty = 6850
	// Note ID is empty
	AppCodeNoteIDEmpty = 6888
	// Collection is empty
	AppCodeNoteColEmpty = 6900
	// Note doesn't exist for this user
	AppCodeNoteNotFound = 6901
	// Error adding note
	AppCodeAddFailed = 6906
)

// ErrorObject struct represents an individual error
type ErrorObject struct {
	// Application specific error code
	Code int `json:"code"`
	// Human readable text explaining the problem
	Text string `json:"text"`
	// Hint to help the user to resolve the issue
	Hint string `json:"hint"`
	// Additional info URL
	Info string `json:"info"`
}

// ErrorsPayload to hold array of errors
type ErrorsPayload struct {
	Errors []ErrorObject `json:"errors"`
}

// NewError function that returns an ErrorObject for convenience
func NewError(code int, text, hint, info string) ErrorObject {
	return ErrorObject{
		Code: code,
		Text: text,
		Hint: hint,
		Info: info,
	}
}
