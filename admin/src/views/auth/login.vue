<template>
  <div
    class="flex items-center justify-center h-screen w-full bg-[url('/background.png')] bg-cover bg-center"
  >
    <a-card class="shadow-sm w-96">
      <div class="flex justify-center">
        <img
          class="my-6 w-32"
          :src="
            systemStore.themeMode === 'dark'
              ? '/logo-dark.png'
              : '/logo-light.png'
          "
        />
      </div>
      <a-form
        :model="LoginForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="Login"
      >
        <a-form-item
          label="用户名"
          name="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
        >
          <a-input v-model:value="LoginForm.username">
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
          <a-input-password v-model:value="LoginForm.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-form-item name="remember" no-style>
            <a-checkbox class="text-sm">记住我</a-checkbox>
          </a-form-item>
          <span class="float-right space-x-2 w-32">
            <a class="text-sm" href="">忘记密码</a>
            <a class="text-sm" href="">立即注册</a>
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
<script setup lang="ts">
import type { LoginReq } from "@/types/user";
import { userLoginAPI } from "@/api/user";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import { useUserStore, useSystemStore } from "@/store";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";

const router = useRouter();
const userStore = useUserStore();
const systemStore = useSystemStore();
const LoginForm: LoginReq = reactive({
  username: "",
  password: "",
});

const Login = async () => {
  const res = await userLoginAPI(LoginForm);
  const { access_token, refresh_token } = res.data;
  userStore.setAccessToken(access_token);
  userStore.setRefreshToken(refresh_token);
  message.success("登录成功");
  router.push("/");
};

const disabled = computed(() => {
  return !(LoginForm.username && LoginForm.password);
});
</script>
