package exception

type exception struct {
	err     error
	code    int
	message string
}

type Exception interface {
	Error() string
	GetCode() int
	GetMessage() string
	GetError() error
}

func New(code int, msg string, err error) Exception {
	return &exception{
		code:    code,
		message: msg,
		err:     err,
	}
}

func (e *exception) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e *exception) GetCode() int {
	return e.code
}

func (e *exception) GetMessage() string {
	return e.message
}

func (e *exception) GetError() error {
	return e.err
}

func getString(val string, def string) string {
	if val == "" {
		return def
	}

	return val
}

func firstOne(str []string) string {
	if len(str) > 0 {
		return str[0]
	}

	return ""
}
