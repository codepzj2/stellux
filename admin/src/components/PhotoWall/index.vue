<template>
  <div>
    <!-- 操作按钮 -->
    <div class="w-full flex justify-between items-center my-4">
      <div
        class="flex gap-2 min-h-[40px]"
        :class="{ invisible: !(props.type === 'display' && props.list.length) }"
      >
        <a-button type="primary" @click="handleToggleSelectAll">
          {{ isAllSelected ? "取消全选" : "全选" }}
        </a-button>
        <slot name="left"></slot>
      </div>
      <div class="flex gap-2">
        <slot name="right"></slot>
      </div>
    </div>

    <!-- 图片墙 -->
    <div
      v-if="photosWall.length"
      class="w-full grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-8 overflow-hidden"
    >
      <div
        v-for="photo in photosWall"
        :key="photo.url"
        class="relative w-full aspect-[16/9] overflow-hidden rounded-md"
        @click="handleClick(photo.url)"
        @dblclick.stop="handlePreview(photo)"
      >
        <img
          v-lazy="photo.url"
          :alt="photo.name"
          :class="[
            'w-full h-full object-contain cursor-pointer border transition-opacity duration-300',
            systemStore.themeMode === 'light'
              ? 'border-zinc-200'
              : 'border-zinc-800',
            loadedImages[photo.url] ? 'opacity-100' : 'opacity-0',
          ]"
          @load="onImageLoad(photo.url)"
        />

        <div
          v-if="selectedPhotos.includes(photo.url)"
          class="absolute top-1 right-1 z-10"
        >
          <CheckCircleFilled
            class="!text-green-500 !bg-white rounded-full text-xl shadow-md"
          />
        </div>
      </div>
    </div>

    <div v-else>
      <a-empty description="暂无图片" />
    </div>

    <!-- 预览弹窗 -->
    <a-modal
      :open="previewVisible"
      :title="previewTitle"
      :footer="null"
      @cancel="handleCancel"
      width="575px"
    >
      <div class="w-full h-full flex justify-center my-8">
        <img
          :alt="previewTitle"
          class="w-full max-w-[320px] max-h-[180px] object-contain mx-auto"
          :src="previewImage"
        />
      </div>
      <a-collapse>
        <a-collapse-panel header="复制图片地址">
          <a-typography-paragraph copyable>{{
            previewImage
          }}</a-typography-paragraph>
        </a-collapse-panel>
      </a-collapse>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, watch } from "vue";
import type { FileVO } from "@/types/file";
import { API_BASE_URL } from "@/constant";
import { CheckCircleFilled } from "@ant-design/icons-vue";
import { useSystemStore } from "@/store";

// Props & Emits
const props = defineProps<{
  type: "display" | "select";
  list: FileVO[];
}>();

const emit = defineEmits<{
  (e: "selected-photos", photos: string[]): void;
}>();

// 全局状态
const systemStore = useSystemStore();

// 数据状态
const selectedPhotos = ref<string[]>([]);
const loadedImages = ref<Record<string, boolean>>({});
const previewVisible = ref(false);
const previewImage = ref("");
const previewTitle = ref("");

// 构建图片列表
const photosWall = computed(() =>
  props.list.map(item => ({
    name: item.file_name,
    url: API_BASE_URL + item.url,
  }))
);

// 图片加载完成
const onImageLoad = (url: string) => {
  loadedImages.value[url] = true;
};

// 判断是否全选
const isAllSelected = computed(
  () =>
    photosWall.value.length > 0 &&
    selectedPhotos.value.length === photosWall.value.length
);

// 全选/取消全选切换
const handleToggleSelectAll = () => {
  if (props.type !== "display") return;
  selectedPhotos.value = isAllSelected.value
    ? []
    : photosWall.value.map(photo => photo.url);
  emit("selected-photos", selectedPhotos.value);
};

// 点击选中逻辑（带双击保护）
let clickTimeout: ReturnType<typeof setTimeout> | null = null;

const handleClick = (url: string) => {
  if (clickTimeout) {
    clearTimeout(clickTimeout);
    clickTimeout = null;
    return; // 忽略双击触发
  }

  clickTimeout = setTimeout(() => {
    clickTimeout = null;
    toggleSelection(url);
  }, 250);
};

// 切换选择状态
const toggleSelection = (url: string) => {
  if (props.type === "display") {
    const idx = selectedPhotos.value.indexOf(url);
    if (idx > -1) {
      selectedPhotos.value.splice(idx, 1);
    } else {
      selectedPhotos.value.push(url);
    }
  } else {
    selectedPhotos.value = selectedPhotos.value.includes(url) ? [] : [url];
  }
  emit("selected-photos", selectedPhotos.value);
};

// 预览逻辑
const handlePreview = async (file: any) => {
  previewImage.value =
    file.url || (await getBase64(file.originFileObj as File)) || "";
  previewTitle.value = file.name || file.url?.split("/").pop() || "";
  previewVisible.value = true;
};

const handleCancel = () => {
  previewVisible.value = false;
  previewTitle.value = "";
};

const getBase64 = (file: File) =>
  new Promise<string>((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = reject;
  });

// 数据变化时清空选中
watch(
  () => props.list,
  () => {
    selectedPhotos.value = [];
  },
  { deep: true }
);
</script>
