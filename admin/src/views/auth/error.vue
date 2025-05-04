<template>
  <div class="flex items-center justify-center h-screen w-full">
    <a-result status="404" :title="title" :sub-title="subTitle">
      <template #extra>
        <a-button type="primary" @click="$router.push('/')">返回首页</a-button>
      </template>
    </a-result>
  </div>
</template>

<script setup lang="ts">
const route = useRoute();
const title = ref("");
const subTitle = ref("");

const statusMap = {
  "401": "未授权访问",
  "403": "禁止访问",
  "404": "未找到",
  "500": "服务器错误",
};

onMounted(() => {
  title.value = route.name as string;
  subTitle.value = statusMap[title.value as unknown as keyof typeof statusMap];
});

onUpdated(() => {
  console.log(route.name);
});
</script>
