<template>
  <div class="flex h-full w-full items-center justify-between">
    <img
      :src="themeStore.tailwindTheme === 'dark' ? darkLogo : lightLogo"
      alt="logo"
      class="w-32"
    />
    <div class="flex items-center gap-2">
      <a-dropdown>
        <button
          class="p-2 rounded-full bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
          aria-label="User menu"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8"
            />
          </svg>
        </button>
        <template #overlay>
          <a-menu @click="onClick">
            <a-menu-item key="1">退出登录</a-menu-item>
            <a-menu-item key="2">设置</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>

      <!-- 主题切换按钮 -->
      <button
        @click="themeStore.toggleTheme"
        class="p-2 rounded-full bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
        aria-label="Toggle theme"
      >
        <svg
          class="w-6 h-6 hidden dark:block"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
          />
        </svg>
        <svg
          class="w-6 h-6 block dark:hidden"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import darkLogo from "@/assets/logo/logo-dark.png";
import lightLogo from "@/assets/logo/logo-light.png";
import { useThemeStore } from "@/store/theme";
import type { MenuProps } from "ant-design-vue";
import { clearStorage } from "@/utils/clearStorage";
import { message } from "ant-design-vue";
import { useRouter } from "vue-router";

const themeStore = useThemeStore();
const router = useRouter();

const onClick: MenuProps["onClick"] = ({ key }) => {
  switch (key) {
    case "1":
      clearStorage();
      message.warning("退出成功，请重新登录");
      router.push("/login");
      break;
    case "2":
      break;
  }
};
</script>
