package web

import (
	"net/http"
	"time"

	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/pkg/utils"
	"github.com/codepzj/stellux/server/internal/user/internal/domain"
	"github.com/codepzj/stellux/server/internal/user/internal/service"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(serv service.IUserService) *UserHandler {
	return &UserHandler{
		serv: serv,
	}
}

type UserHandler struct {
	serv service.IUserService
}

func (h *UserHandler) RegisterGinRoutes(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/refresh", apiwrap.Wrap(h.RefreshToken))
		userGroup.POST("/login", apiwrap.WrapWithBody(h.Login))
	}
	adminGroup := engine.Group("/admin-api/user")
	{
		adminGroup.POST("/create", apiwrap.WrapWithBody(h.AdminCreateUser))
		adminGroup.PUT("/update", apiwrap.WrapWithBody(h.AdminUpdateUser))
		adminGroup.DELETE("/delete/:id", apiwrap.Wrap(h.AdminDeleteUser))
		adminGroup.GET("/list", apiwrap.WrapWithBody(h.AdminGetUserList))
	}
}

func (h *UserHandler) Login(c *gin.Context, userRequest UserRequest) (*apiwrap.Response[any], error) {
	user := domain.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	exist, id := h.serv.CheckUserExist(c, &user)
	if !exist {
		return apiwrap.FailWithMsg(http.StatusBadRequest, "用户名或密码错误"), nil
	}
	accessToken, err := utils.GenerateAccessToken(id)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	refreshToken, err := utils.GenerateRefreshToken(id)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	loginVO := LoginVO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return apiwrap.SuccessWithMsg[any](loginVO, "登录成功"), nil
}

func (h *UserHandler) RefreshToken(c *gin.Context) (*apiwrap.Response[any], error) {
	// 校验refresh_token是否有效
	refreshToken := c.Query("refresh_token")
	if refreshToken == "" {
		return apiwrap.FailWithMsg(http.StatusBadRequest, "refresh_token不能为空"), nil
	}
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusUnauthorized, "refresh_token已过期,请重新登录"), err
	}
	accessToken, err := utils.GenerateAccessToken(claims.ID)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	// 若refresh_token临近过期，则生成新的refresh_token
	if claims.ExpiresAt.Before(time.Now().Add(time.Hour * 24 * 7)) {
		var err error
		refreshToken, err = utils.GenerateRefreshToken(claims.ID)
		if err != nil {
			return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
		}
	}
	return apiwrap.SuccessWithMsg[any](LoginVO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, "刷新token成功"), nil
}

func (h *UserHandler) AdminCreateUser(c *gin.Context, userRequest UserRequest) (*apiwrap.Response[any], error) {
	user := domain.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	err := h.serv.AdminCreate(c, &user)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	return apiwrap.SuccessWithMsg[any](nil, "创建用户成功"), nil
}

func (h *UserHandler) AdminUpdateUser(c *gin.Context, updateUserRequest UpdateUserRequest) (*apiwrap.Response[any], error) {
	user := domain.User{
		ID:       updateUserRequest.ID,
		Username: updateUserRequest.Username,
		Password: updateUserRequest.Password,
		RoleId:   updateUserRequest.RoleId,
		Avatar:   updateUserRequest.Avatar,
		Email:    updateUserRequest.Email,
		Sex:      updateUserRequest.Sex,
		Company:  updateUserRequest.Company,
		Hobby:    updateUserRequest.Hobby,
	}
	err := h.serv.AdminUpdate(c, &user)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	return apiwrap.SuccessWithMsg[any](nil, "更新用户成功"), nil
}

func (h *UserHandler) AdminDeleteUser(c *gin.Context) (*apiwrap.Response[any], error) {
	id := c.Param("id")
	err := h.serv.AdminDelete(c, id)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	return apiwrap.SuccessWithMsg[any](nil, "删除用户成功"), nil
}

func (h *UserHandler) AdminGetUserList(c *gin.Context, page apiwrap.Page) (*apiwrap.Response[any], error) {
	users, count, err := h.serv.GetUserList(c, &domain.Page{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
	})
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	return apiwrap.SuccessWithMsg[any](apiwrap.ToPageVO(page.PageNo, page.PageSize, count, h.DomainToVOList(users)), "获取用户列表成功"), nil
}
