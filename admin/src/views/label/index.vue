<template>
  <div class="h-full">
    <!-- 分类/标签切换 -->

    <!-- 表格与操作区域 -->
    <a-page-header class="!px-0">
      <div class="flex justify-between">
        <a-segmented v-model:value="activeKey" :options="data" />
        <a-button type="primary" @click="onHandleCreate">新增</a-button>
      </div>
    </a-page-header>

    <a-table
      :columns="columns"
      :data-source="labelList"
      :loading="loading"
      rowKey="id"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <div>
            <a-input
              v-if="editableMap[record.id]"
              v-model:value="editableMap[record.id].name"
              style="width: 200px"
            />
            <template v-else>
              <a-tag>{{ record.name }}</a-tag>
            </template>
          </div>
        </template>

        <template v-else-if="column.key === 'action'">
          <span v-if="editableMap[record.id]">
            <a-typography-link @click="onSave(record)">保存</a-typography-link>
            <a-divider type="vertical" />
            <a-button type="link" size="small" @click="onCancel(record.id)"
              >取消</a-button
            >
          </span>
          <span v-else>
            <a-button size="small" type="link" @click="onEdit(record)"
              >编辑</a-button
            >
            <a-divider type="vertical" />
            <a-popconfirm
              placement="bottomRight"
              :title="`确定删除该${activeKey === 'category' ? '分类' : '标签'}吗？`"
              ok-text="确定"
              cancel-text="取消"
              @confirm="onHandleDelete(record)"
            >
              <a-button size="small" type="link" danger>删除</a-button>
            </a-popconfirm>
          </span>
        </template>
      </template>
    </a-table>

    <!-- 新增弹窗 -->
    <a-modal
      v-model:open="createModalOpen"
      :title="`新增${activeKey === 'category' ? '分类' : '标签'}`"
      @ok="handleCreateOk"
    >
      <a-form ref="createFormRef" :model="createForm" :rules="createRules">
        <a-form-item label="类型" name="label_type">
          <a-input
            disabled
            :value="activeKey === 'category' ? '分类' : '标签'"
          />
        </a-form-item>
        <a-form-item
          :label="`${activeKey === 'category' ? '分类名称' : '标签名称'}`"
          name="name"
        >
          <a-input v-model:value="createForm.name" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";
import { message, type FormInstance } from "ant-design-vue";
import {
  createLabelAPI,
  editLabelAPI,
  deleteLabelAPI,
  queryLabelListAPI,
} from "@/api/label";
import type { CreateLabelReq, LabelVO } from "@/types/label";

// 当前激活 tab（分类 or 标签）
const activeKey = ref<"category" | "tag">("category");
const data = reactive([
  { label: "文章分类", value: "category" },
  { label: "文章标签", value: "tag" },
]);

// 数据列表和加载状态
const labelList = ref<LabelVO[]>([]);
const loading = ref(false);
const editableMap = reactive<Record<string, LabelVO>>({});

const getLabelList = async () => {
  loading.value = true;
  const res = await queryLabelListAPI({
    page_no: 1,
    page_size: 10,
    label_type: activeKey.value,
  });
  labelList.value = res.data.list;
  loading.value = false;
};

// 监听 activeKey 切换刷新数据
watch(activeKey, getLabelList, { immediate: true });

// 新增逻辑
const createModalOpen = ref(false);
const createFormRef = ref<FormInstance>();
const createForm = ref<CreateLabelReq>({
  label_type: activeKey.value,
  name: "",
});
const createRules = {
  name: [
    { required: true, message: "请输入名称" },
    {
      validator: (_rule: any, value: string, callback: any) => {
        if (activeKey.value === "category" && value.length > 4) {
          callback(new Error("分类名称不能超过4个字符"));
        } else if (activeKey.value === "tag" && value.length > 10) {
          callback(new Error("标签名称不能超过10个字符"));
        } else {
          callback();
        }
      },
    },
  ],
};

const onHandleCreate = () => {
  createForm.value = {
    label_type: activeKey.value,
    name: "",
  };
  createModalOpen.value = true;
};

const handleCreateOk = async () => {
  await createFormRef.value?.validate();
  const res = await createLabelAPI(createForm.value);
  message.success(res.msg);
  createModalOpen.value = false;
  await getLabelList();
};

// 编辑逻辑
const onEdit = (record: LabelVO) => {
  editableMap[record.id] = { ...record };
};

const onCancel = (id: string) => {
  delete editableMap[id];
};

const onSave = async (record: LabelVO) => {
  const edited = editableMap[record.id];
  if (!edited) return;

  if (edited.name === record.name) {
    message.warning("没有修改");
    delete editableMap[record.id];
    return;
  }

  if (edited.label_type === "category" && edited.name.length > 4) {
    message.warning("分类名称不能超过4个字符");
    return;
  }

  if (edited.label_type === "tag" && edited.name.length > 10) {
    message.warning("标签名称不能超过10个字符");
    return;
  }

  await editLabelAPI(edited);
  message.success("修改成功");
  delete editableMap[record.id];
  await getLabelList();
};

// 删除逻辑
const onHandleDelete = async (record: LabelVO) => {
  await deleteLabelAPI(record.id);
  message.success("删除成功");
  await getLabelList();
};

// 表格列定义
const columns = [
  { title: "名称", key: "name" },
  { title: "操作", key: "action", width: 150, fixed: "right" },
];
</script>
