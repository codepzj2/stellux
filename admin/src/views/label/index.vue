<template>
  <div>
    <a-tabs v-model:activeKey="activeKey" type="card" @change="getLabelList">
      <a-tab-pane tab="文章分类" key="category" />
      <a-tab-pane tab="文章标签" key="tag" />
    </a-tabs>
    <PageHeader>
      <template #right>
        <a-button type="primary" @click="onHandleCreate">新增</a-button>
      </template>
    </PageHeader>

    <a-table :columns="columns" :data-source="labelList" rowKey="id">
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
            :value="createForm.label_type === 'category' ? '分类' : '标签'"
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
import { ref, onMounted, reactive } from "vue";
import { message, type FormInstance } from "ant-design-vue";
import PageHeader from "@/components/PageHeader.vue";
import {
  createLabelAPI,
  editLabelAPI,
  deleteLabelAPI,
  queryLabelListAPI,
} from "@/api/label";
import type { CreateLabelReq, LabelVO } from "@/types/label";
import { Code } from "@/global";

// 当前选中的标签类型
const activeKey = ref<"category" | "tag">("category");

// 标签/分类列表
const labelList = ref<LabelVO[]>([]);

// 可编辑的项
const editableMap = reactive<Record<string, LabelVO>>({});

// 获取列表
const getLabelList = async () => {
  const res = await queryLabelListAPI({
    page_no: 1,
    page_size: 10,
    label_type: activeKey.value,
  });

  labelList.value = res.data.list;
};

// 新增相关
const createModalOpen = ref(false);
const createFormRef = ref<FormInstance>();
const createForm = ref<CreateLabelReq>({
  label_type: activeKey.value,
  name: "",
});
const createRules = {
  label_type: [{ required: true, message: "请输入类型" }],
  name: [
    { required: true, message: "请输入名称" },
    {
      validator: (_rule: any, value: string, callback: any) => {
        if (value.length > 4) {
          callback(new Error("名称不能超过4个字符"));
        }
        callback();
      },
    },
  ],
};

const onHandleCreate = () => {
  createForm.value = { label_type: activeKey.value, name: "" };
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

// 删除
const onHandleDelete = async (record: LabelVO) => {
  await deleteLabelAPI(record.id);
  message.success("删除成功");
  await getLabelList();
};

onMounted(() => {
  getLabelList();
});

// 表格列
const columns = [
  { title: "名称", dataIndex: "name", key: "name" },
  { title: "操作", key: "action", width: 300 },
];
</script>

<style lang="scss" scoped></style>
