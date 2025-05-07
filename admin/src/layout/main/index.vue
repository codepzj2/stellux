<template>
  <div>
    <router-view v-slot="{ Component }">
      <template v-if="Component">
        <keep-alive :include="cacheNames">
          <component :is="Component" />
        </keep-alive>
      </template>
    </router-view>
  </div>
</template>

<script lang="ts" setup>
import { cacheNames } from "@/router";
</script>
<style lang="scss" scoped></style>
