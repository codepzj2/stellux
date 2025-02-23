<template>
  <div class="">
    <a-layout>
      <a-layout-header
        :style="{
          height: '56px',
          background:
            themeStore.tailwindTheme === 'dark' ? '#00151d' : '#f2f3f4',
          padding: '0 15px',
        }"
      >
        <app-header></app-header>
      </a-layout-header>
      <a-layout style="min-height: calc(100vh - 56px); padding: 10px 15px">
        <a-layout-sider
          :style="{
            background:
              themeStore.tailwindTheme === 'dark' ? '#1d1d1d' : '#fff',
          }"
        >
          <Sidebar></Sidebar>
        </a-layout-sider>
        <a-layout>
          <a-layout-content class="ml-4">
            <tab></tab>
            <app-main></app-main>
          </a-layout-content>
          <a-layout-footer>
            <app-footer></app-footer>
          </a-layout-footer>
        </a-layout>
      </a-layout>
    </a-layout>
  </div>
</template>

<script setup lang="ts">
import Sidebar from "./components/SideBar/index.vue";
import Tab from "./components/Tab/index.vue";
import AppHeader from "./components/AppHeader.vue";
import AppMain from "./components/AppMain.vue";
import AppFooter from "./components/AppFooter.vue";
import { useThemeStore } from "@/store/theme";
const themeStore = useThemeStore();
</script>

<style lang="scss" scoped></style>
