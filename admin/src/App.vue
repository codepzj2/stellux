<template>
  <a-config-provider
    :theme="{
      algorithm:
        themeStore.tailwindTheme === 'dark'
          ? theme.darkAlgorithm
          : theme.defaultAlgorithm,
      token: {
        fontSize: 15,
        borderRadius: 6,
        wireframe: false,
      },
    }"
  >
    <div class="h-screen dark:bg-gray-900">
      <router-view></router-view>
    </div>
  </a-config-provider>
</template>

<script setup lang="ts">
import { useThemeStore } from "@/store/theme";

import { theme } from "ant-design-vue";
import { onMounted } from "vue";
const themeStore = useThemeStore();
onMounted(() => {
  themeStore.initTheme();
  console.log(themeStore.tailwindTheme);
});
</script>

<style lang="scss"></style>
