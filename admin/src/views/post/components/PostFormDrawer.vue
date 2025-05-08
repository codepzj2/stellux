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
              <a-input
                v-model:value="postForm.thumbnail"
                placeholder="请输入封面地址"
              />
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
      emit("update:open", false);
      router.push({ name: "PostList" });
    } else {
      await updatePostAPI(postForm.value);
      message.success("编辑成功");
      emit("update:open", false);
      router.push({ name: "PostList" });
    }
  });
};

onMounted(() => {
  getCategories();
  getTags();
});
</script>
