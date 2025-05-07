<template>
  <a-table
    :columns="columns"
    :data-source="postList"
    :scroll="{ x: 500 }"
    :pagination="false"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'created_at'">
        {{ formatTime(record.created_at) }}
      </template>
      <template v-else-if="column.key === 'updated_at'">
        {{ formatTime(record.updated_at) }}
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
      <template v-else-if="column.key === 'is_top'">
        <a-switch v-model:checked="record.is_top" />
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
import { getPostDetailListAPI } from "@/api/post";
import { formatTime } from "@/utils/time";
import type { PostDetailVO } from "@/types/post";
import ImgFallback from "@/assets/png/img-fallback.png";

const postList = ref<PostDetailVO[]>([]);
const router = useRouter();

const getPostList = async () => {
  const res = await getPostDetailListAPI({
    page_no: 1,
    page_size: 10,
  });
  postList.value = res.data.list;
};

onMounted(() => {
  getPostList();
});

const onHandleEdit = (record: PostDetailVO) => {
  router.push({
    name: "PostEdit",
    params: { id: record.id },
  });
};

const onHandleDelete = (record: PostDetailVO) => {
  console.log(record);
};

const columns = [
  {
    title: "标题",
    dataIndex: "title",
    key: "title",
    width: 200,
    ellipsis: true,
  },
  {
    title: "作者",
    dataIndex: "author",
    key: "author",
    width: 100,
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
    ellipsis: true,
  },
  {
    title: "创建时间",
    dataIndex: "created_at",
    key: "created_at",
    width: 150,
  },
  {
    title: "更新时间",
    dataIndex: "updated_at",
    key: "updated_at",
    width: 150,
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
    title: "是否置顶",
    dataIndex: "is_top",
    key: "is_top",
    width: 100,
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
