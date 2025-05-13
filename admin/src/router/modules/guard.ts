import type { Router } from "vue-router";
import NProgress from "nprogress";
import { useUserStore } from "@/store";

NProgress.configure({ showSpinner: false });

const createRouterGuard = (router: Router) => {
  router.beforeEach((to, _from, next) => {
    NProgress.start();
    const userStore = useUserStore();

    const isLoginPage = to.name === "Login";
    const isAuthenticated = !!userStore.access_token;

    if (!isAuthenticated && !isLoginPage) {
      next({ name: "Login" });
    } else if (isAuthenticated && isLoginPage) {
      // 如果已登录还访问 login 页面，可选择跳转到首页或允许访问
      next({ name: "Dashboard" });
    } else {
      next();
    }
  });

  router.afterEach(() => {
    NProgress.done();
  });
};

export { createRouterGuard };
