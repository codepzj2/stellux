import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import Layout from "@/layout/index.vue";
import Dashboard from "@/views/dashboard/index.vue";
import UserList from "@/views/user/list.vue";
import CreateContent from "@/views/content/create.vue";
import Login from "@/views/auth/login.vue";
import UserRole from "@/views/user/role.vue";
import LabelCategory from "@/views/label/category.vue";
import LabelTag from "@/views/label/tag.vue";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Layout,
    children: [
      {
        path: "",
        name: "Dashboard",
        component: Dashboard,
        meta: { title: "仪表盘", icon: "HomeOutlined" },
      },
      {
        path: "/user",
        name: "User",
        meta: { title: "用户管理", icon: "UserOutlined" },
        children: [
          {
            path: "list",
            component: UserList,
            name: "UserList",
            meta: { title: "用户列表", icon: "UnorderedListOutlined" },
          },
          {
            path: "role",
            component: UserRole,
            name: "UserRole",
            meta: { title: "用户角色", icon: "UsergroupAddOutlined" },
          },
        ],
      },
      {
        path: "/content",
        name: "Content",
        meta: { title: "内容管理", icon: "AppstoreOutlined" },
        children: [
          {
            path: "create",
            component: CreateContent,
            name: "CreateContent",
            meta: {
              title: "发布文章",
              icon: "FileTextOutlined",
              keepAlive: true,
            },
          },
        ],
      },
      {
        path: "/label",
        name: "Label",
        meta: { title: "标签分类", icon: "TagOutlined", hidden: true },
        children: [
          {
            path: "category",
            component: LabelCategory,
            name: "LabelCategory",
            meta: { title: "分类管理", icon: "CategoryOutlined", hidden: true },
          },
          {
            path: "tag",
            component: LabelTag,
            name: "LabelTag",
            meta: { title: "标签管理", icon: "TagOutlined" },
          },
        ],
      },
      {
        path: "/test/global",
        component: () => import("@/views/test/global/index.vue"),
        name: "GlobalTest",
        meta: { title: "全局测试", icon: "ExperimentOutlined" },
      },
    ],
  },
  {
    path: "/login",
    component: Login,
    name: "Login",
  },
  {
    path: "/error",
    redirect: { name: "404" },
    children: ["401", "403", "404", "500"].map(code => {
      return {
        path: code,
        name: code,
        component: () => import(`@/views/auth/error.vue`),
      };
    }),
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

function getCacheNames(routes: RouteRecordRaw[]): string[] {
  return routes.reduce((names: string[], route) => {
    if (route.meta?.keepAlive) {
      names.push(route.name as string);
    }
    if (route.children) {
      names.push(...getCacheNames(route.children));
    }
    return names;
  }, []);
}

export const cacheNames = getCacheNames(routes);

export default router;
