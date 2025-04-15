package web

import (
	"net/http"

	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
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
		userGroup.POST("/login", apiwrap.WrapWithBody(h.Login))
	}
	adminGroup := engine.Group("/admin-api/user")
	{
		adminGroup.POST("/create", apiwrap.WrapWithBody(h.AdminCreateUser))
		adminGroup.PUT("/update", apiwrap.WrapWithBody(h.AdminUpdateUser))
		adminGroup.DELETE("/delete/:id", apiwrap.Wrap(h.AdminDeleteUser))
	}
}

func (h *UserHandler) Login(c *gin.Context, userRequest UserRequest) (*apiwrap.Response[any], error) {
	user := domain.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	exist := h.serv.CheckUserExist(c, &user)
	if !exist {
		return apiwrap.FailWithMsg(http.StatusBadRequest, "用户名或密码错误"), nil
	}
	return apiwrap.SuccessWithMsg[any](nil, "登录成功"), nil
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
