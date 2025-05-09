<template>
  <a-layout class="h-screen flex flex-col overflow-hidden">
    <a-layout-sider
      width="216"
      :trigger="null"
      collapsible
      v-model:collapsed="sidebarStore.collapse"
    >
      <!-- 侧边栏内容 -->
      <SideBar :collapsed="sidebarStore.collapse" />
    </a-layout-sider>

    <a-layout class="flex-1">
      <a-layout-header class="!h-16 !px-0">
        <Header />
      </a-layout-header>
      <a-layout-content class="h-full px-4 pt-4 overflow-y-auto overflow-x-hidden">
        <Main />
        <Footer />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import SideBar from "./sidebar/index.vue";
import Header from "./header/index.vue";
import Main from "./main/index.vue";
import Footer from "./footer/index.vue";
import { useSidebarStore, useUserStore } from "@/store";
import { getUserInfoAPI } from "@/api/user";
import { useWindowSize } from "@vueuse/core";

// 判断是否为移动端设备
const { width } = useWindowSize();

const userStore = useUserStore();
const { userInfo } = storeToRefs(userStore);

const sidebarStore = useSidebarStore();

// 加载界面时初始化用户信息
const getUserInfo = async () => {
  const res = await getUserInfoAPI();
  userInfo.value = res.data;
};

onMounted(async () => {
  await getUserInfo();
});

watch(width, newWidth => {
  sidebarStore.setCollapse(newWidth < 768);
});
</script>

<style lang="scss"></style>
