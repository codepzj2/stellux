<template>
  <div class="my-2 mr-2 flex justify-between">
    <div class="flex-1">
      <slot name="left"></slot>
    </div>
    <div class="flex justify-end">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped></style>
