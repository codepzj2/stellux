<template>
  <div class="h-screen font-sans" :class="{ dark: darkMode }">
    <!-- 降低antd的权重, 避免和tailwind冲突 -->
    <a-config-provider :theme="appTheme" :locale="zhCN">
      <a-style-provider
        hash-priority="low"
        :transformers="[legacyLogicalPropertiesTransformer]"
      >
        <router-view></router-view>
      </a-style-provider>
    </a-config-provider>
  </div>
</template>

<script setup lang="ts">
import { legacyLogicalPropertiesTransformer, theme } from "ant-design-vue";
import { useSystemStore } from "@/store";
import { storeToRefs } from "pinia";

import zhCN from "ant-design-vue/es/locale/zh_CN";
import dayjs from "dayjs";
import "dayjs/locale/zh-cn";

dayjs.locale("zh-cn");
const systemStore = useSystemStore();
const { darkMode } = storeToRefs(systemStore);
const appTheme = computed(() => {
  return {
    algorithm: darkMode.value ? theme.darkAlgorithm : theme.defaultAlgorithm,
    token: {
      fontSize: 15,
      fontFamily: "var(--font-sans)",
    },
  };
});
</script>
