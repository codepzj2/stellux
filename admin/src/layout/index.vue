<template>
  <div class="">
    <a-layout>
      <a-layout-header class="stellux-header">
        <app-header></app-header>
      </a-layout-header>
      <a-layout>
        <a-layout-sider>
          <Sidebar></Sidebar>
        </a-layout-sider>
        <a-layout-content>
          <tab></tab>
          <app-main></app-main>
        </a-layout-content>
        <a-layout-footer>
          <app-footer></app-footer>
        </a-layout-footer>
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
</script>

<style lang="scss" scoped>
.stellux-header {
  background: #fff !important;
  padding: 0 20px !important;
}
.dark .stellux-header {
  background: #000 !important;
}
</style>
