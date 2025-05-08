<template>
  <div>
    <a-tabs v-model:activeKey="activeKey" type="card">
      <a-tab-pane tab="文章分类" key="category" />
      <a-tab-pane tab="文章标签" key="tag" />
    </a-tabs>
    <LabelTable :activeKey="activeKey" />
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import LabelTable from "./LabelTable.vue";

const activeKey = ref<"category" | "tag">("category");
</script>
