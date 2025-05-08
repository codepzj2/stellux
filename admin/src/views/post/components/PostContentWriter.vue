<template>
  <div>
    <div class="flex justify-between items-center my-4">
      <div class="w-[300px] md:w-[50%]">
        <a-input
          v-model:value="postForm.title"
          placeholder="请输入标题"
          addon-before="标题"
          show-count
          :maxlength="50"
        />
      </div>
      <div class="flex gap-2">
        <a-button type="primary" @click="open = true"> 发布文章 </a-button>
      </div>
    </div>

    <div class="md-editor markdown-body">
      <Editor
        :value="postForm.content"
        :locale="zhHans"
        :plugins="mdPlugins"
        @change="postForm.content = $event"
      />
    </div>

    <PostFormDrawer
      :mode="props.mode"
      v-model:open="open"
      v-model:post-form="postForm"
    />
  </div>
</template>

<script lang="ts" setup>
// @ts-ignore
import { Editor } from "@bytemd/vue-next";
import zhHans from "bytemd/locales/zh_Hans.json";

import gfm from "@bytemd/plugin-gfm";
import gemoji from "@bytemd/plugin-gemoji";
import highlight from "@bytemd/plugin-highlight";
import frontmatter from "@bytemd/plugin-frontmatter";
import mediumZoom from "@bytemd/plugin-medium-zoom";
import mermaid from "@bytemd/plugin-mermaid";
import breaks from "@bytemd/plugin-breaks";

import PostFormDrawer from "./PostFormDrawer.vue";
import type { PostReq } from "@/types/post";

const mdPlugins = ref([
  gfm(),
  gemoji(),
  highlight(),
  frontmatter(),
  mediumZoom(),
  breaks(),
  mermaid(),
]);

const open = ref(false);

const props = defineProps<{
  mode: "create" | "edit";
  postForm: PostReq;
}>();

const postForm = computed(() => props.postForm);
</script>

<style lang="scss" scoped></style>
