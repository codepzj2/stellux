<template>
  <div>
    <PostContentWriter mode="edit" :postForm="postForm" />
  </div>
</template>

<script setup lang="ts">
import { getPostByIdAPI } from "@/api/post";
import type { PostReq } from "@/types/post";

import PostContentWriter from "./components/PostContentWriter.vue";
defineOptions({
  name: "PostEdit",
});

const route = useRoute();
const postId = route.params.id;

const postForm = ref<PostReq>({
  id: "", // 更新携带文章id
  created_at: "",
  title: "",
  content: "",
  author: "",
  description: "",
  category_id: "",
  tags_id: [],
  is_publish: false,
  is_top: false,
  thumbnail: "",
});

const getPost = async () => {
  const res = await getPostByIdAPI(postId as string);
  postForm.value = res.data;
};

onMounted(() => {
  getPost();
});
</script>

<style lang="scss" scoped></style>
