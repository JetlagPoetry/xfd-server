package xerr

import (
	"fmt"
)

var _ XErr = (*xErr)(nil)

type XErr interface {
	Error() string
	Code() XCode
}
type xErr struct {
	code XCode
	err  error
}

func (t xErr) Error() string {
	return fmt.Sprintf("%v", t.err.Error())
}

func (t xErr) Code() XCode {
	return t.code
}

func WithCode(code XCode, err error) XErr {
	if err == nil {
		return nil
	}
	return &xErr{
		code: code,
		err:  err,
	}
}
