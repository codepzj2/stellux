import type { LoginForm, LoginVO, UserListVO } from "@/api/interfaces/user";
import type { Response } from "@/api/interfaces/resp";
import request from "@/utils/axios";

export const userLogin: (data: LoginForm) => Promise<Response<LoginVO>> = (
  data
) => {
  return request.post("/user/login", data);
};

export const getUserList: () => Promise<Response<UserListVO>> = () => {
  return request.get("/user/list");
};
