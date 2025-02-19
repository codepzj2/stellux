<script lang="ts" setup>
import darkLogo from "@/assets/logo/logo-dark.png";
import lightLogo from "@/assets/logo/logo-light.png";
import { useThemeStore } from "@/store/theme";
import Sun from "@/assets/icons/sun.svg";
import Moon from "@/assets/icons/moon.svg";
const themeStore = useThemeStore();
</script>

<template>
  <div class="flex h-full w-full items-center justify-between">
    <img
      :src="themeStore.tailwindTheme === 'dark' ? darkLogo : lightLogo"
      alt="logo"
      class="w-32"
    />
    <div class="flex items-center gap-2">
      <a-avatar
        :size="28"
        @click="themeStore.toggleTheme"
        :src="themeStore.tailwindTheme === 'dark' ? Sun : Moon"
        :style="{
          backgroundColor:
            themeStore.tailwindTheme === 'dark' ? '#222' : '#ddd',
          margin: '0 8px',
        }"
      />
    </div>
  </div>
</template>
