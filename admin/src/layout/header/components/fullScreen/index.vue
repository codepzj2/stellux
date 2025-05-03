<template>
  <Tooltip :title="getTitle" placement="bottom" :mouse-enter-delay="0.5">
    <span class="cursor-pointer" @click="toggle">
      <FullscreenOutlined class="hover:bg-zinc-200 rounded-md p-1" v-if="!isFullscreen" />
      <FullscreenExitOutlined class="hover:bg-zinc-200 rounded-md p-1" v-else />
    </span>
  </Tooltip>
</template>
<script lang="ts" setup>
import { computed, unref } from "vue";
import { useFullscreen } from "@vueuse/core";

import {
  FullscreenExitOutlined,
  FullscreenOutlined,
} from "@ant-design/icons-vue";
import { Tooltip } from "ant-design-vue";
const { toggle, isFullscreen } = useFullscreen();

const getTitle = computed(() => {
  return unref(isFullscreen) ? "退出全屏" : "进入全屏";
});
</script>
