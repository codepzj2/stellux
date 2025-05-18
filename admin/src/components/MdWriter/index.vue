<template>
  <div class="h-full">
    <div class="md-editor markdown-body">
      <Editor
        :value="content"
        :locale="zhHans"
        :plugins="mdPlugins"
        @change="content = $event"
        :mode="props.mode"
        placeholder="请输入内容..."
      />
    </div>
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
import { useVModel } from "@vueuse/core";

const props = withDefaults(defineProps<{
  content: string;
  mode?: "split" | "tab" | "auto";
}>(), {
  mode: "auto",
});

const emit = defineEmits<{
  (e: "update:content", value: string): void;
  (e: "update:mode", value: "split" | "tab" | "auto"): void;
}>();

const content = useVModel(props, "content", emit);

const mdPlugins = ref([
  gfm(),
  gemoji(),
  highlight(),
  frontmatter(),
  mediumZoom(),
  breaks(),
  mermaid(),
]);

</script>

<style lang="scss" scoped></style>
