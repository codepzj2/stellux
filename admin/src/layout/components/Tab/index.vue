<template>
  <div>
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      type="editable-card"
      @edit="tabStore.editTab"
      @change="tabStore.changeTab"
    >
      <a-tab-pane
        v-for="pane in panes"
        :key="pane.key"
        :tab="pane.title"
        :closable="pane.closable"
      >
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts" setup>
import { useTabStore } from "@/store/tab";
import { storeToRefs } from "pinia";
const tabStore = useTabStore();

const { panes, activeKey } = storeToRefs(tabStore);
</script>
