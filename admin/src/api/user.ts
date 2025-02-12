import request from "@/utils/axios";

export interface LoginReq {
  username: string;
  password: string;
}

export interface LoginVO {
  token: string;
  session: string;
}

export const login = (data: LoginReq) =>
  request({
    url: "/api/user/login",
    method: "post",
    data,
  });
