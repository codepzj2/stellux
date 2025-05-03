<template>
  <a-tabs
    v-model:activeKey="activeKey"
    hide-add
    animated
    type="editable-card"
    size="small"
    @edit="tabStore.editTab"
    @change="tabStore.changeTab"
  >
    <a-tab-pane
      v-for="pane in panes"
      :key="pane.key"
      :tab="pane.title"
      :closable="pane.closable"
    >
      <router-view v-slot="{ Component }">
        <keep-alive :include="cacheNames">
          <component :is="Component" />
        </keep-alive>
      </router-view>
    </a-tab-pane>
  </a-tabs>
</template>
<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { cacheNames } from "@/router";
import { useTabStore } from "@/store";

const tabStore = useTabStore();
const { panes, activeKey } = storeToRefs(tabStore);
</script>
<style lang="scss" scoped></style>
