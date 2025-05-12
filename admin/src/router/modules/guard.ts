// 路由守卫
import type { Router } from "vue-router";
import NProgress from "nprogress";

NProgress.configure({ showSpinner: false });

const createRouterGuard = (router: Router) => {
  router.beforeEach((_to, _from, next) => {
    NProgress.start();
    next();
  });

  router.afterEach(() => {
    NProgress.done();
  });
};

export { createRouterGuard };
