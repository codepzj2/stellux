<script lang="ts" setup>
import darkLogo from "@/assets/logo/logo-dark.png";
import lightLogo from "@/assets/logo/logo-light.png";
import { useThemeStore } from "@/store/theme";
const themeStore = useThemeStore();
</script>

<template>
  <div class="flex h-full w-full items-center justify-between">
    <img
      :src="themeStore.tailwindTheme === 'dark' ? darkLogo : lightLogo"
      alt="logo"
      class="w-32"
    />
    <a-button type="primary" @click="themeStore.toggleTheme">切换主题</a-button>
  </div>
</template>
