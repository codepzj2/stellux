<template>
  <div
    class="flex items-center justify-center h-screen w-full login-background"
  >
    <a-card class="shadow-md w-96">
      <img class="m-auto w-2/5 h-3/5" :src="Logo" />
      <a-form
        :model="loginForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
      >
        <a-form-item
          label="用户名"
          name="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
        >
          <a-input v-model:value="loginForm.username">
            <template #prefix>
              <UserOutlined class="site-form-item-icon" />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item
          label="密码"
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
        >
          <a-input-password v-model:value="loginForm.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-form-item name="remember" no-style>
            <a-checkbox v-model:checked="loginForm.remember">记住我</a-checkbox>
          </a-form-item>
          <span class="float-right space-x-2 w-32">
            <a class="hover:text-blue-600" href="">忘记密码</a>
            <a class="hover:text-blue-600" href="">立即注册</a>
          </span>
        </a-form-item>

        <a-form-item class="flex justify-center">
          <a-button
            :disabled="disabled"
            type="primary"
            html-type="submit"
            class="w-48"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<script lang="ts" setup>
import { reactive, computed } from "vue";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import $message from "@/utils/message";
import type { LoginForm } from "@/api/interfaces/user";

import Logo from "@/assets/login/logo.svg";

const loginForm: LoginForm = reactive({
  username: "",
  password: "",
  remember: true,
});
const onFinish = (values: LoginForm) => {
  $message.success("登录成功");
  console.log("Success:", values);
};

const onFinishFailed = (errorInfo: unknown) => {
  console.error("登录失败", errorInfo);
  $message.error("登录失败，请重试");
};
const disabled = computed(() => {
  return !(loginForm.username && loginForm.password);
});
</script>
<style lang="scss" scoped>
.login-background {
  background: url("@/assets/login/background.png") no-repeat center center /
    cover;
}
</style>
