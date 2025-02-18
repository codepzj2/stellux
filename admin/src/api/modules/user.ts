import type { LoginForm, LoginVO } from "@/api/interfaces/user";
import type { Response } from "@/api/interfaces/resp";
import request from "@/utils/axios";

export const userLogin: (data: LoginForm) => Promise<Response<LoginVO>> = (
  data
) => {
  return request.post("/user/login", data);
};
