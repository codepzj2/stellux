<template>
  <MdEditor class="h-[calc(100vh-180px)]" v-model="text" :theme="theme" />
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useThemeStore } from "@/store/theme";

const themeStore = useThemeStore();
const theme = computed(() => themeStore.tailwindTheme);
const text = ref("Hello Editor!");
</script>

<style lang="scss">
.md-editor-toolbar-wrapper {
  height: 40px;
  .md-editor-toolbar {
    height: 100%;
    .md-editor-toolbar-left,
    .md-editor-toolbar-right {
      height: 100%;
      .md-editor-toolbar-item {
        height: 100%;
        svg {
          width: 32px;
          height: 32px;
        }
      }
    }
  }
}
</style>
