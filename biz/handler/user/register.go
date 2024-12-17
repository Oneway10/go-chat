package user

import (
	"chat/biz/model/user"
	"chat/common/errorx"
	"chat/common/tools"
	"chat/dal"
	"chat/dal/dao"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RegisterHandler(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	resp := user.NewRegisterResponse()
	if req.Name == "" {
		return nil, errorx.New("用户名不能为空")
	}
	if req.Password == "" || req.ConfirmPassword == "" {
		return nil, errorx.New("密码不能为空")
	}
	if req.GetPassword() != req.GetConfirmPassword() {
		return nil, errorx.New("两次输入密码不一致")
	}

	err := dao.Q.Transaction(func(tx *dao.Query) error {
		// 用户名是否存在
		exist, err := dal.UserRepo.IsNameExist(ctx, req.GetName(), tx)
		if err != nil {
			hlog.CtxErrorf(ctx, "[IsNameExist] fail err=%v", err)
			return errorx.New("用户信息查询失败")
		}
		if exist {
			return errorx.New("用户名已存在")
		}
		if err != nil {
			return err
		}
		password := tools.SHA256(req.GetPassword())
		err = dal.UserRepo.CreateUser(ctx, req.GetName(), password, tx)
		if err != nil {
			hlog.CtxErrorf(ctx, "[CreateUser] fail err=%v", err)
			return errorx.New("用户注册失败")
		}
		return nil
	})
	return resp, err
}
