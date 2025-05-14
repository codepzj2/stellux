<template>
  <div>
    <PostContentWriter mode="create" v-model:post-form="postForm" />
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: "PostCreate" });
import PostContentWriter from "./components/PostContentWriter.vue";
import type { PostReq } from "@/types/post";

const postForm = ref<PostReq>({
  created_at: new Date().toISOString(),
  title: "",
  content: "",
  description: "",
  author: "",
  category_id: "",
  tags_id: [],
  is_publish: true,
  is_top: false,
  thumbnail: "",
});
</script>

<style lang="scss" scoped></style>
