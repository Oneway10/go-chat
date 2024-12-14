package resp_util

import (
	"chat/biz/model/base"
	"chat/common/errorx"
	"errors"
)

var Success = base.BaseResp{}

func GenBaseResp(err error) *base.BaseResp {
	code, msg := getErrCode(err)
	return &base.BaseResp{
		StatusCode:    code,
		StatusMessage: msg,
	}
}

func getErrCode(err error) (code int32, msg string) {
	if err == nil {
		return Success.StatusCode, Success.StatusMessage
	}
	var errCode *errorx.Error
	if errors.As(err, &errCode) {
		return errCode.Code, errCode.Message
	}
	return 1000, err.Error()
}
