import type {
  CreateUserReq,
  UpdateUserReq,
  LoginReq,
  LoginVO,
  UserInfoVO,
  UpdatePasswordReq,
} from "@/types/user";
import type { Response, PageReq, PageResponse } from "@/types/dto";
import request from "@/utils/request";

// 用户登录
export const userLoginAPI: (
  data: LoginReq
) => Promise<Response<LoginVO>> = data => {
  return request.post("/user/login", data);
};

// 通过access_token找到用户
export const getUserInfoAPI: () => Promise<Response<UserInfoVO>> = () => {
  return request.get("/admin-api/user/info");
};

// 更新用户密码
export const updateUserPasswordAPI: (
  data: UpdatePasswordReq
) => Promise<Response<UserInfoVO>> = data => {
  return request.put("/admin-api/user/update-password", data);
};

// 获取用户列表
export const getUserListAPI: (
  params: PageReq
) => Promise<PageResponse<UserInfoVO>> = params => {
  return request.get("/admin-api/user/list", { params });
};

// 新增用户
export const createUserAPI: (
  data: CreateUserReq
) => Promise<Response<UserInfoVO>> = data => {
  return request.post("/admin-api/user/create", data);
};

// 更新用户
export const updateUserAPI: (
  data: UpdateUserReq
) => Promise<Response<UserInfoVO>> = data => {
  return request.put("/admin-api/user/update", data);
};

// 删除用户
export const deleteUserAPI: (
  id: string
) => Promise<Response<UserInfoVO>> = id => {
  return request.delete(`/admin-api/user/delete/${id}`);
};
