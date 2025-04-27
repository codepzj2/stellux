<template>
  <div class="page">
    <div class="page-header">
      <slot name="header"></slot>
    </div>
    <div class="page-content">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.page {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  .page-header {
    height: 50px;
  }
  .page-content {
    flex: 1;
  }
}
</style>
