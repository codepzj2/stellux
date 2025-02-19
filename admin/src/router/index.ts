import router from "./router";
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

export default router;
