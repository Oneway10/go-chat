package user

import (
	"chat/biz/model/user"
	"chat/common/auth"
	"chat/common/errorx"
	"chat/dal"
	"chat/dal/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/duke-git/lancet/v2/pointer"
)

func GetUserInfoHandler(ctx context.Context, c *app.RequestContext, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp := user.NewGetUserInfoResponse()
	// 用户是否有权限查询
	data, ok := c.Get(auth.IdentityKey)
	userData, okk := data.(*model.User)
	if !ok || !okk {
		hlog.Infof("Invalid Token")
		return resp, errorx.NewWithCode(consts.StatusUnauthorized, "Invalid Token")
	}
	if userData.ID != req.GetID() {
		hlog.Warnf("No Auth, userID: %v, want request user info id: %v", userData.ID, req.GetID())
		return resp, errorx.NewWithCode(consts.StatusForbidden, "无查询权限")
	}
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
