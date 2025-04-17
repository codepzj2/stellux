import type { LoginReq, LoginVO, UserVO } from "@/types/user";
import type { Response, PageReq, PageResponse } from "@/types/response";
import request from "@/utils/request";

// 用户登录
export const userLogin: (
  data: LoginReq
) => Promise<Response<LoginVO>> = data => {
  return request.post("/user/login", data);
};

// 获取用户列表
export const getUserList: (
  params: PageReq
) => Promise<PageResponse<UserVO>> = params => {
  return request.get("/admin-api/user/list", { params });
};
