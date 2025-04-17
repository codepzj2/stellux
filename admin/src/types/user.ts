export interface UserVO {
  id: string;
  username: string;
  role_id: number;
  avatar: string;
  email: string;
  sex: string;
  company: string;
  hobby: string;
}

export interface LoginReq {
  username: string;
  password: string;
}

export interface LoginVO {
  access_token: string; // 访问令牌
  refresh_token: string; // 刷新令牌
}
