import axios from "axios";
import { useUserStore } from "@/store/user";

const baseURL = import.meta.env.VITE_API_URL;

const request = axios.create({
  baseURL: baseURL,
  timeout: 5000,
});

// 请求拦截器
request.interceptors.request.use(
  config => {
    const accessToken = useUserStore().access_token;
    // 设置请求头
    config.headers.Authorization = `Bearer ${accessToken}`;
    return config;
  },
  error => Promise.reject(error)
);

// 响应拦截器
request.interceptors.response.use(
  response => response.data,
  async error => {
    if (!error.response) {
      return Promise.reject(error);
    }

    if (error.response.status === 401) {
      const userStore = useUserStore();
      const refresh_token = userStore.refresh_token;
      try {
        const res = await request.get(
          `/user/refresh?refresh_token=${refresh_token}`
        );
        console.log(res);
        const newAccessToken = res.data.access_token;
        const newRefreshToken = res.data.refresh_token;

        // 更新store中的token
        userStore.setAccessToken(newAccessToken);
        userStore.setRefreshToken(newRefreshToken);
        console.log(newAccessToken, newRefreshToken);
        // 使用新的access token重新发送请求
        error.config.headers["Authorization"] = `Bearer ${newAccessToken}`;
        console.log(error.config);
        return request(error.config);
      } catch (error) {
        return Promise.reject(error);
      }
    }

    let errMessage = "";
    switch (error.response.status) {
      case 400:
        errMessage = error.response.data.msg || "请求错误";
        break;
      case 401:
        errMessage = "登录过期，请重新登录";
        break;
      case 403:
        errMessage = "没有权限";
        break;
      case 500:
        errMessage = error.response.data.msg || "服务器错误";
        break;
      default:
        errMessage = "服务器内部错误";
        break;
    }
    return Promise.reject(error);
  }
);

export default request;
