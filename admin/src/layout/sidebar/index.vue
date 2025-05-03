<template>
  <div class="flex flex-col h-full">
    <!-- 显示应用程序的logo -->
    <img
      v-if="!props.collapsed"
      :src="
        systemStore.themeMode === 'dark' ? '/logo-dark.png' : '/logo-light.png'
      "
      alt="logo"
      class="w-32 my-2 ml-4"
    />
    <img
      v-else
      :src="
        systemStore.themeMode === 'dark'
          ? '/logo-sm-dark.png'
          : '/logo-sm-light.png'
      "
      alt="logo"
      class="size-8 my-2 mx-auto"
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
  </div>
</template>

<script lang="ts" setup>
import { routes } from "@/router";
import { useSidebarStore, useSystemStore } from "@/store";
import type { ItemType } from "ant-design-vue";
import type { RouteRecordRaw } from "vue-router";
import type { VNodeChild } from "vue";

import * as Icons from "@ant-design/icons-vue"; // 动态导入所有图标
const sidebarStore = useSidebarStore();
const systemStore = useSystemStore();

const props = defineProps<{
  collapsed: boolean;
}>();

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
