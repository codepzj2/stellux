<template>
  <div class="my-2 mr-2 flex items-center justify-between">
    <div class="w-[300px] md:w-[50%]">
      <a-input
        v-model:value="post.title"
        placeholder="请输入标题"
        addon-before="标题"
        show-count
        :maxlength="50"
      />
    </div>
    <a-button type="primary" @click="onHandleCreate">发布文章</a-button>
  </div>

  <div class="md-editor markdown-body">
    <Editor
      :value="post.content"
      :locale="zhHans"
      :plugins="mdPlugins"
      @change="post.content = $event"
    />
  </div>

  <a-drawer
    title="发布文章"
    :open="open"
    :width="width < 768 ? '100%' : '768px'"
    :body-style="{ paddingBottom: '80px' }"
    :footer-style="{ textAlign: 'right' }"
    @close="open = false"
  >
    <a-form ref="postFormRef" :model="post" :rules="rules" layout="vertical">
      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item label="标题" name="title">
            <a-input
              v-model:value="post.title"
              placeholder="请输入标题"
              :maxlength="50"
            />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="作者" name="author">
            <a-input
              v-model:value="post.author"
              disabled
              placeholder="作者昵称"
            />
          </a-form-item>
        </a-col>
      </a-row>

      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item label="分类" name="category">
            <a-select v-model:value="post.category" placeholder="请选择分类">
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
          <a-form-item label="标签" name="tags">
            <a-select
              mode="multiple"
              v-model:value="post.tags"
              placeholder="请选择标签"
            >
              <a-select-option v-for="t in tagsType" :key="t.id" :value="t.id">
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
              v-model:value="post.thumbnail"
              placeholder="请输入封面地址"
            />
          </a-form-item>
        </a-col>
      </a-row>

      <a-form-item label="描述" name="description">
        <a-textarea
          v-model:value="post.description"
          placeholder="请输入描述"
          :rows="2"
        />
      </a-form-item>

      <a-form-item label="内容" name="content">
        <a-textarea
          v-model:value="post.content"
          placeholder="请输入内容"
          :rows="8"
        />
      </a-form-item>

      <div class="flex items-center gap-2 my-4">
        <span>发布</span>
        <a-switch v-model:checked="post.is_publish" />
        <a-divider type="vertical" />
        <span>置顶</span>
        <a-switch v-model:checked="post.is_top" />
      </div>
    </a-form>

    <template #extra>
      <a-space>
        <a-button @click="open = false">取消</a-button>
        <a-button type="primary" @click="onSubmit">提交</a-button>
      </a-space>
    </template>
  </a-drawer>
</template>

<script lang="ts" setup>
import Editor from "@/lib/editor";
import zhHans from "bytemd/locales/zh_Hans.json";
import "./md.scss";

import gfm from "@bytemd/plugin-gfm";
import gemoji from "@bytemd/plugin-gemoji";
import highlight from "@bytemd/plugin-highlight";
import frontmatter from "@bytemd/plugin-frontmatter";
import mediumZoom from "@bytemd/plugin-medium-zoom";
import mermaid from "@bytemd/plugin-mermaid";
import breaks from "@bytemd/plugin-breaks";

import { useUserStore } from "@/store";
import { useWindowSize } from "@vueuse/core";
import { createPostAPI } from "@/api/posts";
import { message } from "ant-design-vue";

import { reactive, ref } from "vue";
import { storeToRefs } from "pinia";
import type { FormInstance } from "ant-design-vue";
import type { PostReq } from "@/types/posts";
import type { LabelVO } from "@/types/label";
import { Code } from "@/global";
import { useHeaderStore } from "@/store";
import { useRouter } from "vue-router";
import { getLabelListAPI } from "@/api/label";
defineOptions({ name: "CreateContent" });

const router = useRouter();
const headerStore = useHeaderStore();

const mdPlugins = ref([
  gfm(),
  gemoji(),
  highlight(),
  frontmatter(),
  mediumZoom(),
  breaks(),
  mermaid(),
]);

const { width } = useWindowSize();
const postFormRef = ref<FormInstance>();
const open = ref(false);

const userStore = useUserStore();
const { userInfo } = storeToRefs(userStore);

const post = reactive<PostReq>({
  title: "",
  content: "",
  description: "",
  author: userInfo.value?.nickname ?? "",
  category: undefined,
  tags: [],
  is_publish: true,
  is_top: false,
  thumbnail: "",
});

const rules = {
  title: [{ required: true, message: "请输入标题" }],
  content: [{ required: true, message: "请输入内容" }],
  description: [{ required: true, message: "请输入描述" }],
  thumbnail: [{ required: true, message: "请输入封面" }],
};

const categoriesType = ref<LabelVO[]>([]);
const tagsType = ref<LabelVO[]>([]);

const getCategories = async () => {
  const res = await getLabelListAPI({
    label_type: "category",
    page_no: 1,
    page_size: 20,
  });
  categoriesType.value = res.data.list;
};

const getTags = async () => {
  const res = await getLabelListAPI({
    label_type: "tag",
    page_no: 1,
    page_size: 20,
  });
  tagsType.value = res.data.list;
};

const onHandleCreate = async () => {
  open.value = true;
  await getCategories();
  await getTags();
};

const onSubmit = async () => {
  try {
    const values = await postFormRef.value?.validate();
    const res = await createPostAPI(values as PostReq);
    if (res.code === Code.RequestSuccess) {
      message.success("发布成功");
      open.value = false;
    } else {
      message.error(res.msg);
    }
  } catch (e) {
    console.error("验证失败", e);
  }
};

onActivated(() => {
  headerStore.setRightHeaderActions([
    {
      label: "分类",
      onClick: () => router.push({ name: "LabelCategory" }),
    },
    {
      label: "标签",
      onClick: () => router.push({ name: "LabelTag" }),
    },
  ]);
});
</script>
