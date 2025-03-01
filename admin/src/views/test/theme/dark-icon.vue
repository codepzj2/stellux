<template>
  <a-space>
    <icon :style="{ fontSize: '48px' }">
      <template #component>
        <img :src="SunIcon" alt="sun" />
      </template>
    </icon>
  </a-space>
</template>
<script lang="ts" setup>
import Icon from "@ant-design/icons-vue";
import SunIcon from "@/assets/icons/sun.svg";
</script>
<style scoped>

</style>
