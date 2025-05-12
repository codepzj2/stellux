export interface UserInfoVO {
  id: string;
  username: string;
  nickname: string;
  role_id: number;
  avatar: string;
  email: string;
}

export interface LoginReq {
  username: string;
  password: string;
}

export interface LoginVO {
  access_token: string; // 访问令牌
  refresh_token: string; // 刷新令牌
}

export interface CreateUserReq {
  username: string;
  password: string;
  nickname: string;
  role_id: number;
  avatar: string;
  email: string;
}

export interface UpdateUserReq {
  id: string;
  nickname: string;
  avatar: string;
  email: string;
}

export interface UpdatePasswordReq {
  id: string;
  old_password: string;
  new_password: string;
}
