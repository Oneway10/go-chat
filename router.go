// Code generated by hertz generator.

package main

import (
	"chat/biz/handler"
	"chat/common/auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", auth.JwtMiddleware.MiddlewareFunc(), handler.Ping)
	r.POST("/login", auth.JwtMiddleware.LoginHandler)
}
