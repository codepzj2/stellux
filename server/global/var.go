package global

import (
	"flag"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 全局变量
var (
	Enforcer *casbin.Enforcer
)

// 命令行参数
var (
	Mode = flag.String("mode", "development", "运行模式,eg: development/production")
)

// 常量
var (
	AccessTokenNotFound  = gin.H{"code": http.StatusUnauthorized, "msg": "未携带access_token"}
	AccessTokenExpired   = gin.H{"code": http.StatusUnauthorized, "msg": "access_token已过期"}
	LoadPermissionFailed = gin.H{"code": http.StatusInternalServerError, "msg": "权限加载失败"}
	PermissionDenied     = gin.H{"code": http.StatusForbidden, "msg": "禁止访问"}
)
