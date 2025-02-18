<template>
  <div class="font-bold">
    <div class="flex items-center justify-center">
      <img :src="Stellux" alt="logo" class="w-1/2 my-2" />
    </div>
    <a-menu
      v-model:selectedKeys="state.selectedKeys"
      mode="inline"
      :open-keys="state.openKeys"
      :items="items"
      @openChange="onOpenChange"
    ></a-menu>
  </div>
</template>
<script lang="ts" setup>
import { VueElement, h, reactive } from "vue";
import {
  HomeOutlined,
  UserOutlined,
  AppstoreOutlined,
  InteractionOutlined,
  LineChartOutlined,
  BellOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";
import type { ItemType } from "ant-design-vue";
import Stellux from "@/assets/dashboard/stellux.png";

function getItem(
  label: VueElement | string,
  key: string,
  icon?: any,
  children?: ItemType[],
  type?: "group"
): ItemType {
  return {
    key,
    icon,
    children,
    label,
    type,
  } as ItemType;
}

const items: ItemType[] = reactive([
  getItem("仪表盘", "dashboard", () => h(HomeOutlined)),
  getItem("用户管理", "user", () => h(UserOutlined), [
    getItem("用户列表", "userList", null),
    getItem("用户角色", "userRole", null),
  ]),
  getItem("内容管理", "content", () => h(AppstoreOutlined), [
    getItem("发布文章", "publishArticle", null),
    getItem("文章列表", "articleList", null),
    getItem("草稿箱", "draft", null),
    getItem("素材库", "materialLibrary", null),
    getItem("回收站", "recycleBin", null),
  ]),
  getItem("互动管理", "interaction", () => h(InteractionOutlined), [
    getItem("评论管理", "comment", null),
    getItem("点赞管理", "like", null),
    getItem("私信管理", "message", null),
  ]),
  getItem("数据分析", "data", () => h(LineChartOutlined), [
    getItem("文章统计", "articleStatistics", null),
    getItem("用户统计", "userStatistics", null),
    getItem("互动统计", "interactionStatistics", null),
  ]),
  getItem("通知中心", "notice", () => h(BellOutlined), [
    getItem("系统通知", "systemNotice", null),
    getItem("用户通知", "userNotice", null),
  ]),
  getItem("系统设置", "system", () => h(SettingOutlined), [
    getItem("系统设置", "systemSetting", null),
    getItem("用户设置", "userSetting", null),
    getItem("角色设置", "roleSetting", null),
    getItem("权限设置", "permissionSetting", null),
  ]),
]);

const state = reactive({
  rootSubmenuKeys: [
    "dashboard",
    "user",
    "content",
    "interaction",
    "data",
    "notice",
    "system",
  ],
  openKeys: ["content"],
  selectedKeys: [],
});
const onOpenChange = (openKeys: string[]) => {
  const latestOpenKey = openKeys.find(
    (key) => state.openKeys.indexOf(key) === -1
  );
  if (
    typeof latestOpenKey === "string" &&
    state.rootSubmenuKeys.indexOf(latestOpenKey) === -1
  ) {
    state.openKeys = openKeys;
  } else {
    state.openKeys = latestOpenKey ? [latestOpenKey] : [];
  }
};
</script>
