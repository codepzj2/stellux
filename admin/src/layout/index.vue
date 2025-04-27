<template>
  <a-layout class="h-screen flex flex-col overflow-hidden">
    <a-layout-sider
      class="hidden"
      :class="{
        'md:block': !sidebarStore.collapse, // 根据 sidebarStore 的 collapse 状态动态显示
      }"
      width="224"
    >
      <!-- 侧边栏内容 -->
      <SideBar />
    </a-layout-sider>

    <a-layout class="flex-1 overflow-scroll">
      <a-layout-header class="!h-12 !px-4">
        <Header />
      </a-layout-header>
      <div class="h-[1px]"></div>
      <a-layout class="mt-2 mx-2 my-0">
        <a-layout-content>
          <Main class="h-full overflow-y-scroll"></Main>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
  <!-- 侧边栏抽屉 -->
  <SiderBarDrawer />
</template>

<script setup lang="ts">
import SideBar from "./components/SideBar.vue";
import SiderBarDrawer from "./components/SiderBarDrawer.vue";
import Header from "./components/Header.vue";
import Main from "./components/Main.vue";
import { useSidebarStore, useUserStore } from "@/store";
import { getUserInfoAPI } from "@/api/user";
const sidebarStore = useSidebarStore();

const userStore = useUserStore();
const { userInfo } = storeToRefs(userStore);

// 加载界面时初始化用户信息
const getUserInfo = async () => {
  const res = await getUserInfoAPI();
  userInfo.value = res.data;
};

onMounted(async () => {
  await getUserInfo();
});
</script>

<style lang="scss"></style>
