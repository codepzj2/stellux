<template>
  <div>
    <a-form
      :model="form"
      :label-col="{ span: 4 }"
      :wrapper-col="{ span: 14 }"
      :rules="rules"
      @finish="handleFinish"
      ref="formRef"
    >
      <a-form-item label="旧密码" name="oldPassword">
        <a-input v-model:value="form.oldPassword" />
      </a-form-item>
      <a-form-item label="新密码" name="newPassword">
        <a-input v-model:value="form.newPassword" />
      </a-form-item>
      <a-form-item label="重复新密码" name="confirmPassword">
        <a-input v-model:value="form.confirmPassword" />
      </a-form-item>
      <a-form-item :wrapper-col="{ offset: 4, span: 14 }">
        <div class="flex justify-end">
          <a-button type="primary" @click="handleFinish">重置密码</a-button>
        </div>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { getUserInfoAPI, updateUserPasswordAPI } from "@/api/user";
import { message, type FormInstance } from "ant-design-vue";
import type { UserInfoVO } from "@/types/user";

const formRef = ref<FormInstance>();
const form = ref({
  oldPassword: "",
  newPassword: "",
  confirmPassword: "",
});

const checkNewPassword = (_rule: any, value: string, callback: any) => {
  if (value !== form.value.newPassword) {
    callback(new Error("两次输入的密码不一致"));
  } else {
    callback();
  }
};

const rules = ref({
  oldPassword: [{ required: true, message: "请输入旧密码" }],
  newPassword: [{ required: true, message: "请输入新密码" }],
  confirmPassword: [
    { required: true, message: "请输入重复新密码" },
    { validator: checkNewPassword, trigger: "blur" },
  ],
});

const userInfo = ref<UserInfoVO>();

const handleFinish = async () => {
  await formRef.value?.validate();
  await updateUserPasswordAPI({
    id: userInfo.value?.id as string,
    old_password: form.value.oldPassword,
    new_password: form.value.newPassword,
  });
  message.success("密码重置成功");
};

onMounted(async () => {
  const res = await getUserInfoAPI();
  userInfo.value = res.data;
});
</script>
