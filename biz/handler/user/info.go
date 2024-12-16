package user

import (
	"chat/biz/model/user"
	"chat/common/errorx"
	"chat/dal"
	"context"
	"github.com/duke-git/lancet/v2/pointer"
)

func GetUserInfoHandler(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp := user.NewGetUserInfoResponse()
	u, err := dal.UserRepo.GetUserByID(ctx, req.GetID())
	if err != nil {
		return resp, errorx.New("用户信息查询失败")
	}
	if u == nil {
		return resp, errorx.New("用户不存在")
	}
	userInfo := &user.UserInfo{
		ID:         pointer.Of(u.ID),
		Name:       pointer.Of(u.Name),
		Avatar:     u.Avatar,
		Phone:      u.Phone,
		Email:      u.Email,
		Descrition: u.Description,
	}
	resp.UserInfo = userInfo

	return resp, nil
}
