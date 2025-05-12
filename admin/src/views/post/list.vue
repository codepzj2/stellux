<template>
  <PostList v-model:page="page" type="publish" />
</template>
<script setup lang="ts">
import type { PageReq } from "@/types/dto";
import PostList from "./components/PostList.vue";

const page = reactive<PageReq>({
  page_no: 1,
  page_size: 10,
});
</script>
