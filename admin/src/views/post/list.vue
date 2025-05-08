<template>
  <a-page-header title="文章列表" class="!px-0">
    <template #extra>
      <a-input-search
        v-model:value="page.keyword"
        placeholder="请输入关键字"
        style="width: 200px"
        :loading="searchLoading"
        enter-button
        @search="onSearch"
      />
      <a-button @click="$router.push({ name: 'PostBin' })">回收箱</a-button>
      <a-button>草稿箱</a-button>
      <a-button type="primary" @click="$router.push({ name: 'PostCreate' })"
        >新增文章</a-button
      >
    </template>
  </a-page-header>
  <a-table
    :columns="columns"
    :data-source="postList"
    :scroll="{ x: 500 }"
    :pagination="false"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'title'">
        <span class="text-sm flex items-center gap-2">
          <span class="font-bold">{{ record.title }}</span>
          <a-tag v-if="record.is_top" color="red">置顶</a-tag>
        </span>
      </template>
      <template v-else-if="column.key === 'description'">
        <span class="text-sm line-clamp-2">{{ record.description }}</span>
      </template>
      <template v-else-if="column.key === 'thumbnail'">
        <a-image :width="80" :src="record.thumbnail" :fallback="ImgFallback" />
      </template>
      <template v-else-if="column.key === 'category'">
        <a-tag>
          {{ record.category }}
        </a-tag>
      </template>
      <template v-else-if="column.key === 'tags'">
        <span>
          <a-tag v-for="tag in record.tags" :key="tag">
            {{ tag }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a-button type="link" size="small" @click="onHandleEdit(record)">
            编辑
          </a-button>
          <a-divider type="vertical" />
          <a-popconfirm
            placement="bottomRight"
            title="确定删除该文章吗？"
            ok-text="确定"
            cancel-text="取消"
            @confirm="onHandleDelete(record)"
          >
            <a-button type="link" size="small" danger>删除</a-button>
          </a-popconfirm>
        </span>
      </template>
    </template>
  </a-table>
</template>
<script lang="ts" setup>
import { getPostDetailListAPI, softDeletePostAPI } from "@/api/post";
import type { PostDetailVO } from "@/types/post";
import ImgFallback from "@/assets/png/img-fallback.png";
import type { PageReq } from "@/types/response";
import { message } from "ant-design-vue";

const postList = ref<PostDetailVO[]>([]);
const router = useRouter();
const page = reactive<PageReq>({
  page_no: 1,
  page_size: 10,
  keyword: "",
});

const searchLoading = ref(false);

const getPostList = async () => {
  const res = await getPostDetailListAPI(page);
  postList.value = res.data.list;
};

onMounted(() => {
  getPostList();
});

const onSearch = async () => {
  searchLoading.value = true;
  page.page_no = 1;
  await getPostList();
  searchLoading.value = false;
};

const onHandleEdit = (record: PostDetailVO) => {
  router.push({
    name: "PostEdit",
    params: { id: record.id },
  });
};

const onHandleDelete = async (record: PostDetailVO) => {
  const res = await softDeletePostAPI(record.id);
  message.success(res.msg);
  await getPostList();
};

const columns = [
  {
    title: "标题",
    dataIndex: "title",
    key: "title",
    width: 150,
  },
  {
    title: "封面",
    dataIndex: "thumbnail",
    key: "thumbnail",
    width: 120,
  },
  {
    title: "描述",
    dataIndex: "description",
    key: "description",
    width: 200,
  },
  {
    title: "分类",
    dataIndex: "category",
    key: "category",
    width: 100,
  },
  {
    title: "标签",
    key: "tags",
    dataIndex: "tags",
    width: 200,
  },
  {
    title: "操作",
    key: "action",
    width: 150,
    fixed: "right",
  },
];
</script>

<style lang="scss" scoped></style>
