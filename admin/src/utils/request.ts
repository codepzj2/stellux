import axios from "axios";
import { useUserStore } from "@/store";
import { message } from "ant-design-vue";
import { Code } from "@/global";
import { clearStore } from "./clear";
import router from "@/router";
import { API_BASE_URL } from "@/constant";

const request = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
});

// 请求拦截器
request.interceptors.request.use(
  config => {
    const userStore = useUserStore();
    const accessToken = userStore.access_token;
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
  },
  error => Promise.reject(error)
);

// 响应拦截器
request.interceptors.response.use(
  async response => {
    const userStore = useUserStore();
    const refresh_token = userStore.refresh_token;

    if (response.data.code === Code.RequestAccessTokenExpired) {
      if (refresh_token) {
        try {
          const res = await request.get(
            `/user/refresh?refresh_token=${refresh_token}`
          );
          const newAccessToken = res.data.access_token;
          const newRefreshToken = res.data.refresh_token;

          userStore.setAccessToken(newAccessToken);
          userStore.setRefreshToken(newRefreshToken);

          response.config.headers["Authorization"] = `Bearer ${newAccessToken}`;
          return request(response.config);
        } catch (refreshError) {
          clearStore();
          message.warning("登录已过期，请重新登录");
          router.push("/login");
          return Promise.reject(refreshError);
        }
      } else {
        clearStore();
        message.warning("令牌已过期，请重新登录");
        router.push("/login");
        return Promise.reject(new Error("access_token已过期,请重新登录"));
      }
    }

    if (response.data.code !== Code.RequestSuccess) {
      console.log(response.data);
      const errMessage = response.data.msg || "操作失败";
      message.error(errMessage);
      return Promise.reject(new Error(errMessage));
    }

    return response.data;
  },
  error => {
    if (!error.response) {
      message.error("网络错误，请稍后重试");
    } else {
      message.error(error.response.data.msg || "操作失败");
    }
    return Promise.reject(error);
  }
);

export default request;
