<template>
  <div class="flex flex-col h-full">
    <!-- 显示应用程序的logo -->
    <img
      :src="systemStore.darkMode ? '/logo-dark.png' : '/logo-light.png'"
      alt="logo"
      class="w-32 my-2 ml-4"
    />
    <!-- 侧边栏菜单 -->
    <a-menu
      v-model:selectedKeys="sidebarStore.selectedKeys"
      mode="inline"
      :open-keys="sidebarStore.openKeys"
      :items="menuItems"
      @openChange="onOpenChange"
      @select="onSelect"
    >
    </a-menu>
    <div
      class="mt-auto flex items-center justify-between gap-x-1 px-1 py-3 font-semibold"
    >
      <div class="flex items-center gap-x-1">
        <img
          class="size-12 rounded-full"
          :src="userInfo?.avatar"
          alt="avatar"
        />
        <div v-if="userInfo" class="flex flex-col gap-x-1">
          <span class="text-sm pl-1">{{ userInfo.nickname }}</span>
          <a-tag
            size="small"
            :bordered="false"
            :color="roleColors[userInfo.role_id as keyof typeof roleColors]"
          >
            {{ roleNames[userInfo.role_id as keyof typeof roleNames] }}
          </a-tag>
        </div>
      </div>

      <div class="flex items-center">
        <a-button size="small" type="text" @click="changeTheme">
          <div class="flex items-center justify-center">
            <img
              v-if="systemStore.darkMode"
              src="@/assets/sun.svg"
              alt="theme"
            />
            <img v-else src="@/assets/moon.svg" alt="theme" />
          </div>
        </a-button>
        <a-button size="small" type="text">
          <SettingOutlined />
        </a-button>
        <a-popconfirm
          placement="bottomRight"
          title="确定退出吗？"
          ok-text="确定"
          cancel-text="取消"
          @confirm="backToLogin"
        >
          <a-button size="small" type="text">
            <LogoutOutlined />
          </a-button>
        </a-popconfirm>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { routes } from "@/router";
import { useSidebarStore, useSystemStore, useUserStore } from "@/store";
import type { ItemType } from "ant-design-vue";
import type { RouteRecordRaw } from "vue-router";
import type { VNodeChild } from "vue";
import * as Icons from "@ant-design/icons-vue"; // 动态导入所有图标
import { LogoutOutlined, SettingOutlined } from "@ant-design/icons-vue";
import { roleNames, roleColors } from "@/global";
import { backToLogin } from "@/utils/clear";
const sidebarStore = useSidebarStore();
const systemStore = useSystemStore();
const userStore = useUserStore();

const changeTheme = () => {
  systemStore.setDarkMode(!systemStore.darkMode);
  document.documentElement.classList.toggle("dark", systemStore.darkMode);
};
const { userInfo } = storeToRefs(userStore);

// 定义菜单项类型
type MenuItemType = {
  key: string;
  label: string;
  children?: MenuItemType[];
  icon?: () => VNodeChild;
  type?: "group";
} & Omit<ItemType, "key" | "label" | "children" | "icon" | "type">;

// 创建菜单项的辅助函数
function getItem(
  label: string,
  key: string,
  icon?: string,
  children?: MenuItemType[],
  type?: "group"
): MenuItemType {
  return {
    key,
    icon:
      icon && Icons[icon as keyof typeof Icons]
        ? () => h(Icons[icon as keyof typeof Icons])
        : undefined,
    children,
    label,
    type,
  };
}

// 计算属性：生成菜单项
const menuItems = computed(() => {
  const mainRoute = routes.find((r): r is RouteRecordRaw => r.path === "/");
  if (!mainRoute || !mainRoute.children) return [];

  // 递归生成菜单项
  const generateMenuItems = (route: RouteRecordRaw): MenuItemType | null => {
    if (!route.name || !route.meta?.title || route.meta.hidden === true) {
      return null;
    }

    const hasComponent = !!route.component;
    const childrenRoutes = route.children?.filter(
      (child): child is RouteRecordRaw =>
        !!child.name && !!child.meta?.title && child.meta.hidden !== true
    );

    const itemChildren = childrenRoutes
      ?.map(child => generateMenuItems(child))
      .filter((item): item is MenuItemType => item !== null);

    if (!hasComponent && (!itemChildren || itemChildren.length === 0)) {
      return null;
    }

    return getItem(
      route.meta.title as string,
      route.name.toString(),
      route.meta.icon as string | undefined,
      itemChildren?.length ? itemChildren : undefined
    );
  };

  return mainRoute.children
    .map(route => generateMenuItems(route))
    .filter((item): item is MenuItemType => item !== null);
});

// 计算属性：获取根菜单项的key
const rootSubmenuKeys = computed(() =>
  menuItems.value
    .filter((item): item is MenuItemType => !!item.children)
    .map(item => item.key)
);

const router = useRouter();

// 菜单项被选中时的处理函数
const onSelect = ({ key }: { key: string }) => {
  sidebarStore.setSelectedKeys([key]);
  router.push({ name: key });
};

// 菜单展开状态改变时的处理函数
const onOpenChange = (openKeys: string[]) => {
  const latestOpenKey = openKeys.find(
    key => sidebarStore.openKeys.indexOf(key) === -1
  );
  if (
    typeof latestOpenKey === "string" &&
    rootSubmenuKeys.value.indexOf(latestOpenKey) === -1
  ) {
    sidebarStore.setOpenKeys(openKeys);
  } else {
    sidebarStore.setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
  }
};
</script>
