package api

import (
	"log"
	"server/internal/pkg/http/resp"
	"server/internal/pkg/utils"
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
		userRouter.POST("/login", h.Login)
		userRouter.POST("/create", h.CreateUser)
		userRouter.GET("/all", h.FindAllUsers)
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var loginReq LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	user := h.LoginReqToUser(loginReq)
	if u, isExist := h.serv.FindIsExist(ctx, &user); isExist {
		token, _ := utils.GenerateJwt(u.ID.Hex())
		loginVo := LoginVO{
			User:  toUserVO(u),
			Token: token,
		}
		resp.SuccessWithDetail(ctx, loginVo, "登录成功")
		return
	}
	resp.FailWithMsg(ctx, http.StatusBadRequest, "用户名或者密码错误")
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserReq CreateUserReq
	if err := ctx.ShouldBindJSON(&createUserReq); err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	user := h.CreateUserReqToUser(createUserReq)
	if err := h.serv.CreateUser(ctx, &user); err != nil {
		resp.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessWithMsg(ctx, "创建用户成功")
}

func (h *UserHandler) LoginReqToUser(loginReq LoginReq) domain.User {
	return domain.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}
}

func (h *UserHandler) CreateUserReqToUser(createUserReq CreateUserReq) domain.User {
	return domain.User{
		Username: createUserReq.Username,
		Password: createUserReq.Password,
		RoleId:   createUserReq.RoleId,
	}
}

func (h *UserHandler) FindAllUsers(ctx *gin.Context) {
	val, isExist := ctx.Get("userId")
	if !isExist {
		resp.FailWithMsg(ctx, http.StatusUnauthorized, "请先登录")
		return
	}
	userId := val.(string)
	log.Println("userId:", userId)
	users, err := h.serv.FindAllUsers(ctx)
	if err != nil {
		resp.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessWithDetail(ctx, users, "查询成功")
}
