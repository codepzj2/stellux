<template>
  <div>
    <a-drawer
      :title="props.mode === 'edit' ? '编辑文章' : '发布文章'"
      :open="props.open"
      :width="width < 768 ? '100%' : '768px'"
      :footer-style="{ textAlign: 'right' }"
      @close="emit('update:open', false)"
    >
      <a-form
        ref="postFormRef"
        :model="postForm"
        :rules="rules"
        layout="vertical"
      >
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="标题" name="title">
              <a-input
                v-model:value="postForm.title"
                placeholder="请输入标题"
                :maxlength="50"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="作者" name="author">
              <a-input v-model:value="postForm.author" placeholder="作者昵称" />
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="分类" name="category_id">
              <a-select
                v-model:value="postForm.category_id"
                placeholder="请选择分类"
              >
                <a-select-option
                  v-for="c in categoriesType"
                  :key="c.id"
                  :value="c.id"
                >
                  {{ c.name }}
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>

          <a-col :span="12">
            <a-form-item label="标签" name="tags_id">
              <a-select
                mode="multiple"
                v-model:value="postForm.tags_id"
                placeholder="请选择标签"
              >
                <a-select-option
                  v-for="t in tagsType"
                  :key="t.id"
                  :value="t.id"
                >
                  {{ t.name }}
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="封面" name="thumbnail">
              <div
                v-if="postForm.thumbnail"
                class="w-[200px] h-[112px] flex justify-center relative group"
              >
                <!-- 封面图 -->
                <a-image
                  :src="postForm.thumbnail"
                  class="rounded-md cursor-pointer object-cover"
                  :preview="false"
                  @click="getThumbnailList"
                />

                <!-- 删除按钮，仅在 hover 时显示 -->
                <div
                  class="absolute top-1 right-1 z-10 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                >
                  <a-button
                    type="primary"
                    shape="circle"
                    danger
                    size="small"
                    @click.stop="postForm.thumbnail = ''"
                  >
                    <CloseOutlined />
                  </a-button>
                </div>
              </div>
              <div
                v-else
                class="w-[200px] h-[112px] flex items-center justify-center border-2 border-dashed border-gray-300 rounded-md cursor-pointer text-gray-400 transition-colors"
                @click="getThumbnailList"
              >
                <span class="text-sm">点击上传封面</span>
              </div>
            </a-form-item>
          </a-col>
        </a-row>

        <a-form-item label="描述" name="description">
          <a-textarea
            v-model:value="postForm.description"
            placeholder="请输入描述"
            :rows="2"
          />
        </a-form-item>

        <a-form-item label="内容" name="content">
          <a-textarea
            v-model:value="postForm.content"
            placeholder="请输入内容"
            :rows="6"
          />
        </a-form-item>

        <a-form-item label="发布时间" name="created_at">
          <a-date-picker show-time v-model:value="createdAt" />
        </a-form-item>

        <div class="flex items-center gap-2 my-4">
          <span>发布</span>
          <a-switch v-model:checked="postForm.is_publish" />
          <a-divider type="vertical" />
          <span>置顶</span>
          <a-switch v-model:checked="postForm.is_top" />
        </div>
      </a-form>

      <template #extra>
        <a-space>
          <a-button type="primary" @click="onHandleCreateOrEdit">
            发布
          </a-button>
          <a-button @click="emit('update:open', false)"> 取消 </a-button>
        </a-space>
      </template>
    </a-drawer>

    <a-modal
      width="100%"
      wrap-class-name="full-modal"
      :open="thumbnailModalOpen"
      @update:open="thumbnailModalOpen = $event"
      title="选择封面"
      @ok="thumbnailModalOpen = false"
    >
      <PhotoWall
        mode="picture-card"
        type="select"
        :list="page.list"
        @selected-photos="handleSelectPicture"
      />
      <div class="flex justify-end my-4">
        <a-pagination
          v-model:current="page.page_no"
          :total="page.total_count"
          show-less-items
          @change="getThumbnailList"
        />
      </div>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { queryAllByTypeAPI } from "@/api/label";
import type { LabelVO } from "@/types/label";
import type { PostReq } from "@/types/post";
import { useWindowSize } from "@vueuse/core";
import { message, type FormInstance } from "ant-design-vue";
import { createPostAPI, updatePostAPI } from "@/api/post";
import dayjs from "dayjs";
import PhotoWall from "@/components/PhotoWall/index.vue";
import { queryFileList } from "@/api/file";
import type { FileVO } from "@/types/file";
import type { PageData } from "@/types/dto";
import { CloseOutlined } from "@ant-design/icons-vue";
import { useRouter } from "vue-router";

const props = defineProps<{
  mode: "create" | "edit";
  open: boolean;
  postForm: PostReq;
}>();

const emit = defineEmits<{
  (e: "update:open", value: boolean): void;
  (e: "update:postForm", value: PostReq): void;
}>();

const router = useRouter();
const postFormRef = ref<FormInstance>();
const postForm = computed({
  get() {
    return props.postForm;
  },
  set(value) {
    emit("update:postForm", value);
  },
});

const createdAt = computed({
  get() {
    return dayjs(postForm.value.created_at);
  },
  set(value) {
    postForm.value.created_at = value.toISOString();
  },
});

const thumbnailModalOpen = ref(false);
const page = reactive<PageData<FileVO>>({
  page_no: 1,
  page_size: 10,
  total_count: 0,
  total_page: 0,
  list: [],
});

const getThumbnailList = async () => {
  const res = await queryFileList({
    page_no: page.page_no,
    page_size: page.page_size,
  });
  page.list = res.data.list;
  page.total_count = res.data.total_count;
  thumbnailModalOpen.value = true;
};

const handleSelectPicture = (pictures: string[]) => {
  postForm.value.thumbnail = pictures[0];
};

const rules = {
  title: [{ required: true, message: "请输入标题" }],
  content: [{ required: true, message: "请输入内容" }],
  description: [{ required: true, message: "请输入描述" }],
  thumbnail: [{ required: true, message: "请输入封面" }],
};

const { width } = useWindowSize();
const categoriesType = ref<LabelVO[]>([]);
const tagsType = ref<LabelVO[]>([]);

const getCategories = async () => {
  const res = await queryAllByTypeAPI("category");
  categoriesType.value = res.data;
};

const getTags = async () => {
  const res = await queryAllByTypeAPI("tag");
  tagsType.value = res.data;
};

const onHandleCreateOrEdit = () => {
  postFormRef.value?.validate().then(async () => {
    if (props.mode === "create") {
      await createPostAPI(postForm.value);
      message.success("发布成功");
    } else {
      await updatePostAPI(postForm.value);
      message.success("编辑成功");
    }
    emit("update:open", false);
    router.push({ name: "PostList" });
  });
};

onMounted(() => {
  getCategories();
  getTags();
});
</script>

<style lang="scss">
.full-modal {
  .ant-modal {
    max-width: 100%;
    top: 0;
    padding-bottom: 0;
    margin: 0;
  }
  .ant-modal-content {
    display: flex;
    flex-direction: column;
    height: calc(100vh);
  }
  .ant-modal-body {
    flex: 1;
    overflow-y: auto;
  }
}

::-webkit-scrollbar {
  width: 8px;
}
</style>
