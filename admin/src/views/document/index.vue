<template>
  <a-card class="h-full">
    <div class="h-full flex">
      <div class="w-44 h-full overflow-y-auto">
        <a-tree
          v-model:expandedKeys="expandedKeys"
          v-model:selectedKeys="selectedKeys"
          :tree-data="treeData"
          @rightClick="handleRightClick"
          @select="onHandleSelect"
        >
        </a-tree>
      </div>
      <div class="flex-1">
        <router-view />
      </div>
    </div>
  </a-card>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import type { TreeProps } from "ant-design-vue";

const router = useRouter();
const expandedKeys = ref<string[]>(["0-0-0", "0-0-1", "0-0-2"]);
const selectedKeys = ref<string[]>([]);
const treeData: TreeProps["treeData"] = [
  {
    title: "stellux文档",
    key: "0-0",
    children: [
      {
        title: "前言",
        key: "0-0-0",
        children: [
          {
            title: "项目介绍",
            key: "0-0-0-0",
          },
          {
            title: "项目结构",
            key: "0-0-0-1",
          },
          {
            title: "项目配置",
            key: "0-0-0-2",
          },
        ],
      },
      {
        title: "项目结构",
        key: "0-0-1",
        children: [
          {
            title: "项目结构",
            key: "0-0-1-0",
          },
        ],
      },
      {
        title: "项目配置",
        key: "0-0-2",
        children: [
          {
            title: "项目配置",
            key: "0-0-2-0",
          },
          {
            title: "项目配置",
            key: "0-0-2-1",
          },
        ],
      },
    ],
  },
];
const handleRightClick = ({ event }: { event: PointerEvent }) => {
  console.log(event.clientX, event.clientY);
};

const onHandleSelect = (selectedKeys: string[]) => {
  console.log(selectedKeys);
  router.push({
    name: "DocumentContent",
    params: { id: selectedKeys[0] },
  });
};
</script>
<style lang="scss" scoped>
:deep(.ant-card-body) {
  height: 100%;
}
</style>
