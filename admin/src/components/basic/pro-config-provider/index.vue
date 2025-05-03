<script setup lang="ts">
  import { configProviderProps } from 'ant-design-vue/es/config-provider/context';
  import { ConfigProvider } from 'ant-design-vue';

  defineOptions({
    name: 'ProConfigProvider',
  });

  const props = defineProps(configProviderProps());
</script>

<template>
  <ConfigProvider v-bind="props">
    <slot />
  </ConfigProvider>
</template>
