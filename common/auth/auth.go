package auth

import (
	"chat/biz/model/user"
	"chat/common/tools"
	"chat/dal"
	"chat/dal/model"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

const JwtKey = "go-chat-server"
const IdentityKey = "go-chat-client"

var JwtMiddleware *jwt.HertzJWTMiddleware

func Init() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "jwt auth",
		Key:         []byte(JwtKey),
		Timeout:     time.Hour * 2,
		MaxRefresh:  time.Hour * 2,
		IdentityKey: IdentityKey,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			req := user.NewLoginRequest()
			if err := c.BindAndValidate(req); err != nil {
				return nil, err
			}
			password := tools.SHA256(req.GetPassword())
			data, err := dal.UserRepo.Login(context.Background(), req.GetName(), password)
			if err != nil {
				return nil, err
			}
			if data == nil {
				return nil, errors.New("用户不存在")
			}
			return data, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"user_id":   v.ID,
					"user_name": v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"token":      message,
				"expireTime": expire.Format(time.RFC3339),
				"BaseResp": utils.H{
					"StatusCode":    0,
					"StatusMessage": "",
				},
			})
		},
		TokenLookup:   "header:Authorization,query:token,cookie:jwt",
		TokenHeadName: "Bearer",
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "JWT biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"BaseResp": utils.H{
					"BaseCode": code,
					"Message":  message,
				},
			})
		},
		// IdentityHandler：用于设置获取身份信息的函数，在 demo 中，此处提取 token 的负载，并配合 IdentityKey 将用户名存入上下文信息。
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userID := claims["user_id"].(float64)
			id, _ := convertor.ToInt(userID)
			return &model.User{
				ID:   id,
				Name: claims["user_name"].(string),
			}
		},
	})
	if err != nil {
		panic(err)
	}
}
