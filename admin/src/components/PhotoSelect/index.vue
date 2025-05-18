<!-- 选择图片组件，用于选择图片或上传并返回图片的url -->
<template>
  <div>
    <a-modal
      width="100%"
      wrap-class-name="full-modal"
      :open="thumbnailModalOpen"
      @update:open="fetchThumbnails"
      title="选择图片"
      @ok="thumbnailModalOpen = false"
      @cancel="thumbnailModalOpen = false"
    >
      <div class="flex justify-end my-4">
        <Uploader @close="fetchThumbnails" />
      </div>
      <PhotoWall
        mode="picture-card"
        type="select"
        :list="page.list"
        @selected-photos="handleSelectPicture"
      />
      <div class="flex justify-end my-4">
        <!-- 分页器 只有当总页数大于1时才显示 -->
        <a-pagination
          v-if="page.total_count > page.page_size"
          v-model:current="page.page_no"
          :total="page.total_count"
          show-less-items
          @change="fetchThumbnails"
        />
      </div>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { queryFileList } from "@/api/file";
import PhotoWall from "@/components/PhotoWall/index.vue";
import type { PageData } from "@/types/dto";
import type { FileVO } from "@/types/file";
import { useVModel } from "@vueuse/core";
import Uploader from "@/components/Uploader/index.vue";

const props = defineProps<{
  open: boolean;
}>();

const emit = defineEmits<{
  (e: "update:open", open: boolean): void;
  (e: "selected-picture", picture: string): void;
}>();

const thumbnailModalOpen = useVModel(props, "open", emit);

const page = reactive<PageData<FileVO>>({
  page_no: 1,
  page_size: 10,
  total_count: 0,
  total_page: 0,
  list: [],
});

const handleSelectPicture = (pictures: string[]) => {
  emit("selected-picture", pictures[0]);
};

const fetchThumbnails = async () => {
  const res = await queryFileList({
    page_no: page.page_no,
    page_size: page.page_size,
  });
  page.list = res.data.list;
  page.total_count = res.data.total_count;
};

onMounted(() => {
  fetchThumbnails();
});
</script>