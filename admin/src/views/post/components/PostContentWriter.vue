<template>
  <div>
    <PageHeader>
      <template #left>
        <div class="w-[300px] md:w-[50%]">
          <a-input
            v-model:value="postForm.title"
            placeholder="请输入标题"
            addon-before="标题"
            show-count
            :maxlength="50"
          />
        </div>
      </template>
      <template #right>
        <a-button type="primary" @click="open = true">
          {{ props.mode === "edit" ? "保存文章" : "发布文章" }}
        </a-button>
      </template>
    </PageHeader>

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
      v-model:is-confirm-leave="isConfirmLeave"
    />
  </div>
</template>

<script lang="ts" setup>
// @ts-ignore
import { Editor } from "@bytemd/vue-next";
import zhHans from "bytemd/locales/zh_Hans.json";
import { Modal } from "ant-design-vue";
import { onBeforeRouteLeave } from "vue-router";
import PageHeader from "@/components/PageHeader.vue";

import gfm from "@bytemd/plugin-gfm";
import gemoji from "@bytemd/plugin-gemoji";
import highlight from "@bytemd/plugin-highlight";
import frontmatter from "@bytemd/plugin-frontmatter";
import mediumZoom from "@bytemd/plugin-medium-zoom";
import mermaid from "@bytemd/plugin-mermaid";
import breaks from "@bytemd/plugin-breaks";

import PostFormDrawer from "./PostFormDrawer.vue";
import type { PostReq } from "@/types/post";
import { useSidebarStore } from "@/store";

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

// 编辑模式下拦截路由跳转，防止误操作
const isConfirmLeave = ref(false);

onBeforeRouteLeave((_to, _from, next) => {
  if (props.mode !== "edit") {
    next();
    return;
  }

  if (isConfirmLeave.value) {
    next();
    return;
  }

  const sidebarStore = useSidebarStore();
  sidebarStore.setSelectedKeys([]);
  Modal.confirm({
    title: "确定要离开吗？",
    content: "当前文章内容尚未保存，离开将导致内容丢失。",
    okText: "确认离开",
    cancelText: "取消",
    onOk() {
      isConfirmLeave.value = true;
      next();
    },
    onCancel() {
      next(false);
    },
  });
});
</script>

<style lang="scss" scoped></style>
