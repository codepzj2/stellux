<template>
  <div>
    <div class="flex justify-end items-center rounded-md px-4 py-2 my-2">
      <a-button type="primary" @click="onHandleCreate">新增用户</a-button>
    </div>
    <a-table :columns="columns" :data-source="userList">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'avatar'">
          <a-avatar :src="record.avatar" />
        </template>
        <template v-else-if="column.key === 'role_id'">
          <a-tag :color="roleColors[record.role_id as keyof typeof roleColors]">
            {{ roleNames[record.role_id as keyof typeof roleNames] }}
          </a-tag>
        </template>
        <template v-else-if="column.key === 'action'">
          <span>
            <a-button type="link" size="small" @click="onHandleEdit(record)">
              编辑
            </a-button>
            <a-divider type="vertical" />
            <a-popconfirm
              placement="bottomRight"
              title="确定删除该用户吗？"
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

    <!-- 新增弹窗 -->
    <a-modal
      v-model:open="createModalOpen"
      title="新增用户"
      @ok="handleCreateOk"
    >
      <a-form ref="createFormRef" :model="createForm" :rules="createRules">
        <a-form-item label="用户名" name="username">
          <a-input v-model:value="createForm.username" />
        </a-form-item>
        <a-form-item label="密码" name="password">
          <a-input-password v-model:value="createForm.password" />
        </a-form-item>
        <a-form-item label="角色" name="role_id">
          <a-select v-model:value="createForm.role_id">
            <a-select-option :value="0">管理员</a-select-option>
            <a-select-option :value="1">普通用户</a-select-option>
            <a-select-option :value="2">游客</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="昵称" name="nickname">
          <a-input v-model:value="createForm.nickname" />
        </a-form-item>
        <a-form-item label="邮箱" name="email">
          <a-input v-model:value="createForm.email" />
        </a-form-item>
        <a-form-item label="性别" name="sex">
          <a-select v-model:value="createForm.sex">
            <a-select-option value="男">男</a-select-option>
            <a-select-option value="女">女</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="爱好" name="hobby">
          <a-input v-model:value="createForm.hobby" />
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 编辑弹窗 -->
    <a-modal v-model:open="editModalOpen" title="编辑用户" @ok="handleEditOk">
      <a-form ref="editFormRef" :model="editForm" :rules="editRules">
        <a-form-item label="用户名" name="username">
          <a-input v-model:value="editForm.username" />
        </a-form-item>
        <a-form-item label="角色" name="role_id">
          <a-select v-model:value="editForm.role_id">
            <a-select-option :value="0">管理员</a-select-option>
            <a-select-option :value="1">普通用户</a-select-option>
            <a-select-option :value="2">游客</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="昵称" name="nickname">
          <a-input v-model:value="editForm.nickname" />
        </a-form-item>
        <a-form-item label="邮箱" name="email">
          <a-input v-model:value="editForm.email" />
        </a-form-item>
        <a-form-item label="性别" name="sex">
          <a-select v-model:value="editForm.sex">
            <a-select-option value="男">男</a-select-option>
            <a-select-option value="女">女</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="爱好" name="hobby">
          <a-input v-model:value="editForm.hobby" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { message, type FormInstance } from "ant-design-vue";
import {
  createUserAPI,
  editUserAPI,
  getUserListAPI,
  deleteUserAPI,
} from "@/api/user";
import type { CreateUserReq, EditUserReq, UserInfoVO } from "@/types/user";
import { Code, roleNames, roleColors } from "@/global";

const userList = ref<UserInfoVO[]>([]);
const createModalOpen = ref(false);
const createFormRef = ref<FormInstance>();
const createForm = ref<CreateUserReq>({
  username: "",
  password: "",
  nickname: "",
  role_id: 0,
  avatar: "",
  email: "",
  sex: "",
  hobby: "",
});
const createRules = ref({
  username: [{ required: true, message: "请输入用户名" }],
  password: [{ required: true, message: "请输入密码" }],
  role_id: [{ required: true, message: "请选择角色" }],
});

const editModalOpen = ref(false);
const editFormRef = ref<FormInstance>();
const editForm = ref<EditUserReq>({
  id: "",
  username: "",
  nickname: "",
  role_id: 0,
  avatar: "",
  email: "",
  sex: "",
  hobby: "",
});
const originalEditForm = ref<EditUserReq>({ ...editForm.value });
const editRules = ref({
  username: [{ required: true, message: "请输入用户名" }],
  role_id: [{ required: true, message: "请选择角色" }],
});

const getUserList = async () => {
  const res = await getUserListAPI({ page_no: 1, page_size: 10 });
  userList.value = res.data.list;
};

const onHandleCreate = () => {
  createModalOpen.value = true;
};
const onHandleEdit = (record: UserInfoVO) => {
  editForm.value = { ...record };
  originalEditForm.value = { ...record };
  editModalOpen.value = true;
};
const onHandleDelete = async (record: UserInfoVO) => {
  const res = await deleteUserAPI(record.id);
  if (res.code === Code.RequestSuccess) {
    message.success(res.msg);
    await getUserList();
  }
};
const handleCreateOk = () => {
  createFormRef.value?.validate().then(async () => {
    const res = await createUserAPI(createForm.value);
    if (res.code === Code.RequestSuccess) {
      message.success(res.msg);
      createModalOpen.value = false;
      clearCreateForm();
      await getUserList();
    }
  });
};

const clearCreateForm = () => {
  createForm.value = {
    username: "",
    password: "",
    nickname: "",
    role_id: 0,
    avatar: "",
    email: "",
    sex: "",
    hobby: "",
  };
};
const clearEditForm = () => {
  editForm.value = {
    id: "",
    username: "",
    nickname: "",
    role_id: 0,
    avatar: "",
    email: "",
    sex: "",
    hobby: "",
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
    const res = await editUserAPI(editForm.value);
    message.success(res.msg);
    editModalOpen.value = false;
    clearEditForm();
    await getUserList();
  });
};

onMounted(() => {
  getUserList();
});

const columns = [
  { title: "头像", dataIndex: "avatar", key: "avatar" },
  { title: "用户名", dataIndex: "username", key: "username" },
  { title: "昵称", dataIndex: "nickname", key: "nickname" },
  { title: "角色", dataIndex: "role_id", key: "role_id" },
  { title: "邮箱", dataIndex: "email", key: "email" },
  { title: "操作", key: "action", width: 300, fixed: "right" },
];
</script>
