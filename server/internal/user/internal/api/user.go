package api

import (
	"server/internal/pkg/http/resp"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	serv service.IUserService
}

func NewUserHandler(serv service.IUserService) *UserHandler {
	return &UserHandler{serv}
}

func (h *UserHandler) RegisterGinRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/create", h.CreateUser)
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var loginReq LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	user := domain.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
		RoleId:   1, // 默认创建普通用户
	}
	if err := h.serv.CreateUser(ctx, &user); err != nil {
		resp.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessWithMsg(ctx, "创建用户成功")
}
