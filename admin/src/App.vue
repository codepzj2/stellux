<template>
  <a-config-provider
    :theme="{
      algorithm:
        themeStore.tailwindTheme === 'dark'
          ? theme.darkAlgorithm
          : theme.defaultAlgorithm,
      token: {
        fontSize: 16,
        wireframe: true,
      },
    }"
  >
    <a-style-provider hash-priority="high">
      <div class="h-screen bg-white dark:bg-gray-900">
        <router-view></router-view>
      </div>
    </a-style-provider>
  </a-config-provider>
</template>

<script setup lang="ts">
import { useThemeStore } from "@/store/theme";
import { theme } from "ant-design-vue";
const themeStore = useThemeStore();
themeStore.initTheme();
</script>

<style lang="scss"></style>
