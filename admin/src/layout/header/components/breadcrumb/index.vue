<script setup lang="ts">
import { computed } from "vue";
import { useRoute, useRouter, type RouteRecordRaw } from "vue-router";

defineOptions({
  name: "BreadCrumb",
});

const router = useRouter();
const route = useRoute();

// 点击菜单
// 点击菜单
const clickMenuItem = (menuItem: RouteRecordRaw) => {
  const { isExt, extOpenMode, type } = menuItem?.meta || {};

  // 如果点击的是首页并且没有子路由，跳转到仪表盘
  if (menuItem.name === "__index") {
    router.push({ name: "Dashboard" });
    return;
  }

  if (type === 0 && !menuItem.redirect) return;

  if (isExt && extOpenMode === 1) {
    window.open(menuItem.path);
  } else {
    const to =
      typeof menuItem.redirect === "string" ? menuItem.redirect : menuItem;
    router.push(to);
  }
};

// 使用 route.matched 构建面包屑
const menus = computed(() => {
  return [
    ...route.matched
      .filter(r => r.meta?.title && r.meta.hidden !== true)
      .map(routeItem => {
        // 对 children 也做 hidden 过滤
        const children = (routeItem.children || []).filter(
          c => c.meta?.hidden !== true
        );
        return { ...routeItem, children };
      }),
  ];
});

// 获取选中子菜单项
const getSelectKeys = (routeIndex: number) => {
  return [menus.value[routeIndex + 1]?.name] as string[];
};
</script>

<template>
  <div>
    <a-breadcrumb>
      <template v-for="(routeItem, routeIndex) in menus" :key="routeItem?.name">
        <a-breadcrumb-item>
          <span class="cursor-pointer">
            {{ routeItem?.meta?.title }}
          </span>
          <!-- 下拉菜单（如果有子路由） -->
          <template v-if="routeItem?.children?.length" #overlay>
            <a-menu :selected-keys="getSelectKeys(routeIndex)">
              <template
                v-for="childItem in routeItem?.children"
                :key="childItem.name"
              >
                <a-menu-item
                  v-if="
                    !childItem.meta?.hideInMenu &&
                    !childItem.meta?.hideInBreadcrumb
                  "
                  :key="childItem.name"
                  @click="clickMenuItem(childItem)"
                >
                  {{ childItem.meta?.title }}
                </a-menu-item>
              </template>
            </a-menu>
          </template>
        </a-breadcrumb-item>
      </template>
    </a-breadcrumb>
  </div>
</template>

<style scoped lang="less"></style>
