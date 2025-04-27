<template>
  <a-drawer
    width="224"
    placement="left"
    :visible="visible"
    @close="onClose"
    :closable="false"
    :bodyStyle="{ padding: '0' }"
    :destroyOnClose="true"
  >
    <SideBar></SideBar>
  </a-drawer>
</template>
<script setup lang="ts">
import { useSidebarStore } from "@/store";
import SideBar from "./SideBar.vue";
import { useWindowSize } from "@vueuse/core";
const { width } = useWindowSize();
const sidebarStore = useSidebarStore();

const visible = computed(() => sidebarStore.collapse && width.value < 768);

const onClose = () => {
  console.log("onClose");
  sidebarStore.setCollapse(false);
};

watch(width, newVal => {
  if (newVal > 768) {
    sidebarStore.setCollapse(false);
  } else {
    sidebarStore.setCollapse(true);
  }
});
</script>
