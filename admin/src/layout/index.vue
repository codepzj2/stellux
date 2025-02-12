<template>
  <div class="w-full flex px-4 py-6">
    <div class="w-1/5">
      <Sidebar></Sidebar>
    </div>
    <div class="w-4/5 h-screen">
      <Tab></Tab>
      <app-main></app-main>
      <app-footer></app-footer>
    </div>
  </div>
</template>

<script setup lang="ts">
import Sidebar from "./components/SideBar/index.vue";
import Tab from "./components/Tab/index.vue";
import AppMain from "./components/AppMain.vue";
import AppFooter from "./components/AppFooter.vue";
</script>

<style scoped></style>
