import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import Layout from "@/layout/index.vue";
import Dashboard from "@/views/dashboard/index.vue";
import CreateContent from "@/views/content/create.vue";
import Login from "@/views/auth/login.vue";
import LabelCategory from "@/views/label/category.vue";
import LabelTag from "@/views/label/tag.vue";
import {
  TagOutlined,
  HomeOutlined,
  UserOutlined,
  AppstoreOutlined,
  FileTextOutlined,
  ExperimentOutlined,
} from "@ant-design/icons-vue";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Layout,
    children: [
      {
        path: "",
        name: "Dashboard",
        component: Dashboard,
        meta: { title: "仪表盘", icon: () => h(HomeOutlined) },
      },
      {
        path: "/user",
        name: "UserManagement",
        meta: { title: "用户管理", icon: () => h(UserOutlined) },
        component: () => import("@/views/user/list.vue"),
      },
      {
        path: "/content",
        name: "Content",
        meta: { title: "内容管理", icon: () => h(AppstoreOutlined) },
        children: [
          {
            path: "create",
            component: CreateContent,
            name: "CreateContent",
            meta: {
              title: "发布文章",
              icon: () => h(FileTextOutlined),
              keepAlive: true,
            },
          },
        ],
      },
      {
        path: "label",
        name: "Label",
        meta: { title: "标签管理", icon: () => h(TagOutlined) },
        children: [
          {
            path: "category",
            component: LabelCategory,
            name: "LabelCategory",
            meta: { title: "文章分类", icon: () => h(TagOutlined) },
          },
          {
            path: "tag",
            component: LabelTag,
            name: "LabelTag",
            meta: { title: "文章标签", icon: () => h(TagOutlined) },
          },
        ],
      },

      {
        path: "/test",
        component: () => import("@/views/test/index.vue"),
        name: "Test",
        meta: { title: "测试", icon: () => h(ExperimentOutlined) },
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
