package user

import (
	"chat/biz/model/user"
	"chat/common/errorx"
	"chat/common/tools"
	"chat/dal"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/pointer"
	"gorm.io/gorm"
)

func LoginHandler(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	if req.GetName() == "" {
		return nil, errorx.New("用户名不能为空")
	}
	if req.GetPassword() == "" {
		return nil, errorx.New("密码不能为空")
	}

	password := tools.SHA256(req.GetPassword())

	data, err := dal.UserRepo.Login(ctx, req.GetName(), password)
	if err != nil {
		hlog.CtxErrorf(ctx, "user find fail, err=%s", err.Error())
		return nil, errorx.NewWithCode(500, "用户查询失败")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) || data == nil {
		return nil, errorx.New("用户名或密码错误")
	}

	userInfo := &user.UserInfo{
		ID:         pointer.Of(data.ID),
		Name:       pointer.Of(data.Name),
		Avatar:     data.Avatar,
		Phone:      data.Phone,
		Email:      data.Email,
		Descrition: data.Description,
	}
	resp.UserInfo = userInfo

	return resp, nil
}
