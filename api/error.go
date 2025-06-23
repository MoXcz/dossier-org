package api

type APIServerError struct {
	Status int
	Msg    string
}

func (e APIServerError) Error() string {
	return e.Msg
}

type APIValidateUserError struct {
	Status      int
	NameMsg     string
	PasswordMsg string
	EmailMsg    string
}

func (e APIValidateUserError) Error() string {
	if e.NameMsg != "" {
		return e.NameMsg
	}
	if e.PasswordMsg != "" {
		return e.PasswordMsg
	}
	if e.EmailMsg != "" {
		return e.EmailMsg
	}
	return "No error"
}
