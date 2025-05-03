<template>
  <div
    class="flex h-full w-full items-center justify-between px-4 border-b"
    :class="{
      'border-b-gray-200': systemStore.themeMode === 'light',
      'border-b-gray-800': systemStore.themeMode === 'dark',
    }"
  >
    <div class="w-full">
      <div class="w-full flex justify-between">
        <div class="flex items-center gap-4">
          <MenuFoldOutlined
            class="cursor-pointer hover:bg-zinc-200 rounded-md p-1"
            v-if="sidebarStore.collapse"
            @click="sidebarStore.setCollapse(false)"
          />
          <MenuUnfoldOutlined
            class="cursor-pointer hover:bg-zinc-200 rounded-md p-1"
            v-else
            @click="sidebarStore.setCollapse(true)"
          />
          <Breadcrumb />
        </div>
      </div>
    </div>
    <div class="flex items-center gap-2 md:gap-4">
      <Search />
      <FullScreen />
      <Dropdown placement="bottomRight">
        <Avatar
          class="cursor-pointer"
          :src="userStore.userInfo?.avatar"
          :alt="userStore.userInfo?.username"
        >
          {{ userStore.userInfo?.nickname }}
        </Avatar>
        <template #overlay>
          <Menu>
            <Menu.Item>
              <div @click.prevent="Logout">
                <PoweroffOutlined />
                退出登录
              </div>
            </Menu.Item>
          </Menu>
        </template>
      </Dropdown>
      <Setting />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Breadcrumb, FullScreen, Search, Setting } from "./components";
import { useSystemStore } from "@/store/modules/system";
import { useUserStore } from "@/store/modules/user";
import { Menu, Dropdown, Avatar, message } from "ant-design-vue";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PoweroffOutlined,
} from "@ant-design/icons-vue";
import { clearStore } from "@/utils/clear";
import { useRouter } from "vue-router";
import { useSidebarStore } from "@/store/modules/sidebar";
const systemStore = useSystemStore();
const sidebarStore = useSidebarStore();
const userStore = useUserStore();
const router = useRouter();

function Logout() {
  clearStore();
  message.info("退出成功,请重新登录");
  router.push("/login");
}
</script>
