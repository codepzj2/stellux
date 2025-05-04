import { createRouterGuard } from "./modules/guard";
import router, { routes, cacheNames } from "./modules/router";
import "nprogress/nprogress.css"; // 进度条样式

// 创建路由守卫
createRouterGuard(router);

export { routes, cacheNames };
export default router;
