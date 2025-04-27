import { message } from "ant-design-vue";

export const clearStorage = () => {
  localStorage.clear();
};

export const backToLogin = () => {
  clearStorage();
  window.location.href = "/login";
  message.warn("请重新登录");
};
