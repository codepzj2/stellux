<template>
  <div>
    <a-form
      :model="form"
      :label-col="{ span: 4 }"
      :wrapper-col="{ span: 14 }"
      :rules="rules"
      ref="formRef"
    >
      <a-form-item label="昵称" name="nickname">
        <a-input v-model:value="form.nickname" />
      </a-form-item>

      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="form.email" />
      </a-form-item>

      <a-form-item label="头像" name="avatar">
        <div v-if="form.avatar" class="w-[100px] h-[100px] relative group">
          <img
            :src="form.avatar"
            class="w-[100px] h-[100px] object-contain cursor-pointer"
            @click="getThumbnailList"
          />

          <div
            class="absolute top-1 right-1 z-10 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
          >
            <a-button
              type="primary"
              shape="circle"
              danger
              size="small"
              class="!scale-80"
              @click.stop="form.avatar = ''"
            >
              <CloseOutlined />
            </a-button>
          </div>
        </div>
        <div
          v-else
          class="w-[100px] h-[100px] flex items-center justify-center border-2 border-dashed border-gray-300 rounded-md cursor-pointer text-gray-400 transition-colors"
          @click="getThumbnailList"
        >
          <span class="text-sm">选择头像</span>
        </div>
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 4, span: 14 }">
        <div class="flex justify-end">
          <a-button type="primary" @click="handleSubmit">提交</a-button>
        </div>
      </a-form-item>
    </a-form>
    <a-modal
      width="100%"
      wrap-class-name="full-modal"
      :open="thumbnailModalOpen"
      @update:open="thumbnailModalOpen = $event"
      title="选择头像"
      @ok="thumbnailModalOpen = false"
    >
      <PhotoWall
        mode="picture-card"
        type="select"
        :list="page.list"
        @selected-photos="handleSelectAvatar"
      />
      <div class="flex justify-end my-4">
        <a-pagination
          v-model:current="page.page_no"
          :total="page.total_count"
          show-less-items
          @change="getThumbnailList"
        />
      </div>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { CloseOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { getUserInfoAPI, updateUserAPI } from "@/api/user";
import type { FormInstance, Rule } from "ant-design-vue/es/form";
import { useUserStore } from "@/store";

import type { PageData } from "@/types/dto";
import type { FileVO } from "@/types/file";
import { queryFileList } from "@/api/file";

const formRef = ref<FormInstance>();

const thumbnailModalOpen = ref(false);
const page = reactive<PageData<FileVO>>({
  page_no: 1,
  page_size: 10,
  total_count: 0,
  total_page: 0,
  list: [],
});

const getThumbnailList = async () => {
  const res = await queryFileList({
    page_no: page.page_no,
    page_size: page.page_size,
  });
  page.list = res.data.list;
  page.total_count = res.data.total_count;
  thumbnailModalOpen.value = true;
};

const handleSelectAvatar = (pictures: string[]) => {
  form.value.avatar = pictures[0];
};

const form = ref({
  nickname: "",
  avatar: "",
  email: "",
});

const userStore = useUserStore();

const rules: Record<string, Rule[]> = {
  nickname: [{ required: true, message: "请输入昵称" }],
  email: [{ required: true, message: "请输入邮箱" }],
  avatar: [{ required: true, message: "请上传头像" }],
};

const handleSubmit = async () => {
  await formRef.value?.validate();
  await updateUserAPI({
    id: userStore.userInfo?.id as string,
    nickname: form.value.nickname,
    avatar: form.value.avatar,
    email: form.value.email,
  });
  await reloadUserInfo();
  message.success("更新成功");
};

const reloadUserInfo = async () => {
  const res = await getUserInfoAPI();
  userStore.setUserInfo(res.data);
};

onMounted(async () => {
  const res = await getUserInfoAPI();
  form.value = res.data;
});
</script>
