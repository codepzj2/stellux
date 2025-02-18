<template>
  <div class="w-full flex p-4 space-x-4">
    <div class="w-48">
      <Sidebar></Sidebar>
    </div>
    <div class="h-full">
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
