package user

import (
	"chat/biz/model/user"
	"chat/common/errorx"
	"chat/common/tools"
	"chat/dal/dao"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func LoginHandle(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	if req.GetName() == "" {
		return nil, errorx.New("用户名不能为空")
	}
	if req.GetPassword() == "" {
		return nil, errorx.New("密码不能为空")
	}

	password := tools.SHA256(req.GetPassword())

	u := dao.User
	conditions := []gen.Condition{
		u.Name.Eq(req.GetName()),
		u.Password.Eq(password),
		u.IsDeleted.Eq(0),
	}
	data, err := u.WithContext(ctx).Where(conditions...).First()
	if err != nil {
		hlog.CtxErrorf(ctx, "user find fail, err=%s", err.Error())
		return nil, errorx.NewWithCode(500, "用户查询失败")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) || data == nil {
		return nil, errorx.New("用户名或密码错误")
	}

	return resp, nil
}
