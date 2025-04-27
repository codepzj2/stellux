export const Code = {
  RequestSuccess: 0, // 请求成功
  RequestBadRequest: 1, // 请求参数错误
  RequestUnauthorized: 2, // 未授权
  RequestForbidden: 3, // 禁止访问
  RequestNotFound: 4, // 未找到
  RequestInternalServerError: 5, // 服务器错误
  RequestAccessTokenExpired: 6, // access_token已过期
  RequestAccessTokenNotFound: 7, // access_token未找到
  RequestRefreshTokenExpired: 8, // refresh_token已过期
  RequestRefreshTokenNotFound: 9, // refresh_token未找到
  RequestPermissionDenied: 10, // 权限不足
  RequestLoadPermissionFailed: 11, // 权限加载失败
};
