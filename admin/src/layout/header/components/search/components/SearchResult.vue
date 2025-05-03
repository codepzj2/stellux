<template>
  <div>
    <div class="pb-3">
      <template v-for="item in options" :key="item.name">
        <div
          class="bg-[#e5e7eb] h-12 px-4 rounded-md flex items-center justify-between"
          style="cursor: pointer"
          :style="{
            background: item.name === active ? systemStore.themeColor : '',
            color: item.name === active ? '#fff' : '',
          }"
          @click="handleTo"
          @mouseenter="handleMouse(item)"
        >
          <div class="flex items-center gap-2">
            <BookOutlined />
            <span>{{ item.meta?.title }}</span>
          </div>
          <EnterOutlined class="icon text-lg p-0.5 mr-1" />
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { EnterOutlined, BookOutlined } from "@ant-design/icons-vue";
import type { RouteRecordRaw } from "vue-router";
import { useSystemStore } from "@/store/modules/system";

interface Props {
  value: string;
  options: RouteRecordRaw[];
}

interface Emits {
  (e: "update:value", val: string): void;
  (e: "enter"): void;
}

const systemStore = useSystemStore();
const props = withDefaults(defineProps<Props>(), {});
const emit = defineEmits<Emits>();

const active = computed({
  get() {
    return props.value;
  },
  set(val: string) {
    emit("update:value", val);
  },
});
/** 鼠标移入 */
async function handleMouse(item: RouteRecordRaw) {
  active.value = item.name as string;
}

function handleTo() {
  emit("enter");
}
</script>
<style lang="scss" scoped></style>
