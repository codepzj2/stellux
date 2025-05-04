<template>
  <div>
    <a-table :columns="columns" :data-source="labelList">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <a-tag>
            {{ record.name }}
          </a-tag>
        </template>
        <template v-else-if="column.key === 'action'">
          <span>
            <a-button type="link" size="small" @click="onHandleEdit(record)"
              >编辑</a-button
            >
            <a-divider type="vertical" />
            <a-popconfirm
              placement="bottomRight"
              title="确定删除该分类吗？"
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
      title="新增分类"
      @ok="handleCreateOk"
    >
      <a-form ref="createFormRef" :model="createForm" :rules="createRules">
        <a-form-item label="分类类型" name="label_type">
          <a-input disabled v-model:value="createForm.label_type" />
        </a-form-item>
        <a-form-item label="分类名称" name="name">
          <a-input v-model:value="createForm.name" />
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 编辑弹窗 -->
    <a-modal v-model:open="editModalOpen" title="编辑分类" @ok="handleEditOk">
      <a-form ref="editFormRef" :model="editForm" :rules="editRules">
        <a-form-item label="分类类型" name="label_type">
          <a-input disabled v-model:value="editForm.label_type" />
        </a-form-item>
        <a-form-item label="分类名称" name="name">
          <a-input v-model:value="editForm.name" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { message, type FormInstance } from "ant-design-vue";
import {
  createLabelAPI,
  editLabelAPI,
  getLabelListAPI,
  deleteLabelAPI,
} from "@/api/label";
import type { CreateLabelReq, EditLabelReq, LabelVO } from "@/types/label";
import { Code } from "@/global";
const labelList = ref<LabelVO[]>([]);
const createModalOpen = ref(false);
const createFormRef = ref<FormInstance>();
const createForm = ref<CreateLabelReq>({
  label_type: "category",
  name: "",
});
const createRules = ref({
  label_type: [{ required: true, message: "请输入分类类型" }],
  name: [{ required: true, message: "请输入分类名称" }],
});

const editModalOpen = ref(false);
const editFormRef = ref<FormInstance>();
const editForm = ref<EditLabelReq>({
  id: "",
  label_type: "category",
  name: "",
});
const originalEditForm = ref<EditLabelReq>({ ...editForm.value });
const editRules = ref({
  label_type: [{ required: true, message: "请输入分类类型" }],
  name: [{ required: true, message: "请输入分类名称" }],
});

const getLabelList = async () => {
  const res = await getLabelListAPI({
    page_no: 1,
    page_size: 10,
    label_type: "category",
  });
  labelList.value = res.data.list;
};

const onHandleCreate = () => {
  createModalOpen.value = true;
};
const onHandleEdit = (record: LabelVO) => {
  editForm.value = { ...record };
  originalEditForm.value = { ...record };
  editModalOpen.value = true;
};
const onHandleDelete = async (record: LabelVO) => {
  const res = await deleteLabelAPI(record.id);
  if (res.code === Code.RequestSuccess) {
    message.success(res.msg);
    await getLabelList();
  }
};
const handleCreateOk = () => {
  createFormRef.value?.validate().then(async () => {
    const res = await createLabelAPI(createForm.value);
    if (res.code === Code.RequestSuccess) {
      message.success(res.msg);
      createModalOpen.value = false;
      clearCreateForm();
      await getLabelList();
    }
  });
};

const clearCreateForm = () => {
  createForm.value = {
    label_type: "",
    name: "",
  };
};
const clearEditForm = () => {
  editForm.value = {
    id: "",
    label_type: "",
    name: "",
  };
  originalEditForm.value = { ...editForm.value };
};
const handleEditOk = () => {
  if (
    JSON.stringify(editForm.value) === JSON.stringify(originalEditForm.value)
  ) {
    message.warning("没有修改");
    return;
  }
  editFormRef.value?.validate().then(async () => {
    const res = await editLabelAPI(editForm.value);
    message.success(res.msg);
    editModalOpen.value = false;
    clearEditForm();
    await getLabelList();
  });
};

onMounted(() => {
  getLabelList();
});

const columns = [
  { title: "分类名称", dataIndex: "name", key: "name" },
  { title: "操作", key: "action", width: 300, fixed: "right" },
];
</script>
