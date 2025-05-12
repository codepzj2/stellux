<template>
  <PostList v-model:page="page" type="bin" />
</template>
<script setup lang="ts">
import PostList from "./components/PostList.vue";
import type { PageReq } from "@/types/dto";

const page = reactive<PageReq>({
  page_no: 1,
  page_size: 10,
});
</script>
