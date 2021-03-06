package optional

var ErrorValueIsNotPresentMessage = `Value is not present`
var ErrorValueIsPresentMessage = `Value is present`

// ErrorValueIsNotPresent
type ErrorValueIsNotPresent struct {
	message string
}

func CreateErrorValueIsNotPresent() *ErrorValueIsNotPresent {
	return &ErrorValueIsNotPresent{
		message: ErrorValueIsNotPresentMessage,
	}
}

func (err *ErrorValueIsNotPresent) Error() string {
	return err.message
}

// ErrorValueIsPresent
type ErrorValueIsPresent struct {
	message string
}

func CreateErrorValueIsPresent() *ErrorValueIsPresent {
	return &ErrorValueIsPresent{
		message: ErrorValueIsPresentMessage,
	}
}

func (err *ErrorValueIsPresent) Error() string {
	return err.message
}
