import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import Login from "@/views/auth/login.vue";
import { errorView } from "@/router/modules/errorView";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: () => import("@/layout/index.vue"),
    name: "Layout",
    children: [
      {
        path: "",
        component: () => import("@/views/dashboard/index.vue"),
        name: "Dashboard",
      },
      {
        path: "user/list",
        component: () => import("@/views/user/list.vue"),
        name: "UserList",
      },
    ],
  },
  {
    path: "/login",
    component: Login,
    name: "Login",
  },
  {
    path: "/test",
    component: () => import("@/test/index.vue"),
    name: "Test",
  },
  {
    path: "/error",
    redirect: { name: "404" },
    children: errorView,
  },
  {
    path: "/:pathMatch(.*)",
    redirect: { name: "404" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
