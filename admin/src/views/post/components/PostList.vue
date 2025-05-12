<template>
  <div>
    <a-page-header
      :title="$route.meta.title"
      class="!px-0"
      @back="() => $router.back()"
      :backIcon="props.type !== 'publish' ? backIcon : false"
    >
      <template #extra>
        <a-input-search
          v-model:value="page.keyword"
          placeholder="搜索标题 / 描述 / 标签"
          style="width: 240px"
          :loading="searchLoading"
          enter-button
          @search="onSearch"
        />
      </template>

      <div class="flex justify-between items-center mt-2">
        <div class="flex items-center gap-2">
          <span v-if="selectedIDList.length" class="text-gray-500 text-sm">
            已选中 {{ selectedIDList.length }} 篇文章
          </span>

          <a-popconfirm
            title="确定删除选中的文章吗？"
            ok-text="确定"
            cancel-text="取消"
            @confirm="onHandleDeleteSelected"
          >
            <a-button type="primary" danger :disabled="!selectedIDList.length">
              <template #icon><DeleteOutlined /></template>
              批量删除
            </a-button>
          </a-popconfirm>

          <a-popconfirm
            v-if="props.type === 'bin'"
            title="确定恢复选中的文章吗？"
            ok-text="确定"
            cancel-text="取消"
            @confirm="onHandleRestoreSelected"
          >
            <a-button type="primary" :disabled="!selectedIDList.length">
              <template #icon><UndoOutlined /></template>
              批量恢复
            </a-button>
          </a-popconfirm>
        </div>

        <div class="flex gap-2">
          <a-tooltip title="查看回收站">
            <a-button
              v-if="props.type !== 'bin'"
              @click="$router.push({ name: 'PostBin' })"
            >
              回收站
            </a-button>
          </a-tooltip>
          <a-tooltip title="查看草稿箱">
            <a-button
              v-if="props.type !== 'draft'"
              @click="$router.push({ name: 'PostDraft' })"
            >
              草稿箱
            </a-button>
          </a-tooltip>
          <a-tooltip title="创建新文章">
            <a-button
              v-if="props.type !== 'publish'"
              type="primary"
              @click="$router.push({ name: 'PostCreate' })"
            >
              新增文章
            </a-button>
          </a-tooltip>
        </div>
      </div>
    </a-page-header>

    <a-table
      :columns="columns"
      :data-source="postList"
      :scroll="{ x: 1000 }"
      :pagination="false"
      :loading="loading"
      :rowKey="(record: PostDetailVO) => record.id"
      :rowSelection="rowSelection"
      class="mt-4"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'title'">
          <div class="text-sm font-medium flex items-center gap-2">
            {{ record.title }}
            <a-tag v-if="record.is_top" color="red">置顶</a-tag>
          </div>
        </template>

        <template v-else-if="column.key === 'description'">
          <span class="text-sm text-gray-500 line-clamp-2">
            {{ record.description }}
          </span>
        </template>

        <template v-else-if="column.key === 'thumbnail'">
          <a-image
            :width="80"
            :src="record.thumbnail"
            :fallback="ImgFallback"
          />
        </template>

        <template v-else-if="column.key === 'category'">
          <a-tag>{{ record.category }}</a-tag>
        </template>

        <template v-else-if="column.key === 'tags'">
          <div class="flex flex-wrap gap-1">
            <a-tag v-for="tag in record.tags" :key="tag">{{ tag }}</a-tag>
          </div>
        </template>

        <template v-else-if="column.key === 'action'">
          <div class="flex gap-1">
            <template v-if="props.type !== 'bin'">
              <a-button
                type="link"
                size="small"
                @click="() => onHandlePublish(record.id)"
              >
                {{ props.type === "publish" ? "下架" : "上架" }}
              </a-button>
              <a-button
                type="link"
                size="small"
                @click="() => onHandleEdit(record)"
              >
                编辑
              </a-button>
              <a-popconfirm
                placement="bottomRight"
                title="确定删除该文章吗？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="() => onHandleSoftDelete(record)"
              >
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </template>
            <template v-else>
              <a-button
                type="link"
                size="small"
                @click="() => onHandleRestore(record.id)"
              >
                恢复
              </a-button>
              <a-popconfirm
                placement="bottomRight"
                title="永久删除该文章？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="() => onHandleDelete(record)"
              >
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </template>
          </div>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script lang="ts" setup>
import {
  getPublishPostDetailListAPI,
  getDraftPostDetailListAPI,
  getBinPostDetailListAPI,
  softDeletePostAPI,
  updatePostPublishStatusAPI,
  restorePostAPI,
  deletePostAPI,
  softDeletePostBatchAPI,
  deletePostBatchAPI,
  restorePostBatchAPI,
} from "@/api/post";
import type { PostDetailVO } from "@/types/post";
import ImgFallback from "@/assets/png/img-fallback.png";
import { message, type TableColumnType, type TableProps } from "ant-design-vue";
import {
  DeleteOutlined,
  UndoOutlined,
  ArrowLeftOutlined,
} from "@ant-design/icons-vue";
import type { Key } from "ant-design-vue/es/table/interface";
import type { PageResponse, PageReq } from "@/types/dto";

const loading = ref(false);
const searchLoading = ref(false);
const postList = ref<PostDetailVO[]>([]);
const router = useRouter();

const props = defineProps<{
  page: PageReq;
  type: "publish" | "draft" | "bin";
}>();

const emit = defineEmits<{
  (e: "update:page", value: PageReq): void;
}>();

const backIcon = h(ArrowLeftOutlined);
const page = toRef(props, "page");

const selectedIDList = ref<string[]>([]);
const rowSelection: TableProps<PostDetailVO>["rowSelection"] = {
  onChange: (_keys: Key[], selectedRows) => {
    selectedIDList.value = selectedRows.map(row => row.id);
  },
};

const getPostList = async () => {
  loading.value = true;
  try {
    let res: PageResponse<PostDetailVO>;
    switch (props.type) {
      case "publish":
        res = await getPublishPostDetailListAPI(page.value);
        break;
      case "draft":
        res = await getDraftPostDetailListAPI(page.value);
        break;
      case "bin":
        res = await getBinPostDetailListAPI(page.value);
        break;
    }
    postList.value = res.data.list;
    emit("update:page", page.value);
  } finally {
    loading.value = false;
  }
};

const onHandlePublish = async (id: string) => {
  await updatePostPublishStatusAPI(id, props.type !== "publish");
  message.success(props.type === "publish" ? "下架成功" : "上架成功");
  await getPostList();
};

const onHandleRestore = async (id: string) => {
  await restorePostAPI(id);
  message.success("恢复成功，请到草稿箱查看");
  await getPostList();
};

const onHandleDeleteSelected = async () => {
  if (!selectedIDList.value.length) return;
  if (props.type === "bin") {
    await deletePostBatchAPI(selectedIDList.value);
  } else {
    await softDeletePostBatchAPI(selectedIDList.value);
  }
  message.success("批量操作成功");
  selectedIDList.value = [];
  await getPostList();
};

const onHandleRestoreSelected = async () => {
  if (!selectedIDList.value.length) return;
  await restorePostBatchAPI(selectedIDList.value);
  message.success("批量恢复成功");
  selectedIDList.value = [];
  await getPostList();
};

const onSearch = async () => {
  searchLoading.value = true;
  page.value.page_no = 1;
  await getPostList();
  searchLoading.value = false;
};

const onHandleEdit = (record: PostDetailVO) => {
  router.push({ name: "PostEdit", params: { id: record.id } });
};

const onHandleSoftDelete = async (record: PostDetailVO) => {
  await softDeletePostAPI(record.id);
  message.success("删除成功");
  await getPostList();
};

const onHandleDelete = async (record: PostDetailVO) => {
  await deletePostAPI(record.id);
  message.success("已永久删除");
  await getPostList();
};

const columns: TableColumnType<PostDetailVO>[] = [
  { title: "标题", key: "title", width: 200 },
  { title: "封面", key: "thumbnail", width: 100 },
  { title: "描述", key: "description", width: 220 },
  { title: "分类", key: "category", width: 100 },
  { title: "标签", key: "tags", width: 200 },
  { title: "操作", key: "action", width: 180, fixed: "right" },
];

onMounted(() => {
  getPostList();
});
</script>

<style scoped lang="scss"></style>
