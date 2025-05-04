<template>
  <span class="cursor-pointer" @click="toggle">
    <FullscreenOutlined class="rounded-md p-1" v-if="!isFullscreen" />
    <FullscreenExitOutlined class="rounded-md p-1" v-else />
  </span>
</template>
<script lang="ts" setup>
import { useFullscreen } from "@vueuse/core";

import {
  FullscreenExitOutlined,
  FullscreenOutlined,
} from "@ant-design/icons-vue";
const { toggle, isFullscreen } = useFullscreen();
</script>
