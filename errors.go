package types

import (
	"fmt"
)

const (
	EUNKNOWN uint16 = iota + 1
	EEXIST
	ENOTEXIST
)

const (
	DYNAMIC_ERRORS_INDEX    = 256
	UINT16_MAX_VALUE        = 65535
	MSG_TEMPLATE_MAX_LENGTH = 23
)

var CURRENT_ERROR_INDEX uint16 = DYNAMIC_ERRORS_INDEX + 1

type ErrorInfo struct {
	Code    uint16 `json:"-"`
	Message string `json:"error,omitempty" validate:"required"`
}
type ErrorInfos map[string]uint16

func (e ErrorInfos) GetCodeByMsg(msgTpl string) uint16 {
	if v, ok := e[msgTpl]; ok {
		return v
	}
	return EUNKNOWN
}

func (e ErrorInfos) GetMsgTplByCode(code uint16) string {
	for m, v := range e {
		if v == code {
			return m
		}
	}
	return ""
}

var _ error = &AppError{}

var ErMap ErrorInfos = make(ErrorInfos)

func (err *ErrorInfos) AddError(code uint16, msgTpl string) {
	if len(msgTpl) > MSG_TEMPLATE_MAX_LENGTH {
		panic(fmt.Sprintf("Error message can't be longer than 30 symbols. Message length is %d", len(msgTpl)))
	}
	if code > UINT16_MAX_VALUE {
		panic("Error limit exceeded")
	}
	if _, ok := (*err)[msgTpl]; !ok {
		(*err)[msgTpl] = code
	}
}

func (err AppError) Error() (e string) {

	return fmt.Sprintf("Code: %d, Message: %s", err.Code, err.Message)
}

func ErrorByMsg(msgTpl string, args ...any) ErrorInfo {
	e := GetAppError(msgTpl, args...)
	return e
}

func Err(msgTpl string, args ...any) (e ErrorInfo) {
	e = ErrorByMsg(msgTpl, args...)
	if e.Code == EUNKNOWN {
		CURRENT_ERROR_INDEX = CURRENT_ERROR_INDEX + 1
		ErMap.AddError(CURRENT_ERROR_INDEX, msgTpl)
	}
	e = ErrorByMsg(msgTpl, args...)
	return

}

func GetAppError(msgTpl string, args ...any) ErrorInfo {
	if code, ok := ErMap[msgTpl]; ok {
		return ErrorInfo{
			Message: fmt.Sprintf(msgTpl, args...),
			Code:    code,
		}
	} else {
		msgTpl = ErMap.GetMsgTplByCode(EUNKNOWN)
		return ErrorInfo{
			Message: fmt.Sprintf(msgTpl, args...),
			Code:    EUNKNOWN,
		}
	}
}

func init() {
	ErMap.AddError(EUNKNOWN, "EUNKNOWN: %s/%#v")
	ErMap.AddError(ENOTEXIST, "File not found")
	ErMap.AddError(EEXIST, "File exist")

}
