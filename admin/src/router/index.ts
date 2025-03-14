import router from "./router";
import { useUserStore } from "@/store/user";
import { useTabStore } from "@/store/tab";
import { routeNames } from "./modules/name";

// 监听标签栏
router.beforeEach((to, _from, next) => {
  const tabStore = useTabStore();
  const index = Object.keys(routeNames).indexOf(to.name as string);
  if (index !== -1) {
    tabStore.addTab({
      key: to.name as string,
      title: routeNames[to.name as keyof typeof routeNames],
    });
  }
  next();
});

// 监听用户是否登录
router.beforeEach((to, _from, next) => {
  const userStore = useUserStore();
  if (to.path === "/login") {
    next();
  } else {
    if (userStore.token) {
      next();
    } else {
      next({ path: "/login" });
    }
  }
});

export default router;
