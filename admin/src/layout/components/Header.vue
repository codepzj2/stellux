<template>
  <div class="flex h-full w-full items-center justify-between">
    <div class="w-full">
      <a-button
        type="text"
        @click="sidebarStore.setCollapse(!sidebarStore.collapse)"
        class="md:!hidden"
      >
        <MenuOutlined />
      </a-button>
      <div class="w-full flex justify-between">
        <div
          v-if="headerStore.leftHeaderActions.length === 0"
          class="flex-grow"
        ></div>
        <div
          v-for="action in headerStore.leftHeaderActions"
          :key="action.label"
          class="px-1"
        >
          <a-button
            :type="action.type ?? 'default'"
            :size="action.size ?? 'default'"
            :icon="action.icon"
            @click="action.onClick"
          >
            {{ action.label }}
          </a-button>
        </div>
        <div
          v-for="action in headerStore.rightHeaderActions"
          :key="action.label"
          class="px-1"
        >
          <a-button
            :type="action.type ?? 'default'"
            :size="action.size ?? 'default'"
            :icon="action.icon"
            @click="action.onClick"
          >
            {{ action.label }}
          </a-button>
        </div>
      </div>
    </div>
    <div class="flex items-center gap-2"></div>
  </div>
</template>

<script lang="ts" setup>
import { MenuOutlined } from "@ant-design/icons-vue";
import { useSidebarStore } from "@/store";
import { useHeaderStore } from "@/store";

const sidebarStore = useSidebarStore();
const headerStore = useHeaderStore();
</script>
