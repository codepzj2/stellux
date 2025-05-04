<template>
  <div class="flex flex-col h-full">
    <!-- logo -->
    <div
      :class="`flex items-center ${props.collapsed ? 'justify-center' : 'justify-start'}`"
    >
      <img
        class="transition-all duration-300 ease-in-out"
        :src="`/logo${props.collapsed ? '-sm' : ''}-${systemStore.themeMode === 'dark' ? 'dark' : 'light'}.png`"
        alt="logo"
        :class="`my-[15px] ${props.collapsed ? 'mx-auto' : 'ml-4'} ${props.collapsed ? 'size-8' : 'w-32'}`"
      />
    </div>

    <!-- 菜单 -->
    <a-menu
      v-model:selectedKeys="sidebarStore.selectedKeys"
      mode="inline"
      :open-keys="sidebarStore.openKeys"
      :items="menuItems"
      @openChange="onOpenChange"
      @select="onSelect"
    />
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useRouter } from "vue-router";
import { routes } from "@/router";
import { useSidebarStore, useSystemStore } from "@/store";
import type { VNodeChild } from "vue";
import type { ItemType } from "ant-design-vue";

const sidebarStore = useSidebarStore();
const systemStore = useSystemStore();

const props = defineProps<{ collapsed: boolean }>();

type MenuItemType = {
  key: string;
  label: string;
  children?: MenuItemType[];
  icon?: () => VNodeChild;
  type?: "group";
} & Omit<ItemType, "key" | "label" | "children" | "icon" | "type">;

function getItem(
  label: string,
  key: string,
  icon?: () => VNodeChild,
  children?: MenuItemType[],
  type?: "group"
): MenuItemType {
  return {
    key,
    icon,
    children,
    label,
    type,
  };
}

const menuItems = computed(() => {
  const mainRoute = routes.find((r): r is typeof r => r.path === "/");
  if (!mainRoute || !mainRoute.children) return [];

  const generateMenuItems = (
    route: (typeof mainRoute.children)[number]
  ): MenuItemType | null => {
    if (!route.name || !route.meta?.title || route.meta.hidden === true) {
      return null;
    }

    const childrenRoutes = route.children?.filter(
      (child): child is typeof child =>
        !!child.name && !!child.meta?.title && child.meta.hidden !== true
    );

    const itemChildren = childrenRoutes
      ?.map(child => generateMenuItems(child))
      .filter((item): item is MenuItemType => item !== null);

    return getItem(
      route.meta.title as string,
      route.name.toString(),
      route.meta.icon as (() => VNodeChild) | undefined,
      itemChildren?.length ? itemChildren : undefined
    );
  };

  return mainRoute.children
    .map(route => generateMenuItems(route))
    .filter((item): item is MenuItemType => item !== null);
});

const rootSubmenuKeys = computed(() =>
  menuItems.value.filter(item => !!item.children).map(item => item.key)
);

const router = useRouter();

const onSelect = ({ key }: { key: string }) => {
  sidebarStore.setSelectedKeys([key]);
  router.push({ name: key });
};

const onOpenChange = (openKeys: string[]) => {
  const latestOpenKey = openKeys.find(
    key => sidebarStore.openKeys.indexOf(key) === -1
  );
  if (latestOpenKey && !rootSubmenuKeys.value.includes(latestOpenKey)) {
    sidebarStore.setOpenKeys(openKeys);
  } else {
    sidebarStore.setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
  }
};
</script>
