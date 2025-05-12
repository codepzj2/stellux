<template>
  <div class="w-full mx-auto">
    <PhotoWall
      :list="page.list"
      type="display"
      @selected-photos="photos => (selectedPhotos = photos)"
    >
      <template #left>
        <a-button
          type="primary"
          danger
          :disabled="selectedPhotos.length === 0"
          @click="handleDelete"
          >删除</a-button
        >
      </template>
      <template #right>
        <Uploader @close="getList" />
      </template>
    </PhotoWall>
    <div class="flex justify-end my-4">
      <a-pagination
        v-model:current="page.page_no"
        v-model:page-size="page.page_size"
        :total="page.total_count"
        @change="getList"
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import PhotoWall from "@/components/PhotoWall/index.vue";
import Uploader from "@/components/Uploader/index.vue";
import { deleteFilesByIdListAPI, queryFileList } from "@/api/file";
import type { FileVO } from "@/types/file";
import { message } from "ant-design-vue";
import type { PageData } from "@/types/dto";
import { API_BASE_URL } from "@/constant";

const page = reactive<PageData<FileVO>>({
  page_no: 1,
  page_size: 24,
  total_count: 0,
  total_page: 0,
  list: [],
});

// 选中的图片
const selectedPhotos = ref<string[]>([]);
const selectedPhotosIdList = computed(() => {
  const relativeUrl = selectedPhotos.value.map(
    item => item.split(API_BASE_URL)[1]
  );
  return page.list
    .filter(item => relativeUrl.includes(item.url))
    ?.map(item => item.id);
});

const getList = async () => {
  try {
    const res = await queryFileList({
      page_no: page.page_no,
      page_size: page.page_size,
    });
    page.list = res.data.list;
    page.total_count = res.data.total_count;
  } catch (error: any) {
    message.error({
      content: error + "，图片列表获取失败",
    });
  }
};
// 异步获取图片列表
await getList();

const handleDelete = async () => {
  if (selectedPhotosIdList.value.length === 0) {
    message.info("请单击图片删除");
    return;
  }
  await deleteFilesByIdListAPI({
    id_list: selectedPhotosIdList.value,
  });
  message.success("删除成功");
  await getList();
};
</script>
<style scoped lang="scss"></style>
