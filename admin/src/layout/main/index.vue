<template>
  <div v-bind="$attrs">
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      animated
      type="editable-card"
      size="small"
      @edit="tabStore.editTab"
      @change="tabStore.changeTab"
    >
      <a-tab-pane
        v-for="pane in panes"
        :key="pane.key"
        :tab="pane.title"
        :closable="pane.closable"
      >
      </a-tab-pane>
    </a-tabs>
    <router-view v-slot="{ Component }">
      <template v-if="Component">
        <Transition name="slide-fade">
          <keep-alive :include="cacheNames">
            <component :is="Component" />
          </keep-alive>
        </Transition>
      </template>
    </router-view>
  </div>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { cacheNames } from "@/router";
import { useTabStore } from "@/store";

const tabStore = useTabStore();
const { panes, activeKey } = storeToRefs(tabStore);
</script>
<style lang="scss" scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>
