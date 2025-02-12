import axios from "axios";
import SystemMessage from "./message";
const baseURL = import.meta.env.VITE_API_URL;

const request = axios.create({
  baseURL: baseURL,
  timeout: 10000,
});

// 添加请求拦截器
request.interceptors.request.use(
  function (config) {
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
request.interceptors.response.use(
  function (response) {
    console.log(response.data);

    return response.data;
  },
  function (error) {
    let errMessage = "";
    switch (error.response.status) {
      case 401:
        errMessage = "登录过期，请重新登录";
        break;
      case 403:
        errMessage = "没有权限";
        break;
      case 500:
        errMessage = error.response.data.msg;
        break;
      default:
        errMessage = "网络错误，请重试";
        break;
    }
    SystemMessage.error(errMessage);
    return Promise.reject(error);
  }
);

export default request;
