<template>
  <div class="h-full">
    <a-card class="h-full overflow-scroll" :bordered="false">
      <div class="h-full flex">
        <SplitPanel>
          <template #left-content>
            <a-tree
              v-model:expandedKeys="expandedKeys"
              v-model:selectedKeys="selectedKeys"
              :tree-data="treeData"
              @rightClick="handleRightClick"
              @select="onHandleSelect"
            />
          </template>
          <template #right-content>
            <router-view v-slot="{ Component }">
              <component :is="Component" v-model:title="title" />
            </router-view>
          </template>
        </SplitPanel>
      </div>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import SplitPanel from "@/components/SplitPanel/index.vue";
import type { TreeProps } from "ant-design-vue";

const router = useRouter();
const expandedKeys = ref<string[]>(["0-0-0", "0-0-1", "0-0-2"]);
const selectedKeys = ref<string[]>([]);
const title = ref("");
const treeData: TreeProps["treeData"] = [
  {
    title: "stellux文档",
    key: "index",
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

const onHandleSelect = (selectedKeys: string[], event: any) => {
  title.value = event.node.title;
  router.push({
    name: "DocumentContentIndex",
    params: { id: selectedKeys[0] },
  });
  console.log(selectedKeys);
};
</script>
<style lang="scss" scoped>
:deep(.ant-card-body) {
  height: 100%;
  padding: 10px;
}
</style>
